# Running a Node with Docker

This tutorial will walk you through the process of using Docker to run an BOBA Sepolia node, OP Mainnet node and OP Sepolia node. You can find all Docker Compose files [here](https://github.com/bobanetwork/boba/tree/develop/boba-community).

## Execution Client

**op-reth is the only supported execution client.** op-geth and op-erigon reached end-of-life on 2026-05-31 and are no longer supported. If you are still running one, see the [migration guide](https://github.com/bobanetwork/boba/tree/develop/boba-community/scripts/geth-to-reth) to generate a reth database from your existing data.

There is a single compose file per network:

| Compose File | Execution Client | Consensus Client | Optional |
|---|---|---|---|
| `docker-compose-boba-mainnet.yml` | **op-reth** | **op-node** | legacy l2geth (`--profile legacy`) |
| `docker-compose-boba-sepolia.yml` | **op-reth** | **op-node** | legacy l2geth (`--profile legacy`) |

> **op-reth** uses the upstream [OP Labs op-reth](https://github.com/ethereum-optimism/op-reth) image. The Boba chain spec is supplied at runtime via a JSON file mounted from `boba-community/chainspecs/` — no custom op-reth build is required. See [`boba-community/chainspecs/README.md`](https://github.com/bobanetwork/boba/blob/develop/boba-community/chainspecs/README.md) for background on why this is required and how to regenerate the chain spec files.

## Prerequisites

* [docker](https://docs.docker.com/engine/install/)
* [docker-compose](https://docs.docker.com/compose/install/)

## Setup

Clone the `boba` repository to get started

```bash
git clone https://github.com/bobanetwork/boba.git
cd boba
cd boba-community
```

## Configuration

Configuration for the `docker-compose` is handled through environment variables inside of an `.env` file.

### Create an `.env` file

The repository includes a sample environment variable file located at `.env.example` that you can copy and modify to get started. Make a copy of this file and name it `.env`.

```bash
cp .env.example .env
```

### Configure the `.env` file

Open the `.env` in your directory and set the variables inside. Read the descriptions of each variable to understand what they do and how to set them. Read the [software release](software-release) page to set the correct version.

## DB Configuration

### Create a Shared Secret (JWT Token)

```bash
openssl rand -hex 32 > jwt-secret.txt
```

### Seed the Database

A fresh node needs a database before it can sync. The reth snapshots contain a pre-initialized database built from the op-geth state at the Bedrock migration block (block 1149019 for mainnet, block 511 for sepolia), so you do not need a manual `init-state` step — just get the snapshot into the data directory.

By default the database lives in `./boba-mainnet-reth-datadir` or `./boba-sepolia-reth-datadir` (next to the compose file). Set `DATA_DIR` in your `.env` to relocate it.

#### Option A — automated seeder (recommended)

The compose files include a one-shot `seed-database` helper that downloads and extracts the latest published snapshot for you. Run it once against an empty data directory before starting the node:

```bash
# BOBA Mainnet
docker compose -f docker-compose-boba-mainnet.yml --profile seed run --rm seed-database

# BOBA Sepolia
docker compose -f docker-compose-boba-sepolia.yml --profile seed run --rm seed-database
```

The helper verifies the snapshot's sha256 checksum automatically and refuses to overwrite a non-empty data directory. To pin a specific snapshot, set `SNAPSHOT_URL` and `SNAPSHOT_SHA256` in your `.env` (see the [snapshot downloads](snapshot-downloads) page for the values).

> **Note:** the `--profile seed` flag is required — including with `run` — under both `docker compose` and `podman-compose`. podman-compose does not auto-enable a service's profile the way `docker compose run` does.

#### Option B — download manually

Download the snapshot for your network and always verify its sha256sum against the [snapshot downloads](snapshot-downloads) page (`sha256sum <filename>`):

```bash
# BOBA Mainnet
curl -o boba-mainnet-reth-db-20260526.tar.zst -sL https://boba-db.s3.us-east-2.amazonaws.com/mainnet/boba-mainnet-reth-db-20260526.tar.zst

# BOBA Sepolia
curl -o boba-sepolia-reth-db-20260526.tar.zst -sL https://boba-db.s3.us-east-2.amazonaws.com/sepolia/boba-sepolia-reth-db-20260526.tar.zst
```

Extract it into the data directory (create it first if needed):

```bash
mkdir -p boba-mainnet-reth-datadir
tar --zstd -xf boba-mainnet-reth-db-20260526.tar.zst -C boba-mainnet-reth-datadir
```

To regenerate the database from scratch instead of using a snapshot, see the [migration guide](https://github.com/bobanetwork/boba/tree/develop/boba-community/scripts/geth-to-reth).

## Run the Node

Once you've configured your `.env` file, you can run the node using Docker Compose.

```bash
# BOBA Mainnet
docker compose -f docker-compose-boba-mainnet.yml up -d

# BOBA Sepolia
docker compose -f docker-compose-boba-sepolia.yml up -d
```

### Optional: legacy (pre-Anchorage) node

Some RPC methods (e.g. `debug_traceTransaction`) are not available for blocks from before the Anchorage migration. A legacy l2geth node that serves those historical blocks is bundled into the same compose file behind the `legacy` profile, disabled by default. To run it alongside op-reth, download the legacy snapshot from the [snapshot downloads](snapshot-downloads) page, extract it into `LEGACY_DATA_DIR` (default `./legacy-data`), and start with the profile enabled:

```bash
# BOBA Mainnet
docker compose -f docker-compose-boba-mainnet.yml --profile legacy up -d

# BOBA Sepolia
docker compose -f docker-compose-boba-sepolia.yml --profile legacy up -d
```

When you run the legacy node, point op-reth at your local replica so pre-Anchorage queries are served by it instead of the public endpoint. Set the following in your `.env` before starting:

```bash
HISTORICAL_RPC=http://legacy-l2:8545
```

op-reth reaches the legacy node by its service name (`legacy-l2`) on the compose network — `8545` is the legacy node's *in-container* RPC port, independent of the host port below.

op-reth serves its JSON-RPC on `8545` (HTTP) and `8546` (WS); op-node's rollup RPC is on `9545`. The optional legacy node defaults to host ports `8547` (HTTP) / `8548` (WS) so it can run alongside op-reth.

## Operating the Node

### Start

```bash
docker-compose -f [docker-compose-file] up -d
```

Will start the node in a detatched shell (`-d`), meaning the node will continue to run in the background.

### View Logs

```bash
docker-compose logs -f --tail 10
```

To view logs of all containers.

```bash
docker-compose logs <CONTAINER_NAME> -f --tail 10
```

### Stop

```bash
docker-compose -f [docker-compose-file] down
```

### Wipe [DANGER]

```bash
docker-compose -f [docker-compose-file] down -v
```

