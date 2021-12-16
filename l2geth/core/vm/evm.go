// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package vm

import (
	"bytes"
	"fmt"
	"math/big"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rollup/dump"
	"github.com/ethereum/go-ethereum/rollup/rcfg"
	"github.com/ethereum/go-ethereum/rollup/util"
	"golang.org/x/crypto/sha3"

	//"github.com/ethereum/go-ethereum/common/bytes"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"

	// "github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
)

// emptyCodeHash is used by create to ensure deployment is disallowed to already
// deployed contract addresses (relevant after the account abstraction).
var emptyCodeHash = crypto.Keccak256Hash(nil)

type (
	// CanTransferFunc is the signature of a transfer guard function
	CanTransferFunc func(StateDB, common.Address, *big.Int) bool
	// TransferFunc is the signature of a transfer function
	TransferFunc func(StateDB, common.Address, common.Address, *big.Int)
	// GetHashFunc returns the n'th block hash in the blockchain
	// and is used by the BLOCKHASH EVM op code.
	GetHashFunc func(uint64) common.Hash
)

// run runs the given contract and takes care of running precompiles with a fallback to the byte code interpreter.
func run(evm *EVM, contract *Contract, input []byte, readOnly bool) ([]byte, error) {
	if contract.CodeAddr != nil {
		precompiles := PrecompiledContractsHomestead
		if evm.chainRules.IsByzantium {
			precompiles = PrecompiledContractsByzantium
		}
		if evm.chainRules.IsIstanbul {
			precompiles = PrecompiledContractsIstanbul
		}
		if p := precompiles[*contract.CodeAddr]; p != nil {
			return RunPrecompiledContract(p, input, contract)
		}
	}
	for _, interpreter := range evm.interpreters {
		if interpreter.CanRun(contract.Code) {
			if evm.interpreter != interpreter {
				// Ensure that the interpreter pointer is set back
				// to its current value upon return.
				defer func(i Interpreter) {
					evm.interpreter = i
				}(evm.interpreter)
				evm.interpreter = interpreter
			}

			//log.Debug("TURING processing contract", "Address", contract.Address().Hex())

			return interpreter.Run(contract, input, readOnly)
		}
	}
	return nil, ErrNoCompatibleInterpreter
}

// Context provides the EVM with auxiliary information. Once provided
// it shouldn't be modified.
type Context struct {
	// CanTransfer returns whether the account contains
	// sufficient ether to transfer the value
	CanTransfer CanTransferFunc
	// Transfer transfers ether from one account to the other
	Transfer TransferFunc
	// GetHash returns the hash corresponding to n
	GetHash GetHashFunc

	// Message information
	Origin   common.Address // Provides information for ORIGIN
	GasPrice *big.Int       // Provides information for GASPRICE

	// Block information
	Coinbase    common.Address // Provides information for COINBASE
	GasLimit    uint64         // Provides information for GASLIMIT
	BlockNumber *big.Int       // Provides information for NUMBER
	Time        *big.Int       // Provides information for TIME
	Difficulty  *big.Int       // Provides information for DIFFICULTY

	// OVM information
	L1BlockNumber *big.Int // Provides information for L1BLOCKNUMBER
}

// FIXME - should move this somewhere else.
// For now, only caches the most recent result. Can be extended with a map of
// multiple requests, but that needs some logic to expire/purge old entries.
// "key" for now is simply the request URL. May need tighter scope in the future,
// e.g. per contract. That would also allow different expiration thresholds for
// different users.
//
// Another future enhancement could be to allow an external program to pre-load
// results into the cache on a periodic basis (e.g. updating the latest market
// prices for various tokens). Contracts would then be able to access this data
// without the latency of making an off-chain JSON-RPC call. This is similar to
// some of the earlier concepts for a "Turing" mechanism.

var turingCache struct {
	lock    sync.RWMutex
	expires time.Time
	key     common.Hash
	value   []byte
}

