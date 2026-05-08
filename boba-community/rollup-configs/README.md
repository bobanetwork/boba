# Boba Rollup Configs for op-node

These JSON files are op-node rollup configurations for Boba networks, in the format consumed by `op-node --rollup.config=<path>`.

| File | Chain ID | Network |
|---|---|---|
| `boba-mainnet.json` | 288 | Boba Mainnet |
| `boba-sepolia.json` | 28882 | Boba Sepolia Testnet |

Each file describes the rollup's L1/L2 genesis anchor, hardfork timestamps, batch inbox address, sequencer window size, and other parameters that op-node needs to derive L2 blocks from L1 data.

## Why these files exist

These rollup configs let us run the **upstream OP Labs op-node image** (`us-docker.pkg.dev/oplabs-tools-artifacts/images/op-node`) with Boba — Boba is not in op-node's built-in network list, so the rollup config is supplied at runtime via `--rollup.config=<path>` instead of `--network=boba-{mainnet,sepolia}`.

The Boba-built op-node image (`us-docker.pkg.dev/boba-392114/bobanetwork-tools-artifacts/images/op-node`) does include Boba as a built-in network, but switching to upstream removes a Boba-specific build step.

## Regenerating these files

The files in this directory were produced from the boba-built op-node:

```bash
docker run --rm us-docker.pkg.dev/boba-392114/bobanetwork-tools-artifacts/images/op-node:v1.16.5 \
  op-node networks dump-rollup-config --network=boba-mainnet > boba-mainnet.json

docker run --rm us-docker.pkg.dev/boba-392114/bobanetwork-tools-artifacts/images/op-node:v1.16.5 \
  op-node networks dump-rollup-config --network=boba-sepolia > boba-sepolia.json
```

They should not need regenerating unless Boba activates a new hardfork.
