# Bootnodes

Synchronizing your node occurs in two ways.

Firstly, the consensus client (op-node) pulls blocks from the L1 and decodes the data committed there to construct the finalized version of the blockchain.  However, the sequencer only periodically commits the data to the L1 (resulting in cost savings for all users) so the L1 version of the blockchain is always some period of time (usually no more than an hour) behind the sequencer.

Secondly, op-node pulls blocks from P2P which the sequencer has signed, indicating that the sequencer intends to eventually commit those block to the L1.  This P2P process allows for replicas in the network to near-instantly reflect the current state of the sequencer.  Since the block sharing process is P2P, your node needs to discover other nodes in the P2P network.  There are special nodes for facilitating this initial discovery called bootnodes that should be specified in your node config.

Both the execution client (op-reth) and the consensus client (op-node) run their own P2P discovery networks, so each has its own set of bootnodes.

## Consensus layer (op-node)

Consensus-layer bootnodes can be specified to op-node with either the `--p2p.bootnodes` flag, or via the `OP_NODE_P2P_BOOTNODES` environment variable.

### Boba Mainnet

```
enode://ef3def930e1c9be05b2c2b6c0de35555e1bc283feb6edf1f0f4d7385f82a011ae88c14c4691555d6f5f3887c3e8e69d2aee4ccc39b96dc045c9e25181eaad0ec@100.58.145.205:0?discport=30301,enode://cae574d3aada643b6dea04d3f3fa737f42c38631ffee609b63aada6558c19fdf30bfcba86c877678dbc70828bdff281bd8cc5109a69134d630782fb8d93aded1@3.217.241.101:0?discport=30301
```

### Boba Sepolia Testnet

```
enode://b3d3f7d947461138e850b5fa0c417b8c1c498d3d7edb17f662b2e2c99f096b756be238b07002e98a0a373ae23ff87054b68a9c6bb0ca55fb852d9969debfa6cd@52.201.174.220:0?discport=30301,enode://b75a091361d9ed31e2eac8e64b06ae26708f828042dde7dbed21b74869ecfead030ee25570c758e54ae7462d22f61afabec75e24d48a494a990b25bde009d5c5@3.230.114.57:0?discport=30301
```

## Execution layer (op-reth)

Execution-layer bootnodes let op-reth discover peers on its own P2P network. This is required for [execution-layer sync](https://docs.optimism.io/operators/node-operators/management/snap-sync), where op-node tells the execution client to sync toward the tip of the chain (op-node's `--syncmode=execution-layer`) rather than importing every block itself. Execution-layer bootnodes can be specified to op-reth with the `--bootnodes` flag.

### Boba Mainnet

```
enode://ea3452bc663cde356b8b4a6e8c4a4953d4fd10fb8589d5e2955569853121700e8f3c157634253b1238496e923108206fddff2dc55f7ba280a49ddf7e782a09d4@3.214.7.247:0?discport=30301,enode://5f34742bd61a953fde88dfae2d6e0ddded36babe72211019c5988410e1041870bc6542045fa397a280e101ca26b4721e1738a8f9618dedcb45f027b0c066f806@100.50.243.238:0?discport=30301
```

### Boba Sepolia Testnet

```
enode://ac9f697ab2692b0acb5a4624336d14def4c981796e5d351c654135754c646c6d1a31cbbdb06f5d04156489bd10aa1fa8416eec75839a50020064f5a9f530bb3a@54.242.254.87:0?discport=30301,enode://d6adc6552cccd68b5216f5279a6c4f6766ef0ad1b5056d73b5d846260c04c039537a731479d6854d9d275f8e6efe7346d6a5ecc11259ecc655d12f040415eaa8@18.211.99.160:0?discport=30301
```
