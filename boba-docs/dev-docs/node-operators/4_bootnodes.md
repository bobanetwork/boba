# Bootnodes

Synchronizing your node occurs in two ways.

Firstly, the consensus client (op-node) pulls blocks from the L1 and decodes the data committed there to construct the finalized version of the blockchain.  However, the sequencer only periodically commits the data to the L1 (resulting in cost savings for all users) so the L1 version of the blockchain is always some period of time (usually no more than an hour) behind the sequencer.

Secondly, op-node pulls blocks from P2P which the sequencer has signed, indicating that the sequencer intends to eventually commit those block to the L1.  This P2P process allows for replicas in the network to near-instantly reflect the current state of the sequencer.  Since the block sharing process is P2P, your node needs to discover other nodes in the P2P network.  There are special nodes for facilitating this initial discovery called bootnodes that should be specified in your node config.

Both the execution client (op-reth) and the consensus client (op-node) run their own P2P discovery networks, so each has its own set of bootnodes.

## Consensus layer (op-node)

Consensus-layer bootnodes can be specified to op-node with either the `--p2p.bootnodes` flag, or via the `OP_NODE_P2P_BOOTNODES` environment variable.

### Boba Mainnet

```
enode://a38db98391708094dd679aacdd356efdbe10ad7c9ba4b87d8f4b79eb24ac26328e7fdf911f8c5c3b7786c2505bc77d896a0ecd1816bd56107592f728dcff3945@35.153.183.193:0?discport=30301,enode://a92b84bac2893ef868659364a6784a6eeb146cb04ec0bb3b1e9925f5bb895af9e1c50deaf0703ab26db366a3d82272e710e06d925cb0b7245129ed76d54ccac9@13.221.254.11:0?discport=30301
```

### Boba Sepolia Testnet

```
enode://b3d3f7d947461138e850b5fa0c417b8c1c498d3d7edb17f662b2e2c99f096b756be238b07002e98a0a373ae23ff87054b68a9c6bb0ca55fb852d9969debfa6cd@52.201.174.220:0?discport=30301,enode://b75a091361d9ed31e2eac8e64b06ae26708f828042dde7dbed21b74869ecfead030ee25570c758e54ae7462d22f61afabec75e24d48a494a990b25bde009d5c5@3.230.114.57:0?discport=30301
```

## Execution layer (op-reth)

Execution-layer bootnodes let op-reth discover peers on its own P2P network. This is required for [execution-layer sync](https://docs.optimism.io/operators/node-operators/management/snap-sync), where op-node tells the execution client to sync toward the tip of the chain (op-node's `--syncmode=execution-layer`) rather than importing every block itself. Execution-layer bootnodes can be specified to op-reth with the `--bootnodes` flag.

### Boba Sepolia Testnet

```
enode://ac9f697ab2692b0acb5a4624336d14def4c981796e5d351c654135754c646c6d1a31cbbdb06f5d04156489bd10aa1fa8416eec75839a50020064f5a9f530bb3a@54.242.254.87:0?discport=30301,enode://d6adc6552cccd68b5216f5279a6c4f6766ef0ad1b5056d73b5d846260c04c039537a731479d6854d9d275f8e6efe7346d6a5ecc11259ecc655d12f040415eaa8@18.211.99.160:0?discport=30301
```
