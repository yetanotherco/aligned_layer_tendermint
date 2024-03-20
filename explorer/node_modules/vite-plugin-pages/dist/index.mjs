var __defProp = Object.defineProperty;
var __defProps = Object.defineProperties;
var __getOwnPropDescs = Object.getOwnPropertyDescriptors;
var __getOwnPropSymbols = Object.getOwnPropertySymbols;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __propIsEnum = Object.prototype.propertyIsEnumerable;
var __defNormalProp = (obj, key, value) => key in obj ? __defProp(obj, key, { enumerable: true, configurable: true, writable: true, value }) : obj[key] = value;
var __spreadValues = (a, b) => {
  for (var prop in b || (b = {}))
    if (__hasOwnProp.call(b, prop))
      __defNormalProp(a, prop, b[prop]);
  if (__getOwnPropSymbols)
    for (var prop of __getOwnPropSymbols(b)) {
      if (__propIsEnum.call(b, prop))
        __defNormalProp(a, prop, b[prop]);
    }
  return a;
};
var __spreadProps = (a, b) => __defProps(a, __getOwnPropDescs(b));
var __publicField = (obj, key, value) => {
  __defNormalProp(obj, typeof key !== "symbol" ? key + "" : key, value);
  return value;
};
var __accessCheck = (obj, member, msg) => {
  if (!member.has(obj))
    throw TypeError("Cannot " + msg);
};
var __privateGet = (obj, member, getter) => {
  __accessCheck(obj, member, "read from private field");
  return getter ? getter.call(obj) : member.get(obj);
};
var __privateAdd = (obj, member, value) => {
  if (member.has(obj))
    throw TypeError("Cannot add the same private member more than once");
  member instanceof WeakSet ? member.add(obj) : member.set(obj, value);
};
var __privateSet = (obj, member, value, setter) => {
  __accessCheck(obj, member, "write to private field");
  setter ? setter.call(obj, value) : member.set(obj, value);
  return value;
};
var __privateWrapper = (obj, member, setter, getter) => ({
  set _(value) {
    __privateSet(obj, member, value, setter);
  },
  get _() {
    return __privateGet(obj, member, getter);
  }
});

// src/constants.ts
var MODULE_IDS = [
  "~pages",
  "~react-pages",
  "~solid-pages",
  "pages-generated",
  "virtual:generated-pages",
  "virtual:generated-pages-react"
];
var MODULE_ID_VIRTUAL = "/@vite-plugin-pages/generated-pages";
var ROUTE_BLOCK_ID_VIRTUAL = "/@vite-plugin-pages/route-block";
var ROUTE_IMPORT_NAME = "__pages_import_$1__";
var routeBlockQueryRE = /\?vue&type=route/;
var dynamicRouteRE = /^\[(.+)\]$/;
var cacheAllRouteRE = /^\[\.{3}(.*)\]$/;
var replaceDynamicRouteRE = /^\[(?:\.{3})?(.*)\]$/;
var nuxtDynamicRouteRE = /^_(.*)$/;
var nuxtCacheAllRouteRE = /^_$/;
var countSlashRE = /\//g;
var replaceIndexRE = /\/?index$/;

// src/context.ts
import { extname, join as join2, resolve as resolve3 } from "path";

// node_modules/.pnpm/@antfu+utils@0.7.2/node_modules/@antfu/utils/dist/index.mjs
function toArray(array) {
  array = array != null ? array : [];
  return Array.isArray(array) ? array : [array];
}
function slash(str) {
  return str.replace(/\\/g, "/");
}
var Node = class {
  constructor(value) {
    __publicField(this, "value");
    __publicField(this, "next");
    this.value = value;
  }
};
var _head, _tail, _size;
var Queue = class {
  constructor() {
    __privateAdd(this, _head, void 0);
    __privateAdd(this, _tail, void 0);
    __privateAdd(this, _size, void 0);
    this.clear();
  }
  enqueue(value) {
    const node = new Node(value);
    if (__privateGet(this, _head)) {
      __privateGet(this, _tail).next = node;
      __privateSet(this, _tail, node);
    } else {
      __privateSet(this, _head, node);
      __privateSet(this, _tail, node);
    }
    __privateWrapper(this, _size)._++;
  }
  dequeue() {
    const current = __privateGet(this, _head);
    if (!current) {
      return;
    }
    __privateSet(this, _head, __privateGet(this, _head).next);
    __privateWrapper(this, _size)._--;
    return current.value;
  }
  clear() {
    __privateSet(this, _head, void 0);
    __privateSet(this, _tail, void 0);
    __privateSet(this, _size, 0);
  }
  get size() {
    return __privateGet(this, _size);
  }
  *[Symbol.iterator]() {
    let current = __privateGet(this, _head);
    while (current) {
      yield current.value;
      current = current.next;
    }
  }
};
_head = new WeakMap();
_tail = new WeakMap();
_size = new WeakMap();
var VOID = Symbol("p-void");

