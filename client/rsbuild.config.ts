import path from "node:path";
import { defineConfig } from "@rsbuild/core";
import { pluginBabel } from "@rsbuild/plugin-babel";
import { pluginLess } from "@rsbuild/plugin-less";
import { pluginVue } from "@rsbuild/plugin-vue";
import { pluginVueJsx } from "@rsbuild/plugin-vue-jsx";

export default defineConfig({
  html: {
    template: "index.html",
  },
  plugins: [
    pluginBabel({
      include: /\.(?:jsx|tsx)$/,
    }),
    pluginVue(),
    pluginVueJsx(),
    pluginLess(),
  ],
  source: {
    entry: {
      index: path.resolve(__dirname, "src/main.ts"),
    },
    alias: {
      "@": path.resolve(__dirname, "src"),
    },
  },
});
