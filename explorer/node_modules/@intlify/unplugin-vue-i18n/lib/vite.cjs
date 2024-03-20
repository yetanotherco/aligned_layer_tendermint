'use strict';

const index = require('./index.cjs');
require('unplugin');
require('pathe');
require('debug');
require('fast-glob');
require('@intlify/shared');
require('@rollup/pluginutils');
require('@intlify/bundle-utils');
require('@vue/compiler-sfc');
require('json5');
require('js-yaml');
require('fs');
require('picocolors');

const vite = index.unplugin.vite;

module.exports = vite;
