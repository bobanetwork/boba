# Boba Chain Specs for op-reth

These JSON files are op-reth chain specifications for Boba networks, in the format consumed by `op-reth --chain=<path>`.

| File | Chain ID | Network |
|---|---|---|
| `boba.json` | 288 | Boba Mainnet |
| `boba-sepolia.json` | 28882 | Boba Sepolia Testnet |

Each file is a complete `Genesis` JSON containing:
- Chain config: `chainId`, the bedrock block (`bedrockBlock`: the L2 block at which Boba migrated from OVM to EVM — 1149019 for mainnet, 511 for sepolia), all hardfork activation times (canyon/delta/ecotone/fjord/granite/holocene), and the OP Stack `optimism` config (EIP-1559 params)
- OVM-era genesis block 0 metadata (`extraData`, `gasLimit`, etc.)

## bedrockBlock and historical RPC forwarding

`bedrockBlock` marks the L2 block where Boba migrated from the legacy OVM (a custom Solidity-based VM) to the EVM under Bedrock. Pre-bedrock blocks cannot be executed by op-reth's EVM and require a separate l2geth (OVM) process to serve them.

When `--rollup.historicalrpc=<legacy-l2geth-url>` is set, op-reth uses `bedrockBlock` to decide which requests to forward: queries for blocks numbered below `bedrockBlock` are forwarded to the legacy node, while post-bedrock blocks are served locally.

If `bedrockBlock` is set to 0 (or unset), op-reth treats the chain as Bedrock-from-genesis and never forwards. The earlier `paradigmxyz/op-reth:v1.10.2` build of these chain specs had this bug — `dump-genesis` returned `bedrockBlock: 0` for Boba, which broke historical RPC forwarding. This directory contains the corrected values.

## Why these files exist

`op-reth` was originally built and released by Paradigm at `ghcr.io/paradigmxyz/op-reth`, which embedded all superchain-registry chains (including Boba) at compile time. Starting with op-reth v1.11.0, ownership of op-reth moved to OP Labs (`ethereum-optimism/op-reth`), and the upstream image now only embeds OP Mainnet, OP Sepolia, Base Mainnet, Base Sepolia, and a dev chain — Boba is no longer built in.

To run op-reth on Boba with the OP Labs image, the chain spec must be supplied at runtime via `--chain=<path-to-json>`.

## Regenerating these files

The files in this directory were produced by extracting the in-memory chain spec from the last paradigmxyz build that included Boba (`ghcr.io/paradigmxyz/op-reth:v1.10.2`), then **manually correcting `bedrockBlock`** (which paradigm's build defaulted to 0). They should not need to change unless Boba activates a new hardfork.

If you do need to regenerate them (e.g. after a hardfork is added to the superchain-registry), run:

```bash
docker run --rm ghcr.io/paradigmxyz/op-reth:v1.10.2 dump-genesis --chain boba 2>/dev/null \
  | tail -n +2 \
  | jq '.config.bedrockBlock = 1149019' > boba.json

docker run --rm ghcr.io/paradigmxyz/op-reth:v1.10.2 dump-genesis --chain boba-sepolia 2>/dev/null \
  | tail -n +2 \
  | jq '.config.bedrockBlock = 511' > boba-sepolia.json
```

(`tail -n +2` strips the leading log line that `op-reth` writes to stdout before the JSON; the `jq` step fixes the bedrockBlock that paradigm leaves at 0.)

Note: this approach is frozen at v1.10.2 of paradigm's chain spec definitions. For new hardforks added after v1.10.2, the chain spec JSON will need to be edited manually or extracted from a newer source.

## Verifying

To confirm a chain spec produces an identical OpChainSpec to the paradigm built-in:

```bash
docker run --rm ghcr.io/paradigmxyz/op-reth:v1.10.2 dump-genesis --chain boba | tail -n +2 > /tmp/a.json
docker run --rm -v "$PWD:/cs" us-docker.pkg.dev/oplabs-tools-artifacts/images/op-reth:v2.2.1 \
  dump-genesis --chain /cs/boba.json | tail -n +2 > /tmp/b.json
diff /tmp/a.json /tmp/b.json && echo "OK"
```
