<script setup lang="ts">
import { onBeforeMount, ref } from "vue";
import { ResponseStatus, sendHttpRequest } from "frontend-resources";
import type { Page } from "frontend-resources";
import PageLayout from "@/components/Shared/PageLayout.vue";
import DataTable from "@/components/Shared/DataTable.vue";
import EditorSwitch from "@/components/Shared/EditorSwitch.vue";
import PageCreateUpdateForm from "@/components/Pages/PageCreateUpdateForm.vue";
import PageDetailsList from "@/components/Pages/PageDetailsList.vue";
import PageDeleteForm from "@/components/Pages/PageDeleteForm.vue";

const pages = ref<Page[]>([]);
const selectedPageKeys = ref<string[]>([]);

onBeforeMount(() => {
  loadData();
});

const loadData = async () => {
  const response = await sendHttpRequest<undefined, undefined, Page[]>("", "/site/page-list");
  if (response.status !== ResponseStatus.OK) {
    return;
  }

  pages.value = response.data;
};

const onSelect = (selected: string[]) => {
  selectedPageKeys.value = selected;
};
</script>

<template>
  <page-layout
    class="pages"
    title="Pages"
  >
    <template #table>
      <data-table
        :columns="[
          { label: 'Key', dataExtractor: (page) => page.key, align: 'left', isKey: true },
          { label: 'Title', dataExtractor: (page) => page.title, align: 'left' },
          { label: 'URL Pattern', dataExtractor: (page) => page.urlPattern, align: 'left' },
          { label: 'Description', dataExtractor: (page) => page.description , align: 'left', multiline: true },
        ]"
        :data="pages"
        :selected="selectedPageKeys"
        @select="onSelect"
      />
    </template>

    <template #editor>
      <editor-switch
        :editors="[
          { name: 'create', label: 'Create Page', type: 'global', default: true },
          { name: 'details', label: 'Info', type: 'single', default: true },
          { name: 'edit', label: 'Edit', type: 'single' },
          { name: 'delete', label: 'Delete', type: 'single' },
        ]"
        :selected="selectedPageKeys"
        @select="onSelect"
      >
        <template #create="{ changeTo }">
          <page-create-update-form
            @save="changeTo(null); loadData();"
          />
        </template>

        <template #details>
          <page-details-list
            :page-key="selectedPageKeys[0]"
          />
        </template>

        <template #edit="{ changeTo }">
          <page-create-update-form
            :page-key="selectedPageKeys[0]"
            @save="changeTo('details'); loadData();"
          />
        </template>

        <template #delete="{ changeTo }">
          <page-delete-form
            :page-key="selectedPageKeys[0]"
            @delete="changeTo(null); loadData();"
            @cancel="changeTo('details')"
          />
        </template>
      </editor-switch>
    </template>
  </page-layout>
</template>
