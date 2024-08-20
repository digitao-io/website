<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import { ResponseStatus, sendHttpRequest } from "frontend-resources";
import type { Page, PageIdentifier, Response } from "frontend-resources";
import TextButton from "@/components/Shared/TextButton.vue";
import TextInput from "@/components/Shared/TextInput.vue";
import CodeEditor from "@/components/Shared/CodeEditor.vue";

const props = defineProps<{
  pageKey?: string;
}>();

const emit = defineEmits<{
  save: [];
}>();

const key = ref<string>("");
const title = ref<string>("");
const urlPattern = ref<string>("");
const description = ref<string>("");
const details = ref<string>("");

const codeEditor = ref<typeof CodeEditor | null>(null);

const updating = computed(() => !!props.pageKey);

onMounted(async () => {
  if (!updating.value) {
    return;
  }

  const query = {
    key: props.pageKey!,
  };

  const response = await sendHttpRequest<PageIdentifier, undefined, Page>("", "/site/page-get", query);
  if (response.status !== ResponseStatus.OK) {
    return;
  }

  key.value = response.data.key;
  title.value = response.data.title;
  urlPattern.value = response.data.urlPattern;
  description.value = response.data.description;

  details.value = JSON.stringify(response.data.details, null, 2);
  codeEditor.value?.setText(details.value);
});

const onCodeInput = (text: string) => {
  details.value = text;
};

const onSaveClick = async () => {
  const body = {
    key: key.value,
    title: title.value,
    urlPattern: urlPattern.value,
    description: description.value,
    details: JSON.parse(details.value),
  };

  let response: Response<undefined>;

  if (!updating.value) {
    response = await sendHttpRequest<undefined, Page, undefined>("", "/site/page-create", undefined, body);
  } else {
    const query = {
      key: props.pageKey!,
    };
    response = await sendHttpRequest<PageIdentifier, Page, undefined>("", "/site/page-update", query, body);
  }

  if (response.status !== ResponseStatus.OK) {
    return;
  }

  emit("save");
};
</script>

<template>
  <div class="page-create-update-form">
    <text-input
      v-model="key"
      label="Key"
      placeholder="Internal code of the page"
      :validation="{
        type: 'string',
        minLength: 1,
      }"
      error-message="Key cannot be empty"
      :disabled="updating"
    />

    <text-input
      v-model="title"
      label="Title"
      placeholder="A human-readable title for the page"
      :validation="{
        type: 'string',
        minLength: 1,
      }"
      error-message="Title cannot be empty"
    />

    <text-input
      v-model="urlPattern"
      label="Type"
      placeholder="The URL pattern for the page, following the Express format"
      :validation="{
        type: 'string',
        minLength: 1,
      }"
      error-message="Type cannot be empty"
    />

    <text-input
      v-model="description"
      label="Description"
      placeholder="A long description of what the page is used for"
      :validation="{
        type: 'string',
        minLength: 1,
      }"
      error-message="Description cannot be empty"
    />

    <code-editor
      ref="codeEditor"
      label="Details"
      placeholder="Details of the page, should be a JSON object ..."
      lang="json"
      @input="onCodeInput"
    />

    <text-button
      class="page-create-update-button"
      :label="updating ? 'Update' : 'Create'"
      type="primary"
      @click="onSaveClick"
    />
  </div>
</template>

<style scoped>
</style>