// src/options.ts
import { resolve as resolve2 } from "path";

// src/files.ts
import { join } from "path";
import fg from "fast-glob";

// src/utils.ts
import { resolve, win32 } from "path";
import { URLSearchParams } from "url";
import Debug from "debug";
var debug = {
  hmr: Debug("vite-plugin-pages:hmr"),
  routeBlock: Debug("vite-plugin-pages:routeBlock"),
  options: Debug("vite-plugin-pages:options"),
  pages: Debug("vite-plugin-pages:pages"),
  search: Debug("vite-plugin-pages:search"),
  env: Debug("vite-plugin-pages:env"),
  cache: Debug("vite-plugin-pages:cache"),
  resolver: Debug("vite-plugin-pages:resolver")
};
function extsToGlob(extensions) {
  return extensions.length > 1 ? `{${extensions.join(",")}}` : extensions[0] || "";
}
function countSlash(value) {
  return (value.match(countSlashRE) || []).length;
}
function isPagesDir(path, options) {
  for (const page of options.dirs) {
    const dirPath = slash(resolve(options.root, page.dir));
    if (path.startsWith(dirPath))
      return true;
  }
  return false;
}
function isTarget(path, options) {
  return isPagesDir(path, options) && options.extensionsRE.test(path);
}
function isDynamicRoute(routePath, nuxtStyle = false) {
  return nuxtStyle ? nuxtDynamicRouteRE.test(routePath) : dynamicRouteRE.test(routePath);
}
function isCatchAllRoute(routePath, nuxtStyle = false) {
  return nuxtStyle ? nuxtCacheAllRouteRE.test(routePath) : cacheAllRouteRE.test(routePath);
}
function resolveImportMode(filepath, options) {
  const mode = options.importMode;
  if (typeof mode === "function")
    return mode(filepath, options);
  return mode;
}
function invalidatePagesModule(server) {
  const { moduleGraph } = server;
  const mods = moduleGraph.getModulesByFile(MODULE_ID_VIRTUAL);
  if (mods) {
    const seen = /* @__PURE__ */ new Set();
    mods.forEach((mod) => {
      moduleGraph.invalidateModule(mod, seen);
    });
  }
}
function normalizeCase(str, caseSensitive) {
  if (!caseSensitive)
    return str.toLocaleLowerCase();
  return str;
}
function normalizeName(name, isDynamic, nuxtStyle = false) {
  if (!isDynamic)
    return name;
  return nuxtStyle ? name.replace(nuxtDynamicRouteRE, "$1") || "all" : name.replace(replaceDynamicRouteRE, "$1");
}
function buildReactRoutePath(node, nuxtStyle = false) {
  const isDynamic = isDynamicRoute(node, nuxtStyle);
  const isCatchAll = isCatchAllRoute(node, nuxtStyle);
  const normalizedName = normalizeName(node, isDynamic, nuxtStyle);
  if (isDynamic) {
    if (isCatchAll)
      return "*";
    return `:${normalizedName}`;
  }
  return `${normalizedName}`;
}
function buildReactRemixRoutePath(node) {
  const escapeStart = "[";
  const escapeEnd = "]";
  let result = "";
  let rawSegmentBuffer = "";
  let inEscapeSequence = 0;
  let skipSegment = false;
  for (let i = 0; i < node.length; i++) {
    let isNewEscapeSequence2 = function() {
      return !inEscapeSequence && char === escapeStart && lastChar !== escapeStart;
    }, isCloseEscapeSequence2 = function() {
      return inEscapeSequence && char === escapeEnd && nextChar !== escapeEnd;
    }, isStartOfLayoutSegment2 = function() {
      return char === "_" && nextChar === "_" && !rawSegmentBuffer;
    };
    var isNewEscapeSequence = isNewEscapeSequence2, isCloseEscapeSequence = isCloseEscapeSequence2, isStartOfLayoutSegment = isStartOfLayoutSegment2;
    const char = node.charAt(i);
    const lastChar = i > 0 ? node.charAt(i - 1) : void 0;
    const nextChar = i < node.length - 1 ? node.charAt(i + 1) : void 0;
    if (skipSegment) {
      if (char === "/" || char === "." || char === win32.sep)
        skipSegment = false;
      continue;
    }
    if (isNewEscapeSequence2()) {
      inEscapeSequence++;
      continue;
    }
    if (isCloseEscapeSequence2()) {
      inEscapeSequence--;
      continue;
    }
    if (inEscapeSequence) {
      result += char;
      continue;
    }
    if (char === "/" || char === win32.sep || char === ".") {
      if (rawSegmentBuffer === "index" && result.endsWith("index"))
        result = result.replace(replaceIndexRE, "");
      else
        result += "/";
      rawSegmentBuffer = "";
      continue;
    }
    if (isStartOfLayoutSegment2()) {
      skipSegment = true;
      continue;
    }
    rawSegmentBuffer += char;
    if (char === "$") {
      result += typeof nextChar === "undefined" ? "*" : ":";
      continue;
    }
    result += char;
  }
  if (rawSegmentBuffer === "index" && result.endsWith("index"))
    result = result.replace(replaceIndexRE, "");
  return result || void 0;
}
function parsePageRequest(id) {
  const [moduleId, rawQuery] = id.split("?", 2);
  const query = new URLSearchParams(rawQuery);
  const pageId = query.get("id");
  return {
    moduleId,
    query,
    pageId
  };
}