// EVM is the Ethereum Virtual Machine base object and provides
// the necessary tools to run a contract on the given state with
// the provided context. It should be noted that any error
// generated through any of the calls should be considered a
// revert-state-and-consume-all-gas operation, no checks on
// specific errors should ever be performed. The interpreter makes
// sure that any errors generated are to be considered faulty code.
//
// The EVM should never be reused and is not thread safe.
type EVM struct {
	// Context provides auxiliary blockchain related information
	Context
	// StateDB gives access to the underlying state
	StateDB StateDB
	// Depth is the current call stack
	depth int

	// chainConfig contains information about the current chain
	chainConfig *params.ChainConfig
	// chain rules contains the chain rules for the current epoch
	chainRules params.Rules
	// virtual machine configuration options used to initialise the
	// evm.
	vmConfig Config
	// global (to this context) ethereum virtual machine
	// used throughout the execution of the tx.
	interpreters []Interpreter
	interpreter  Interpreter
	// abort is used to abort the EVM calling operations
	// NOTE: must be set atomically
	abort int32
	// callGasTemp holds the gas available for the current call. This is needed because the
	// available gas is calculated in gasCall* according to the 63/64 rule and later
	// applied in opCall*.
	callGasTemp uint64
}

// NewEVM returns a new EVM. The returned EVM is not thread safe and should
// only ever be used *once*.
func NewEVM(ctx Context, statedb StateDB, chainConfig *params.ChainConfig, vmConfig Config) *EVM {
	evm := &EVM{
		Context:      ctx,
		StateDB:      statedb,
		vmConfig:     vmConfig,
		chainConfig:  chainConfig,
		chainRules:   chainConfig.Rules(ctx.BlockNumber),
		interpreters: make([]Interpreter, 0, 1),
	}

	if chainConfig.IsEWASM(ctx.BlockNumber) {
		// to be implemented by EVM-C and Wagon PRs.
		// if vmConfig.EWASMInterpreter != "" {
		//  extIntOpts := strings.Split(vmConfig.EWASMInterpreter, ":")
		//  path := extIntOpts[0]
		//  options := []string{}
		//  if len(extIntOpts) > 1 {
		//    options = extIntOpts[1..]
		//  }
		//  evm.interpreters = append(evm.interpreters, NewEVMVCInterpreter(evm, vmConfig, options))
		// } else {
		// 	evm.interpreters = append(evm.interpreters, NewEWASMInterpreter(evm, vmConfig))
		// }
		panic("No supported ewasm interpreter yet.")
	}

	// vmConfig.EVMInterpreter will be used by EVM-C, it won't be checked here
	// as we always want to have the built-in EVM as the failover option.
	evm.interpreters = append(evm.interpreters, NewEVMInterpreter(evm, vmConfig))
	evm.interpreter = evm.interpreters[0]

	return evm
}

// Cancel cancels any running EVM operation. This may be called concurrently and
// it's safe to be called multiple times.
func (evm *EVM) Cancel() {
	atomic.StoreInt32(&evm.abort, 1)
}

// Cancelled returns true if Cancel has been called
func (evm *EVM) Cancelled() bool {
	return atomic.LoadInt32(&evm.abort) == 1
}

// Interpreter returns the current interpreter
func (evm *EVM) Interpreter() Interpreter {
	return evm.interpreter
}
// In response to an off-chain Turing request, obtain the requested data and
// rewrite the parameters so that the contract can be called a second time.
// FIXME - needs error handling. For now, bails out and lets the contract
// be called a second time with the original parameters. 2nd failure is not intercepted.

