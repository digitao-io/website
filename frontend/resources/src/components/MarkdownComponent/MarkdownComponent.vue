<script setup lang="ts">
import { defineProps, ref } from "vue";
import type { PropType } from "vue";
import { Marked } from 'marked';
import { markedHighlight } from "marked-highlight";
import hljs from 'highlight.js';
import 'highlight.js/styles/github.css'; 

export type MarkdownComponentConfig = {
  markdown: string;
};

const props = defineProps({
  config: { type: Object as PropType<MarkdownComponentConfig>, required: true },
});

const marked = new Marked(
  markedHighlight({
    langPrefix: 'hljs language-',
    highlight(code, lang, info) {
      const language = hljs.getLanguage(lang) ? lang : 'plaintext';
      return hljs.highlight(code, { language }).value;
    }
  })
);


const htmlContent = ref(marked.parse(props.config.markdown));
</script>

<template>
  <div class="markdown-component" v-html="htmlContent">
  </div>
</template>

<style scoped>
.markdown-component {}
</style>
