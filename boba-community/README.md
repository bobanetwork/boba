# Running a replica node

## Docker configuration

Instructions for running a Boba replica node for Mainnet or Testnet.

### Execution client

**op-reth is the only supported execution client.** op-geth and op-erigon reached end-of-life on 2026-05-31 and are no longer supported; if you are still running one, migrate to op-reth using the [migration guide](scripts/geth-to-reth/README.md).

There is now a single compose file per network:

| Compose file | Execution | Consensus | Optional |
|---|---|---|---|
| `docker-compose-boba-mainnet.yml` | **op-reth** | **op-node** | legacy l2geth (`--profile legacy`) |
| `docker-compose-boba-sepolia.yml` | **op-reth** | **op-node** | legacy l2geth (`--profile legacy`) |

> **Note:** The op-reth snapshots contain a pre-initialized reth database built from the op-geth state at the Bedrock migration block. Download, extract, and run — no manual `init-state` step is needed. To regenerate the database from scratch, see the [migration guide](scripts/geth-to-reth/README.md).

### Getting started

See the [running a node with Docker](../boba-docs/dev-docs/node-operators/1_run-node-docker.md) guide for full setup instructions including snapshot downloads, configuration, and starting the node.

### The initial synchronization

During the initial synchronization, you get log messages from the consensus node, and nothing else appears to happen.

```bash
INFO [08-04|16:36:07.150] Advancing bq origin                      origin=df76ff..48987e:8301316 originBehind=false
```

After a few minutes, the node finds the right batch and then it starts synchronizing.

```bash
INFO [08-04|16:36:01.204] Found next batch                         epoch=44e203..fef9a5:8301309 batch_epoch=8301309                batch_timestamp=1,673,567,518
INFO [08-04|16:36:01.205] generated attributes in payload queue    txs=2  timestamp=1,673,567,518
```

### Migrating from op-geth or op-erigon to op-reth

op-geth and op-erigon are no longer supported. See the [migration guide](scripts/geth-to-reth/README.md) for instructions on generating a reth database from an existing op-geth or op-erigon node.

### Optional: Run the legacy node

Due to the Anchorage migration, the modern client does not support some RPC requests for pre-Anchorage blocks, such as `debug_traceTransaction`. The legacy l2geth node is bundled into the same compose file behind the `legacy` profile (disabled by default). Enable it with:

```bash
docker compose -f docker-compose-boba-sepolia.yml --profile legacy up -d
```

The legacy Geth database can be downloaded from the [snapshot page](https://docs.boba.network/for-developers/node-operators/snapshot-downloads); extract it into `LEGACY_DATA_DIR` (default `./legacy-data`). op-reth serves its JSON-RPC on `8545` (HTTP) / `8546` (WS); the legacy node defaults to `8547` (HTTP) / `8548` (WS) so the two can run side by side.
