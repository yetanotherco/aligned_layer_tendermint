import { IBCInfo } from '@chain-registry/types'

export * from '@chain-registry/types'

export interface Entry {
    name: string,
    type: string,
    mtime: string,
    size?: number,
}

export interface IBCPath {
    path: string,
    from?: string,
    to?: string,
}

export type IBCPathInfo = IBCInfo