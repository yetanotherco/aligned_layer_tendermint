# 🍵 name-matcha

**Name Matcha** (*Matcha* - 抹茶; Meaning - 'ground and powdered green tea'; Emoji - 🍵) - a pun on **name-matcher**.

A developer-friendly javascript library that provides a standardized way to easily resolve Cosmos ecosystem name services to wallet addresses (and vice-versa) using just one line of code. 

## Installation

```bash
npm install @leapwallet/name-matcha
```

## Usage

Here's how you can quickly get started with resolving names -

```js
import nameMatcha, { services } from '@leapwallet/name-matcha';

const address = await nameMatcha.resolve('alice.stars', service.stargazeNames);
```

To know more about the library, check out the [Docs](https://leapwallet.github.io/name-matcha/).

## Name Service Support

We support the following name services -

1. [ICNS](https://www.icns.xyz)
2. [Stargaze Names](https://www.stargaze.zone/names)
3. [IBC Domains](https://ibc.domains)
4. [ArchId](https://archid.app)

We also allow you to add your own name service for custom use-cases. You can look at the docs for the same [here](https://leapwallet.github.io/name-matcha/docs/advanced#custom-name-services).