// src/files.ts
function getPageDirs(PageOptions, root, exclude) {
  const dirs = fg.sync(slash(PageOptions.dir), {
    ignore: exclude,
    onlyDirectories: true,
    dot: true,
    unique: true,
    cwd: root
  });
  const pageDirs = dirs.map((dir) => __spreadProps(__spreadValues({}, PageOptions), {
    dir
  }));
  return pageDirs;
}
function getPageFiles(path, options) {
  const {
    exclude,
    extensions
  } = options;
  const ext = extsToGlob(extensions);
  const files = fg.sync(`**/*.${ext}`, {
    ignore: exclude,
    onlyFiles: true,
    cwd: path
  }).map((p) => slash(join(path, p)));
  return files;
}

// src/resolvers/vue.ts
import colors from "picocolors";
import deepEqual from "deep-equal";

// src/stringify.ts
var componentRE = /"(?:component|element)":("(.*?)")/g;
var hasFunctionRE = /"(?:props|beforeEnter)":("(.*?)")/g;
var multilineCommentsRE = /\/\*(.|[\r\n])*?\*\//gm;
var singlelineCommentsRE = /\/\/.*/g;
function replaceFunction(_, value) {
  if (value instanceof Function || typeof value === "function") {
    const fnBody = value.toString().replace(multilineCommentsRE, "").replace(singlelineCommentsRE, "").replace(/(\t|\n|\r|\s)/g, "");
    if (fnBody.length < 8 || fnBody.substring(0, 8) !== "function")
      return `_NuFrRa_${fnBody}`;
    return fnBody;
  }
  return value;
}
function stringifyRoutes(preparedRoutes, options) {
  const importsMap = /* @__PURE__ */ new Map();
  function getImportString(path, importName) {
    var _a, _b;
    const mode = resolveImportMode(path, options);
    return mode === "sync" ? `import ${importName} from "${path}"` : `const ${importName} = ${((_b = (_a = options.resolver.stringify) == null ? void 0 : _a.dynamicImport) == null ? void 0 : _b.call(_a, path)) || `() => import("${path}")`}`;
  }
  function componentReplacer(str, replaceStr, path) {
    var _a, _b;
    let importName = importsMap.get(path);
    if (!importName)
      importName = ROUTE_IMPORT_NAME.replace("$1", `${importsMap.size}`);
    importsMap.set(path, importName);
    importName = ((_b = (_a = options.resolver.stringify) == null ? void 0 : _a.component) == null ? void 0 : _b.call(_a, importName)) || importName;
    return str.replace(replaceStr, importName);
  }
  function functionReplacer(str, replaceStr, content) {
    if (content.startsWith("function"))
      return str.replace(replaceStr, content);
    if (content.startsWith("_NuFrRa_"))
      return str.replace(replaceStr, content.slice(8));
    return str;
  }
  const stringRoutes = JSON.stringify(preparedRoutes, replaceFunction).replace(componentRE, componentReplacer).replace(hasFunctionRE, functionReplacer);
  const imports = Array.from(importsMap).map((args) => getImportString(...args));
  return {
    imports,
    stringRoutes
  };
}
function generateClientCode(routes, options) {
  var _a, _b;
  const { imports, stringRoutes } = stringifyRoutes(routes, options);
  const code = `${imports.join(";\n")};

const routes = ${stringRoutes};

export default routes;`;
  return ((_b = (_a = options.resolver.stringify) == null ? void 0 : _a.final) == null ? void 0 : _b.call(_a, code)) || code;
}

