import ChainRegistryClient from '../src'
import { Entry } from '../src/types'

it('Test Fetch Chain Names', () => {
    const client = new ChainRegistryClient()
    client.fetchChainNames().then(x => {
        expect(x.length).toBeGreaterThan(0);
    })
})

it('Test Fetch Chain Info', () => {
    const client = new ChainRegistryClient()
    client.fetchChainInfo('cosmoshub').then(x => {
        expect(x.chain_name).toBe('cosmoshub');
    })
})

it('Test IBC Paths', () => {
    const client = new ChainRegistryClient()
    client.fetchIBCPaths().then(x => {
        expect(x.length).toBeGreaterThan(0);
    })
})

it('Test IBC Path Info', (done) => {
    const client = new ChainRegistryClient()
    client.fetchIBCPathInfo('cosmoshub-osmosis.json').then(x => {
        expect(x.chain_1.chain_name).toBe('cosmoshub');
        expect(x.chain_2.chain_name).toBe('osmosis');
        done()
    })
})


