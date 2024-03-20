"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.parse = void 0;
const compiler = __importStar(require("@vue/compiler-dom"));
function parse(source) {
    const errors = [];
    const ast = compiler.parse(source, {
        // there are no components at SFC parsing level
        isNativeTag: () => true,
        // preserve all whitespaces
        isPreTag: () => true,
        getTextMode: ({ tag, props }, parent) => {
            if ((!parent && tag !== 'template')
                || (tag === 'template'
                    && props.some(p => p.type === 6 /* compiler.NodeTypes.ATTRIBUTE */ &&
                        p.name === 'lang' &&
                        p.value &&
                        p.value.content &&
                        p.value.content !== 'html'))) {
                return 2 /* compiler.TextModes.RAWTEXT */;
            }
            else {
                return 0 /* compiler.TextModes.DATA */;
            }
        },
        onError: e => {
            errors.push(e);
        },
        comments: true,
    });
    const descriptor = {
        filename: 'anonymous.vue',
        source,
        template: null,
        script: null,
        scriptSetup: null,
        styles: [],
        customBlocks: [],
        cssVars: [],
        slotted: false,
        shouldForceReload: () => false,
    };
    ast.children.forEach(node => {
        if (node.type !== 1 /* compiler.NodeTypes.ELEMENT */) {
            return;
        }
        switch (node.tag) {
            case 'template':
                const templateBlock = (descriptor.template = createBlock(node, source));
                templateBlock.ast = node;
                break;
            case 'script':
                const scriptBlock = createBlock(node, source);
                const isSetup = !!scriptBlock.attrs.setup;
                if (isSetup && !descriptor.scriptSetup) {
                    descriptor.scriptSetup = scriptBlock;
                    break;
                }
                if (!isSetup && !descriptor.script) {
                    descriptor.script = scriptBlock;
                    break;
                }
                break;
            case 'style':
                const styleBlock = createBlock(node, source);
                descriptor.styles.push(styleBlock);
                break;
            default:
                descriptor.customBlocks.push(createBlock(node, source));
                break;
        }
    });
    return {
        descriptor,
        errors,
    };
}
exports.parse = parse;
function createBlock(node, source) {
    const type = node.tag;
    let { start, end } = node.loc;
    let content = '';
    if (node.children.length) {
        start = node.children[0].loc.start;
        end = node.children[node.children.length - 1].loc.end;
        content = source.slice(start.offset, end.offset);
    }
    else {
        const offset = node.loc.source.indexOf(`</`);
        if (offset > -1) {
            start = {
                line: start.line,
                column: start.column + offset,
                offset: start.offset + offset
            };
        }
        end = Object.assign({}, start);
    }
    const loc = {
        source: content,
        start,
        end
    };
    const attrs = {};
    const block = {
        type,
        content,
        loc,
        attrs
    };
    node.props.forEach(p => {
        if (p.type === 6 /* compiler.NodeTypes.ATTRIBUTE */) {
            attrs[p.name] = p.value ? p.value.content || true : true;
            if (p.name === 'lang') {
                block.lang = p.value && p.value.content;
            }
            else if (p.name === 'src') {
                block.src = p.value && p.value.content;
            }
            else if (type === 'style') {
                if (p.name === 'scoped') {
                    block.scoped = true;
                }
                else if (p.name === 'module') {
                    block.module = attrs[p.name];
                }
            }
            else if (type === 'script' && p.name === 'setup') {
                block.setup = attrs.setup;
            }
        }
    });
    return block;
}
//# sourceMappingURL=parseSfc.js.map