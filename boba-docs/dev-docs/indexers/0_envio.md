# Envio

[Envio](https://envio.dev/?utm_source=boba&utm_medium=partner-docs) is a high-performance indexing framework that turns smart contract events into a queryable GraphQL API, with managed hosting on Envio Cloud.

[HyperIndex](https://docs.envio.dev/docs/HyperIndex/overview?utm_source=boba&utm_medium=partner-docs), Envio's indexing framework, natively supports indexing any EVM chain and fully supports Boba. Where a network is on HyperSync, as Boba is, HyperIndex uses it as the default data source for dramatically faster syncs, up to 2000x faster than traditional RPC.

## Why Envio?

- Real-time and historical Boba data through a single GraphQL API
- Blazing-fast backfills powered by HyperSync
- Auto-generate an indexer from any verified contract with `pnpx envio init`
- TypeScript, JavaScript, or ReScript event handlers, with reorg support
- Managed hosting on Envio Cloud, or self-host your own deployment

## Boba Network Details

Boba is supported on HyperSync and HyperRPC.

| Field | Value |
| --- | --- |
| Chain ID | 288 |
| HyperSync | `https://boba.hypersync.xyz` or `https://288.hypersync.xyz` |
| HyperRPC | `https://boba.rpc.hypersync.xyz` or `https://288.rpc.hypersync.xyz` |

## Getting Started

1. Initialise your indexer and import your Boba contract with `pnpx envio init`. See the [Quickstart](https://docs.envio.dev/docs/HyperIndex/quickstart?utm_source=boba&utm_medium=partner-docs).
2. Add Boba to your `config.yaml` using chain ID `288`:

```yaml
name: IndexerName # Specify indexer name
chains:
  - id: 288 # Boba
    start_block: START_BLOCK_NUMBER # Specify the starting block
    contracts:
      - name: ContractName
        address:
          - "0xYourContractAddress"
        events:
          - event: Event # Specify event
```

3. Run your indexer locally with Docker, or deploy it with a git-based workflow on [Envio Cloud](https://docs.envio.dev/docs/HyperIndex/hosted-service-deployment?utm_source=boba&utm_medium=partner-docs).
4. Query your indexed Boba data through the GraphQL API.

Need help setting up your indexer? Reach out on [Discord](https://discord.gg/envio), where the team is always happy to help.

[Website](https://envio.dev/?utm_source=boba&utm_medium=partner-docs) | [Docs](https://docs.envio.dev/?utm_source=boba&utm_medium=partner-docs) | [Supported Networks](https://docs.envio.dev/docs/HyperIndex/supported-networks?utm_source=boba&utm_medium=partner-docs) | [Discord](https://discord.gg/envio) | [X](https://twitter.com/envio_indexer) | [GitHub](https://github.com/enviodev)
