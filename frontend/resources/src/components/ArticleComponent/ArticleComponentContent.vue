<script setup lang="ts">
import { computed } from "vue";

import { Marked } from "marked";
import { markedHighlight } from "marked-highlight";
import hljs from "highlight.js";

import "highlight.js/styles/github.css";

const props = defineProps<{
  content: string;
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

const htmlContent = computed(() => marked.parse(props.content));
</script>

<template>
  <!-- eslint-disable vue/no-v-html -->
  <div
    class="article-content"
    v-html="htmlContent"
  />
  <!-- eslint-enable vue/no-v-html -->
</template>

<style scoped>
:deep(h1),
:deep(h5),
:deep(h6) {
  display: none;
}

:deep(h2) {
  margin: 0.6em 0 0.5em 0;
  font-family: var(--font-raleway);
  font-size: var(--font-size-xxl);
  line-height: var(--line-height-xxl);
}

:deep(h3) {
  margin: 0.6em 0 0.5em 0;
  font-family: var(--font-raleway);
  font-size: var(--font-size-xl);
  line-height: var(--line-height-xl);
}

:deep(h4) {
  margin: 0.6em 0 0.5em 0;
  font-family: var(--font-raleway);
  font-size: var(--font-size-l);
  line-height: var(--line-height-l);
}

:deep(strong) {
  font-weight: 600;
}

:deep(code) {
  font-size: var(--font-size-m);
  background-color: var(--color-secondary);
}
</style>
