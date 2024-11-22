import path from "node:path";
import { defineConfig } from "@rsbuild/core";
import { pluginBabel } from "@rsbuild/plugin-babel";
import { pluginLess } from "@rsbuild/plugin-less";
import { pluginVue } from "@rsbuild/plugin-vue";
import { pluginVueJsx } from "@rsbuild/plugin-vue-jsx";

export default defineConfig({
  performance: {
    bundleAnalyze: process.env.BUNDLE_ANALYZE
      ? {
          analyzerMode: "server",
          openAnalyzer: true,
        }
      : {},
  },
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
  output: {
    filename: {
      js: "static/js/[name].[contenthash:8].js",
      css: "static/css/[name].[contenthash:8].css",
    },
  },
  source: {
    entry: {
      index: path.resolve(__dirname, "src/main.ts"),
    },
    alias: {
      "@": path.resolve(__dirname, "src"),
    },
  },
});