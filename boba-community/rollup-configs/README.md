# Boba Rollup Configs for op-node

These JSON files are op-node rollup configurations for Boba networks, in the format consumed by `op-node --rollup.config=<path>`.

| File | Chain ID | Network |
|---|---|---|
| `boba-mainnet-rollup.json` | 288 | Boba Mainnet |
| `boba-sepolia-rollup.json` | 28882 | Boba Sepolia Testnet |

Each file describes the rollup's L1/L2 genesis anchor, hardfork timestamps, batch inbox address, sequencer window size, and other parameters that op-node needs to derive L2 blocks from L1 data.

## Why these files exist

These rollup configs let us run the **upstream OP Labs op-node image** (`us-docker.pkg.dev/oplabs-tools-artifacts/images/op-node`) with Boba — Boba is not in op-node's built-in network list, so the rollup config is supplied at runtime via `--rollup.config=<path>` instead of `--network=boba-{mainnet,sepolia}`.

For **Boba Sepolia**, supplying this file is now **mandatory**, not just a convenience: the Isthmus and Jovian hardfork activation times have not yet been merged into the [superchain-registry](https://github.com/ethereum-optimism/superchain-registry). Neither the upstream `--network=boba-sepolia` config nor the Boba-built op-node's built-in network carries those fork times, so a node started without `--rollup.config` would derive blocks against a stale fork schedule and diverge. Run with this file until the hardforks land in the registry and the images are rebuilt against it.

## Regenerating these files

The `boba-sepolia-rollup.json` committed here is the canonical rollup config for Boba
Sepolia; run op-node against it directly. It should not need regenerating unless
Boba Sepolia activates a new hardfork.

Both committed files are normalized with keys sorted lexicographically (`jq -S`)
so that regenerating them produces a minimal, reviewable diff. Keep the `-S`
when refreshing.

<!-- Maintainer note: the file is refreshed from the RaaS provider endpoint that
     backs Boba's infrastructure. This URL is an internal workflow detail — keep
     it out of operator-facing docs; point operators at the committed file. -->

```bash
curl -fsSL https://raas-backend.g.alchemy.com/rollups/boba-sepolia/rollup.json | jq -S . > boba-sepolia-rollup.json
```

`boba-mainnet-rollup.json` was produced from the boba-built op-node and should not need regenerating unless Boba Mainnet activates a new hardfork:

```bash
docker run --rm us-docker.pkg.dev/boba-392114/bobanetwork-tools-artifacts/images/op-node:v1.16.5 \
  op-node networks dump-rollup-config --network=boba-mainnet | jq -S . > boba-mainnet-rollup.json
```
