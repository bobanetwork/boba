package immutables

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/bobanetwork/v3-anchorage/boba-bindings/bindings"
	"github.com/bobanetwork/v3-anchorage/boba-bindings/predeploys"
	"github.com/bobanetwork/v3-anchorage/boba-chain-ops/deployer"
	"github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/erigon-lib/common/hexutility"
	"github.com/ledgerwatch/erigon/accounts/abi/bind"
	"github.com/ledgerwatch/erigon/accounts/abi/bind/backends"
	"github.com/ledgerwatch/erigon/common/hexutil"
	"github.com/ledgerwatch/erigon/core/types"

	"github.com/ledgerwatch/log/v3"
)

// ImmutableValues represents the values to be set in immutable code.
// The key is the name of the variable and the value is the value to set in
// immutable code.
type ImmutableValues map[string]any

// ImmutableConfig represents the immutable configuration for the L2 predeploy
// contracts.
type ImmutableConfig map[string]ImmutableValues

// Check does a sanity check that the specific values that
// Optimism uses are set inside of the ImmutableConfig.
func (i ImmutableConfig) Check() error {
	if _, ok := i["L2CrossDomainMessenger"]["otherMessenger"]; !ok {
		return errors.New("L2CrossDomainMessenger otherMessenger not set")
	}
	if _, ok := i["L2StandardBridge"]["otherBridge"]; !ok {
		return errors.New("L2StandardBridge otherBridge not set")
	}
	if _, ok := i["L2ERC721Bridge"]["messenger"]; !ok {
		return errors.New("L2ERC721Bridge messenger not set")
	}
	if _, ok := i["L2ERC721Bridge"]["otherBridge"]; !ok {
		return errors.New("L2ERC721Bridge otherBridge not set")
	}
	if _, ok := i["OptimismMintableERC721Factory"]["bridge"]; !ok {
		return errors.New("OptimismMintableERC20Factory bridge not set")
	}
	if _, ok := i["OptimismMintableERC721Factory"]["remoteChainId"]; !ok {
		return errors.New("OptimismMintableERC20Factory remoteChainId not set")
	}
	if _, ok := i["SequencerFeeVault"]["recipient"]; !ok {
		return errors.New("SequencerFeeVault recipient not set")
	}
	if _, ok := i["L1FeeVault"]["recipient"]; !ok {
		return errors.New("L1FeeVault recipient not set")
	}
	if _, ok := i["BaseFeeVault"]["recipient"]; !ok {
		return errors.New("BaseFeeVault recipient not set")
	}
	return nil
}

// DeploymentResults represents the output of deploying each of the
// contracts so that the immutables can be set properly in the bytecode.
type DeploymentResults map[string]hexutility.Bytes

