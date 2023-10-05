package predeploys

import (
	"github.com/ledgerwatch/erigon-lib/common"
)

// The legacy system has a different set of predeploys
// BobaTuringCredit -> 0x4200000000000000000000000000000000000020
// BobaLegacyTuringCreditImplementation -> 0x4200000000000000000000000000000000000021
// BobaHCHelperImplementation -> 0x4200000000000000000000000000000000000022
// BobaL2 -> 0x4200000000000000000000000000000000000023

// Boba predeployed contracts will be move to the following addresses from the index 1000
// BobaTuringCredit -> 0x42000000000000000000000000000000000003E8
// BobaHCHelper -> 0x42000000000000000000000000000000000003E9

const (
	L2ToL1MessagePasser = "0x4200000000000000000000000000000000000016"
	DeployerWhitelist   = "0x4200000000000000000000000000000000000002"
	// We are different here
	LegacyERC20ETH                = "0x4200000000000000000000000000000000000006"
	WETH9                         = "0xDeadDeAddeAddEAddeadDEaDDEAdDeaDDeAD0000"
	L2CrossDomainMessenger        = "0x4200000000000000000000000000000000000007"
	L2StandardBridge              = "0x4200000000000000000000000000000000000010"
	SequencerFeeVault             = "0x4200000000000000000000000000000000000011"
	OptimismMintableERC20Factory  = "0x4200000000000000000000000000000000000012"
	L1BlockNumber                 = "0x4200000000000000000000000000000000000013"
	GasPriceOracle                = "0x420000000000000000000000000000000000000F"
	L1Block                       = "0x4200000000000000000000000000000000000015"
	LegacyMessagePasser           = "0x4200000000000000000000000000000000000000"
	L2ERC721Bridge                = "0x4200000000000000000000000000000000000014"
	OptimismMintableERC721Factory = "0x4200000000000000000000000000000000000017"
	ProxyAdmin                    = "0x4200000000000000000000000000000000000018"
	BaseFeeVault                  = "0x4200000000000000000000000000000000000019"
	L1FeeVault                    = "0x420000000000000000000000000000000000001a"
	SchemaRegistry                = "0x4200000000000000000000000000000000000020"
	EAS                           = "0x4200000000000000000000000000000000000021"

	// Boba specific
	BobaTuringCredit = "0x42000000000000000000000000000000000003e8"
	BobaHCHelper     = "0x42000000000000000000000000000000000003E9"
	BobaL2           = "0x4200000000000000000000000000000000000023"

	// Boba Legacy address
	BobaLegacyTuringCredit               = "0x4200000000000000000000000000000000000020"
	BobaLegacyTuringCreditImplementation = "0x4200000000000000000000000000000000000021"

	// Special case for Boba mainnet
	BobaTuringCredit288                     = "0xF8D2f1b0292C0Eeef80D8F47661A9DaCDB4b23bf"
	BOBAL2288                               = "0xa18bF3994C0Cc6E3b63ac420308E5383f53120D7"
	BobaLegacyTuringCreditImplementation288 = "0xd8E006702bdCbE2582dF13f900bCF750129bB449"
)

