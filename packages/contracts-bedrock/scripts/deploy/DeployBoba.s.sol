// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Script } from "forge-std/Script.sol";
import { console } from "forge-std/console.sol";

// Interfaces
import { DeployUtils } from "scripts/libraries/DeployUtils.sol";
import { Solarray } from "scripts/libraries/Solarray.sol";
import { BaseDeployIO } from "scripts/deploy/BaseDeployIO.sol";

import { BOBA } from "src/boba/BOBA.sol";

contract DeployBobaOutput is BaseDeployIO {
    BOBA internal _bobaL1;

    function set(bytes4 _sel, address _addr) public {
        require(_addr != address(0), "DeployBobaOutput: cannot set zero address");

        // forgefmt: disable-start
        if (_sel == this.bobaL1.selector) _bobaL1 = BOBA(_addr);
        else revert("DeployBobaOutput: unknown selector");
        // forgefmt: disable-end
    }

    function bobaL1() public view returns (BOBA) {
        DeployUtils.assertValidContractAddress(address(_bobaL1));
        return _bobaL1;
    }
}

contract DeployBoba is Script {
    function run(DeployBobaOutput _dbo) public {
        vm.startBroadcast(msg.sender);
        deployBobaL1(_dbo);
        vm.stopBroadcast();

        address[] memory addrs = Solarray.addresses(address(_dbo.bobaL1()));
        DeployUtils.assertValidContractAddresses(addrs);
    }

    function deployBobaL1(DeployBobaOutput _dbo) public virtual {
        BOBA bobaL1 = new BOBA();

        address owner = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
        bobaL1.transfer(owner, 10000e18);

        _dbo.set(_dbo.bobaL1.selector, address(bobaL1));
    }
}
