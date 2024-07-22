import stylisticEslintPlugin from "@stylistic/eslint-plugin";
import typescriptEslintPlugin from "@typescript-eslint/eslint-plugin";
import typescriptEslintParser from "@typescript-eslint/parser";
import myEslintRules from "my-eslint-rules";

import eslintPluginVue from "eslint-plugin-vue";

export default [
  ...eslintPluginVue.configs["flat/recommended"],

  {
    plugins: {
      "@stylistic": stylisticEslintPlugin,
      "@typescript-eslint": typescriptEslintPlugin,
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
      ...myEslintRules,
    },
  },
];
