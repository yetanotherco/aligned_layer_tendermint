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
exports.decorateLanguageService = exports.createLanguageService = exports.getExternalFiles = exports.decorateLanguageServiceHost = exports.getProgram = void 0;
const typescript_1 = require("@volar/typescript");
const vue = __importStar(require("@vue/language-core"));
var typescript_2 = require("@volar/typescript");
Object.defineProperty(exports, "getProgram", { enumerable: true, get: function () { return typescript_2.getProgram; } });
Object.defineProperty(exports, "decorateLanguageServiceHost", { enumerable: true, get: function () { return typescript_2.decorateLanguageServiceHost; } });
Object.defineProperty(exports, "getExternalFiles", { enumerable: true, get: function () { return typescript_2.getExternalFiles; } });
function createLanguageService(host, vueCompilerOptions, ts, sys) {
    const languageContext = vue.createLanguageContext(host, vue.createLanguages(host.getCompilationSettings(), vueCompilerOptions, ts));
    const languageServiceHost = (0, typescript_1.createLanguageServiceHost)(languageContext, ts, sys, undefined);
    const languageService = ts.createLanguageService(languageServiceHost, (0, typescript_1.getDocumentRegistry)(ts, sys.useCaseSensitiveFileNames, host.workspacePath));
    decorateLanguageService(languageContext.virtualFiles, languageService, false);
    return {
        ...languageService,
        __internal__: {
            context: languageContext,
        },
    };
}
exports.createLanguageService = createLanguageService;
function decorateLanguageService(virtualFiles, ls, isTsPlugin) {
    (0, typescript_1.decorateLanguageService)(virtualFiles, ls, isTsPlugin);
    const getCompletionsAtPosition = ls.getCompletionsAtPosition.bind(ls);
    ls.getCompletionsAtPosition = (fileName, position, options) => {
        const result = getCompletionsAtPosition(fileName, position, options);
        if (result) {
            result.entries = result.entries.filter(entry => entry.name.indexOf('__VLS_') === -1);
        }
        return result;
    };
}
exports.decorateLanguageService = decorateLanguageService;
//# sourceMappingURL=index.js.map