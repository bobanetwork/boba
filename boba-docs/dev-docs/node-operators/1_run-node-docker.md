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

### Download Snapshots

Download the database snapshot for the client and network you wish to run. Always verify snapshots by comparing the sha256sum of the downloaded file to the sha256sum listed on the [snapshot downloads](snapshot-downloads) page.

```bash
sha256sum <filename>
```

* BOBA Mainnet

  ```bash
  curl -o boba-mainnet-reth-db-20260526.tar.zst -sL https://boba-db.s3.us-east-2.amazonaws.com/mainnet/boba-mainnet-reth-db-20260526.tar.zst
  ```

* BOBA Sepolia

  ```bash
  curl -o boba-sepolia-reth-db-20260526.tar.zst -sL https://boba-db.s3.us-east-2.amazonaws.com/sepolia/boba-sepolia-reth-db-20260526.tar.zst
  ```

Extract the snapshot into a `reth-data` directory:

```bash
mkdir -p reth-data
tar --zstd -xf boba-{network}-reth-db-20260526.tar.zst -C reth-data
```

These snapshots contain a pre-initialized reth database built from the op-geth state at the Bedrock migration block (block 1149019 for mainnet, block 511 for sepolia). No manual `init-state` step is needed — just extract and run. To regenerate the database from scratch, see the [migration guide](https://github.com/bobanetwork/boba/tree/develop/boba-community/scripts/geth-to-reth).

Set the `DATA_DIR` in your `.env` to point to this directory (or leave it blank to use the default `./reth-data`).

### Modify Volume Location

The volumes of l2 and op-node should be modified to your file locations.

```yaml
l2:
  volumes:
    - ./jwt-secret.txt:/config/jwt-secret.txt
    - DATA_DIR:/db
op-node:
  volumes:
  	- ./jwt-secret.txt:/config/jwt-secret.txt
```

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

