declare function getLang(filename: string): string;
declare function isTs(lang?: string): boolean | "" | undefined;

export { getLang, isTs };
