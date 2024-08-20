<script setup lang="ts">
import TextButton from "@/components/Shared/TextButton.vue";
import { ResponseStatus, sendHttpRequest } from "frontend-resources";
import type { PageIdentifier } from "frontend-resources";

const props = defineProps<{
  pageKey: string;
}>();

const emit = defineEmits<{
  delete: [];
  cancel: [];
}>();

const onDeleteClick = async () => {
  const query = {
    key: props.pageKey,
  };
  const response = await sendHttpRequest<PageIdentifier, undefined, undefined>("", "/site/page-delete", query);
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
  <div class="page-delete-form">
    <p class="page-delete-form-message">
      Are you sure to delete resource <strong>{{ props.pageKey }}</strong>? This operation cannot be undone.
    </p>

    <div class="page-delete-form-button-area">
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
.page-delete-form-message {
  font-size: var(--font-size-l);
  line-height: var(--line-height-m);
}

.page-delete-form-message strong {
  color: var(--color-text-flavor);
}

.page-delete-form-button-area {
  margin-top: var(--margin-m);
}
</style>
