import { unplugin } from './index.mjs';
import 'unplugin';
import 'pathe';
import 'debug';
import 'fast-glob';
import '@intlify/shared';
import '@rollup/pluginutils';
import '@intlify/bundle-utils';
import '@vue/compiler-sfc';
import 'json5';
import 'js-yaml';
import 'fs';
import 'picocolors';

const vite = unplugin.vite;

export { vite as default };
