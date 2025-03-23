import { readFile, writeFile } from "node:fs/promises";
import { resolve } from "node:path";

import { componentize } from "@bytecodealliance/componentize-js";

// AoT compilation makes use of weval (https://github.com/bytecodealliance/weval)
const enableAot = process.env.ENABLE_AOT == "1";

const jsSource = await readFile("server.js", "utf8");

const { component } = await componentize(jsSource, {
  witPath: resolve("wit"),
  enableAot,
});

let componentName = process.env.COMPONENT_NAME ?? "server.component.wasm";
await writeFile(componentName, component);
