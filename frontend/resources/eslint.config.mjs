import stylistic from "@stylistic/eslint-plugin";
import typescriptEslintParser from "@typescript-eslint/parser";
import stylisticDefault from "stylistic-default";

import eslintPluginVue from "eslint-plugin-vue";

export default [
  ...eslintPluginVue.configs["flat/recommended"],

  {
    files: ["*.ts", "*.mjs", "**/*.ts"],

    languageOptions: {
      parser: typescriptEslintParser,
    },
  },

  {
    files: ["*.ts", "*.mjs", "src/**/*.ts", "src/**/*.vue"],

    plugins: {
      "@stylistic": stylistic,
    },

    rules: {
      ...stylisticDefault,
    },
  },
];
