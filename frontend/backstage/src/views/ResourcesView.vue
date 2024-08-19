<script setup lang="ts">
import { onBeforeMount, ref } from "vue";
import { ResponseStatus, sendHttpRequest } from "frontend-resources";
import type { Resource } from "frontend-resources";
import PageLayout from "@/components/Shared/PageLayout.vue";
import DataTable from "@/components/Shared/DataTable.vue";
import EditorSwitch from "@/components/Shared/EditorSwitch.vue";
import ResourceCreateUpdateForm from "@/components/Resources/ResourceCreateUpdateForm.vue";
import ResourceDetailsList from "@/components/Resources/ResourceDetailsList.vue";
import ResourceDeleteForm from "@/components/Resources/ResourceDeleteForm.vue";

const resources = ref<Resource[]>([]);
const selectedResourceKeys = ref<string[]>([]);

onBeforeMount(() => {
  loadData();
});

const loadData = async () => {
  const response = await sendHttpRequest<undefined, undefined, Resource[]>("", "/site/resource-list");
  if (response.status !== ResponseStatus.OK) {
    return;
  }

  resources.value = response.data;
};

const onSelect = (selected: string[]) => {
  selectedResourceKeys.value = selected;
};
</script>

<template>
  <page-layout
    class="resources"
    title="Resources"
  >
    <template #table>
      <data-table
        :columns="[
          { label: 'Key', dataExtractor: (resource) => resource.key, align: 'left', isKey: true },
          { label: 'Title', dataExtractor: (resource) => resource.title, align: 'left' },
          { label: 'Type', dataExtractor: (resource) => resource.type, align: 'left' },
          { label: 'Description', dataExtractor: (resource) => resource.description, align: 'left', multiline: true },
        ]"
        :data="resources"
        :selected="selectedResourceKeys"
        @select="onSelect"
      />
    </template>
    <template #editor>
      <editor-switch
        :editors="[
          { name: 'create', label: 'Create Resource', type: 'global', default: true },
          { name: 'details', label: 'Info', type: 'single', default: true },
          { name: 'delete', label: 'Delete', type: 'single' },
        ]"
        :selected="selectedResourceKeys"
        @select="onSelect"
      >
        <template #create="{ changeTo }">
          <resource-create-update-form
            @save="changeTo(null); loadData();"
          />
        </template>

        <template #details>
          <resource-details-list
            :resource-key="selectedResourceKeys[0]"
          />
        </template>

        <template #delete="{ changeTo }">
          <resource-delete-form
            :resource-key="selectedResourceKeys[0]"
            @delete="changeTo(null); loadData();"
            @cancel="changeTo('details')"
          />
        </template>
      </editor-switch>
    </template>
  </page-layout>
</template>
