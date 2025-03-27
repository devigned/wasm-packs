import { nodeResolve } from "@rollup/plugin-node-resolve";
import commonjs from "@rollup/plugin-commonjs";

export default {
  input: ["src/server.js"],
  output: {
    file: "dist/index.js",
    format: "esm",
  },
  plugins: [commonjs(), nodeResolve()],
};
