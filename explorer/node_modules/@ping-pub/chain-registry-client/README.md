# `@ping-pub/chain-registry-client`

Chain Registry Client is a client which fetch Cosmos chain registry from https://registry.ping.pub.

## Usage

```typescript
import ChainRegistryClient from '@ping-pub/chain-registry-client';

const client = new ChainRegistryClient()

client.fetchChainNames().then(x => {
    expect(x.length).toBeGreaterThan(0);
})

client.fetchChainInfo('cosmoshub').then(x => {
    expect(x.chain_name).toBe('cosmoshub');
})

client.fetchIBCPaths().then(x => {
    expect(x.length).toBeGreaterThan(0);
})

client.fetchIBCPathInfo('cosmoshub-osmosis.json').then(x => {
    expect(x.chain_1.chain_name).toBe('cosmoshub');
    expect(x.chain_2.chain_name).toBe('osmosis');
    done()
})

```