var (
	L2ToL1MessagePasserAddr           = common.HexToAddress(L2ToL1MessagePasser)
	DeployerWhitelistAddr             = common.HexToAddress(DeployerWhitelist)
	LegacyERC20ETHAddr                = common.HexToAddress(LegacyERC20ETH)
	WETH9Addr                         = common.HexToAddress(WETH9)
	L2CrossDomainMessengerAddr        = common.HexToAddress(L2CrossDomainMessenger)
	L2StandardBridgeAddr              = common.HexToAddress(L2StandardBridge)
	SequencerFeeVaultAddr             = common.HexToAddress(SequencerFeeVault)
	OptimismMintableERC20FactoryAddr  = common.HexToAddress(OptimismMintableERC20Factory)
	L1BlockNumberAddr                 = common.HexToAddress(L1BlockNumber)
	GasPriceOracleAddr                = common.HexToAddress(GasPriceOracle)
	L1BlockAddr                       = common.HexToAddress(L1Block)
	LegacyMessagePasserAddr           = common.HexToAddress(LegacyMessagePasser)
	L2ERC721BridgeAddr                = common.HexToAddress(L2ERC721Bridge)
	OptimismMintableERC721FactoryAddr = common.HexToAddress(OptimismMintableERC721Factory)
	ProxyAdminAddr                    = common.HexToAddress(ProxyAdmin)
	BaseFeeVaultAddr                  = common.HexToAddress(BaseFeeVault)
	L1FeeVaultAddr                    = common.HexToAddress(L1FeeVault)
	SchemaRegistryAddr                = common.HexToAddress(SchemaRegistry)
	EASAddr                           = common.HexToAddress(EAS)

	// Boba specific
	BobaTuringCreditAddr = common.HexToAddress(BobaTuringCredit)
	BobaHCHelperAddr     = common.HexToAddress(BobaHCHelper)
	BobaL2Addr           = common.HexToAddress(BobaL2)

	// Boba legacy
	BobaLegacyTuringCreditAddr               = common.HexToAddress(BobaLegacyTuringCredit)
	BobaLegacyTuringCreditImplementationAddr = common.HexToAddress(BobaLegacyTuringCreditImplementation)

	// Special case for Boba mainnet
	BobaTuringCredit288Addr                     = common.HexToAddress(BobaTuringCredit288)
	BobaLegacyTuringCreditImplementation288Addr = common.HexToAddress(BobaLegacyTuringCreditImplementation288)
	BOBAL2288Addr                               = common.HexToAddress(BOBAL2288)

	Predeploys                    = make(map[string]*common.Address)
	LegacyBobaProxy               = make(map[string]*common.Address)
	LegacyBobaProxyImplementation = make(map[string]*common.Address)
)

// IsProxied returns true for predeploys that will sit behind a proxy contract
func IsProxied(predeployAddr common.Address) bool {
	switch predeployAddr {
	case LegacyERC20ETHAddr:
	case WETH9Addr:
	case BobaL2Addr:
	default:
		return true
	}
	return false
}

func init() {
	Predeploys["L2ToL1MessagePasser"] = &L2ToL1MessagePasserAddr
	Predeploys["DeployerWhitelist"] = &DeployerWhitelistAddr
	Predeploys["LegacyERC20ETH"] = &LegacyERC20ETHAddr
	Predeploys["WETH9"] = &WETH9Addr
	Predeploys["L2CrossDomainMessenger"] = &L2CrossDomainMessengerAddr
	Predeploys["L2StandardBridge"] = &L2StandardBridgeAddr
	Predeploys["SequencerFeeVault"] = &SequencerFeeVaultAddr
	Predeploys["OptimismMintableERC20Factory"] = &OptimismMintableERC20FactoryAddr
	Predeploys["L1BlockNumber"] = &L1BlockNumberAddr
	Predeploys["GasPriceOracle"] = &GasPriceOracleAddr
	Predeploys["L1Block"] = &L1BlockAddr
	Predeploys["LegacyMessagePasser"] = &LegacyMessagePasserAddr
	Predeploys["L2ERC721Bridge"] = &L2ERC721BridgeAddr
	Predeploys["OptimismMintableERC721Factory"] = &OptimismMintableERC721FactoryAddr
	Predeploys["ProxyAdmin"] = &ProxyAdminAddr
	Predeploys["BaseFeeVault"] = &BaseFeeVaultAddr
	Predeploys["L1FeeVault"] = &L1FeeVaultAddr
	Predeploys["SchemaRegistry"] = &SchemaRegistryAddr
	Predeploys["EAS"] = &EASAddr

	// Boba specific
	Predeploys["BobaTuringCredit"] = &BobaTuringCreditAddr
	Predeploys["BobaHCHelper"] = &BobaHCHelperAddr
	Predeploys["BobaL2"] = &BobaL2Addr

	// Boba legacy
	LegacyBobaProxy["BobaLegacyTuringCredit"] = &BobaLegacyTuringCreditAddr
	LegacyBobaProxyImplementation["BobaLegacyTuringCreditImplementation"] = &BobaLegacyTuringCreditImplementationAddr
}
