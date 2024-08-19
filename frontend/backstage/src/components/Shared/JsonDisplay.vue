
<script setup lang="ts">
/* eslint-disable @typescript-eslint/no-explicit-any */
import { computed } from "vue";

import hljs from "highlight.js";
import json from "highlight.js/lib/languages/json";
import "highlight.js/styles/vs.min.css";

hljs.registerLanguage("json", json);

const props = defineProps<{
  jsonObject: any;
}>();

const jsonString = computed(() => JSON.stringify(props.jsonObject, null, 2));

const jsonHtml = computed(() => hljs.highlight(jsonString.value, { language: "json" }).value);

/* eslint-enable @typescript-eslint/no-explicit-any */
</script>

<template>
  <div class="json-display">
    <!-- eslint-disable vue/no-v-html -->
    <pre><code
      class="language-html"
      v-html="jsonHtml"
    /></pre>
    <!-- eslint-enable vue/no-v-html -->
  </div>
</template>
