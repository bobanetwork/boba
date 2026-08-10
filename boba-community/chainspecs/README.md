# Boba Chain Specs for op-reth

These JSON files are op-reth chain specifications for Boba networks, in the format consumed by `op-reth --chain=<path>`.

| File | Chain ID | Network |
|---|---|---|
| `boba.json` | 288 | Boba Mainnet |
| `boba-sepolia.json` | 28882 | Boba Sepolia Testnet |

Each file is a genesis configuration — the chain config and OVM-era genesis block 0 header metadata that op-reth needs to construct an `OpChainSpec`. These are *not* full L2 genesis files: they do not contain an `alloc` section with initial account states. The L2 state is imported separately via `op-reth init-state` from a state dump (see [`../scripts/geth-to-reth/`](../scripts/geth-to-reth/)).

For **`boba.json` (mainnet)**, the only field we override versus the upstream OP Labs op-reth's built-in spec is `bedrockBlock` (the L2 block at which Boba migrated from OVM to EVM — 1149019). Everything else (chainId, hardfork timestamps, OP Stack `optimism` config, OVM-era genesis block 0 metadata) is left at the upstream default.

For **`boba-sepolia.json`**, in addition to `bedrockBlock` (511) we also carry the **Isthmus, Jovian, and Prague hardfork activation times** (`isthmusTime` / `jovianTime` / `pragueTime`). These forks have not yet been merged into the [superchain-registry](https://github.com/ethereum-optimism/superchain-registry), so the upstream image's built-in `--chain=boba-sepolia` spec does **not** contain them. Supplying this file is therefore **mandatory** on Sepolia — a node started with `--chain=boba-sepolia` would run against a stale fork schedule and diverge at the fork boundary. The committed `boba-sepolia.json` is the canonical spec; run op-reth against it directly (see [Regenerating these files](#regenerating-these-files) for how it is produced).

## Why these files exist

Boba is bundled in the upstream OP Labs op-reth image (`us-docker.pkg.dev/oplabs-tools-artifacts/images/op-reth`) — its `superchain-configs.tar` includes the Boba mainnet and sepolia configs. On mainnet, `--chain=boba` works out of the box for everything except `bedrockBlock`. On sepolia, `--chain=boba-sepolia` is additionally missing the pending Isthmus/Jovian/Prague fork times (see above), so the bundled spec cannot be used at all until those land in the registry.

`bedrockBlock` marks the L2 block where Boba migrated from the legacy OVM to the EVM under Bedrock. Pre-bedrock blocks cannot be executed by op-reth's EVM and require a separate l2geth (OVM) process to serve them. When `--rollup.historicalrpc=<legacy-l2geth-url>` is set, op-reth uses `bedrockBlock` to decide which requests to forward: requests for blocks before `bedrockBlock` go to the legacy node, the rest are served locally.

Upstream op-reth in `crates/chainspec/src/superchain/chain_metadata.rs` hardcodes `bedrock_block` to `Some(105235063)` only for OP Mainnet (chain ID 10) and `Some(0)` for every other chain. Boba (and any other migrated chain that isn't OP Mainnet) therefore gets `bedrockBlock=0` from the built-in spec, which silently disables historical RPC forwarding. A fix is proposed at <https://github.com/ethereum-optimism/optimism/pull/20638>.

Until that fix lands and the upstream image is rebuilt against it, this directory holds Boba chain specs with the correct `bedrockBlock` supplied at runtime via `--chain=<path-to-json>`.

### Why we don't also override the L1 hardfork blocks and regolithTime

`bobanetwork/op-geth`'s `params/superchain.go` carries additional Boba-specific values that the superchain-registry doesn't express: `berlinBlock` / `londonBlock` / `arrowGlacierBlock` / `grayGlacierBlock` / `mergeNetsplitBlock` (all = bedrock block, since Boba activated those at the migration rather than during the OVM era) and `regolithTime` (= bedrock activation timestamp).

We **don't** override these here, because **changing them invalidates an existing database**. Moving `londonBlock` from 0 to a non-zero value removes `baseFeePerGas` from the genesis block 0 header, which changes the genesis block hash. op-reth refuses to open a database whose stored genesis hash doesn't match the chainspec's computed hash:

```
ERROR shutting down due to error err=genesis hash in the storage does not match the specified chainspec
```

For a running node started against the upstream defaults (all those fields at 0), this would require a full re-init. The runtime behavior of post-bedrock blocks is unaffected by where these hardforks are declared as activating — they're all "in the past" relative to any post-bedrock block — so the override would be documentary correctness at the cost of operational disruption.

If/when the upstream registry is extended to express these fields and upstream op-reth reads them, new databases will use the correct values from the start and the override would be unnecessary.

### Path to deleting these files

**`boba.json` (mainnet):** if the upstream PR linked above lands and we then push the missing `bedrockBlock` mapping into the [superchain-registry](https://github.com/ethereum-optimism/superchain-registry) (it's already there as `genesis.l2.number` — the PR just teaches op-reth to read it), upstream op-reth will produce the correct chain spec from `--chain=boba` directly and this JSON file can be retired.

**`boba-sepolia.json`:** in addition to the above, the Isthmus/Jovian/Prague fork times must first be merged into the superchain-registry and the upstream op-reth image rebuilt against it. Until then this file cannot be retired.

## Regenerating these files

Both committed files are normalized with keys sorted lexicographically (`jq -S`)
so that regenerating them produces a minimal, reviewable diff. Keep the `-S`
when refreshing.

### boba-sepolia.json

The committed `boba-sepolia.json` is the canonical Boba Sepolia chain spec (it is
already in the `--chain=<path>` genesis format op-reth expects, with an empty
`alloc`). It should not need regenerating unless Boba Sepolia activates a new
hardfork.

<!-- Maintainer note: the file is refreshed from the RaaS provider endpoint that
     backs Boba's infrastructure. This URL is an internal workflow detail — keep
     it out of operator-facing docs; point operators at the committed file. -->

```bash
curl -fsSL https://raas-backend.g.alchemy.com/rollups/boba-sepolia/genesis.json | jq -S . > boba-sepolia.json
```

### boba.json (mainnet)

To regenerate the mainnet spec from the upstream bundled spec:

```bash
docker run --rm us-docker.pkg.dev/oplabs-tools-artifacts/images/op-reth:v2.3.1 dump-genesis --chain boba 2>/dev/null \
  | tail -n +2 \
  | jq -S '.config.bedrockBlock = 1149019' > boba.json
```

(`tail -n +2` strips the leading log line that `op-reth` writes to stdout before the JSON; the `jq` step patches in the only field that the upstream bundled config leaves at 0.)

## Verifying (mainnet)

To diff the patched mainnet chain spec against the upstream built-in spec and confirm the only difference is `bedrockBlock`:

```bash
docker run --rm us-docker.pkg.dev/oplabs-tools-artifacts/images/op-reth:v2.3.1 dump-genesis --chain boba 2>/dev/null \
  | tail -n +2 > /tmp/upstream.json
docker run --rm -v "$PWD:/cs" us-docker.pkg.dev/oplabs-tools-artifacts/images/op-reth:v2.3.1 \
  dump-genesis --chain /cs/boba.json 2>/dev/null | tail -n +2 > /tmp/ours.json
diff /tmp/upstream.json /tmp/ours.json
# Expected: a single hunk replacing "bedrockBlock": 0 with "bedrockBlock": 1149019
```
