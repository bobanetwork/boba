# Running a Node with Docker

This tutorial walks you through using Docker to run a Boba Mainnet or Boba Sepolia node. You can find all Docker Compose files [here](https://github.com/bobanetwork/boba/tree/develop/boba-community).

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

The helper verifies the snapshot's sha256 checksum automatically and refuses to overwrite a non-empty data directory. It stages the download inside the data directory itself (not the container's `/tmp`), so only the data disk needs free space — roughly twice the snapshot size during extraction. To pin a specific snapshot, set `SNAPSHOT_URL` and `SNAPSHOT_SHA256` in your `.env` (see the [snapshot downloads](snapshot-downloads) page for the values).

> **Note:** the `--profile seed` flag is required — including with `run` — under both `docker compose` and `podman-compose`. podman-compose does not auto-enable a service's profile the way `docker compose run` does.

> **Note:** files in the data directory are written by the container and are therefore owned by the container's user, not your host user. Under rootless podman this is a mapped sub-UID, so to inspect or delete the data directory directly you'll need `podman unshare rm -rf <dir>` (or `sudo`). The node itself reads and writes it fine.

#### Option B — download manually

Snapshots are rotated regularly, so there is no stable URL — always take the current URL and sha256 for your network from the [snapshot downloads](snapshot-downloads) page (do not hardcode a dated URL). Download and verify it:

```bash
# Replace <SNAPSHOT_URL> with the current URL from the snapshot downloads page
curl -o snapshot.tar.zst -sL <SNAPSHOT_URL>
sha256sum snapshot.tar.zst   # must match the checksum on the snapshot downloads page
```

Extract it into the data directory (create it first if needed):

```bash
mkdir -p boba-mainnet-reth-datadir
tar --zstd -xf snapshot.tar.zst -C boba-mainnet-reth-datadir
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

For peer-to-peer sync, op-reth also uses the devp2p port `30303` — **TCP** for RLPx and **UDP** for discovery — which the compose file publishes. Outbound discovery and dialing work without any firewall changes (a node behind NAT syncs fine by dialing out); only if you also want to *accept inbound* peers do you need to forward TCP+UDP `30303` through your firewall/NAT.

> **Running both networks on one host:** the mainnet and sepolia compose files both publish `30303` (and the RPC ports). If you co-locate them, remap one network's host ports — e.g. change the `ports:` entries and the matching `--port` / `--discovery.port` / `--discovery.v5.port` flags so the two nodes don't collide.

## Operating the Node

### Start

```bash
docker compose -f [docker-compose-file] up -d
```

Will start the node in a detatched shell (`-d`), meaning the node will continue to run in the background.

### View Logs

```bash
docker compose logs -f --tail 10
```

To view logs of all containers.

```bash
docker compose logs <CONTAINER_NAME> -f --tail 10
```

### Stop

```bash
docker compose -f [docker-compose-file] down
```

### Wipe [DANGER]

```bash
docker compose -f [docker-compose-file] down -v
```

## Troubleshooting

### op-reth reports 0 peers / discovery finds nothing

Execution-layer (op-reth) peering is separate from op-node: your node still **syncs via op-node** even with 0 EL peers, so this is not fatal — but a healthy node should discover EL peers on its own.

op-reth discovers EL peers over **discv5** (UDP). If you run the compose stack under **rootless Podman**, the Podman project *bridge* network adds a second NAT hop on top of rootless networking's `passt` layer. That double translation does not preserve the UDP source port, which breaks discv5's handshake — so the node never finds peers (0 peers), even though everything else looks fine. **Rootful Docker is not affected** (its bridge is a single NAT), so the default compose files are correct for Docker users.

If you are on rootless Podman, give the execution client a single NAT by putting it on `pasta` and having op-node share its network namespace — op-node then reaches the engine over `localhost`, and all default ports are preserved:

```yaml
services:
  l2:
    network_mode: pasta          # single NAT instead of the compose bridge's double NAT
    # ...otherwise unchanged; move op-node's RPC port (8545) into this service's `ports:`
  op-node:
    network_mode: "service:l2"   # share l2's netns
    # change  --l2=http://l2:8551  ->  --l2=http://127.0.0.1:8551
```

To confirm discovery is healthy, run op-reth with `--log.stdout.filter info,discv5=debug` and look for `discv5 ... Session established` with the bootnodes, and a non-zero `admin_peers` count on the RPC.