func bobaTuringRandom(input []byte) hexutil.Bytes {

	var ret hexutil.Bytes

	rest := input[4:]

	//some things are easier with a hex string
	inputHexUtil := hexutil.Bytes(input)

	/* The input and calldata have a well defined structure
	    1/ The methodID (4 bytes)
	    2/ The rType (32 bytes)
	    3/ The return placeholder uint256
	*/

	// If things fail, we'll return an integer parameter which should fail a
	// "require" in the contract without generating another TURING marker.
	// FIXME - would be cleaner to return nil and put better error handling
	// into l2geth to avoid that second call into the contract.
	
	methodID := make([]byte, 4)
	copy(methodID, inputHexUtil[0:4])

	// Check the rType
	// 1 for Request, 2 for Response, integer >= 10 for various failures
	rType := int(rest[31])
	if rType != 1 {
		log.Warn("TURING-1 bobaTuringRandom:Wrong state (rType != 1)", "rType", rType)
		return append(methodID, hexutil.MustDecode(fmt.Sprintf("0x%064x", 10))...) // wrong input state
	}

	rlen := len(rest) 
	if rlen < 2*32 {
		log.Warn("TURING-2 bobaTuringRandom:Calldata too short", "len < 2*32", rlen)
		return append(methodID, hexutil.MustDecode(fmt.Sprintf("0x%064x", 11))...) // calldata too short
	}

	//generate a Uint64 random number
	randomUint64 := rand.Uint64()

	log.Debug("TURING-3 bobaTuringRandom:Random number",
		"randomUint32", fmt.Sprintf("0x%064x", randomUint64))

	// build the calldata
	ret = append(methodID, hexutil.MustDecode(fmt.Sprintf("0x%064x", 2))...) // the usual prefix and the rType, now changed to 2
	ret = append(ret, hexutil.MustDecode(fmt.Sprintf("0x%064x", randomUint64))...)

	log.Debug("TURING-4 bobaTuringRandom:Modified parameters",
		"newValue", ret)

	return ret
}

// In response to an off-chain Turing request, obtain the requested data and
// rewrite the parameters so that the contract can be called a second time.
// FIXME - needs error handling. For now, bails out and lets the contract
// be called a second time with the original parameters. 2nd failure is not intercepted.