// BuildOptimism will deploy the L2 predeploys so that their immutables are set
// correctly.
func BuildOptimism(immutable ImmutableConfig) (DeploymentResults, error) {
	if err := immutable.Check(); err != nil {
		return DeploymentResults{}, err
	}

	deployments := []deployer.Constructor{
		{
			Name: "GasPriceOracle",
		},
		{
			Name: "L1Block",
		},
		{
			Name: "L2CrossDomainMessenger",
			Args: []interface{}{
				immutable["L2CrossDomainMessenger"]["otherMessenger"],
			},
		},
		{
			Name: "L2StandardBridge",
			Args: []interface{}{
				immutable["L2StandardBridge"]["otherBridge"],
			},
		},
		{
			Name: "L2ToL1MessagePasser",
		},
		{
			Name: "SequencerFeeVault",
			Args: []interface{}{
				immutable["SequencerFeeVault"]["recipient"],
				immutable["SequencerFeeVault"]["minimumWithdrawalAmount"],
				immutable["SequencerFeeVault"]["withdrawalNetwork"],
			},
		},
		{
			Name: "BaseFeeVault",
			Args: []interface{}{
				immutable["BaseFeeVault"]["recipient"],
				immutable["BaseFeeVault"]["minimumWithdrawalAmount"],
				immutable["BaseFeeVault"]["withdrawalNetwork"],
			},
		},
		{
			Name: "L1FeeVault",
			Args: []interface{}{
				immutable["L1FeeVault"]["recipient"],
				immutable["L1FeeVault"]["minimumWithdrawalAmount"],
				immutable["L1FeeVault"]["withdrawalNetwork"],
			},
		},
		{
			Name: "OptimismMintableERC20Factory",
		},
		{
			Name: "DeployerWhitelist",
		},
		{
			Name: "LegacyMessagePasser",
		},
		{
			Name: "L1BlockNumber",
		},
		{
			Name: "L2ERC721Bridge",
			Args: []interface{}{
				predeploys.L2CrossDomainMessengerAddr,
				immutable["L2ERC721Bridge"]["otherBridge"],
			},
		},
		{
			Name: "OptimismMintableERC721Factory",
			Args: []interface{}{
				predeploys.L2ERC721BridgeAddr,
				immutable["OptimismMintableERC721Factory"]["remoteChainId"],
			},
		},
		{
			Name: "LegacyERC20ETH",
		},
		{
			Name: "BobaL2",
			Args: []interface{}{
				immutable["BobaL2"]["l2Bridge"],
				immutable["BobaL2"]["l1Token"],
				immutable["BobaL2"]["_name"],
				immutable["BobaL2"]["_symbol"],
				immutable["BobaL2"]["_decimals"],
			},
		},
		{
			Name: "BobaTuringCredit",
		},
		{
			Name: "BobaGasPriceOracle",
		},
	}
	return BuildL2(deployments)
}

// BuildL2 will deploy contracts to a simulated backend so that their immutables
// can be properly set. The bytecode returned in the results is suitable to be
// inserted into the state via state surgery.
func BuildL2(constructors []deployer.Constructor) (DeploymentResults, error) {
	deployments, err := deployer.Deploy(deployer.NewBackend(), constructors, l2Deployer)
	if err != nil {
		return nil, err
	}
	results := make(DeploymentResults)
	for _, dep := range deployments {
		results[dep.Name] = dep.Bytecode
	}
	return results, nil
}

