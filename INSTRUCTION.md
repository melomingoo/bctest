# Coin Exchange Rate Service

Please create a service that can inquire the exchange rate of each token's currency and convert it.

## Description

https://coinmarketcap.com/api/

The above website provides real-time/past exchange rate of cryptocurrencies.

Use this site to provide real-time or historical information about exchange rate and conversion prices of currencies listed on the website:

## Requirements

- Code management: please use this repository
- Please provide `README.md` file to describe `how-to-use`
- Program language: please use Python, Javascript or Golang
- Store exchange rate and information in a database per hour.
- Use the ticker to search and/or convert an exchange rate.
- No restriction for a platform (CLI, Web, ...)
- Use a free API key. You do not have to purchase a paid one.

## Interface Examples

### Search latest data
- [INPUT]
  - base currency: KRW
  - target currency: BTC, ETH
- [OUTPUT]
  - data at: the current or the latest timestamp
  - BTC currency at now or the latest timestamp
  - ETH currency at now or the latest timestamp

### Search history data
- [INPUT]
  - base currency: USD
  - target currency: BTC, USDT(테더)
  - starting at: 2021-03-05 10:11:12
  - ending at: 2021-03-05 11:22:33
- [OUTPUT]
  - BTC currencies between <starting_at> and <ending_at>
  - USDT currencies between <starting_at> and <ending_at>

* Fetch historical data from your own DB as the free version of API doesn't support historical price data.

### Convert price
- [INPUT]
  - from: XRP
  - to: ATOM
  - at: now
- [OUTPUT]
  - show how much ATOM for 1 XRP

## Advantage Points
- Supports REST APIs
- Supports HMAC
- Applies or develops reusable caching APIs which could be open in public as a opensource
