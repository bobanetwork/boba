---
title: DIA Oracles
sidebar_label: DIA
description: Read DIA asset prices on Boba Network using asset-specific oracle adapters.
---

# DIA Oracles

[DIA](https://www.diadata.org/) provides asset price feeds for applications on Boba Network. Use an asset-specific adapter to read the latest USD price and its update timestamp from your smart contract.

## Oracle contract addresses

The following oracle contracts are listed in [DIA's Boba Network guide](https://www.diadata.org/docs/guides/chain-specific-guide/boba-network).

| Network | Oracle contract |
| --- | --- |
| Boba Ethereum Mainnet | [`0x5612599CF48032d7428399d5Fcb99eDcc75c06A7`](https://bobascan.com/address/0x5612599CF48032d7428399d5Fcb99eDcc75c06A7) |
| Boba Ethereum Sepolia | [`0xb4C652a1fF2022D98Ad9406FaB42B242313B7d1a`](https://testnet.bobascan.com/address/0xb4C652a1fF2022D98Ad9406FaB42B242313B7d1a) |

For the `latestRoundData()` integration below, call the asset's **adapter address** from the next table.

:::note Testnet feeds
The Sepolia feed is for testing. Updates can be less frequent and less consistent than mainnet. The asset adapters below are for **Boba Ethereum Mainnet**; obtain the appropriate testnet adapter details from DIA before using this example on Sepolia.
:::

## Mainnet asset adapters

Each adapter exposes `latestRoundData()` and `decimals()`. The listed feeds are denominated in USD. Read `decimals()` from the adapter when interpreting a price.

| Feed | Adapter on Boba Ethereum Mainnet |
| --- | --- |
| BTC/USD | [`0xcdf63e59d432De724c8afDc8064A81aA03cEeAFd`](https://bobascan.com/address/0xcdf63e59d432De724c8afDc8064A81aA03cEeAFd) |
| ETH/USD | [`0x78B31334F73C50363433ffC551c7CbB2556C5eCD`](https://bobascan.com/address/0x78B31334F73C50363433ffC551c7CbB2556C5eCD) |
| USDC/USD | [`0x61187Bc50429e7B509b31924aB26c923FbE3102A`](https://bobascan.com/address/0x61187Bc50429e7B509b31924aB26c923FbE3102A) |
| USDT/USD | [`0xBAEaA74999819176364bB7005dCf079545c1D46F`](https://bobascan.com/address/0xBAEaA74999819176364bB7005dCf079545c1D46F) |
| DAI/USD | [`0xf9b5cCb73d3dEf96EfEd61aB4b2D587ecEb6E3aA`](https://bobascan.com/address/0xf9b5cCb73d3dEf96EfEd61aB4b2D587ecEb6E3aA) |
| BOBA/USD | [`0xB2545be654E053cffa5501C08258eaeaBa94170f`](https://bobascan.com/address/0xB2545be654E053cffa5501C08258eaeaBa94170f) |

## Read a price in Solidity

This example reads the ETH/USD adapter on Boba Ethereum Mainnet. Set `maxPriceAgeSeconds` during deployment to the maximum age your application accepts, in seconds. Choose that limit based on the feed's update configuration and your application's requirements.

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

interface IDIAAggregatorV3 {
    function decimals() external view returns (uint8);

    function latestRoundData() external view returns (
        uint80 roundId,
        int256 answer,
        uint256 startedAt,
        uint256 updatedAt,
        uint80 answeredInRound
    );
}

contract DIAETHPriceConsumer {
    // ETH/USD adapter on Boba Ethereum Mainnet.
    IDIAAggregatorV3 public constant ETH_FEED =
        IDIAAggregatorV3(0x78B31334F73C50363433ffC551c7CbB2556C5eCD);

    uint256 public immutable maxPriceAge;

    constructor(uint256 maxPriceAgeSeconds) {
        require(maxPriceAgeSeconds > 0, "Invalid maximum age");
        maxPriceAge = maxPriceAgeSeconds;
    }

    function latestPrice() external view returns (
        int256 price,
        uint8 priceDecimals,
        uint256 updatedAt
    ) {
        (, price, , updatedAt, ) = ETH_FEED.latestRoundData();

        require(price > 0, "Invalid price");
        require(
            updatedAt > 0 && updatedAt <= block.timestamp,
            "Invalid timestamp"
        );
        require(block.timestamp - updatedAt <= maxPriceAge, "Stale price");

        priceDecimals = ETH_FEED.decimals();
    }
}
```

The returned `price` is an integer scaled by `10 ** priceDecimals`, and `updatedAt` is a Unix timestamp in seconds. For example, an answer of `250000000000` with 8 decimals represents **2,500 USD**. Preserve the scale when using the price in calculations.

The example rejects non-positive prices, invalid timestamps, and prices older than the configured maximum age. Decide how your application should handle an unavailable or stale feed before using it in production.

## Feed configuration

DIA documents the following configuration for Boba:

| Setting | Value |
| --- | --- |
| Pricing methodology | [VWAPIR](https://www.diadata.org/docs/guides/methodologies/pricing-methodologies/vwapir-volume-weighted-average-price-with-interquartile-range-filter) |
| Price deviation threshold | 0.2% |
| Heartbeat | 1 hour |

These settings describe the intended update triggers. Applications should still check the returned timestamp; the heartbeat is not a guarantee that a fresh value is always available.

## Resources

- [DIA's Boba Network guide](https://www.diadata.org/docs/guides/chain-specific-guide/boba-network)
- [DIA integration guide](https://www.diadata.org/docs/nexus/how-to-guides/migrate-to-dia)
- [DIA support on Discord](https://discord.gg/ZvGjVY5uvs)
- [DIA support on Telegram](https://t.me/diadata_org)
