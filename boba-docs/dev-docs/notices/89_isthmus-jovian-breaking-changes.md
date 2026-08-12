# Preparing for Isthmus and Jovian breaking changes

The Isthmus and Jovian upgrades for Boba Sepolia require mandatory upgrades for node operators. Both are activated by timestamp:

* Isthmus (Boba Sepolia): 1787227200 (Thursday, August 20, 2026 at 12:00:00 UTC)
* Jovian (Boba Sepolia): 1787659200 (Tuesday, August 25, 2026 at 12:00:00 UTC)

Isthmus also activates the Prague EVM (`pragueTime`) at the same timestamp as Isthmus.

Boba Mainnet activation dates are not yet scheduled and will be announced in a future update to this notice.

:::warning
These hardforks are not yet included in the [superchain-registry](https://github.com/ethereum-optimism/superchain-registry). Until they are, the activation times are **not** present in the Boba Sepolia config bundled into the upstream `op-reth` / `op-node` images. Running with the built-in `--chain=boba-sepolia` / `--network=boba-sepolia` will derive against a stale fork schedule and diverge at the fork boundary. You **must** supply the Boba Sepolia chain spec and rollup config JSON files at runtime — see [Supply the Boba Sepolia config files](#supply-the-boba-sepolia-config-files-mandatory) below.
:::

## For Node Operators

Node operators are required to upgrade before the activation dates to avoid chain divergence.

### Update to the latest release

* op-reth at `v2.3.1` — image `us-docker.pkg.dev/oplabs-tools-artifacts/images/op-reth:v2.3.1`
* op-node at `v1.19.0` — image `us-docker.pkg.dev/oplabs-tools-artifacts/images/op-node:v1.19.0`

:::note
op-geth and op-erigon reached end-of-life on 2026-05-31 and do not support these upgrades. Operators still running them must migrate to op-reth; see the [op-geth/op-erigon to op-reth migration guide](https://github.com/bobanetwork/boba/blob/develop/boba-community/scripts/geth-to-reth/README.md).
:::

### Supply the Boba Sepolia config files (mandatory)

Because these forks are not yet in the superchain-registry, the activation times must be supplied at runtime rather than relying on the images' built-in Boba Sepolia config:

* op-reth: `--chain=<path>` pointing at [`boba-community/chainspecs/boba-sepolia-chainspec.json`](https://github.com/bobanetwork/boba/blob/develop/boba-community/chainspecs/boba-sepolia-chainspec.json)
* op-node: `--rollup.config=<path>` pointing at [`boba-community/rollup-configs/boba-sepolia-rollup.json`](https://github.com/bobanetwork/boba/blob/develop/boba-community/rollup-configs/boba-sepolia-rollup.json)

The bundled [`docker-compose-boba-sepolia.yml`](https://github.com/bobanetwork/boba/blob/develop/boba-community/docker-compose-boba-sepolia.yml) in `boba-community/` already mounts these files. Once the forks land in the superchain-registry and the upstream images are rebuilt against it, the built-in config may be used again.

### Verify your configuration

Both clients log their fork schedule at startup.

* Check that Isthmus is set to `1787227200` in the op-reth and op-node startup logs
* Check that Jovian is set to `1787659200` in the op-reth and op-node startup logs
