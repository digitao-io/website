<script setup lang="ts">
import { onBeforeMount, ref } from "vue";
import { ResponseStatus, sendHttpRequest } from "frontend-resources";
import type { Tag } from "frontend-resources";
import PageLayout from "@/components/Shared/PageLayout.vue";
import DataTable from "@/components/Shared/DataTable.vue";

const tags = ref<Tag[]>([]);
const selectedTagKeys = ref<string[]>([]);

onBeforeMount(async () => {
  const response = await sendHttpRequest<undefined, undefined, Tag[]>("", "/data/tag-list");
  if (response.status !== ResponseStatus.OK) {
    return;
  }

  tags.value = response.data;
});

const onSelect = (selected: string[]) => {
  selectedTagKeys.value = selected;
};
</script>

<template>
  <page-layout
    class="tags"
    title="Tags"
  >
    <template #table>
      <data-table
        :columns="[
          { label: 'Key', dataExtractor: (tag) => tag.key, align: 'left', isKey: true },
          { label: 'Name', dataExtractor: (tag) => tag.name, align: 'left' },
        ]"
        :data="tags"
        :selected="selectedTagKeys"
        @select="onSelect"
      />
    </template>
  </page-layout>
</template>

<style scoped>
</style>
