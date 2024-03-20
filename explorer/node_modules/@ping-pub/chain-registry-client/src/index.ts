import fetch from "cross-fetch"
import {Chain, Entry, AssetList, IBCPath, IBCPathInfo} from './types'

export default class ChainRegistryClient{

    readonly endpoint: string

    constructor(endpoint = 'https://registry.ping.pub') {
        this.endpoint = endpoint
    }

    async getAny(url : string): Promise<any> {
        return fetch(`${this.endpoint}${url}`).then((res) => res.json())
    }
    async get<T>(url : string): Promise<T> {
        return fetch(`${this.endpoint}${url}`).then((res) => res.json())
    }
    async fetchChainNames(): Promise<string[]> {
        const entris = await this.get<Entry[]>('/')
        return entris.filter( i => i.type === 'directory' && i.name !== 'testnet' && !i.name.startsWith('_')).map(x => x.name)
    }
    async fetchChainInfo(chainName: string) {
        return this.get<Chain>(`/${chainName}/chain.json`)
    }
    async fetchAssetsList(chainName: string) {
        return this.get<AssetList>(`/${chainName}/assetlist.json`)
    }
    async fetchIBCPaths() {
        const entries = await this.get<Entry[]>('/_IBC/')
        const re = /([\w]+)-([\w]+)\.json/;
        return entries.map(x => {
            const matches = x.name.match(re)
            const bridge = {} as IBCPath
            bridge.path = x.name
            bridge.from = matches[1]
            bridge.to = matches[2]
            return bridge
        })
    }
    async fetchIBCPathInfo(path: string) {
        const info = await this.get<IBCPathInfo>(`/_IBC/${path}`)
        return info
    }

}