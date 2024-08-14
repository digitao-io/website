<script setup lang="ts">
import { ref } from "vue";
import { ResponseStatus, sendHttpRequest } from "frontend-resources";
import type { Tag } from "frontend-resources";
import TextButton from "@/components/Shared/TextButton.vue";
import TextInput from "@/components/Shared/TextInput.vue";

const emit = defineEmits<{
  create: [];
}>();

const key = ref<string>("");
const name = ref<string>("");

const onCreateClick = async () => {
  const body = {
    key: key.value,
    name: name.value,
  };

  const fileCreateRes = await sendHttpRequest<undefined, Tag, undefined>("", "/data/tag-create", undefined, body);
  if (fileCreateRes.status !== ResponseStatus.OK) {
    return;
  }

  reset();

  emit("create");
};

const reset = () => {
  key.value = "";
  name.value = "";
};
</script>

<template>
  <div class="tag-create-form">
    <text-input
      v-model="key"
      label="Key"
      placeholder="Internal key of the tag"
      :validation="{
        type: 'string',
        minLength: 1,
      }"
      error-message="Key cannot be empty"
    />

    <text-input
      v-model="name"
      label="Name"
      placeholder="Displayed name of the tag"
      :validation="{
        type: 'string',
        minLength: 1,
      }"
      error-message="Name cannot be empty"
    />

    <text-button
      class="tag-create-button"
      label="Create"
      type="primary"
      @click="onCreateClick"
    />
  </div>
</template>

<style scoped>
.tag-create-form {
  width: 80%;
}

.tag-create-button {
  margin-top: var(--margin-m);
}
</style>
