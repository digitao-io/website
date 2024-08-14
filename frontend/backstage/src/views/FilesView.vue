<script setup lang="ts">
import { onBeforeMount, ref } from "vue";
import { ResponseStatus, sendHttpRequest } from "frontend-resources";
import type { File, FileSearchParams } from "frontend-resources";
import PageLayout from "@/components/Shared/PageLayout.vue";
import DataTable from "@/components/Shared/DataTable.vue";

const files = ref<File[]>([]);
const selectedFileId = ref<string | null>(null);

onBeforeMount(async () => {
  const query: FileSearchParams = {
    sort: "createdAt",
    order: "DESC",
    take: 20,
    skip: 0,
  };

  const response = await sendHttpRequest<FileSearchParams, undefined, File[]>("", "/data/file-list", query);
  if (response.status !== ResponseStatus.OK) {
    return;
  }

  files.value = response.data;
});

const onRowSelected = (selected: string | number) => {
  selectedFileId.value = selected as string;
};
</script>

<template>
  <page-layout
    class="files"
    title="Files"
  >
    <template #table>
      <data-table
        :columns="[
          { label: 'ID', dataExtractor: (file) => file.id, align: 'left', isKey: true },
          { label: 'Title', dataExtractor: (file) => file.title, align: 'left' },
          { label: 'Type', dataExtractor: (file) => file.mimeType, align: 'left' },
        ]"
        :data="files"
        :selected="selectedFileId"
        @row-select="onRowSelected"
      />
    </template>
  </page-layout>
</template>

<style scoped>
.files {
  display: flex;
  flex-wrap: wrap;
  margin: var(--margin-m);
}

.files-table-area {
  margin-right: var(--margin-m);
}
</style>