// src/customBlock.ts
import fs from "fs";
import JSON5 from "json5";
import { parse as YAMLParser } from "yaml";
import { importModule } from "local-pkg";
import extractComments from "extract-comments";
var routeJSXReg = /^[\n\s]+(route)[\n\s]+/gm;
function parseJSX(code) {
  return extractComments(code).slice(0, 1).filter((comment) => routeJSXReg.test(comment.value) && comment.value.includes(":") && comment.loc.start.line === 1);
}
function parseYamlComment(code, path) {
  return code.reduce((memo, item) => {
    const { value } = item;
    const v = value.replace(routeJSXReg, "");
    debug.routeBlock(`use ${v} parser`);
    try {
      const yamlResult = YAMLParser(v);
      return __spreadValues(__spreadValues({}, memo), yamlResult);
    } catch (err) {
      throw new Error(`Invalid YAML format of comment in ${path}
${err.message}`);
    }
  }, {});
}
async function parseSFC(code) {
  try {
    const { parse } = await importModule("@vue/compiler-sfc");
    return parse(code, {
      pad: "space"
    }).descriptor;
  } catch (e) {
    throw new Error(`[vite-plugin-pages] Vue3's "@vue/compiler-sfc" is required.`);
  }
}
function parseCustomBlock(block, filePath, options) {
  var _a;
  const lang = (_a = block.lang) != null ? _a : options.routeBlockLang;
  debug.routeBlock(`use ${lang} parser`);
  if (lang === "json5") {
    try {
      return JSON5.parse(block.content);
    } catch (err) {
      throw new Error(`Invalid JSON5 format of <${block.type}> content in ${filePath}
${err.message}`);
    }
  } else if (lang === "json") {
    try {
      return JSON.parse(block.content);
    } catch (err) {
      throw new Error(`Invalid JSON format of <${block.type}> content in ${filePath}
${err.message}`);
    }
  } else if (lang === "yaml" || lang === "yml") {
    try {
      return YAMLParser(block.content);
    } catch (err) {
      throw new Error(`Invalid YAML format of <${block.type}> content in ${filePath}
${err.message}`);
    }
  }
}
async function getRouteBlock(path, options) {
  const content = fs.readFileSync(path, "utf8");
  const parsedSFC = await parseSFC(content);
  const blockStr = parsedSFC == null ? void 0 : parsedSFC.customBlocks.find((b) => b.type === "route");
  const parsedJSX = parseJSX(content);
  if (!blockStr && parsedJSX.length === 0)
    return;
  let result;
  if (blockStr)
    result = parseCustomBlock(blockStr, path, options);
  if (parsedJSX.length > 0)
    result = parseYamlComment(parsedJSX, path);
  return result;
}

