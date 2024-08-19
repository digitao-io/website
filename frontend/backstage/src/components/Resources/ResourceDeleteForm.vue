<script setup lang="ts">
import TextButton from "@/components/Shared/TextButton.vue";
import { ResponseStatus, sendHttpRequest } from "frontend-resources";
import type { ResourceIdentifier } from "frontend-resources";

const props = defineProps<{
  resourceKey: string;
}>();

const emit = defineEmits<{
  delete: [];
  cancel: [];
}>();

const onDeleteClick = async () => {
  const query = {
    key: props.resourceKey,
  };
  const response = await sendHttpRequest<ResourceIdentifier, undefined, undefined>("", "/site/resource-delete", query);
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
  <div class="resource-delete-form">
    <p class="resource-delete-form-message">
      Are you sure to delete resource <strong>{{ props.resourceKey }}</strong>? This operation cannot be undone.
    </p>

    <div class="resource-delete-form-button-area">
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
.resource-delete-form-message {
  font-size: var(--font-size-l);
  line-height: var(--line-height-m);
}

.resource-delete-form-message strong {
  color: var(--color-text-flavor);
}

.resource-delete-form-button-area {
  margin-top: var(--margin-m);
}
</style>
