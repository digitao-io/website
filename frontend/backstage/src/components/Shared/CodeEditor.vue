<script setup lang="ts">
import { onMounted, ref } from "vue";
import { EditorView, minimalSetup } from "codemirror";
import { placeholder as cmPlaceholder } from "@codemirror/view";
import { json } from "@codemirror/lang-json";
import { markdown } from "@codemirror/lang-markdown";

const props = defineProps<{
  label: string;
  lang: "json" | "markdown";
  placeholder: string;
}>();

const model = defineModel<string>();

const mountPoint = ref<Element | null>(null);

onMounted(() => {
  const editors = {
    "json": json,
    "markdown": markdown,
  };

  new EditorView({
    doc: model.value,
    extensions: [
      minimalSetup,
      editors[props.lang](),
      cmPlaceholder(props.placeholder),
      EditorView.updateListener.of((e) => {
        if (e.docChanged) {
          model.value = e.state.doc.toString();
        }
      }),
    ],
    parent: mountPoint.value!,
  });
});
</script>

<template>
  <label class="code-editor">
    <div class="code-editor-label">
      {{ props.label }}
    </div>
    <div
      ref="mountPoint"
      class="code-editor-editor"
    />
  </label>
</template>

<style scoped>
:deep(.cm-editor) {
  border: var(--border-xs) solid var(--color-text);
  padding: 4px 6px 4px 2px;
}

:deep(.cm-editor.cm-focused) {
  border: var(--border-xs) solid var(--color-text-flavor);
  outline: none;
}

.code-editor {
  display: block;
  margin-bottom: var(--margin-l);
}
</style>
