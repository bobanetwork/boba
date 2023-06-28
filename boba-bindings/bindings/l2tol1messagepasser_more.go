// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/bobanetwork/v3-anchorage/boba-bindings/solc"
)

const L2ToL1MessagePasserStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/L2ToL1MessagePasser.sol:L2ToL1MessagePasser\",\"label\":\"sentMessages\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1001,\"contract\":\"contracts/L2/L2ToL1MessagePasser.sol:L2ToL1MessagePasser\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint240\"}],\"types\":{\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"}}}"

var L2ToL1MessagePasserStorageLayout = new(solc.StorageLayout)

var L2ToL1MessagePasserDeployedBin = "0x6080604052600436106100695760003560e01c806382e3702d1161004357806382e3702d146100f6578063c2b3e5ac14610136578063ecc704281461014957600080fd5b80633f827a5a1461009257806344df8e70146100bf57806354fd4d50146100d457600080fd5b3661008d5761008b33620186a0604051806020016040528060008152506101ae565b005b600080fd5b34801561009e57600080fd5b506100a7600181565b60405161ffff90911681526020015b60405180910390f35b3480156100cb57600080fd5b5061008b610372565b3480156100e057600080fd5b506100e96103aa565b6040516100b6919061068c565b34801561010257600080fd5b506101266101113660046106a6565b60006020819052908152604090205460ff1681565b60405190151581526020016100b6565b61008b6101443660046106ee565b6101ae565b34801561015557600080fd5b506101a06001547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b6040519081526020016100b6565b60006102446040518060c001604052806102086001547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b815233602082015273ffffffffffffffffffffffffffffffffffffffff871660408201523460608201526080810186905260a00184905261044d565b600081815260208190526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055905073ffffffffffffffffffffffffffffffffffffffff8416336102df6001547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b7f02a52367d10742d8032712c1bb8e0144ff1ec5ffda1ed7d70bb05a27449550543487878760405161031494939291906107f2565b60405180910390a45050600180547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8082168301167fffff0000000000000000000000000000000000000000000000000000000000009091161790555050565b4761037c8161049a565b60405181907f7967de617a5ac1cc7eba2d6f37570a0135afa950d8bb77cdd35f0d0b4e85a16f90600090a250565b60606103d57f00000000000000000000000000000000000000000000000000000000000000006104c9565b6103fe7f00000000000000000000000000000000000000000000000000000000000000006104c9565b6104277f00000000000000000000000000000000000000000000000000000000000000006104c9565b60405160200161043993929190610822565b604051602081830303815290604052905090565b80516020808301516040808501516060860151608087015160a0880151935160009761047d979096959101610898565b604051602081830303815290604052805190602001209050919050565b806040516104a790610606565b6040518091039082f09050801580156104c4573d6000803e3d6000fd5b505050565b60608160000361050c57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b811561053657806105208161091e565b915061052f9050600a83610985565b9150610510565b60008167ffffffffffffffff811115610551576105516106bf565b6040519080825280601f01601f19166020018201604052801561057b576020820181803683370190505b5090505b84156105fe57610590600183610999565b915061059d600a866109b0565b6105a89060306109c4565b60f81b8183815181106105bd576105bd6109dc565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506105f7600a86610985565b945061057f565b949350505050565b600880610a0c83390190565b60005b8381101561062d578181015183820152602001610615565b8381111561063c576000848401525b50505050565b6000815180845261065a816020860160208601610612565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061069f6020830184610642565b9392505050565b6000602082840312156106b857600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008060006060848603121561070357600080fd5b833573ffffffffffffffffffffffffffffffffffffffff8116811461072757600080fd5b925060208401359150604084013567ffffffffffffffff8082111561074b57600080fd5b818601915086601f83011261075f57600080fd5b813581811115610771576107716106bf565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156107b7576107b76106bf565b816040528281528960208487010111156107d057600080fd5b8260208601602083013760006020848301015280955050505050509250925092565b8481528360208201526080604082015260006108116080830185610642565b905082606083015295945050505050565b60008451610834818460208901610612565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610870816001850160208a01610612565b6001920191820152835161088b816002840160208801610612565b0160020195945050505050565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a08301526108e360c0830184610642565b98975050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361094f5761094f6108ef565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261099457610994610956565b500490565b6000828210156109ab576109ab6108ef565b500390565b6000826109bf576109bf610956565b500690565b600082198211156109d7576109d76108ef565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfe608060405230fffea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L2ToL1MessagePasserStorageLayoutJSON), L2ToL1MessagePasserStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2ToL1MessagePasser"] = L2ToL1MessagePasserStorageLayout
	deployedBytecodes["L2ToL1MessagePasser"] = L2ToL1MessagePasserDeployedBin
}