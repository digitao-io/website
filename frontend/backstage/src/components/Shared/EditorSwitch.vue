<script setup lang="ts">
import { computed, ref, watch } from "vue";
import TextButton from "./TextButton.vue";

type EditorDefinition = {
  name: string;
  label: string;
  type: "global" | "single" | "multiple";
  default?: true;
};

const props = defineProps<{
  editors: EditorDefinition[];
  selected: string[];
}>();

const emit = defineEmits<{
  select: [string[]];
}>();

const currentEditor = ref<string | null>(null);

const globalEditors = computed(() => props.editors.filter((editor) => editor.type === "global"));
const singleEditors = computed(() => props.editors.filter((editor) => editor.type === "single"));
const multipleEditors = computed(() => props.editors.filter((editor) => editor.type === "multiple"));

watch(() => props.selected, (newValue) => {
  if (newValue.length === 0) {
    currentEditor.value = globalEditors.value.find((editor) => editor.default)?.name ?? null;
  } else if (newValue.length === 1) {
    currentEditor.value = singleEditors.value.find((editor) => editor.default)?.name ?? null;
  } else {
    currentEditor.value = multipleEditors.value.find((editor) => editor.default)?.name ?? null;
  }
}, { immediate: true });

const selectEditor = (editor: EditorDefinition) => {
  if (editor.type === "global") {
    emit("select", []);
  }

  currentEditor.value = editor.name;
};

const switchTo = (name: string | null) => {
  const editor = props.editors.find((editor) => editor.name === name);
  if (!editor) {
    emit("select", []);
    return;
  }

  selectEditor(editor);
};
</script>

<template>
  <div class="editor-switch">
    <div class="editor-switch-global-actions">
      <text-button
        v-for="editor of globalEditors"
        :key="editor.name"
        :label="editor.label"
        :type="currentEditor === editor.name ? 'primary' : 'secondary'"
        @click="selectEditor(editor)"
      />
    </div>

    <div
      v-if="props.selected.length === 1"
      class="editor-switch-single-actions"
    >
      <text-button
        v-for="editor of singleEditors"
        :key="editor.name"
        :label="editor.label"
        :type="currentEditor === editor.name ? 'primary' : 'secondary'"
        @click="selectEditor(editor)"
      />
    </div>

    <div
      v-if="props.selected.length > 1"
      class="editor-switch-multiple-actions"
    >
      <text-button
        v-for="editor of multipleEditors"
        :key="editor.name"
        :label="editor.label"
        :type="currentEditor === editor.name ? 'primary' : 'secondary'"
        @click="selectEditor(editor)"
      />
    </div>

    <div class="editor-switch-editor-container">
      <template
        v-for="editor of props.editors"
        :key="editor.name"
      >
        <slot
          v-if="editor.name === currentEditor"
          :name="editor.name"
          :change-to="switchTo"
        />
      </template>
    </div>
  </div>
</template>

<style scoped>
.editor-switch-single-actions,
.editor-switch-multiple-actions,
.editor-switch-editor-container {
  margin-top: var(--margin-m);
}
</style>
