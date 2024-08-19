<script setup lang="ts">
import { computed, ref } from "vue";
import { ResponseStatus, sendHttpRequest } from "frontend-resources";
import type { Resource } from "frontend-resources";
import TextButton from "@/components/Shared/TextButton.vue";
import TextInput from "@/components/Shared/TextInput.vue";
import CodeEditor from "@/components/Shared/CodeEditor.vue";

const props = defineProps<{
  resourceKey?: string;
}>();

const emit = defineEmits<{
  save: [];
}>();

const key = ref<string>("");
const title = ref<string>("");
const type = ref<string>("");
const description = ref<string>("");
const details = ref<string>("");

const updating = computed(() => !!props.resourceKey);

const onSaveClick = async () => {
  const body = {
    key: key.value,
    title: title.value,
    type: type.value,
    description: description.value,
    details: JSON.parse(details.value),
  };

  const response = await sendHttpRequest<undefined, Resource, undefined>("", "/site/resource-create", undefined, body);

  if (response.status !== ResponseStatus.OK) {
    return;
  }

  emit("save");
};
</script>

<template>
  <div class="resource-create-update-form">
    <text-input
      v-model="key"
      label="Key"
      placeholder="Internal code of the resource"
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
      placeholder="A human-readable title for the resource"
      :validation="{
        type: 'string',
        minLength: 1,
      }"
      error-message="Title cannot be empty"
    />

    <text-input
      v-model="type"
      label="Type"
      placeholder="The type of the resource"
      :validation="{
        type: 'string',
        minLength: 1,
      }"
      error-message="Type cannot be empty"
    />

    <text-input
      v-model="description"
      label="Description"
      placeholder="A long description of what is the resource used for"
      :validation="{
        type: 'string',
        minLength: 1,
      }"
      error-message="Description cannot be empty"
    />

    <code-editor
      v-model="details"
      label="Details"
      placeholder="Details of the resource, should be JSON object or JSON array ..."
      lang="json"
    />

    <text-button
      class="resource-create-update-button"
      :label="updating ? 'Update' : 'Create'"
      type="primary"
      @click="onSaveClick"
    />
  </div>
</template>

<style scoped>
</style>
