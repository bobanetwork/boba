# Boba Chain Specs for op-reth

These JSON files are op-reth chain specifications for Boba networks, in the format consumed by `op-reth --chain=<path>`.

| File | Chain ID | Network |
|---|---|---|
| `boba.json` | 288 | Boba Mainnet |
| `boba-sepolia.json` | 28882 | Boba Sepolia Testnet |

Each file is a complete `Genesis` JSON containing:
- Chain config: `chainId`, the bedrock block (`bedrockBlock`: the L2 block at which Boba migrated from OVM to EVM — 1149019 for mainnet, 511 for sepolia), the pre-Bedrock L1-style hardfork activation blocks (`berlinBlock`, `londonBlock`, `arrowGlacierBlock`, `grayGlacierBlock`, `mergeNetsplitBlock` — all equal to the bedrock block for Boba, since Boba activated those at the migration rather than during the OVM era), the post-Bedrock hardfork activation timestamps (canyon/delta/ecotone/fjord/granite/holocene, plus `regolithTime` at the bedrock activation time), and the OP Stack `optimism` config (EIP-1559 params)
- OVM-era genesis block 0 metadata (`extraData`, `gasLimit`, etc.)

The L1-style hardfork blocks and `regolithTime` come from [`bobanetwork/op-geth`'s `params/superchain.go`](https://github.com/bobanetwork/op-geth/blob/optimism/params/superchain.go) — they are not currently expressed in the upstream superchain-registry, so op-reth defaults them to 0 (for L1 blocks) or 0 (for `regolithTime`) when reading the bundled config. Patching them in the JSON keeps op-reth's view of the chain consistent with op-geth's.

## Why these files exist

Boba is bundled in the upstream OP Labs op-reth image (`us-docker.pkg.dev/oplabs-tools-artifacts/images/op-reth`) — its `superchain-configs.tar` includes the Boba mainnet and sepolia configs, and `--chain=boba` / `--chain=boba-sepolia` work out of the box.

What does *not* work out of the box is the chain's **hardfork activation history**:

1. **`bedrockBlock`** marks the L2 block where Boba migrated from the legacy OVM to the EVM under Bedrock. Pre-bedrock blocks cannot be executed by op-reth's EVM and require a separate l2geth (OVM) process to serve them. When `--rollup.historicalrpc=<legacy-l2geth-url>` is set, op-reth uses `bedrockBlock` to decide which requests to forward.

   Upstream op-reth in `crates/chainspec/src/superchain/chain_metadata.rs` hardcodes `bedrock_block` to `Some(105235063)` only for OP Mainnet (chain ID 10) and `Some(0)` for every other chain. Boba (and any other migrated chain that isn't OP Mainnet) therefore gets `bedrockBlock=0` from the built-in spec, which silently disables historical RPC forwarding. A fix is proposed at <https://github.com/ethereum-optimism/optimism/pull/20638>.

2. **`berlinBlock` / `londonBlock` / `arrowGlacierBlock` / `grayGlacierBlock` / `mergeNetsplitBlock`** — for Boba, all of these L1-style EVM hardforks activated at the bedrock migration block (1149019 for mainnet, 511 for sepolia). Upstream op-reth doesn't carry these values; they're not currently expressed in the superchain-registry. The canonical source is `bobanetwork/op-geth`'s `params/superchain.go`.

3. **`regolithTime`** — Boba's Regolith hardfork activated at the bedrock migration timestamp (1713302879 for mainnet, 1705600788 for sepolia). Upstream op-reth hardcodes this to 0.

Until these are pushed into the upstream superchain-registry, this directory holds Boba chain specs with the correct values baked in, supplied at runtime via `--chain=<path-to-json>`.

### Path to deleting these files

If the upstream PR linked above lands and we then push the missing fields into the [superchain-registry](https://github.com/ethereum-optimism/superchain-registry) (adding `regolith_time` to `[hardforks]` and the L1-style hardfork blocks somewhere in the schema), upstream op-reth will produce the correct chain spec from `--chain=boba` directly and these JSON files can be retired.

## Regenerating these files

These files were produced by dumping the chain spec from a version of op-reth that bundles Boba and then patching `bedrockBlock`. They should not need to change unless Boba activates a new hardfork.

To regenerate:

```bash
docker run --rm us-docker.pkg.dev/oplabs-tools-artifacts/images/op-reth:v2.2.1 dump-genesis --chain boba 2>/dev/null \
  | tail -n +2 \
  | jq '.config.bedrockBlock = 1149019
      | .config.berlinBlock = 1149019
      | .config.londonBlock = 1149019
      | .config.arrowGlacierBlock = 1149019
      | .config.grayGlacierBlock = 1149019
      | .config.mergeNetsplitBlock = 1149019
      | .config.regolithTime = 1713302879' > boba.json

docker run --rm us-docker.pkg.dev/oplabs-tools-artifacts/images/op-reth:v2.2.1 dump-genesis --chain boba-sepolia 2>/dev/null \
  | tail -n +2 \
  | jq '.config.bedrockBlock = 511
      | .config.berlinBlock = 511
      | .config.londonBlock = 511
      | .config.arrowGlacierBlock = 511
      | .config.grayGlacierBlock = 511
      | .config.mergeNetsplitBlock = 511
      | .config.regolithTime = 1705600788' > boba-sepolia.json
```

(`tail -n +2` strips the leading log line that `op-reth` writes to stdout before the JSON. The `jq` step patches in the fields the upstream bundled config leaves at 0 — these values come from [`bobanetwork/op-geth`'s `params/superchain.go`](https://github.com/bobanetwork/op-geth/blob/optimism/params/superchain.go).)

## Verifying

To diff the patched chain spec against the upstream built-in spec and see exactly which fields we override:

```bash
docker run --rm us-docker.pkg.dev/oplabs-tools-artifacts/images/op-reth:v2.2.1 dump-genesis --chain boba 2>/dev/null \
  | tail -n +2 > /tmp/upstream.json
docker run --rm -v "$PWD:/cs" us-docker.pkg.dev/oplabs-tools-artifacts/images/op-reth:v2.2.1 \
  dump-genesis --chain /cs/boba.json 2>/dev/null | tail -n +2 > /tmp/ours.json
diff /tmp/upstream.json /tmp/ours.json
# Expected: hunks replacing bedrockBlock, berlinBlock, londonBlock, arrowGlacierBlock,
# grayGlacierBlock, mergeNetsplitBlock, and regolithTime (all 0 in upstream, correct values in ours)
```
