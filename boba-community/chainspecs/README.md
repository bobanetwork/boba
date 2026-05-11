# Boba Chain Specs for op-reth

These JSON files are op-reth chain specifications for Boba networks, in the format consumed by `op-reth --chain=<path>`.

| File | Chain ID | Network |
|---|---|---|
| `boba.json` | 288 | Boba Mainnet |
| `boba-sepolia.json` | 28882 | Boba Sepolia Testnet |

Each file is a complete `Genesis` JSON containing:
- Chain config: `chainId`, the bedrock block (`bedrockBlock`: the L2 block at which Boba migrated from OVM to EVM — 1149019 for mainnet, 511 for sepolia), all hardfork activation times (canyon/delta/ecotone/fjord/granite/holocene), and the OP Stack `optimism` config (EIP-1559 params)
- OVM-era genesis block 0 metadata (`extraData`, `gasLimit`, etc.)

## Why these files exist (the upstream bedrockBlock bug)

Boba is bundled in the upstream OP Labs op-reth image (`us-docker.pkg.dev/oplabs-tools-artifacts/images/op-reth`) — its `superchain-configs.tar` includes the Boba mainnet and sepolia configs, and `--chain=boba` / `--chain=boba-sepolia` work out of the box.

What does *not* work out of the box is the **bedrock block**, which marks the L2 block where Boba migrated from the legacy OVM to the EVM under Bedrock. Pre-bedrock blocks cannot be executed by op-reth's EVM and require a separate l2geth (OVM) process to serve them. When `--rollup.historicalrpc=<legacy-l2geth-url>` is set, op-reth uses `bedrockBlock` to decide which requests to forward: queries for blocks numbered below `bedrockBlock` go to the legacy node; post-bedrock blocks are served locally.

Upstream op-reth has a bug: in `optimism/rust/op-reth/crates/chainspec/src/superchain/chain_metadata.rs`, `bedrock_block` is hardcoded to `Some(105235063)` only for OP Mainnet (chain ID 10) and `Some(0)` for every other chain. Boba (and any other migrated chain that isn't OP Mainnet) therefore gets `bedrockBlock=0` from the built-in spec, which silently disables historical RPC forwarding. The same bug is present in the older `ghcr.io/paradigmxyz/op-reth:v1.10.2` build.

The fix in upstream is a one-line change: each superchain-registry config already carries the migration block at `genesis.l2.number` (1149019 for Boba mainnet, 511 for Boba sepolia, 0 for born-bedrock chains like Base). The code just needs to use that field instead of the hardcoded OP-Mainnet-only special case.

Until upstream is fixed, this directory holds Boba chain specs with the correct `bedrockBlock` baked in, supplied at runtime via `--chain=<path-to-json>`.

## Regenerating these files

These files were produced by dumping the chain spec from a version of op-reth that bundles Boba and then patching `bedrockBlock`. They should not need to change unless Boba activates a new hardfork.

To regenerate:

```bash
docker run --rm us-docker.pkg.dev/oplabs-tools-artifacts/images/op-reth:v2.2.1 dump-genesis --chain boba 2>/dev/null \
  | tail -n +2 \
  | jq '.config.bedrockBlock = 1149019' > boba.json

docker run --rm us-docker.pkg.dev/oplabs-tools-artifacts/images/op-reth:v2.2.1 dump-genesis --chain boba-sepolia 2>/dev/null \
  | tail -n +2 \
  | jq '.config.bedrockBlock = 511' > boba-sepolia.json
```

(`tail -n +2` strips the leading log line that `op-reth` writes to stdout before the JSON; the `jq` step fixes the `bedrockBlock` field that the built-in spec leaves at 0.)

Once the upstream fix lands, these JSON overrides should become unnecessary and `--chain=boba` against the upstream image should work directly. Until then, mount the JSON.

## Verifying

To confirm the patched chain spec only differs from the upstream built-in spec in the `bedrockBlock` field:

```bash
docker run --rm us-docker.pkg.dev/oplabs-tools-artifacts/images/op-reth:v2.2.1 dump-genesis --chain boba 2>/dev/null \
  | tail -n +2 > /tmp/upstream.json
docker run --rm -v "$PWD:/cs" us-docker.pkg.dev/oplabs-tools-artifacts/images/op-reth:v2.2.1 \
  dump-genesis --chain /cs/boba.json 2>/dev/null | tail -n +2 > /tmp/ours.json
diff /tmp/upstream.json /tmp/ours.json
# Expected: a single hunk replacing "bedrockBlock": 0 with "bedrockBlock": 1149019
```
