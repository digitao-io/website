<script setup lang="ts">
import { computed } from "vue";

import { Marked } from "marked";
import { markedHighlight } from "marked-highlight";
import hljs from "highlight.js";

import "highlight.js/styles/github.css";

export type MarkdownComponentConfig = {
  markdown: string;
};

const props = defineProps<{
  config: MarkdownComponentConfig;
}>();

const marked = new Marked(
  markedHighlight({
    langPrefix: "hljs language-",
    highlight: (code, lang) => {
      const language = hljs.getLanguage(lang) ? lang : "plaintext";
      return hljs.highlight(code, { language }).value;
    },
  }),
);

const htmlContent = computed(() => marked.parse(props.config.markdown));
</script>

<template>
  <!-- eslint-disable vue/no-v-html -->
  <div
    class="markdown-component"
    v-html="htmlContent"
  />
  <!-- eslint-enable vue/no-v-html -->
</template>

<style scoped>
</style>
