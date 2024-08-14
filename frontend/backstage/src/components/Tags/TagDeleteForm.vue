<script setup lang="ts">
import TextButton from "@/components/Shared/TextButton.vue";
import { ResponseStatus, sendHttpRequest } from "frontend-resources";
import type { TagIdentifier } from "frontend-resources";

const props = defineProps<{
  tagKey: string;
}>();

const emit = defineEmits<{
  delete: [];
  cancel: [];
}>();

const onDeleteClick = async () => {
  const query = {
    key: props.tagKey,
  };
  const response = await sendHttpRequest<TagIdentifier, undefined, undefined>("", "/data/tag-delete", query);
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
  <div class="tag-delete-form">
    <p class="tag-delete-form-message">
      Are you sure to delete tag <strong>{{ props.tagKey }}</strong>? This operation cannot be undone.
    </p>

    <div class="tag-delete-form-button-area">
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
.tag-delete-form-message {
  font-size: var(--font-size-l);
  line-height: var(--line-height-m);
}

.tag-delete-form-message strong {
  color: var(--color-text-flavor);
}

.tag-delete-form-button-area {
  margin-top: var(--margin-m);
}
</style>
