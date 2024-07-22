<script setup lang="ts">
import { defineProps } from "vue";
import type { PropType } from "vue";

export type MarkdownComponentConfig = {
  markdown: string;
};

const props = defineProps({
  config: { type: Object as PropType<MarkdownComponentConfig>, required: true },
});
</script>

<template>
  <div class="markdown-component">{{ config.markdown }}</div>
</template>

<style scoped>
.markdown-component {}
</style>
