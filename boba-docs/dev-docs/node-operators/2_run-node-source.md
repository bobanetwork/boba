# Running a Node from Source

Running a Boba node with the [Docker images](1_run-node-docker.md) is the recommended path. Boba runs the **upstream** op-reth and op-node binaries unmodified — the network is configured entirely through runtime files, not a custom build — so there is normally no reason to build from source. Build from source only if you have a specific need, such as targeting an unusual architecture.

If you do, build the upstream binaries from their own documentation, then apply the Boba-specific runtime configuration below.

## Build the binaries (upstream)

Boba does not fork either client, so a stock upstream build works:

- **op-reth** (execution client) — build from the OP Labs [op-reth](https://github.com/ethereum-optimism/op-reth) repository, following its build instructions (`cargo build --bin op-reth --release`).
- **op-node** (rollup / consensus client) — build from the [Optimism monorepo](https://github.com/ethereum-optimism/optimism), following the [op-node README](https://github.com/ethereum-optimism/optimism/tree/develop/op-node) (`make op-node`).

## Configure for Boba

The upstream images include **no** built-in Boba network, and the built-in `--chain=boba-sepolia` / `--chain=boba` configs are stale and will make your node **diverge** from the canonical chain. Supply the Boba configuration at runtime, exactly as the compose files in [`boba-community/`](https://github.com/bobanetwork/boba/tree/develop/boba-community) do — those files are the canonical reference for the full command line:

- **op-reth `--chain=<path>`** → `boba-community/chainspecs/boba-sepolia-chainspec.json` (Sepolia) or `boba-mainnet-chainspec.json` (Mainnet). See [`chainspecs/README.md`](https://github.com/bobanetwork/boba/tree/develop/boba-community/chainspecs) for why this is required.
- **op-node `--rollup.config=<path>`** → `boba-community/rollup-configs/boba-sepolia-rollup.json` or `boba-mainnet-rollup.json`.
- **Bootnodes** — pass the Boba EL bootnodes to op-reth (`--bootnodes`) and the CL bootnodes to op-node (`--p2p.bootnodes`); see the [Bootnodes](4_bootnodes.md) page.
- **Remaining flags** — JWT secret, `--l1`, `--l1.beacon`, `--syncmode=execution-layer`, `--rollup.sequencer-http`, and the RPC/p2p ports: copy the values straight from the compose files.

Finally, seed the database from a [snapshot](5_snapshot-downloads.md) before starting, the same as for the Docker setup.