func bobaTuringCall(input []byte) hexutil.Bytes {

	var responseStringEnc string
	var responseString []byte
	var ret hexutil.Bytes

	rest         := input[4:]

	//some things are easier with a hex string
	inputHexUtil := hexutil.Bytes(input)
	restHexUtil  := inputHexUtil[4:]

	/* The input and calldata have a well defined structure
	    1/ The methodID (4 bytes)
	    2/ The rType (32 bytes)
	    3/ Data offset 1 - beginning of URL string (32 bytes)
	    4/ Data offset 2 - beginning of payload (32 bytes)
	    5/ URL string length (32 bytes)
	    6/ URL string - either 32 or 64 bytes
	    7/ Payload length (32 bytes)
	    8/ Payload data - variable (at least 32 bytes)
	    This means that the calldata are allways >= 8*32
	*/

	// If things fail, we'll return an integer parameter which should fail a
	// "require" in the contract without generating another TURING marker.
	// FIXME - would be cleaner to return nil and put better error handling
	// into l2geth to avoid that second call into the contract.
	
	methodID := make([]byte, 4)
	copy(methodID, inputHexUtil[0:4])

	// Check the rType
	// 1 for Request, 2 for Response, integer >= 10 for various failures
	rType := int(rest[31])
	if rType != 1 {
		log.Warn("TURING-1 bobaTuringCall:Wrong state (rType != 1)", "rType", rType)
		return append(methodID, hexutil.MustDecode(fmt.Sprintf("0x%064x", 10))...) //wrong input state
	}

	rlen := len(rest) 
	if rlen < 8*32 {
		log.Warn("TURING-2 bobaTuringCall:Calldata too short", "len < 8*32", rlen)
		return append(methodID, hexutil.MustDecode(fmt.Sprintf("0x%064x", 11))...) //calldata too short
	}

	// Check the URL length
	// Note: we do not handle URLs that are longer than 64 characters

	// A micro-ABI decoder... this works because we know that all these numbers can never exceed 256
	// Since the rType is 32 bytes and the three headers are 32 bytes each, the max possible value 
	// of any of these numbers is 32 + 32 + 32 + 32 + 64 = 192
	startIDXurl     := int(rest[ 63]) + 32 //the +32 means that we are going right for the actual string
	startIDXpayload := int(rest[ 95])
	lengthURL       := int(rest[127])
	
	if lengthURL > 64 {
		log.Warn("TURING-3 bobaTuringCall:URL too long", "urlLength", lengthURL)
		return append(methodID, hexutil.MustDecode(fmt.Sprintf("0x%064x", 12))...) //URL string > 64 bytes
	}

	// The url we are going to query
	url := string(rest[startIDXurl:startIDXurl+lengthURL])

	// At this point, we have the API endpoint and the payload that needs to go there...
	payload := restHexUtil[startIDXpayload:]
	
	log.Debug("TURING-4 bobaTuringCall:Have URL and payload",
		"url", url,
		"payload", payload)

	// don't recalculate unless actually needed...
	// if we have data and if the time is right,
	// replace the calldata with the cached value and return that
	turingCache.lock.Lock()
	if crypto.Keccak256Hash(rest) == turingCache.key && time.Now().Before(turingCache.expires) {
		ret = turingCache.value
	}
	turingCache.lock.Unlock()

	if ret != nil {
		log.Debug("TURING-5 bobaTuringCall:Found cached result; returning that",
			"key", crypto.Keccak256Hash(rest),
			"cached", ret)
		return ret
	}

	client, err := rpc.Dial(url)

	if client != nil {
		log.Debug("TURING-6 bobaTuringCall:Calling off-chain client at", "url", url)
		if err := client.Call(&responseStringEnc, "turing", payload); err != nil {
			log.Warn("TURING-7 bobaTuringCall:Client error", "err", err)
			return append(methodID, hexutil.MustDecode(fmt.Sprintf("0x%064x", 13))...) //Client Error
		}
		responseString, err = hexutil.Decode(responseStringEnc)
		if err != nil {
			log.Warn("TURING-8 bobaTuringCall:Error decoding responseString", "err", err)
			return append(methodID, hexutil.MustDecode(fmt.Sprintf("0x%064x", 14))...) //Client Response Decode Error
		}
	} else {
		log.Warn("TURING-9 bobaTuringCall:Failed to create client for off-chain request", "err", err)
		return append(methodID, hexutil.MustDecode(fmt.Sprintf("0x%064x", 15))...) //Could not create client
	}

	log.Debug("TURING-10 bobaTuringCall:Have valid response from offchain API",
		"Target", url,
		"Payload", payload,
		"responseStringEnc", responseStringEnc,
		"Response", responseString)

	// build the calldata:
	ret = append(methodID, hexutil.MustDecode(fmt.Sprintf("0x%064x", 2))...) // the usual prefix and the rType, now changed to 2
	ret = append(ret, restHexUtil[32:startIDXpayload]...)         // the unmodified offsets and the first dynamic data type
	ret = append(ret, responseString...)                          // and the data themselves

	log.Debug("TURING-11 bobaTuringCall:Modified parameters",
		"newValue", ret)

	// saving the modified calldata in the cache
	turingCache.lock.Lock()
	turingCache.key = crypto.Keccak256Hash(rest)
	turingCache.expires = time.Now().Add(2 * time.Second)
	turingCache.value = ret
	turingCache.lock.Unlock()

	log.Debug("TURING-12 bobaTuringCall:TuringCache entry stored for",
		"key", crypto.Keccak256Hash(rest))

	return ret
}

