// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.15;

import "forge-std/Script.sol";
import { GameType, OutputRoot, Claim, GameStatus, Hash } from "src/dispute/lib/Types.sol";
import { Proxy } from "src/universal/Proxy.sol";
import { AnchorStateRegistry } from "src/dispute/AnchorStateRegistry.sol";

contract UpdateASR is Script {
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

    function run() external returns (address[1] memory) {
        address deployAddr = vm.envAddress("DEPLOY_ADDR");
        uint256 deployerPrivateKey = vm.envUint("DEPLOY_PRIVKEY");

        vm.startBroadcast(deployerPrivateKey);

        AnchorStateRegistry asImpl = AnchorStateRegistry(0xe5113600eEed9F73B5a63425d757f653a773718B);
        address SuperchainConfigProxy = 0x029A23c6E9D3026f984cd1Fd9C47906e4F5327F3;

        AnchorStateRegistry.StartingAnchorRoot[] memory startingAnchorRoots =
            new AnchorStateRegistry.StartingAnchorRoot[](4);

        startingAnchorRoots[0] = AnchorStateRegistry.StartingAnchorRoot({
            gameType: makeGame(0),
            outputRoot: OutputRoot({
                root: makeHash(bytes32(0xfb200fe732c3fb9969667c3ad904313db85e19804b0fd3e7000e0a578a4c1321)),
                l2BlockNumber: 17625240
            })
        });
        startingAnchorRoots[1] = AnchorStateRegistry.StartingAnchorRoot({
            gameType: makeGame(1),
            outputRoot: OutputRoot({
                root: makeHash(bytes32(0x87d724ad1668b86c08608a022fc09366329a7a3545e47fe678b3c69ead6a91f8)),
                l2BlockNumber: 8515831
            })
        });
        startingAnchorRoots[2] = AnchorStateRegistry.StartingAnchorRoot({
            gameType: makeGame(2),
            outputRoot: OutputRoot({
                root: makeHash(bytes32(0x87d724ad1668b86c08608a022fc09366329a7a3545e47fe678b3c69ead6a91f8)),
                l2BlockNumber: 8515831
            })
        });
        startingAnchorRoots[3] = AnchorStateRegistry.StartingAnchorRoot({
            gameType: makeGame(255),
            outputRoot: OutputRoot({
                root: makeHash(bytes32(0x87d724ad1668b86c08608a022fc09366329a7a3545e47fe678b3c69ead6a91f8)),
                l2BlockNumber: 8515831
            })
        });

        Proxy asProxy = new Proxy(deployAddr);

        asProxy.upgradeToAndCall(
            address(asImpl),
            abi.encodeWithSelector(AnchorStateRegistry.initialize.selector, startingAnchorRoots, SuperchainConfigProxy)
        );

        vm.stopBroadcast();
        return [address(asProxy)];
    }
}
