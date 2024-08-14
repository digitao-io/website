<script setup lang="ts">
import { onBeforeMount, ref } from "vue";
import { ResponseStatus, sendHttpRequest } from "frontend-resources";
import type { File, FileSearchParams } from "frontend-resources";
import PageLayout from "@/components/Shared/PageLayout.vue";
import DataTable from "@/components/Shared/DataTable.vue";
import EditorSwitch from "@/components/Shared/EditorSwitch.vue";
import FileUploadForm from "@/components/Files/FileUploadForm.vue";
import FileDetailsList from "@/components/Files/FileDetailsList.vue";
import FileDeleteForm from "@/components/Files/FileDeleteForm.vue";

const files = ref<File[]>([]);
const selectedFileIds = ref<string[]>([]);

onBeforeMount(() => {
  loadData();
});

const loadData = async () => {
  const query: FileSearchParams = {
    sort: "createdAt",
    order: "DESC",
    take: 100,
    skip: 0,
  };

  const response = await sendHttpRequest<FileSearchParams, undefined, File[]>("", "/data/file-list", query);
  if (response.status !== ResponseStatus.OK) {
    return;
  }

  files.value = response.data;
};

const onSelect = (selected: string[]) => {
  selectedFileIds.value = selected;
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
        :selected="selectedFileIds"
        @select="onSelect"
      />
    </template>

    <template #editor>
      <editor-switch
        :editors="[
          { name: 'upload', label: 'Upload File', type: 'global', default: true },
          { name: 'details', label: 'Details', type: 'single', default: true },
          { name: 'delete', label: 'Delete', type: 'single' },
        ]"
        :selected="selectedFileIds"
        @select="onSelect"
      >
        <template
          #upload="{ changeTo }"
        >
          <file-upload-form
            @create="changeTo(null); loadData();"
          />
        </template>

        <template #details>
          <file-details-list
            :file-id="selectedFileIds[0]"
          />
        </template>

        <template #delete="{ changeTo }">
          <file-delete-form
            :file-id="selectedFileIds[0]"
            @delete="changeTo(null); loadData();"
            @cancel="changeTo('details')"
          />
        </template>
      </editor-switch>
    </template>
  </page-layout>
</template>

<style scoped>
</style>
