<script setup lang="ts">
import TextButton from "@/components/Shared/TextButton.vue";
import { ResponseStatus, sendHttpRequest } from "frontend-resources";
import type { FileIdentifier } from "frontend-resources";

const props = defineProps<{
  fileId: string;
}>();

const emit = defineEmits<{
  delete: [];
  cancel: [];
}>();

const onDeleteClick = async () => {
  const query = {
    id: props.fileId,
  };
  const response = await sendHttpRequest<FileIdentifier, undefined, undefined>("", "/data/file-delete", query);
  if (response.status !== ResponseStatus.OK) {
    return;
  }

  emit("delete");
};

const onCancelClick = () => {
  emit("cancel");
};
</script>

<template>
  <div class="file-delete-form">
    <p class="file-delete-form-message">
      Are you sure to delete file <strong>{{ props.fileId }}</strong>? This operation cannot be undone.
    </p>

    <div class="file-delete-form-button-area">
      <text-button
        label="Delete"
        type="primary"
        @click="onDeleteClick"
      />

      <text-button
        label="Cancel"
        type="secondary"
        @click="onCancelClick"
      />
    </div>
  </div>
</template>

<style scoped>
.file-delete-form-message {
  font-size: var(--font-size-l);
  line-height: var(--line-height-m);
}

.file-delete-form-message strong {
  color: var(--color-text-flavor);
}

.file-delete-form-button-area {
  margin-top: var(--margin-m);
}
</style>
