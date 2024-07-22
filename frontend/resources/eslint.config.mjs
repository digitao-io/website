import stylistic from "@stylistic/eslint-plugin";
import typescriptEslintParser from "@typescript-eslint/parser";
import stylisticDefault from "stylistic-default";

import eslintPluginVue from "eslint-plugin-vue";

export default [
  ...eslintPluginVue.configs["flat/recommended"],

  {
    plugins: {
      "@stylistic": stylistic,
    },

    files: [
      "*.mjs",
      "*.ts",
      "src/**/*",
    ],

    languageOptions: {
      parserOptions: {
        parser: typescriptEslintParser,
        sourceType: "module",
      },
    },

    rules: {
      ...stylisticDefault,
    },
  },
];
