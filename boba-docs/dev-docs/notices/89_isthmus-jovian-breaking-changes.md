# Preparing for Isthmus and Jovian breaking changes

The Isthmus and Jovian upgrades for Boba Mainnet and Boba Sepolia require mandatory upgrades for node operators. Both are activated by timestamp:

* Isthmus (Boba Mainnet): 1796058000 (Monday, November 30, 2026 at 17:00:00 UTC)
* Jovian (Boba Mainnet): 1796230800 (Wednesday, December 2, 2026 at 17:00:00 UTC)
* Isthmus (Boba Sepolia): 1787227200 (Thursday, August 20, 2026 at 12:00:00 UTC)
* Jovian (Boba Sepolia): 1787659200 (Tuesday, August 25, 2026 at 12:00:00 UTC)

Isthmus also activates the Prague EVM (`pragueTime`) at the same timestamp as Isthmus.

:::warning
These hardforks are not yet included in the [superchain-registry](https://github.com/ethereum-optimism/superchain-registry). Until they are, the activation times are **not** present in the Boba configs bundled into the upstream `op-reth` / `op-node` images. Running with the built-in `--chain=boba` / `--chain=boba-sepolia` or `--network=boba-mainnet` / `--network=boba-sepolia` will derive against a stale fork schedule and diverge at the fork boundary. You **must** supply the Boba chain spec and rollup config JSON files at runtime — see [Supply the Boba config files](#supply-the-boba-config-files-mandatory) below.
:::

## For Node Operators

Node operators are required to upgrade before the activation dates to avoid chain divergence.

### Update to a supported release

These forks are supported from **op-reth `v2.3.1`** / **op-node `v1.19.0`** onward — run at least these versions. The current recommended versions are **op-reth `v2.8.1`** / **op-node `v1.20.1`**; always check the [Node Software Releases](../node-operators/6_software-release.md) page, which is kept up to date.

:::note
op-geth and op-erigon reached end-of-life on 2026-05-31 and do not support these upgrades. Operators still running them must migrate to op-reth; see the [op-geth/op-erigon to op-reth migration guide](https://github.com/bobanetwork/boba/blob/develop/boba-community/scripts/geth-to-reth/README.md).
:::

### Supply the Boba config files (mandatory)

Because these forks are not yet in the superchain-registry, the activation times must be supplied at runtime rather than relying on the images' built-in Boba configs:

| Network | op-reth `--chain=<path>` | op-node `--rollup.config=<path>` |
|---|---|---|
| Boba Mainnet | [`boba-mainnet-chainspec.json`](https://github.com/bobanetwork/boba/blob/develop/boba-community/chainspecs/boba-mainnet-chainspec.json) | [`boba-mainnet-rollup.json`](https://github.com/bobanetwork/boba/blob/develop/boba-community/rollup-configs/boba-mainnet-rollup.json) |
| Boba Sepolia | [`boba-sepolia-chainspec.json`](https://github.com/bobanetwork/boba/blob/develop/boba-community/chainspecs/boba-sepolia-chainspec.json) | [`boba-sepolia-rollup.json`](https://github.com/bobanetwork/boba/blob/develop/boba-community/rollup-configs/boba-sepolia-rollup.json) |

Mainnet operators who already mount the mainnet files must **re-download them**: earlier versions did not contain the fork times.

The bundled [`docker-compose-boba-mainnet.yml`](https://github.com/bobanetwork/boba/blob/develop/boba-community/docker-compose-boba-mainnet.yml) and [`docker-compose-boba-sepolia.yml`](https://github.com/bobanetwork/boba/blob/develop/boba-community/docker-compose-boba-sepolia.yml) in `boba-community/` already mount these files. Once the forks land in the superchain-registry and the upstream images are rebuilt against it, the built-in config may be used again.

### Verify your configuration

Both clients log their fork schedule at startup.

* Check that Isthmus is set to `1796058000` (Mainnet) / `1787227200` (Sepolia) in the op-reth and op-node startup logs
* Check that Jovian is set to `1796230800` (Mainnet) / `1787659200` (Sepolia) in the op-reth and op-node startup logs
