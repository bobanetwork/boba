// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.15;

import "forge-std/Script.sol";

// for tag op-contracts/v1.6.0

import { GameType, OutputRoot, Claim, GameStatus, Hash } from "src/dispute/lib/Types.sol";
import { LibHash } from "src/dispute/lib/LibUDT.sol";

import { ProxyAdmin } from "src/universal/ProxyAdmin.sol";
import { Proxy } from "src/universal/Proxy.sol";
import { AddressManager } from "src/legacy/AddressManager.sol";
import { IAddressManager } from "scripts/interfaces/IAddressManager.sol";
import { IDisputeGameFactory } from "src/dispute/interfaces/IDisputeGameFactory.sol";
import { IAnchorStateRegistry } from "src/dispute/interfaces/IAnchorStateRegistry.sol";

import { AnchorStateRegistry } from "src/dispute/AnchorStateRegistry.sol";
import { SuperchainConfig } from "src/L1/SuperchainConfig.sol";



contract Deploy1 is Script {

    function makeHash(bytes32 _root) internal pure returns (Hash root_) {
        assembly {
            root_ := _root
        }
    }
    function makeGame(uint32 _game) internal pure returns (GameType game_) {
        assembly {
            game_ := _game
        }
    }

    function run() external returns (address[3] memory) {
        address deployAddr = vm.envAddress("DEPLOY_ADDR");
        uint256 deployerPrivateKey = vm.envUint("DEPLOY_PRIVKEY");

        address guardian = vm.envAddress("DEPLOY1_GUARDIAN");
        require(guardian != address(0), "DEPLOY1_GUARDIAN not set");
        address proxyOwner = vm.envAddress("DEPLOY1_PROXYOWNER");
        require(proxyOwner != address(0), "DEPLOY1_PROXYOWNER not set");

        //uint256 deploySalt = vm.envOr("DEPLOY_SALT", uint256(0));

        vm.startBroadcast(deployerPrivateKey);

        //        address addressManager = new AddressManager{salt: salt_val}();
        AddressManager addressManager = new AddressManager();

        ProxyAdmin proxyAdmin = new ProxyAdmin(deployAddr);
        proxyAdmin.setAddressManager(addressManager);

        SuperchainConfig scImpl = new SuperchainConfig();

        Proxy scProxy = new Proxy(address(proxyAdmin));
        
        proxyAdmin.upgradeAndCall(
            payable(scProxy),
            address(scImpl),
            abi.encodeWithSignature("initialize(address,bool)", guardian, false));

        address dGF = 0x29Bd67B23cAC0E6bbDe1373E3859Dd25510f3331;
        AnchorStateRegistry asImpl = new AnchorStateRegistry(IDisputeGameFactory(dGF));
        
        Proxy asProxy = new Proxy(address(proxyAdmin));
      
        AnchorStateRegistry.StartingAnchorRoot[] memory _startingAnchorRoots = new AnchorStateRegistry.StartingAnchorRoot[](3);
                
        _startingAnchorRoots[0] = AnchorStateRegistry.StartingAnchorRoot({gameType:makeGame(0), outputRoot:OutputRoot({root:makeHash(bytes32(0xfb200fe732c3fb9969667c3ad904313db85e19804b0fd3e7000e0a578a4c1321)), l2BlockNumber: 17625240})});
        _startingAnchorRoots[1] = AnchorStateRegistry.StartingAnchorRoot({gameType:makeGame(1), outputRoot:OutputRoot({root:makeHash(bytes32(0x87d724ad1668b86c08608a022fc09366329a7a3545e47fe678b3c69ead6a91f8)), l2BlockNumber: 8515831})});
        _startingAnchorRoots[2] = AnchorStateRegistry.StartingAnchorRoot({gameType:makeGame(2), outputRoot:OutputRoot({root:makeHash(bytes32(0x87d724ad1668b86c08608a022fc09366329a7a3545e47fe678b3c69ead6a91f8)), l2BlockNumber: 8515831})});

        proxyAdmin.upgradeAndCall(
            payable(asProxy),
            address(asImpl),
            abi.encodeWithSelector(
                AnchorStateRegistry.initialize.selector,
                _startingAnchorRoots,
                address(scProxy)));

        addressManager.transferOwnership(proxyOwner);
        proxyAdmin.transferOwnership(proxyOwner);

        vm.stopBroadcast();
        return [address(proxyAdmin), address(scProxy), address(asProxy)];
    }
}