// src/resolvers/vue.ts
function prepareRoutes(ctx, routes, parent) {
  var _a, _b, _c, _d;
  for (const route of routes) {
    if (route.name)
      route.name = route.name.replace(RegExp(`${ctx.options.routeNameSeparator}index$`), "");
    if (parent)
      route.path = (_a = route.path) == null ? void 0 : _a.replace(/^\//, "");
    if (route.children)
      route.children = prepareRoutes(ctx, route.children, route);
    if ((_b = route.children) == null ? void 0 : _b.find((c) => c.name === route.name))
      delete route.name;
    route.props = true;
    delete route.rawRoute;
    if (route.customBlock) {
      Object.assign(route, route.customBlock || {});
      delete route.customBlock;
    }
    Object.assign(route, ((_d = (_c = ctx.options).extendRoute) == null ? void 0 : _d.call(_c, route, parent)) || {});
  }
  return routes;
}
async function computeVueRoutes(ctx, customBlockMap) {
  var _a, _b;
  const { routeStyle, caseSensitive, routeNameSeparator } = ctx.options;
  const pageRoutes = [...ctx.pageRouteMap.values()].sort((a, b) => countSlash(a.route) - countSlash(b.route));
  const routes = [];
  pageRoutes.forEach((page) => {
    const pathNodes = page.route.split("/");
    const component = page.path.replace(ctx.root, "");
    const customBlock = customBlockMap.get(page.path);
    const route = {
      name: "",
      path: "",
      component,
      customBlock,
      rawRoute: page.route
    };
    let parentRoutes = routes;
    let dynamicRoute = false;
    for (let i = 0; i < pathNodes.length; i++) {
      const node = pathNodes[i];
      const nuxtStyle = routeStyle === "nuxt";
      const isDynamic = isDynamicRoute(node, nuxtStyle);
      const isCatchAll = isCatchAllRoute(node, nuxtStyle);
      const normalizedName = normalizeName(node, isDynamic, nuxtStyle);
      const normalizedPath = normalizeCase(normalizedName, caseSensitive);
      if (isDynamic)
        dynamicRoute = true;
      route.name += route.name ? `${routeNameSeparator}${normalizedName}` : normalizedName;
      const parent = parentRoutes.find((parent2) => {
        return pathNodes.slice(0, i + 1).join("/") === parent2.rawRoute;
      });
      if (parent) {
        parent.children = parent.children || [];
        parentRoutes = parent.children;
        route.path = "";
      } else if (normalizedPath === "index") {
        if (!route.path)
          route.path = "/";
      } else if (normalizedPath !== "index") {
        if (isDynamic) {
          route.path += `/:${normalizedName}`;
          if (isCatchAll) {
            if (i === 0)
              route.path += "(.*)*";
            else
              route.path += "(.*)";
          } else if (nuxtStyle && i === pathNodes.length - 1) {
            const isIndexFound = pageRoutes.find(({ route: route2 }) => {
              return route2 === page.route.replace(pathNodes[i], "index");
            });
            if (!isIndexFound)
              route.path += "?";
          }
        } else {
          route.path += `/${normalizedPath}`;
        }
      }
    }
    if (dynamicRoute)
      parentRoutes.push(route);
    else
      parentRoutes.unshift(route);
  });
  let finalRoutes = prepareRoutes(ctx, routes);
  finalRoutes = await ((_b = (_a = ctx.options).onRoutesGenerated) == null ? void 0 : _b.call(_a, finalRoutes)) || finalRoutes;
  return finalRoutes;
}
async function resolveVueRoutes(ctx, customBlockMap) {
  var _a, _b;
  const finalRoutes = await computeVueRoutes(ctx, customBlockMap);
  let client = generateClientCode(finalRoutes, ctx.options);
  client = await ((_b = (_a = ctx.options).onClientGenerated) == null ? void 0 : _b.call(_a, client)) || client;
  return client;
}
function vueResolver() {
  const customBlockMap = /* @__PURE__ */ new Map();
  async function checkCustomBlockChange(ctx, path) {
    var _a;
    const exitsCustomBlock = customBlockMap.get(path);
    let customBlock;
    try {
      customBlock = await getRouteBlock(path, ctx.options);
    } catch (error) {
      (_a = ctx.logger) == null ? void 0 : _a.error(colors.red(`[vite-plugin-pages] ${error.message}`));
      return;
    }
    if (!exitsCustomBlock && !customBlock)
      return;
    if (!customBlock) {
      customBlockMap.delete(path);
      ctx.debug.routeBlock("%s deleted", path);
      return;
    }
    if (!exitsCustomBlock || !deepEqual(exitsCustomBlock, customBlock)) {
      ctx.debug.routeBlock("%s old: %O", path, exitsCustomBlock);
      ctx.debug.routeBlock("%s new: %O", path, customBlock);
      customBlockMap.set(path, customBlock);
      ctx.onUpdate();
    }
  }
  return {
    resolveExtensions() {
      return ["vue", "ts", "js"];
    },
    resolveModuleIds() {
      return ["~pages", "pages-generated", "virtual:generated-pages"];
    },
    async resolveRoutes(ctx) {
      return resolveVueRoutes(ctx, customBlockMap);
    },
    async getComputedRoutes(ctx) {
      return computeVueRoutes(ctx, customBlockMap);
    },
    hmr: {
      added: async (ctx, path) => checkCustomBlockChange(ctx, path),
      changed: async (ctx, path) => checkCustomBlockChange(ctx, path),
      removed: async (_ctx, path) => {
        customBlockMap.delete(path);
      }
    }
  };
}

// src/resolvers/react.ts
function prepareRoutes2(routes, options, parent) {
  var _a, _b;
  for (const route of routes) {
    if (parent)
      route.path = (_a = route.path) == null ? void 0 : _a.replace(/^\//, "");
    if (route.children)
      route.children = prepareRoutes2(route.children, options, route);
    delete route.rawRoute;
    Object.assign(route, ((_b = options.extendRoute) == null ? void 0 : _b.call(options, route, parent)) || {});
  }
  return routes;
}
async function computeReactRoutes(ctx) {
  var _a, _b;
  const { routeStyle, caseSensitive } = ctx.options;
  const nuxtStyle = routeStyle === "nuxt";
  const pageRoutes = [...ctx.pageRouteMap.values()].sort((a, b) => countSlash(a.route) - countSlash(b.route));
  const routes = [];
  pageRoutes.forEach((page) => {
    const pathNodes = page.route.split("/");
    const element = page.path.replace(ctx.root, "");
    let parentRoutes = routes;
    for (let i = 0; i < pathNodes.length; i++) {
      const node = pathNodes[i];
      const route = {
        caseSensitive,
        path: "",
        rawRoute: pathNodes.slice(0, i + 1).join("/")
      };
      if (i === pathNodes.length - 1)
        route.element = element;
      const isIndexRoute = normalizeCase(node, caseSensitive).endsWith("index");
      if (!route.path && isIndexRoute) {
        route.path = "/";
      } else if (!isIndexRoute) {
        if (routeStyle === "remix")
          route.path = buildReactRemixRoutePath(node);
        else
          route.path = buildReactRoutePath(node, nuxtStyle);
      }
      const parent = parentRoutes.find((parent2) => {
        return pathNodes.slice(0, i).join("/") === parent2.rawRoute;
      });
      if (parent) {
        parent.children = parent.children || [];
        parentRoutes = parent.children;
      }
      const exits = parentRoutes.some((parent2) => {
        return pathNodes.slice(0, i + 1).join("/") === parent2.rawRoute;
      });
      if (!exits)
        parentRoutes.push(route);
    }
  });
  let finalRoutes = prepareRoutes2(routes, ctx.options);
  finalRoutes = await ((_b = (_a = ctx.options).onRoutesGenerated) == null ? void 0 : _b.call(_a, finalRoutes)) || finalRoutes;
  return finalRoutes;
}
async function resolveReactRoutes(ctx) {
  var _a, _b;
  const finalRoutes = await computeReactRoutes(ctx);
  let client = generateClientCode(finalRoutes, ctx.options);
  client = await ((_b = (_a = ctx.options).onClientGenerated) == null ? void 0 : _b.call(_a, client)) || client;
  return client;
}
function reactResolver() {
  return {
    resolveModuleIds() {
      return ["~react-pages", "virtual:generated-pages-react"];
    },
    resolveExtensions() {
      return ["tsx", "jsx", "ts", "js"];
    },
    async resolveRoutes(ctx) {
      return resolveReactRoutes(ctx);
    },
    async getComputedRoutes(ctx) {
      return computeReactRoutes(ctx);
    },
    stringify: {
      component: (path) => `React.createElement(${path})`,
      dynamicImport: (path) => `React.lazy(() => import("${path}"))`,
      final: (code) => `import React from "react";
${code}`
    }
  };
}

// src/resolvers/solid.ts
function prepareRoutes3(options, routes, parent) {
  var _a, _b;
  for (const route of routes) {
    if (parent)
      route.path = (_a = route.path) == null ? void 0 : _a.replace(/^\//, "");
    if (route.children)
      route.children = prepareRoutes3(options, route.children, route);
    delete route.rawRoute;
    Object.assign(route, ((_b = options.extendRoute) == null ? void 0 : _b.call(options, route, parent)) || {});
  }
  return routes;
}
async function computeSolidRoutes(ctx) {
  var _a, _b;
  const { routeStyle, caseSensitive } = ctx.options;
  const nuxtStyle = routeStyle === "nuxt";
  const pageRoutes = [...ctx.pageRouteMap.values()].sort((a, b) => countSlash(a.route) - countSlash(b.route));
  const routes = [];
  pageRoutes.forEach((page) => {
    const pathNodes = page.route.split("/");
    const component = page.path.replace(ctx.root, "");
    const element = page.path.replace(ctx.root, "");
    let parentRoutes = routes;
    for (let i = 0; i < pathNodes.length; i++) {
      const node = pathNodes[i];
      const normalizedPath = normalizeCase(node, caseSensitive);
      const route = {
        path: "",
        rawRoute: pathNodes.slice(0, i + 1).join("/")
      };
      const parent = parentRoutes.find(
        (parent2) => pathNodes.slice(0, i).join("/") === parent2.rawRoute
      );
      if (parent) {
        parent.children = parent.children || [];
        parentRoutes = parent.children;
      }
      if (i === pathNodes.length - 1) {
        route.element = element;
        route.component = component;
      }
      if (normalizedPath === "index") {
        if (!route.path)
          route.path = "/";
      } else if (normalizedPath !== "index") {
        if (routeStyle === "remix")
          route.path = buildReactRemixRoutePath(node) || "";
        else
          route.path = buildReactRoutePath(node, nuxtStyle) || "";
      }
      const exist = parentRoutes.some((parent2) => {
        return pathNodes.slice(0, i + 1).join("/") === parent2.rawRoute;
      });
      if (!exist)
        parentRoutes.push(route);
    }
  });
  let finalRoutes = prepareRoutes3(ctx.options, routes);
  finalRoutes = await ((_b = (_a = ctx.options).onRoutesGenerated) == null ? void 0 : _b.call(_a, finalRoutes)) || finalRoutes;
  return finalRoutes;
}
async function resolveSolidRoutes(ctx) {
  var _a, _b;
  const finalRoutes = await computeSolidRoutes(ctx);
  let client = generateClientCode(finalRoutes, ctx.options);
  client = await ((_b = (_a = ctx.options).onClientGenerated) == null ? void 0 : _b.call(_a, client)) || client;
  return client;
}
function solidResolver() {
  return {
    resolveModuleIds() {
      return ["~solid-pages"];
    },
    resolveExtensions() {
      return ["tsx", "jsx", "ts", "js"];
    },
    async resolveRoutes(ctx) {
      return resolveSolidRoutes(ctx);
    },
    async getComputedRoutes(ctx) {
      return computeSolidRoutes(ctx);
    },
    stringify: {
      dynamicImport: (path) => `Solid.lazy(() => import("${path}"))`,
      final: (code) => `import * as Solid from "solid-js";
${code}`
    }
  };
}

// src/options.ts
function resolvePageDirs(dirs, root, exclude) {
  dirs = toArray(dirs);
  return dirs.flatMap((dir) => {
    const option = typeof dir === "string" ? { dir, baseRoute: "" } : dir;
    option.dir = slash(resolve2(root, option.dir)).replace(`${root}/`, "");
    option.baseRoute = option.baseRoute.replace(/^\//, "").replace(/\/$/, "");
    return getPageDirs(option, root, exclude);
  });
}
var syncIndexResolver = (filepath, options) => {
  for (const page of options.dirs) {
    if (page.baseRoute === "" && filepath.startsWith(`/${page.dir}/index`))
      return "sync";
  }
  return "async";
};
var getResolver = (originalResolver) => {
  let resolver = originalResolver || "vue";
  if (typeof resolver !== "string")
    return resolver;
  switch (resolver) {
    case "vue":
      resolver = vueResolver();
      break;
    case "react":
      resolver = reactResolver();
      break;
    case "solid":
      resolver = solidResolver();
      break;
    default:
      throw new Error(`Unsupported resolver: ${resolver}`);
  }
  return resolver;
};
function resolveOptions(userOptions, viteRoot) {
  var _a;
  const {
    dirs = userOptions.pagesDir || ["src/pages"],
    routeBlockLang = "json5",
    exclude = ["node_modules", ".git", "**/__*__/**"],
    caseSensitive = false,
    syncIndex = true,
    routeNameSeparator = "-",
    extendRoute,
    onRoutesGenerated,
    onClientGenerated
  } = userOptions;
  const root = viteRoot || slash(process.cwd());
  const importMode = userOptions.importMode || (syncIndex ? syncIndexResolver : "async");
  const resolver = getResolver(userOptions.resolver);
  const extensions = userOptions.extensions || resolver.resolveExtensions();
  const extensionsRE = new RegExp(`\\.(${extensions.join("|")})$`);
  const resolvedDirs = resolvePageDirs(dirs, root, exclude);
  const routeStyle = userOptions.nuxtStyle ? "nuxt" : userOptions.routeStyle || "next";
  const moduleIds = userOptions.moduleId ? [userOptions.moduleId] : ((_a = resolver.resolveModuleIds) == null ? void 0 : _a.call(resolver)) || MODULE_IDS;
  const resolvedOptions = {
    dirs: resolvedDirs,
    routeStyle,
    routeBlockLang,
    moduleIds,
    root,
    extensions,
    importMode,
    exclude,
    caseSensitive,
    resolver,
    extensionsRE,
    extendRoute,
    onRoutesGenerated,
    onClientGenerated,
    routeNameSeparator
  };
  return resolvedOptions;
}

// src/context.ts
var PageContext = class {
  constructor(userOptions, viteRoot = process.cwd()) {
    this._pageRouteMap = /* @__PURE__ */ new Map();
    this.rawOptions = userOptions;
    this.root = slash(viteRoot);
    debug.env("root", this.root);
    this.options = resolveOptions(userOptions, this.root);
    debug.options(this.options);
  }
  setLogger(logger) {
    this.logger = logger;
  }
  setupViteServer(server) {
    if (this._server === server)
      return;
    this._server = server;
    this.setupWatcher(server.watcher);
  }
  setupWatcher(watcher) {
    watcher.on("unlink", async (path) => {
      path = slash(path);
      if (!isTarget(path, this.options))
        return;
      await this.removePage(path);
      this.onUpdate();
    });
    watcher.on("add", async (path) => {
      path = slash(path);
      if (!isTarget(path, this.options))
        return;
      const page = this.options.dirs.find((i) => path.startsWith(slash(resolve3(this.root, i.dir))));
      await this.addPage(path, page);
      this.onUpdate();
    });
    watcher.on("change", async (path) => {
      var _a, _b;
      path = slash(path);
      if (!isTarget(path, this.options))
        return;
      const page = this._pageRouteMap.get(path);
      if (page)
        await ((_b = (_a = this.options.resolver.hmr) == null ? void 0 : _a.changed) == null ? void 0 : _b.call(_a, this, path));
    });
  }
  async addPage(path, pageDir) {
    var _a, _b;
    debug.pages("add", path);
    for (const p of toArray(path)) {
      const pageDirPath = slash(resolve3(this.root, pageDir.dir));
      const route = slash(join2(pageDir.baseRoute, p.replace(`${pageDirPath}/`, "").replace(extname(p), "")));
      this._pageRouteMap.set(p, {
        path: p,
        route
      });
      await ((_b = (_a = this.options.resolver.hmr) == null ? void 0 : _a.added) == null ? void 0 : _b.call(_a, this, p));
    }
  }
  async removePage(path) {
    var _a, _b;
    debug.pages("remove", path);
    this._pageRouteMap.delete(path);
    await ((_b = (_a = this.options.resolver.hmr) == null ? void 0 : _a.removed) == null ? void 0 : _b.call(_a, this, path));
  }
  onUpdate() {
    if (!this._server)
      return;
    invalidatePagesModule(this._server);
    debug.hmr("Reload generated pages.");
    this._server.ws.send({
      type: "full-reload"
    });
  }
  async resolveRoutes() {
    return this.options.resolver.resolveRoutes(this);
  }
  async searchGlob() {
    const pageDirFiles = this.options.dirs.map((page) => {
      const pagesDirPath = slash(resolve3(this.options.root, page.dir));
      const files = getPageFiles(pagesDirPath, this.options);
      debug.search(page.dir, files);
      return __spreadProps(__spreadValues({}, page), {
        files: files.map((file) => slash(file))
      });
    });
    for (const page of pageDirFiles)
      await this.addPage(page.files, page);
    debug.cache(this.pageRouteMap);
  }
  get debug() {
    return debug;
  }
  get pageRouteMap() {
    return this._pageRouteMap;
  }
};

// src/index.ts
function pagesPlugin(userOptions = {}) {
  let ctx;
  return {
    name: "vite-plugin-pages",
    enforce: "pre",
    async configResolved(config) {
      if (!userOptions.resolver && config.plugins.find((i) => i.name.includes("vite:react")))
        userOptions.resolver = "react";
      if (!userOptions.resolver && config.plugins.find((i) => i.name.includes("solid")))
        userOptions.resolver = "solid";
      ctx = new PageContext(userOptions, config.root);
      ctx.setLogger(config.logger);
      await ctx.searchGlob();
    },
    api: {
      getResolvedRoutes() {
        return ctx.options.resolver.getComputedRoutes(ctx);
      }
    },
    configureServer(server) {
      ctx.setupViteServer(server);
    },
    resolveId(id) {
      if (ctx.options.moduleIds.includes(id))
        return `${MODULE_ID_VIRTUAL}?id=${id}`;
      if (routeBlockQueryRE.test(id))
        return ROUTE_BLOCK_ID_VIRTUAL;
      return null;
    },
    async load(id) {
      const {
        moduleId,
        pageId
      } = parsePageRequest(id);
      if (moduleId === MODULE_ID_VIRTUAL && pageId && ctx.options.moduleIds.includes(pageId))
        return ctx.resolveRoutes();
      if (id === ROUTE_BLOCK_ID_VIRTUAL) {
        return {
          code: "export default {};",
          map: null
        };
      }
      return null;
    }
  };
}
var src_default = pagesPlugin;
export {
  PageContext,
  src_default as default,
  reactResolver,
  solidResolver,
  syncIndexResolver,
  vueResolver
};
//# sourceMappingURL=index.mjs.map