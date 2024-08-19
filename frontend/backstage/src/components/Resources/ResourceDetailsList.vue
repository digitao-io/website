<script setup lang="ts">
import { computed, ref, watch } from "vue";
import { ResponseStatus, sendHttpRequest } from "frontend-resources";
import type { Resource, ResourceIdentifier } from "frontend-resources";
import DetailsList from "@/components/Shared/DetailsList.vue";
import JsonDisplay from "@/components/Shared/JsonDisplay.vue";

const props = defineProps<{
  resourceKey: string;
}>();

const resource = ref<Resource>();

const details = computed(() => {
  if (!resource.value) {
    return [];
  }

  return [
    { label: "Key", data: resource.value.key },
    { label: "Title", data: resource.value.title },
    { label: "Type", data: resource.value.type },
    { label: "Description", data: resource.value.description },
  ];
});

watch(() => props.resourceKey, async () => {
  const query = { key: props.resourceKey };
  const response = await sendHttpRequest<ResourceIdentifier, undefined, Resource>("", "/site/resource-get", query);
  if (response.status !== ResponseStatus.OK) {
    return;
  }

  resource.value = response.data;
}, { immediate: true });
</script>

<template>
  <div class="resource-details-list">
    <details-list
      :details="details"
    />

    <div class="resource-details-extra-title">
      Details
    </div>
    <json-display
      :json-object="resource?.details ?? 'loading ...'"
    />
  </div>
</template>

<style scoped>
.resource-details-extra-title {
  font-size: var(--font-size-s);
  line-height: var(--line-height-m);
  color: var(--color-text-flavor);
}
</style>
