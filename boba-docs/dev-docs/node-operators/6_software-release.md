# Node Software Releases

This page provides a list of the necessary versions of node software and instructions on how to keep them updated.

Our latest releases, notes and changelogs can be found on Github. `op-node` releases can be found [here](https://github.com/bobanetwork/boba/tags) and `op-reth` releases can be found [here](https://github.com/ethereum-optimism/op-reth/releases).

## Required Version by Network

These are the minimal required versions for node software by network. **op-reth is the only supported execution client.** op-geth and op-erigon reached end-of-life on 2026-05-31 and are no longer supported.

| Network          | op-node                                                      | op-reth                                                        |
| ---------------- | ------------------------------------------------------------ | -------------------------------------------------------------- |
| Boba Mainnet | [v1.16.5](https://github.com/bobanetwork/boba/releases/tag/op-node/v1.16.5) | [v2.2.1](https://github.com/ethereum-optimism/op-reth/releases) |
| Boba Sepolia | [v1.16.5](https://github.com/bobanetwork/boba/releases/tag/op-node/v1.16.5) | [v2.2.1](https://github.com/ethereum-optimism/op-reth/releases) |

> **Note:** op-reth and op-node are both published by OP Labs (`us-docker.pkg.dev/oplabs-tools-artifacts/images/{op-reth,op-node}`). Neither image includes Boba as a built-in network — the Boba chain spec for op-reth is supplied via a mounted JSON file, and the Boba rollup config for op-node likewise. See [`boba-community/chainspecs/`](https://github.com/bobanetwork/boba/tree/develop/boba-community/chainspecs) and [`boba-community/rollup-configs/`](https://github.com/bobanetwork/boba/tree/develop/boba-community/rollup-configs) for the JSON files and their regeneration procedures.

## [op-node v1.14.1](https://github.com/bobanetwork/boba/releases/tag/op-node/v1.14.1)

**Description**

This is a mandatory release for node operators on Boba Networks to support the L1 Fusaka upgrade.

**Required Action**

Upgrade your `op-node` software.

## op-geth and op-erigon (end-of-life)

op-geth and op-erigon reached end-of-life on 2026-05-31 and are no longer supported. Migrate to op-reth — see the [migration guide](https://github.com/bobanetwork/boba/tree/develop/boba-community/scripts/geth-to-reth) for generating a reth database from an existing op-geth or op-erigon node, and the [running a node](1_run-node-docker.md) guide for setup instructions.
