# Node Software Releases

This page provides a list of the necessary versions of node software and instructions on how to keep them updated.

Our latest releases, notes and changelogs can be found on Github. `op-node` releases can be found [here](https://github.com/bobanetwork/boba/tags) and `op-geth` release can be found [here](https://github.com/bobanetwork/op-geth/releases).

## Required Version by Network

These are the minimal required versions for node software by network. **op-reth is the recommended execution client.** op-geth is deprecated and will be removed on 2026-05-31. op-erigon is no longer supported.

| Network          | op-node                                                      | op-reth | op-geth (deprecated)                                         |
| ---------------- | ------------------------------------------------------------ | ------- | ------------------------------------------------------------ |
| Boba Mainnet | [v1.16.5](https://github.com/bobanetwork/boba/releases/tag/op-node/v1.16.5) | [v2.2.1](https://github.com/ethereum-optimism/op-reth/releases) | [v1.101500.0](https://github.com/bobanetwork/op-geth/releases/tag/v1.101503.0) |
| Boba Sepolia | [v1.16.5](https://github.com/bobanetwork/boba/releases/tag/op-node/v1.16.5) | [v2.2.1](https://github.com/ethereum-optimism/op-reth/releases) | [v1.101603.1](https://github.com/bobanetwork/op-geth/releases/tag/v1.101603.1) |

> **Note:** op-reth and op-node are both published by OP Labs (`us-docker.pkg.dev/oplabs-tools-artifacts/images/{op-reth,op-node}`). Neither image includes Boba as a built-in network — the Boba chain spec for op-reth is supplied via a mounted JSON file, and the Boba rollup config for op-node likewise. See [`boba-community/chainspecs/`](https://github.com/bobanetwork/boba/tree/develop/boba-community/chainspecs) and [`boba-community/rollup-configs/`](https://github.com/bobanetwork/boba/tree/develop/boba-community/rollup-configs) for the JSON files and their regeneration procedures.

## [op-node v1.14.1](https://github.com/bobanetwork/boba/releases/tag/op-node/v1.14.1)

**Description**

This is a mandatory release for node operators on Boba Networks to support the L1 Fusaka upgrade.

**Required Action**

Upgrade your `op-node` software.

## op-erigon

op-erigon is no longer supported. Migrate to op-reth. See the [running a node](1_run-node-docker.md) guide for setup instructions.

## [op-geth v1.101603.1](https://github.com/ethereum-optimism/op-geth/releases/tag/v1.101603.1)

**Description**

This is a mandatory release for node operators on Boba Networks to support the L1 Fusaka upgrade.

**Required Action**

* Upgrade your `op-geth` software.
