import path from "node:path";
import { defineConfig } from "@rsbuild/core";
import { pluginBabel } from "@rsbuild/plugin-babel";
import { pluginLess } from "@rsbuild/plugin-less";
import { pluginVue } from "@rsbuild/plugin-vue";
import { pluginVueJsx } from "@rsbuild/plugin-vue-jsx";
import { pluginImageCompress } from "@rsbuild/plugin-image-compress";
import pluginAutoImport from "unplugin-auto-import/rspack";
import { AntDesignVueResolver } from "unplugin-vue-components/resolvers";
import pluginComponents from "unplugin-vue-components/rspack";
import Dotenv from "dotenv-webpack";

export default defineConfig({
  performance: {
    chunkSplit: {
      strategy: "split-by-experience",
      forceSplitting: {
        "vue3-apexcharts": /node_modules[\\/]vue3-apexcharts/,
      },
    },
    bundleAnalyze:
      process.env.BUNDLE_ANALYZE === "true"
        ? {
            analyzerMode: "server",
            openAnalyzer: true,
          }
        : { analyzerMode: "disabled" },
  },
  html: {
    template: "index.html",
  },
  tools: {
    rspack: {
      plugins: [
        new Dotenv(),
        pluginComponents({
          resolvers: [AntDesignVueResolver({ importStyle: false })],
          directoryAsNamespace: true,
        }),
        pluginAutoImport({
          imports: [
            "vue",
            "vue-router",
            "pinia",
            {
              "@/gettext": ["$gettext", "$pgettext", "$ngettext", "$npgettext"],
            },
          ],
          vueTemplate: true,
        }),
      ],
    },
  },
  plugins: [
    pluginBabel({
      include: /\.(?:jsx|tsx)$/,
    }),
    pluginVue(),
    pluginVueJsx(),
    pluginLess(),
    pluginImageCompress(),
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
