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
exports.createProgram = void 0;
const vue = __importStar(require("@vue/language-core"));
const vueTs = __importStar(require("@vue/typescript"));
const shared_1 = require("./shared");
const windowsPathReg = /\\/g;
function createProgram(options) {
    if (!options.options.noEmit && !options.options.emitDeclarationOnly)
        throw toThrow('js emit is not supported');
    if (!options.options.noEmit && options.options.noEmitOnError)
        throw toThrow('noEmitOnError is not supported');
    if (options.options.extendedDiagnostics || options.options.generateTrace)
        throw toThrow('--extendedDiagnostics / --generateTrace is not supported, please run `Write Virtual Files` in VSCode to write virtual files and use `--extendedDiagnostics` / `--generateTrace` via tsc instead of vue-tsc to debug.');
    if (!options.host)
        throw toThrow('!options.host');
    const ts = require('typescript');
    let program = options.oldProgram;
    if (shared_1.state.hook) {
        program = shared_1.state.hook.program;
        program.__vue.options = options;
    }
    else if (!program) {
        const ctx = {
            projectVersion: 0,
            options,
            get languageHost() {
                return languageHost;
            },
            get vueCompilerOptions() {
                return vueCompilerOptions;
            },
            get languageService() {
                return vueTsLs;
            },
        };
        const vueCompilerOptions = getVueCompilerOptions();
        const scripts = new Map();
        const languageHost = {
            workspacePath: ctx.options.host.getCurrentDirectory().replace(windowsPathReg, '/'),
            rootPath: ctx.options.host.getCurrentDirectory().replace(windowsPathReg, '/'),
            getCompilationSettings: () => ctx.options.options,
            getScriptFileNames: () => {
                return ctx.options.rootNames;
            },
            getScriptSnapshot,
            getProjectVersion: () => {
                return ctx.projectVersion.toString();
            },
            getProjectReferences: () => ctx.options.projectReferences,
            getCancellationToken: ctx.options.host.getCancellationToken ? () => ctx.options.host.getCancellationToken() : undefined,
        };
        const vueTsLs = vueTs.createLanguageService(languageHost, vueCompilerOptions, ts, ts.sys);
        program = vueTs.getProgram(ts, vueTsLs.__internal__.context, vueTsLs, ts.sys);
        program.__vue = ctx;
        function getVueCompilerOptions() {
            const tsConfig = ctx.options.options.configFilePath;
            if (typeof tsConfig === 'string') {
                return vue.createParsedCommandLine(ts, ts.sys, tsConfig).vueOptions;
            }
            return {};
        }
        function getScriptSnapshot(fileName) {
            return getScript(fileName)?.scriptSnapshot;
        }
        function getScript(fileName) {
            const script = scripts.get(fileName);
            if (script?.projectVersion === ctx.projectVersion) {
                return script;
            }
            const modifiedTime = ts.sys.getModifiedTime?.(fileName)?.valueOf() ?? 0;
            if (script?.modifiedTime === modifiedTime) {
                return script;
            }
            if (ctx.options.host.fileExists(fileName)) {
                const fileContent = ctx.options.host.readFile(fileName);
                if (fileContent !== undefined) {
                    const script = {
                        projectVersion: ctx.projectVersion,
                        modifiedTime,
                        scriptSnapshot: ts.ScriptSnapshot.fromString(fileContent),
                        version: ctx.options.host.createHash?.(fileContent) ?? fileContent,
                    };
                    scripts.set(fileName, script);
                    return script;
                }
            }
        }
    }
    else {
        const ctx = program.__vue;
        ctx.options = options;
        ctx.projectVersion++;
    }
    const vueCompilerOptions = program.__vue.vueCompilerOptions;
    if (vueCompilerOptions?.hooks) {
        const index = (shared_1.state.hook?.index ?? -1) + 1;
        if (index < vueCompilerOptions.hooks.length) {
            const hookPath = vueCompilerOptions.hooks[index];
            const hook = require(hookPath);
            shared_1.state.hook = {
                program,
                index,
                worker: (async () => await hook(program))(),
            };
            throw 'hook';
        }
    }
    for (const rootName of options.rootNames) {
        // register file watchers
        options.host.getSourceFile(rootName, ts.ScriptTarget.ESNext);
    }
    return program;
}
exports.createProgram = createProgram;
function toThrow(msg) {
    console.error(msg);
    return msg;
}
//# sourceMappingURL=index.js.map