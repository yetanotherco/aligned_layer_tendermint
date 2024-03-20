/**
 * Code generator for i18n json/json5 resource
 */
/// <reference types="node" />
import type { JSONProgram } from 'jsonc-eslint-parser/lib/parser/ast';
import type { CodeGenOptions, CodeGenResult } from './codegen';
/**
 * @internal
 */
export declare function generate(targetSource: string | Buffer, { type, bridge, exportESM, filename, inSourceMap, locale, isGlobal, sourceMap, env, forceStringify, onError, useClassComponent }: CodeGenOptions, injector?: () => string): CodeGenResult<JSONProgram>;
