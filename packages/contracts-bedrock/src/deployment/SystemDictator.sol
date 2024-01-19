// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { L2OutputOracle } from "../L1/L2OutputOracle.sol";
import { OptimismPortal } from "../L1/OptimismPortal.sol";
import { L1CrossDomainMessenger } from "../L1/L1CrossDomainMessenger.sol";
import { L1ERC721Bridge } from "../L1/L1ERC721Bridge.sol";
import { L1StandardBridge } from "../L1/L1StandardBridge.sol";
import { L1ChugSplashProxy } from "../legacy/L1ChugSplashProxy.sol";
import { AddressManager } from "../legacy/AddressManager.sol";
import { Proxy } from "../universal/Proxy.sol";
import { ProxyAdmin } from "../universal/ProxyAdmin.sol";
import { OptimismMintableERC20Factory } from "../universal/OptimismMintableERC20Factory.sol";
import { PortalSender } from "./PortalSender.sol";
import { SystemConfig } from "../L1/SystemConfig.sol";
import { ProtocolVersions, ProtocolVersion } from "../L1/ProtocolVersions.sol";
import { ResourceMetering } from "../L1/ResourceMetering.sol";
import { SuperchainConfig } from "../L1/SuperchainConfig.sol";
import { Constants } from "../libraries/Constants.sol";

/**
 * @title SystemDictator
 * @notice The SystemDictator is responsible for coordinating the deployment of a full Bedrock
 *         system. The SystemDictator is designed to support both fresh network deployments and
 *         upgrades to existing pre-Bedrock systems.
 */