func l2Deployer(backend *backends.SimulatedBackend, opts *bind.TransactOpts, deployment deployer.Constructor) (types.Transaction, error) {
	var tx types.Transaction
	var recipient common.Address
	var minimumWithdrawalAmount *big.Int
	var withdrawalNetwork uint8
	var err error
	var addr common.Address
	switch deployment.Name {
	case "GasPriceOracle":
		_, tx, _, err = bindings.DeployGasPriceOracle(opts, backend)
	case "L1Block":
		// No arguments required for the L1Block contract
		_, tx, _, err = bindings.DeployL1Block(opts, backend)
	case "L2CrossDomainMessenger":
		otherMessenger, ok := deployment.Args[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for otherMessenger")
		}
		_, tx, _, err = bindings.DeployL2CrossDomainMessenger(opts, backend, otherMessenger)
	case "L2StandardBridge":
		otherBridge, ok := deployment.Args[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for otherBridge")
		}
		_, tx, _, err = bindings.DeployL2StandardBridge(opts, backend, otherBridge)
	case "L2ToL1MessagePasser":
		// No arguments required for L2ToL1MessagePasser
		_, tx, _, err = bindings.DeployL2ToL1MessagePasser(opts, backend)
	case "SequencerFeeVault":
		recipient, minimumWithdrawalAmount, withdrawalNetwork, err = prepareFeeVaultArguments(deployment)
		if err != nil {
			return nil, err
		}
		_, tx, _, err = bindings.DeploySequencerFeeVault(opts, backend, recipient, minimumWithdrawalAmount, withdrawalNetwork)
	case "BaseFeeVault":
		recipient, minimumWithdrawalAmount, withdrawalNetwork, err = prepareFeeVaultArguments(deployment)
		if err != nil {
			return nil, err
		}
		_, tx, _, err = bindings.DeployBaseFeeVault(opts, backend, recipient, minimumWithdrawalAmount, withdrawalNetwork)
	case "L1FeeVault":
		recipient, minimumWithdrawalAmount, withdrawalNetwork, err = prepareFeeVaultArguments(deployment)
		if err != nil {
			return nil, err
		}
		_, tx, _, err = bindings.DeployL1FeeVault(opts, backend, recipient, minimumWithdrawalAmount, withdrawalNetwork)
	case "OptimismMintableERC20Factory":
		_, tx, _, err = bindings.DeployOptimismMintableERC20Factory(opts, backend, predeploys.L2StandardBridgeAddr)
	case "DeployerWhitelist":
		_, tx, _, err = bindings.DeployDeployerWhitelist(opts, backend)
	case "LegacyMessagePasser":
		_, tx, _, err = bindings.DeployLegacyMessagePasser(opts, backend)
	case "L1BlockNumber":
		_, tx, _, err = bindings.DeployL1BlockNumber(opts, backend)
	case "L2ERC721Bridge":
		// TODO(tynes): messenger should be hardcoded in the contract
		messenger, ok := deployment.Args[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for messenger")
		}
		otherBridge, ok := deployment.Args[1].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for otherBridge")
		}
		_, tx, _, err = bindings.DeployL2ERC721Bridge(opts, backend, messenger, otherBridge)
	case "OptimismMintableERC721Factory":
		bridge, ok := deployment.Args[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for bridge")
		}
		remoteChainId, ok := deployment.Args[1].(*big.Int)
		if !ok {
			return nil, fmt.Errorf("invalid type for remoteChainId")
		}
		_, tx, _, err = bindings.DeployOptimismMintableERC721Factory(opts, backend, bridge, remoteChainId)
	case "LegacyERC20ETH":
		_, tx, _, err = bindings.DeployLegacyERC20ETH(opts, backend)
	case "BobaTuringCredit":
		addr, tx, _, err = bindings.DeployBobaTuringCredit(opts, backend, big.NewInt(10))
		log.Info("BobaTuringCredit Deployment", "err", err, "addr", addr)
	case "BobaL2":
		l2Bridge, ok := deployment.Args[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for l2Bridge")
		}
		l1Token, ok := deployment.Args[1].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for l1Token")
		}
		_name, ok := deployment.Args[2].(string)
		if !ok {
			return nil, fmt.Errorf("invalid type for _name")
		}
		_symbol, ok := deployment.Args[3].(string)
		if !ok {
			return nil, fmt.Errorf("invalid type for _symbol")
		}
		_decimals, ok := deployment.Args[4].(uint8)
		if !ok {
			return nil, fmt.Errorf("invalid type for _decimals")
		}
		addr, tx, _, err = bindings.DeployL2GovernanceERC20(
			opts,
			backend,
			l2Bridge,
			l1Token,
			_name,
			_symbol,
			uint8(_decimals),
		)
		log.Info("BobaL2 Deployment", "err", err, "addr", addr)
	case "BobaGasPriceOracle":
		addr, tx, _, err = bindings.DeployBobaGasPriceOracle(opts, backend)
		log.Info("BobaGasPriceOracle Deployment", "err", err, "addr", addr)
	default:
		return tx, fmt.Errorf("unknown contract: %s", deployment.Name)
	}

	return tx, err
}

func prepareFeeVaultArguments(deployment deployer.Constructor) (common.Address, *big.Int, uint8, error) {
	fmt.Println("deployment", deployment)
	recipient, ok := deployment.Args[0].(common.Address)
	if !ok {
		return common.Address{}, nil, 0, fmt.Errorf("invalid type for recipient")
	}
	minimumWithdrawalAmountHex, ok := deployment.Args[1].(*hexutil.Big)
	if !ok {
		return common.Address{}, nil, 0, fmt.Errorf("invalid type for minimumWithdrawalAmount")
	}
	withdrawalNetwork, ok := deployment.Args[2].(uint8)
	if !ok {
		return common.Address{}, nil, 0, fmt.Errorf("invalid type for withdrawalNetwork")
	}
	return recipient, minimumWithdrawalAmountHex.ToInt(), withdrawalNetwork, nil
}
