import { Plugin } from 'vite';

/**
 * Plugin options.
 */
interface Options {
    /**
     * Relative path to the directory to search for page components.
     * @default 'src/layouts'
     */
    layoutsDirs: string | string[];
    /**
     * Valid file extensions for page components.
     * @default ['vue']
     */
    extensions: string[];
    /**
     * List of path globs to exclude when resolving pages.
     */
    exclude: string[];
    /**
     * Filename of default layout (".vue" is not needed)
     * @default 'default'
     */
    defaultLayout: string;
    /**
     * Mode for importing layouts
     */
    importMode: (name: string) => 'sync' | 'async';
}
declare type FileContainer = {
    path: string;
    files: string[];
};
declare type UserOptions = Partial<Options>;
interface ResolvedOptions extends Options {
}

declare function defaultImportMode(name: string): "sync" | "async";
declare function layoutPlugin(userOptions?: UserOptions): Plugin;

export { FileContainer, ResolvedOptions, UserOptions, layoutPlugin as default, defaultImportMode };
