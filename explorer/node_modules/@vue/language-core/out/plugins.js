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
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.getDefaultVueLanguagePlugins = void 0;
const file_html_1 = __importDefault(require("./plugins/file-html"));
const file_md_1 = __importDefault(require("./plugins/file-md"));
const file_vue_1 = __importDefault(require("./plugins/file-vue"));
const vue_sfc_customblocks_1 = __importDefault(require("./plugins/vue-sfc-customblocks"));
const vue_sfc_scripts_1 = __importDefault(require("./plugins/vue-sfc-scripts"));
const vue_sfc_styles_1 = __importDefault(require("./plugins/vue-sfc-styles"));
const vue_sfc_template_1 = __importDefault(require("./plugins/vue-sfc-template"));
const vue_template_html_1 = __importDefault(require("./plugins/vue-template-html"));
const vue_tsx_1 = __importDefault(require("./plugins/vue-tsx"));
const CompilerDOM = __importStar(require("@vue/compiler-dom"));
const CompilerVue2 = __importStar(require("./utils/vue2TemplateCompiler"));
function getDefaultVueLanguagePlugins(ts, compilerOptions, vueCompilerOptions, codegenStack) {
    const plugins = [
        file_md_1.default,
        file_html_1.default,
        file_vue_1.default,
        vue_template_html_1.default,
        vue_sfc_styles_1.default,
        vue_sfc_customblocks_1.default,
        vue_sfc_scripts_1.default,
        vue_sfc_template_1.default,
        vue_tsx_1.default,
        ...vueCompilerOptions.plugins,
    ];
    const pluginCtx = {
        modules: {
            '@vue/compiler-dom': vueCompilerOptions.target < 3
                ? {
                    ...CompilerDOM,
                    compile: CompilerVue2.compile,
                }
                : CompilerDOM,
            typescript: ts,
        },
        compilerOptions,
        vueCompilerOptions,
        codegenStack,
    };
    const pluginInstances = plugins
        .map(plugin => plugin(pluginCtx))
        .sort((a, b) => {
        const aOrder = a.order ?? 0;
        const bOrder = b.order ?? 0;
        return aOrder - bOrder;
    });
    return pluginInstances.filter((plugin) => {
        const valid = plugin.version >= 1 && plugin.version < 2;
        if (!valid) {
            console.warn(`Plugin ${JSON.stringify(plugin.name)} API version incompatible, expected 1.x but got ${JSON.stringify(plugin.version)}`);
        }
        return valid;
    });
}
exports.getDefaultVueLanguagePlugins = getDefaultVueLanguagePlugins;
//# sourceMappingURL=plugins.js.map