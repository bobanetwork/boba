#!/bin/bash

export DEPLOY_CONFIG_PATH=./deploy-config/boba-mainnet-faultgame.json

export RPC_URL=

# Addresses:
#        address _proxyAdmin,
#        address _systemOwnerSafe,
#        address _superchainConfigProxy,
#        address _disputeGameFactoryImpl,
#        address _delayedWethImpl,
#        address _preimageOracleImpl,
#        address _mipsImpl,
#        address _optimismPortal2Impl
forge script \
  --rpc-url=$RPC_URL \
  scripts/upgrades/v1.3.0/v1.8.0-permissioned/DeployUpgrade.s.sol:DeployUpgrade \
  --sig "deploy(address,address,address,address,address,address,address,address)" \
  0x6e598cec2701FfAA3c06175dc3Af0317a749a0Dc \
  0x56121a8612474C3eB65D69a3b871f284705b9bC4 \
  0x996ffD627901f10C80A7d4B72A12316D2e77c076 \
  0xc641A33cab81C559F2bd4b21EA34C290E2440C2B \
  0x71e966Ae981d1ce531a7b6d23DC0f27B38409087 \
  0x9c065e11870B891D214Bc2Da7EF1f9DDFA1BE277 \
  0x5fE03a12C1236F9C22Cb6479778DDAa4bce6299C \
  0xe2F826324b2faf99E513D16D266c3F80aE87832B \
  $@
