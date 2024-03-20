"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.loadModule = exports.getVueI18nVersion = exports.checkVueI18nBridgeInstallPackage = exports.checkInstallPackage = void 0;
// eslint-disable-next-line @typescript-eslint/ban-types
function checkInstallPackage(pkg, debug) {
    let installedVueI18n = false;
    try {
        debug(`vue-i18n load path: ${require.resolve('vue-i18n')}`);
        installedVueI18n = true;
    }
    catch (e) {
        debug(`cannot find 'vue-i18n'`, e);
    }
    let installedPetiteVueI18n = false;
    try {
        debug(`petite-vue-i18n load path: ${require.resolve('petite-vue-i18n')}`);
        installedPetiteVueI18n = true;
    }
    catch (e) {
        debug(`cannot find 'petite-vue-i18n'`, e);
    }
    if (installedVueI18n) {
        return 'vue-i18n';
    }
    if (installedPetiteVueI18n) {
        return 'petite-vue-i18n';
    }
    throw new Error(`${pkg} requires 'vue-i18n' or 'petite-vue-i18n' to be present in the dependency tree.`);
}
exports.checkInstallPackage = checkInstallPackage;
// eslint-disable-next-line @typescript-eslint/ban-types
function checkVueI18nBridgeInstallPackage(debug) {
    let ret = false;
    try {
        debug(`vue-i18n-bridge load path: ${require.resolve('vue-i18n-bridge')}`);
        ret = true;
    }
    catch (e) {
        debug(`cannot find 'vue-i18n-bridge'`, e);
    }
    return ret;
}
exports.checkVueI18nBridgeInstallPackage = checkVueI18nBridgeInstallPackage;
function getVueI18nVersion(debug) {
    const VueI18n = loadModule('vue-i18n', debug);
    if (VueI18n == null) {
        return '';
    }
    if (VueI18n.version && VueI18n.version.startsWith('8.')) {
        return '8';
    }
    if (VueI18n.VERSION && VueI18n.VERSION.startsWith('9.')) {
        return '9';
    }
    return 'unknown';
}
exports.getVueI18nVersion = getVueI18nVersion;
// eslint-disable-next-line @typescript-eslint/no-explicit-any
function loadModule(moduleName, debug) {
    try {
        return require(moduleName);
    }
    catch (e) {
        debug(`cannot load '${moduleName}'`, e);
        return null;
    }
}
exports.loadModule = loadModule;