// Call executes the contract associated with the addr with the given input as
// parameters. It also handles any necessary value transfer required and takes
// the necessary steps to create accounts and reverses the state in case of an
// execution error or failed value transfer.
func (evm *EVM) Call(caller ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {

	// log.Debug("TURING entering Call",
	// 	"depth", evm.depth,
	// 	"addr", addr,
	// 	"input", hexutil.Bytes(input),
	// 	"gas", gas)

	if evm.vmConfig.NoRecursion && evm.depth > 0 {
		return nil, gas, nil
	}

	// Fail if we're trying to execute above the call depth limit
	if evm.depth > int(params.CallCreateDepth) {
		return nil, gas, ErrDepth
	}
	// Fail if we're trying to transfer more than the available balance
	if !evm.Context.CanTransfer(evm.StateDB, caller.Address(), value) {
		return nil, gas, ErrInsufficientBalance
	}

	var (
		to       = AccountRef(addr)
		snapshot = evm.StateDB.Snapshot()
	)
	if !evm.StateDB.Exist(addr) {
		precompiles := PrecompiledContractsHomestead
		if evm.chainRules.IsByzantium {
			precompiles = PrecompiledContractsByzantium
		}
		if evm.chainRules.IsIstanbul {
			precompiles = PrecompiledContractsIstanbul
		}
		if precompiles[addr] == nil && evm.chainRules.IsEIP158 && value.Sign() == 0 {
			// Calling a non existing account, don't do anything, but ping the tracer
			if evm.vmConfig.Debug && evm.depth == 0 {
				evm.vmConfig.Tracer.CaptureStart(caller.Address(), addr, false, input, gas, value)
				evm.vmConfig.Tracer.CaptureEnd(ret, 0, 0, nil)
			}
			return nil, gas, nil
		}
		evm.StateDB.CreateAccount(addr)
	}
	evm.Transfer(evm.StateDB, caller.Address(), to.Address(), value)
	// Initialise a new contract and set the code that is to be used by the EVM.
	// The contract is a scoped environment for this execution context only.
	contract := NewContract(caller, to, value, gas)
	contract.SetCallCode(&addr, evm.StateDB.GetCodeHash(addr), evm.StateDB.GetCode(addr))

	// Even if the account has no code, we need to continue because it might be a precompile
	start := time.Now()

	// Capture the tracer start/end events in debug mode
	if evm.vmConfig.Debug && evm.depth == 0 {
		evm.vmConfig.Tracer.CaptureStart(caller.Address(), addr, false, input, gas, value)

		defer func() { // Lazy evaluation of the parameters
			evm.vmConfig.Tracer.CaptureEnd(ret, gas-contract.Gas, time.Since(start), err)
		}()
	}
	ret, err = run(evm, contract, input, false)

	// log.Debug("TURING evm.go run",
	// 	"contract", contract.CodeAddr,
	// 	"ret", hexutil.Bytes(ret),
	// 	"err", err)

	if err != nil {

		//let's see if this is a Turing Compute Request
		isTuring := bytes.Contains(ret, []byte("TURING_"))

		//let's see if this is a Turing Compute Request
		isGetRand := bytes.Contains(ret, []byte("RANDOM_"))

		if isTuring || isGetRand {

			log.Debug("TCR_",
				"err", err,
				"ret", hexutil.Bytes(ret),
				"input", hexutil.Bytes(input),
				"contract", contract.CodeAddr)

			// at this point we have a call that has reverted AND the ret includes TURING_ or RANDOM_ somewhere in it
			log.Debug("TURING-M1 calling with", "input", input)

			// bobaTuringCall takes the original calldata, figures out what needs
			// to be done, and then synthesizes a 'new_in' calldata that does not 
			// lead to the revert (by changing rType from 1 to 2)
			var updated_input hexutil.Bytes

			if isTuring {
				updated_input = bobaTuringCall(input)
			} else if isGetRand {
				updated_input = bobaTuringRandom(input)
			}
			//evm.StateDB.RevertToSnapshot(snapshot) // thoughts?

			log.Debug("TURING-M2 replay with modified calldata", "modifiedCalldata", updated_input)

			// and then rerun the call with the modified calldata
			// this returns the API response to the caller 
			ret, err = run(evm, contract, updated_input, false)

			log.Debug("TURING-M3 received replay response",
				"err", err,
				"ret", hexutil.Bytes(ret))
		}
	}

	// When an error was returned by the EVM or when setting the creation code
	// above we revert to the snapshot and consume any gas remaining. Additionally
	// when we're in homestead this also counts for code storage gas errors.
	if err != nil {
		evm.StateDB.RevertToSnapshot(snapshot)
		log.Debug("TURING evm.go errExecutionReverted")
		if err != errExecutionReverted {
			contract.UseGas(contract.Gas)
		}
	}

	// log.Debug("TURING exiting Call",
	// 	"depth", evm.depth,
	// 	"addr", addr,
	// 	"ret", hexutil.Bytes(ret),
	// 	"err", err)

	return ret, contract.Gas, err
}

// CallCode executes the contract associated with the addr with the given input
// as parameters. It also handles any necessary value transfer required and takes
// the necessary steps to create accounts and reverses the state in case of an
// execution error or failed value transfer.
//
// CallCode differs from Call in the sense that it executes the given address'
// code with the caller as context.
func (evm *EVM) CallCode(caller ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {
	if evm.vmConfig.NoRecursion && evm.depth > 0 {
		return nil, gas, nil
	}

	// Fail if we're trying to execute above the call depth limit
	if evm.depth > int(params.CallCreateDepth) {
		return nil, gas, ErrDepth
	}
	// Fail if we're trying to transfer more than the available balance
	if !evm.CanTransfer(evm.StateDB, caller.Address(), value) {
		return nil, gas, ErrInsufficientBalance
	}

	var (
		snapshot = evm.StateDB.Snapshot()
		to       = AccountRef(caller.Address())
	)
	// Initialise a new contract and set the code that is to be used by the EVM.
	// The contract is a scoped environment for this execution context only.
	contract := NewContract(caller, to, value, gas)
	contract.SetCallCode(&addr, evm.StateDB.GetCodeHash(addr), evm.StateDB.GetCode(addr))

	ret, err = run(evm, contract, input, false)
	if err != nil {
		evm.StateDB.RevertToSnapshot(snapshot)
		if err != errExecutionReverted {
			contract.UseGas(contract.Gas)
		}
	}
	return ret, contract.Gas, err
}

// DelegateCall executes the contract associated with the addr with the given input
// as parameters. It reverses the state in case of an execution error.
//
// DelegateCall differs from CallCode in the sense that it executes the given address'
// code with the caller as context and the caller is set to the caller of the caller.
func (evm *EVM) DelegateCall(caller ContractRef, addr common.Address, input []byte, gas uint64) (ret []byte, leftOverGas uint64, err error) {
	if evm.vmConfig.NoRecursion && evm.depth > 0 {
		return nil, gas, nil
	}
	// Fail if we're trying to execute above the call depth limit
	if evm.depth > int(params.CallCreateDepth) {
		return nil, gas, ErrDepth
	}

	var (
		snapshot = evm.StateDB.Snapshot()
		to       = AccountRef(caller.Address())
	)

	// Initialise a new contract and make initialise the delegate values
	contract := NewContract(caller, to, nil, gas).AsDelegate()
	contract.SetCallCode(&addr, evm.StateDB.GetCodeHash(addr), evm.StateDB.GetCode(addr))

	ret, err = run(evm, contract, input, false)
	if err != nil {
		evm.StateDB.RevertToSnapshot(snapshot)
		if err != errExecutionReverted {
			contract.UseGas(contract.Gas)
		}
	}
	return ret, contract.Gas, err
}

// StaticCall executes the contract associated with the addr with the given input
// as parameters while disallowing any modifications to the state during the call.
// Opcodes that attempt to perform such modifications will result in exceptions
// instead of performing the modifications.
func (evm *EVM) StaticCall(caller ContractRef, addr common.Address, input []byte, gas uint64) (ret []byte, leftOverGas uint64, err error) {
	if evm.vmConfig.NoRecursion && evm.depth > 0 {
		return nil, gas, nil
	}
	// Fail if we're trying to execute above the call depth limit
	if evm.depth > int(params.CallCreateDepth) {
		return nil, gas, ErrDepth
	}

	var (
		to       = AccountRef(addr)
		snapshot = evm.StateDB.Snapshot()
	)
	// Initialise a new contract and set the code that is to be used by the EVM.
	// The contract is a scoped environment for this execution context only.
	contract := NewContract(caller, to, new(big.Int), gas)
	contract.SetCallCode(&addr, evm.StateDB.GetCodeHash(addr), evm.StateDB.GetCode(addr))

	// We do an AddBalance of zero here, just in order to trigger a touch.
	// This doesn't matter on Mainnet, where all empties are gone at the time of Byzantium,
	// but is the correct thing to do and matters on other networks, in tests, and potential
	// future scenarios
	evm.StateDB.AddBalance(addr, bigZero)

	// When an error was returned by the EVM or when setting the creation code
	// above we revert to the snapshot and consume any gas remaining. Additionally
	// when we're in Homestead this also counts for code storage gas errors.
	ret, err = run(evm, contract, input, true)
	if err != nil {
		evm.StateDB.RevertToSnapshot(snapshot)
		if err != errExecutionReverted {
			contract.UseGas(contract.Gas)
		}
	}
	return ret, contract.Gas, err
}

type codeAndHash struct {
	code []byte
	hash common.Hash
}

func (c *codeAndHash) Hash() common.Hash {
	if c.hash == (common.Hash{}) {
		c.hash = crypto.Keccak256Hash(c.code)
	}
	return c.hash
}

// create creates a new contract using code as deployment code.
func (evm *EVM) create(caller ContractRef, codeAndHash *codeAndHash, gas uint64, value *big.Int, address common.Address) ([]byte, common.Address, uint64, error) {
	// Depth check execution. Fail if we're trying to execute above the
	// limit.
	if evm.depth > int(params.CallCreateDepth) {
		return nil, common.Address{}, gas, ErrDepth
	}
	if !evm.CanTransfer(evm.StateDB, caller.Address(), value) {
		return nil, common.Address{}, gas, ErrInsufficientBalance
	}
	if rcfg.UsingOVM {
		// Make sure the creator address should be able to deploy.
		if !evm.AddressWhitelisted(caller.Address()) {
			// Try to encode this error as a Solidity error message so it's more clear to end-users
			// what's going on when a contract creation fails.
			solerr := fmt.Errorf("deployer address not whitelisted: %s", caller.Address().Hex())
			ret, err := util.EncodeSolidityError(solerr)
			if err != nil {
				// If we're unable to properly encode the error then just return the original message.
				return nil, common.Address{}, gas, solerr
			}
			return ret, common.Address{}, gas, errExecutionReverted
		}
	}
	nonce := evm.StateDB.GetNonce(caller.Address())
	evm.StateDB.SetNonce(caller.Address(), nonce+1)

	// Ensure there's no existing contract already at the designated address
	contractHash := evm.StateDB.GetCodeHash(address)
	if evm.StateDB.GetNonce(address) != 0 || (contractHash != (common.Hash{}) && contractHash != emptyCodeHash) {
		return nil, common.Address{}, 0, ErrContractAddressCollision
	}
	// Create a new account on the state
	snapshot := evm.StateDB.Snapshot()
	evm.StateDB.CreateAccount(address)
	if evm.chainRules.IsEIP158 {
		evm.StateDB.SetNonce(address, 1)
	}
	evm.Transfer(evm.StateDB, caller.Address(), address, value)

	// Initialise a new contract and set the code that is to be used by the EVM.
	// The contract is a scoped environment for this execution context only.
	contract := NewContract(caller, AccountRef(address), value, gas)
	contract.SetCodeOptionalHash(&address, codeAndHash)

	if evm.vmConfig.NoRecursion && evm.depth > 0 {
		return nil, address, gas, nil
	}

	if evm.vmConfig.Debug && evm.depth == 0 {
		evm.vmConfig.Tracer.CaptureStart(caller.Address(), address, true, codeAndHash.code, gas, value)
	}
	start := time.Now()

	ret, err := run(evm, contract, nil, false)

	// check whether the max code size has been exceeded
	maxCodeSizeExceeded := evm.chainRules.IsEIP158 && len(ret) > params.MaxCodeSize
	// if the contract creation ran successfully and no errors were returned
	// calculate the gas required to store the code. If the code could not
	// be stored due to not enough gas set an error and let it be handled
	// by the error checking condition below.
	if err == nil && !maxCodeSizeExceeded {
		createDataGas := uint64(len(ret)) * params.CreateDataGas
		if contract.UseGas(createDataGas) {
			evm.StateDB.SetCode(address, ret)
		} else {
			err = ErrCodeStoreOutOfGas
		}
	}

	// When an error was returned by the EVM or when setting the creation code
	// above we revert to the snapshot and consume any gas remaining. Additionally
	// when we're in homestead this also counts for code storage gas errors.
	if maxCodeSizeExceeded || (err != nil && (evm.chainRules.IsHomestead || err != ErrCodeStoreOutOfGas)) {
		evm.StateDB.RevertToSnapshot(snapshot)
		if err != errExecutionReverted {
			contract.UseGas(contract.Gas)
		}
	}
	// Assign err if contract code size exceeds the max while the err is still empty.
	if maxCodeSizeExceeded && err == nil {
		err = errMaxCodeSizeExceeded
	}
	if evm.vmConfig.Debug && evm.depth == 0 {
		evm.vmConfig.Tracer.CaptureEnd(ret, gas-contract.Gas, time.Since(start), err)
	}
	return ret, address, contract.Gas, err

}

// Create creates a new contract using code as deployment code.
func (evm *EVM) Create(caller ContractRef, code []byte, gas uint64, value *big.Int) (ret []byte, contractAddr common.Address, leftOverGas uint64, err error) {
	contractAddr = crypto.CreateAddress(caller.Address(), evm.StateDB.GetNonce(caller.Address()))
	return evm.create(caller, &codeAndHash{code: code}, gas, value, contractAddr)
}

// Create2 creates a new contract using code as deployment code.
//
// The different between Create2 with Create is Create2 uses sha3(0xff ++ msg.sender ++ salt ++ sha3(init_code))[12:]
// instead of the usual sender-and-nonce-hash as the address where the contract is initialized at.
func (evm *EVM) Create2(caller ContractRef, code []byte, gas uint64, endowment *big.Int, salt *big.Int) (ret []byte, contractAddr common.Address, leftOverGas uint64, err error) {
	codeAndHash := &codeAndHash{code: code}
	contractAddr = crypto.CreateAddress2(caller.Address(), common.BigToHash(salt), codeAndHash.Hash().Bytes())
	return evm.create(caller, codeAndHash, gas, endowment, contractAddr)
}

// ChainConfig returns the environment's chain configuration
func (evm *EVM) ChainConfig() *params.ChainConfig { return evm.chainConfig }

func (evm *EVM) AddressWhitelisted(addr common.Address) bool {
	// First check if the owner is address(0), which implicitly disables the whitelist.
	ownerKey := common.Hash{}
	owner := evm.StateDB.GetState(dump.OvmWhitelistAddress, ownerKey)
	if (owner == common.Hash{}) {
		return true
	}

	// Next check if the user is whitelisted by resolving the position where the
	// true/false value would be.
	position := common.Big1
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(common.LeftPadBytes(addr.Bytes(), 32))
	hasher.Write(common.LeftPadBytes(position.Bytes(), 32))
	digest := hasher.Sum(nil)
	key := common.BytesToHash(digest)
	isWhitelisted := evm.StateDB.GetState(dump.OvmWhitelistAddress, key)
	return isWhitelisted != common.Hash{}
}
