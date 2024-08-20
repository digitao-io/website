<script setup lang="ts">
import { computed, ref, watch } from "vue";
import { ResponseStatus, sendHttpRequest } from "frontend-resources";
import type { Page, PageIdentifier } from "frontend-resources";
import DetailsList from "@/components/Shared/DetailsList.vue";
import JsonDisplay from "@/components/Shared/JsonDisplay.vue";

const props = defineProps<{
  pageKey: string;
}>();

const page = ref<Page>();

const details = computed(() => {
  if (!page.value) {
    return [];
  }

  return [
    { label: "Key", data: page.value.key },
    { label: "Title", data: page.value.title },
    { label: "URL Pattern", data: page.value.urlPattern },
    { label: "Description", data: page.value.description },
  ];
});

watch(() => props.pageKey, async () => {
  const query = { key: props.pageKey };
  const response = await sendHttpRequest<PageIdentifier, undefined, Page>("", "/site/page-get", query);
  if (response.status !== ResponseStatus.OK) {
    return;
  }

  page.value = response.data;
}, { immediate: true });
</script>

<template>
  <div class="page-details-list">
    <details-list
      :details="details"
    />

    <div class="page-details-extra-title">
      Details
    </div>
    <json-display
      :json-object="page?.details ?? 'loading ...'"
    />
  </div>
</template>

<style scoped>
.page-details-extra-title {
  margin-top: var(--margin-m);
  font-size: var(--font-size-s);
  line-height: var(--line-height-m);
  color: var(--color-text-flavor);
}
</style>