contract SystemDictator is OwnableUpgradeable {
    /**
     * @notice Basic system configuration.
     */
    struct GlobalConfig {
        AddressManager addressManager;
        ProxyAdmin proxyAdmin;
        SuperchainConfig superchainConfig;
        address controller;
        address finalOwner;
    }

    /**
     * @notice Set of proxy addresses.
     */
    struct ProxyAddressConfig {
        address l2OutputOracleProxy;
        address optimismPortalProxy;
        address l1CrossDomainMessengerProxy;
        address l1StandardBridgeProxy;
        address optimismMintableERC20FactoryProxy;
        address l1ERC721BridgeProxy;
        address systemConfigProxy;
        address protocolVersionsProxy;
    }

    /**
     * @notice Set of implementation addresses.
     */
    struct ImplementationAddressConfig {
        L2OutputOracle l2OutputOracleImpl;
        OptimismPortal optimismPortalImpl;
        L1CrossDomainMessenger l1CrossDomainMessengerImpl;
        L1StandardBridge l1StandardBridgeImpl;
        OptimismMintableERC20Factory optimismMintableERC20FactoryImpl;
        L1ERC721Bridge l1ERC721BridgeImpl;
        PortalSender portalSenderImpl;
        SystemConfig systemConfigImpl;
        ProtocolVersions protocolVersionsImpl;
    }

    /**
     * @notice Dynamic L2OutputOracle config.
     */
    struct L2OutputOracleDynamicConfig {
        uint256 l2OutputOracleStartingBlockNumber;
        uint256 l2OutputOracleStartingTimestamp;
        address l2OutputOracleProposer;
        address l2OutputOracleChallenger;
    }

    /**
     * notice SystemConfig Address config.
     */
    struct SystemConfigAddressConfig {
        address l1CrossDomainMessenger;
        address l1ERC721Bridge;
        address l1StandardBridge;
        address l2OutputOracle;
        address optimismPortal;
        address optimismMintableERC20Factory;
    }

    /**
     * @notice Values for the system config contract.
     */
    struct SystemConfigConfig {
        address owner;
        uint256 overhead;
        uint256 scalar;
        bytes32 batcherHash;
        uint64 gasLimit;
        address unsafeBlockSigner;
        ResourceMetering.ResourceConfig resourceConfig;
        uint256 opnodeStartBlock;
        address batchInbox;
        SystemConfig.Addresses systemConfigAddressConfig;
    }

    /**
     * @notice Combined system configuration.
     */
    struct DeployConfig {
        GlobalConfig globalConfig;
        ProxyAddressConfig proxyAddressConfig;
        ImplementationAddressConfig implementationAddressConfig;
        SystemConfigConfig systemConfigConfig;
        ProtocolVersionConfig protocolVersionConfig;
    }

    /**
     * @notice OptimismPortal configuration.
     */
    struct OptimismPortalDynamicConfig {
        address l2OutputOracle;
        address portalGuardian;
        address systemConfig;
        bool paused;
    }

    /**
     * @notice ProtocalVersions configuration.
     */
    struct ProtocolVersionConfig {
        ProtocolVersion requiredProtocolVersion;
        ProtocolVersion recommendedProtocolVersion;
    }

    /**
     * @notice Step after which exit 1 can no longer be used.
     */
    uint8 public constant EXIT_1_NO_RETURN_STEP = 3;

    /**
     * @notice Step where proxy ownership is transferred.
     */
    uint8 public constant PROXY_TRANSFER_STEP = 4;

    /**
     * @notice System configuration.
     */
    DeployConfig public config;

    /**
     * @notice Dynamic configuration for the L2OutputOracle.
     */
    L2OutputOracleDynamicConfig public l2OutputOracleDynamicConfig;

    /**
     * @notice Dynamic configuration for the OptimismPortal. Determines
     *         if the system should be paused when initialized.
     */
    OptimismPortalDynamicConfig public optimismPortalDynamicConfig;

    /**
     * @notice Current step;
     */
    uint8 public currentStep;

    /**
     * @notice Whether or not dynamic config has been set.
     */
    bool public dynamicConfigSet;

    /**
     * @notice Whether or not the deployment is finalized.
     */
    bool public finalized;

    /**
     * @notice Whether or not the deployment has been exited.
     */
    bool public exited;

    /**
     * @notice Address of the old L1CrossDomainMessenger implementation.
     */
    address public oldL1CrossDomainMessenger;

    /**
     * @notice Checks that the current step is the expected step, then bumps the current step.
     *
     * @param _step Current step.
     */
    modifier step(uint8 _step) {
        require(!finalized, "Already finalized");
        require(!exited, "Already exited");
        require(currentStep == _step, "Incorrect step");
        _;
        currentStep++;
    }

    /**
     * @notice Constructor required to ensure that the implementation of the SystemDictator is
     *         initialized upon deployment.
     */
    constructor() {
        ResourceMetering.ResourceConfig memory rcfg = Constants.DEFAULT_RESOURCE_CONFIG();

        // Using this shorter variable as an alias for address(0) just prevents us from having to
        // to use a new line for every single parameter.
        address zero = address(0);
        initialize(
            DeployConfig(
                GlobalConfig(AddressManager(zero), ProxyAdmin(zero), SuperchainConfig(zero), zero, zero),
                ProxyAddressConfig(zero, zero, zero, zero, zero, zero, zero, zero),
                ImplementationAddressConfig(
                    L2OutputOracle(zero),
                    OptimismPortal(payable(zero)),
                    L1CrossDomainMessenger(zero),
                    L1StandardBridge(payable(zero)),
                    OptimismMintableERC20Factory(zero),
                    L1ERC721Bridge(zero),
                    PortalSender(zero),
                    SystemConfig(zero),
                    ProtocolVersions(zero)
                ),
                SystemConfigConfig(
                    zero,
                    0,
                    0,
                    bytes32(0),
                    0,
                    zero,
                    rcfg,
                    type(uint256).max,
                    zero,
                    SystemConfig.Addresses({
                        l1CrossDomainMessenger: zero,
                        l1ERC721Bridge: zero,
                        l1StandardBridge: zero,
                        l2OutputOracle: zero,
                        optimismPortal: zero,
                        optimismMintableERC20Factory: zero
                    })
                ),
                ProtocolVersionConfig(ProtocolVersion.wrap(uint256(0)), ProtocolVersion.wrap(uint256(0)))
            )
        );
    }

    /**
     * @param _config System configuration.
     */
    function initialize(DeployConfig memory _config) public initializer {
        config = _config;
        currentStep = 1;
        __Ownable_init();
        _transferOwnership(config.globalConfig.controller);
    }

    /**
     * @notice Allows the owner to update dynamic config.
     *
     * @param _l2OutputOracleDynamicConfig Dynamic L2OutputOracle config.
     * @param _optimismPortalDynamicConfig Dynamic OptimismPortal config.
     */
    function updateDynamicConfig(
        L2OutputOracleDynamicConfig memory _l2OutputOracleDynamicConfig,
        OptimismPortalDynamicConfig memory _optimismPortalDynamicConfig
    )
        external
        onlyOwner
    {
        l2OutputOracleDynamicConfig = _l2OutputOracleDynamicConfig;
        optimismPortalDynamicConfig = _optimismPortalDynamicConfig;
        dynamicConfigSet = true;
    }

    /**
     * @notice Configures the ProxyAdmin contract.
     */
    function step1() public onlyOwner step(1) {
        // Set the AddressManager in the ProxyAdmin.
        config.globalConfig.proxyAdmin.setAddressManager(config.globalConfig.addressManager);

        // Set the L1CrossDomainMessenger to the RESOLVED proxy type.
        config.globalConfig.proxyAdmin.setProxyType(
            config.proxyAddressConfig.l1CrossDomainMessengerProxy, ProxyAdmin.ProxyType.RESOLVED
        );

        // Set the implementation name for the L1CrossDomainMessenger.
        config.globalConfig.proxyAdmin.setImplementationName(
            config.proxyAddressConfig.l1CrossDomainMessengerProxy, "OVM_L1CrossDomainMessenger"
        );

        // Set the L1StandardBridge to the CHUGSPLASH proxy type.
        config.globalConfig.proxyAdmin.setProxyType(
            config.proxyAddressConfig.l1StandardBridgeProxy, ProxyAdmin.ProxyType.CHUGSPLASH
        );

        // Upgrade and initialize the SystemConfig so the Sequencer can start up.
        config.globalConfig.proxyAdmin.upgradeAndCall(
            payable(config.proxyAddressConfig.systemConfigProxy),
            address(config.implementationAddressConfig.systemConfigImpl),
            abi.encodeCall(
                SystemConfig.initialize,
                (
                    config.systemConfigConfig.owner,
                    config.systemConfigConfig.overhead,
                    config.systemConfigConfig.scalar,
                    config.systemConfigConfig.batcherHash,
                    config.systemConfigConfig.gasLimit,
                    config.systemConfigConfig.unsafeBlockSigner,
                    Constants.DEFAULT_RESOURCE_CONFIG()
                )
            )
        );
    }

    /**
     * @notice Pauses the system by shutting down the L1CrossDomainMessenger and setting the
     *         deposit halt flag to tell the Sequencer's DTL to stop accepting deposits.
     */
    function step2() public onlyOwner step(2) {
        // Store the address of the old L1CrossDomainMessenger implementation. We will need this
        // address in the case that we have to exit early.
        oldL1CrossDomainMessenger = config.globalConfig.addressManager.getAddress("OVM_L1CrossDomainMessenger");

        // Temporarily brick the L1CrossDomainMessenger by setting its implementation address to
        // address(0) which will cause the ResolvedDelegateProxy to revert. Better than pausing
        // the L1CrossDomainMessenger via pause() because it can be easily reverted.
        config.globalConfig.addressManager.setAddress("OVM_L1CrossDomainMessenger", address(0));

        // Set the DTL shutoff block, which will tell the DTL to stop syncing new deposits from the
        // CanonicalTransactionChain. We do this by setting an address in the AddressManager
        // because the DTL already has a reference to the AddressManager and this way we don't also
        // need to give it a reference to the SystemDictator.
        config.globalConfig.addressManager.setAddress("DTL_SHUTOFF_BLOCK", address(uint160(block.number)));
    }

    /**
     * @notice Removes deprecated addresses from the AddressManager.
     */
    function step3() public onlyOwner step(EXIT_1_NO_RETURN_STEP) {
        // Remove all deprecated addresses from the AddressManager
        string[17] memory deprecated = [
            "OVM_CanonicalTransactionChain",
            "OVM_L2CrossDomainMessenger",
            "OVM_DecompressionPrecompileAddress",
            "OVM_Sequencer",
            "OVM_Proposer",
            "OVM_ChainStorageContainer-CTC-batches",
            "OVM_ChainStorageContainer-CTC-queue",
            "OVM_CanonicalTransactionChain",
            "OVM_StateCommitmentChain",
            "OVM_BondManager",
            "OVM_ExecutionManager",
            "OVM_FraudVerifier",
            "OVM_StateManagerFactory",
            "OVM_StateTransitionerFactory",
            "OVM_SafetyChecker",
            "OVM_L1MultiMessageRelayer",
            "BondManager"
        ];

        for (uint256 i = 0; i < deprecated.length; i++) {
            config.globalConfig.addressManager.setAddress(deprecated[i], address(0));
        }
    }

    /**
     * @notice Transfers system ownership to the ProxyAdmin.
     */
    function step4() public onlyOwner step(PROXY_TRANSFER_STEP) {
        // Transfer ownership of the AddressManager to the ProxyAdmin.
        config.globalConfig.addressManager.transferOwnership(address(config.globalConfig.proxyAdmin));

        // Transfer ownership of the L1StandardBridge to the ProxyAdmin.
        L1ChugSplashProxy(payable(config.proxyAddressConfig.l1StandardBridgeProxy)).setOwner(
            address(config.globalConfig.proxyAdmin)
        );

        // Transfer ownership of the L1ERC721Bridge to the ProxyAdmin.
        Proxy(payable(config.proxyAddressConfig.l1ERC721BridgeProxy)).changeAdmin(
            address(config.globalConfig.proxyAdmin)
        );
    }

    /**
     * @notice Upgrades and initializes proxy contracts.
     */
    function step5() public onlyOwner step(5) {
        // Dynamic config must be set before we can initialize the L2OutputOracle.
        require(dynamicConfigSet, "Dynamic oracle config is not yet initialized");

        // Upgrade and initialize the L2OutputOracle.
        config.globalConfig.proxyAdmin.upgradeAndCall(
            payable(config.proxyAddressConfig.l2OutputOracleProxy),
            address(config.implementationAddressConfig.l2OutputOracleImpl),
            abi.encodeCall(
                L2OutputOracle.initialize,
                (
                    l2OutputOracleDynamicConfig.l2OutputOracleStartingBlockNumber,
                    l2OutputOracleDynamicConfig.l2OutputOracleStartingTimestamp
                )
            )
        );

        // Upgrade and initialize the OptimismPortal.
        config.globalConfig.proxyAdmin.upgradeAndCall(
            payable(config.proxyAddressConfig.optimismPortalProxy),
            address(config.implementationAddressConfig.optimismPortalImpl),
            abi.encodeCall(OptimismPortal.initialize, (SuperchainConfig(config.globalConfig.superchainConfig)))
        );

        // Upgrade the L1CrossDomainMessenger.
        config.globalConfig.proxyAdmin.upgrade(
            payable(config.proxyAddressConfig.l1CrossDomainMessengerProxy),
            address(config.implementationAddressConfig.l1CrossDomainMessengerImpl)
        );

        // Try to initialize the L1CrossDomainMessenger, only fail if it's already been initialized.
        try L1CrossDomainMessenger(config.proxyAddressConfig.l1CrossDomainMessengerProxy).initialize(
            SuperchainConfig(config.globalConfig.superchainConfig)
        ) {
            // L1CrossDomainMessenger is the one annoying edge case difference between existing
            // networks and fresh networks because in existing networks it'll already be
            // initialized but in fresh networks it won't be. Try/catch is the easiest and most
            // consistent way to handle this because initialized() is not exposed publicly.
        } catch Error(string memory reason) {
            require(
                keccak256(abi.encodePacked(reason)) == keccak256("Initializable: contract is already initialized"),
                string.concat("Unexpected error initializing L1XDM: ", reason)
            );
        } catch {
            revert("Unexpected error initializing L1XDM (no reason)");
        }

        // Transfer ETH from the L1StandardBridge to the OptimismPortal.
        config.globalConfig.proxyAdmin.upgradeAndCall(
            payable(config.proxyAddressConfig.l1StandardBridgeProxy),
            address(config.implementationAddressConfig.portalSenderImpl),
            abi.encodeCall(PortalSender.donate, ())
        );

        // Upgrade the L1StandardBridge (no initializer).
        config.globalConfig.proxyAdmin.upgrade(
            payable(config.proxyAddressConfig.l1StandardBridgeProxy),
            address(config.implementationAddressConfig.l1StandardBridgeImpl)
        );

        // Upgrade the OptimismMintableERC20Factory (no initializer).
        config.globalConfig.proxyAdmin.upgrade(
            payable(config.proxyAddressConfig.optimismMintableERC20FactoryProxy),
            address(config.implementationAddressConfig.optimismMintableERC20FactoryImpl)
        );

        // Upgrade the L1ERC721Bridge (no initializer).
        config.globalConfig.proxyAdmin.upgrade(
            payable(config.proxyAddressConfig.l1ERC721BridgeProxy),
            address(config.implementationAddressConfig.l1ERC721BridgeImpl)
        );

        // Uprade the ProtocolVersions (no initializer).
        config.globalConfig.proxyAdmin.upgrade(
            payable(config.proxyAddressConfig.protocolVersionsProxy),
            address(config.implementationAddressConfig.protocolVersionsImpl)
        );

        // Try to initialize the L1StandardBridge, only fail if it's already been initialized.
        try L1StandardBridge(payable(config.proxyAddressConfig.l1StandardBridgeProxy)).initialize(
            SuperchainConfig(config.globalConfig.superchainConfig)
        ) {
            // L1StandardBridge is the one annoying edge case difference between existing
            // networks and fresh networks because in existing networks it'll already be
            // initialized but in fresh networks it won't be. Try/catch is the easiest and most
            // consistent way to handle this because initialized() is not exposed publicly.
        } catch Error(string memory reason) {
            require(
                keccak256(abi.encodePacked(reason)) == keccak256("Initializable: contract is already initialized"),
                string.concat("Unexpected error initializing L1SB: ", reason)
            );
        } catch {
            revert("Unexpected error initializing L1SB (no reason)");
        }

        // Try to set superchainConfig on the L1StandardBridge, only fail if it's already been set.
        // Try to initialize the L1StandardBridge, only fail if it's already been initialized.
        try L1StandardBridge(payable(config.proxyAddressConfig.l1StandardBridgeProxy)).setSuperchainConfig(
            SuperchainConfig(config.globalConfig.superchainConfig)
        ) {
            // L1StandardBridge is the one annoying edge case difference between existing
            // networks and fresh networks because in existing networks it'll already be
            // initialized but in fresh networks it won't be. Try/catch is the easiest and most
            // consistent way to handle this because initialized() is not exposed publicly.
        } catch Error(string memory reason) {
            require(
                keccak256(abi.encodePacked(reason)) == keccak256("SuperchainConfig already set"),
                string.concat("Unexpected error setting superchainConfig L1SB: ", reason)
            );
        } catch {
            revert("Unexpected error setting superchainConfig L1SB (no reason)");
        }

        // Try to initialize the L1ERC721Bridge, only fail if it's already been initialized.
        try L1ERC721Bridge(payable(config.proxyAddressConfig.l1ERC721BridgeProxy)).initialize(
            SuperchainConfig(config.globalConfig.superchainConfig)
        ) {
            // L1ERC721Bridge is the one annoying edge case difference between existing
            // networks and fresh networks because in existing networks it'll already be
            // initialized but in fresh networks it won't be. Try/catch is the easiest and most
            // consistent way to handle this because initialized() is not exposed publicly.
        } catch Error(string memory reason) {
            require(
                keccak256(abi.encodePacked(reason)) == keccak256("Initializable: contract is already initialized"),
                string.concat("Unexpected error initializing L1ERC721Bridge: ", reason)
            );
        } catch {
            revert("Unexpected error initializing L1ERC721Bridge (no reason)");
        }

        // Try to initialize the ProtocolVersions, only fail if it's already been initialized.
        try ProtocolVersions(config.proxyAddressConfig.protocolVersionsProxy).initialize(
            address(config.globalConfig.proxyAdmin),
            config.protocolVersionConfig.requiredProtocolVersion,
            config.protocolVersionConfig.recommendedProtocolVersion
        ) {
            // ProtocolVersions is the one annoying edge case difference between existing
            // networks and fresh networks because in existing networks it'll already be
            // initialized but in fresh networks it won't be. Try/catch is the easiest and most
            // consistent way to handle this because initialized() is not exposed publicly.
        } catch Error(string memory reason) {
            require(
                keccak256(abi.encodePacked(reason)) == keccak256("Initializable: contract is already initialized"),
                string.concat("Unexpected error initializing ProtocolVersions: ", reason)
            );
        } catch {
            revert("Unexpected error initializing ProtocolVersions (no reason)");
        }

        // Try to initialize the SuperchainConfig, only fail if it's already been initialized.
        try SuperchainConfig(config.globalConfig.superchainConfig).initialize(
            optimismPortalDynamicConfig.portalGuardian, false
        ) { } catch Error(string memory reason) {
            require(
                keccak256(abi.encodePacked(reason)) == keccak256("Initializable: contract is already initialized"),
                string.concat("Unexpected error initializing ProtocolVersions: ", reason)
            );
        } catch {
            revert("Unexpected error initializing SuperchainConfig (no reason)");
        }
    }

    /**
     * @notice Calls the first 2 steps of the migration process.
     */
    function phase1() external onlyOwner {
        step1();
        step2();
    }

    /**
     * @notice Calls the remaining steps of the migration process, and finalizes.
     */
    function phase2() external onlyOwner {
        step3();
        step4();
        step5();
        finalize();
    }

    /**
     * @notice Tranfers admin ownership to the final owner.
     */
    function finalize() public onlyOwner {
        // Transfer ownership of the ProxyAdmin to the final owner.
        config.globalConfig.proxyAdmin.transferOwnership(config.globalConfig.finalOwner);

        // Optionally also transfer AddressManager and L1StandardBridge if we still own it. Might
        // happen if we're exiting early.
        if (currentStep <= PROXY_TRANSFER_STEP) {
            // Transfer ownership of the AddressManager to the final owner.
            config.globalConfig.addressManager.transferOwnership(address(config.globalConfig.finalOwner));

            // Transfer ownership of the L1StandardBridge to the final owner.
            L1ChugSplashProxy(payable(config.proxyAddressConfig.l1StandardBridgeProxy)).setOwner(
                address(config.globalConfig.finalOwner)
            );

            // Transfer ownership of the L1ERC721Bridge to the final owner.
            Proxy(payable(config.proxyAddressConfig.l1ERC721BridgeProxy)).changeAdmin(
                address(config.globalConfig.finalOwner)
            );
        }

        // Mark the deployment as finalized.
        finalized = true;
    }

    /**
     * @notice First exit point, can only be called before step 3 is executed.
     */
    function exit1() external onlyOwner {
        require(currentStep == EXIT_1_NO_RETURN_STEP, "can only exit1 before step 3 is executed");

        // Reset the L1CrossDomainMessenger to the old implementation.
        config.globalConfig.addressManager.setAddress("OVM_L1CrossDomainMessenger", oldL1CrossDomainMessenger);

        // Unset the DTL shutoff block which will allow the DTL to sync again.
        config.globalConfig.addressManager.setAddress("DTL_SHUTOFF_BLOCK", address(0));

        // Mark the deployment as exited.
        exited = true;
    }
}
