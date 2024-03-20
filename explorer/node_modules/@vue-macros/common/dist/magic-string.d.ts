import MagicStringBase, { OverwriteOptions } from 'magic-string';
import { Node } from '@babel/types';

declare class MagicString extends MagicStringBase {
    removeNode(node: Node, { offset }?: {
        offset?: number;
    }): this;
    moveNode(node: Node, index: number, { offset }?: {
        offset?: number;
    }): this;
    sliceNode(node: Node, { offset }?: {
        offset?: number;
    }): string;
    sliceNodes(nodes: Node[], { offset }?: {
        offset?: number;
    }): string;
    overwriteNode(node: Node, content: string | Node, { offset, ...options }?: OverwriteOptions & {
        offset?: number;
    }): this;
}

export { MagicString };
