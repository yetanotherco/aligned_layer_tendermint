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
var __exportStar = (this && this.__exportStar) || function(m, exports) {
    for (var p in m) if (p !== "default" && !Object.prototype.hasOwnProperty.call(exports, p)) __createBinding(exports, m, p);
};
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.tsCodegen = exports.sharedTypes = exports.scriptRanges = void 0;
__exportStar(require("./generators/template"), exports);
__exportStar(require("./languageModule"), exports);
__exportStar(require("./parsers/scriptSetupRanges"), exports);
__exportStar(require("./plugins"), exports);
__exportStar(require("./sourceFile"), exports);
__exportStar(require("./types"), exports);
__exportStar(require("./utils/ts"), exports);
__exportStar(require("./utils/parseSfc"), exports);
exports.scriptRanges = __importStar(require("./parsers/scriptRanges"));
exports.sharedTypes = __importStar(require("./utils/globalTypes"));
__exportStar(require("./utils/shared"), exports);
var vue_tsx_1 = require("./plugins/vue-tsx");
Object.defineProperty(exports, "tsCodegen", { enumerable: true, get: function () { return vue_tsx_1.tsCodegen; } });
__exportStar(require("@volar/language-core"), exports);
__exportStar(require("@volar/source-map"), exports);
//# sourceMappingURL=index.js.map