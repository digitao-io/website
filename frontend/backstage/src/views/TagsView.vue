<script setup lang="ts">
import { onBeforeMount, ref } from "vue";
import { ResponseStatus, sendHttpRequest } from "frontend-resources";
import type { Tag } from "frontend-resources";
import PageLayout from "@/components/Shared/PageLayout.vue";
import DataTable from "@/components/Shared/DataTable.vue";
import EditorSwitch from "@/components/Shared/EditorSwitch.vue";
import TagCreateForm from "@/components/Tags/TagCreateForm.vue";
import TagInfo from "@/components/Tags/TagInfo.vue";
import TagDeleteForm from "@/components/Tags/TagDeleteForm.vue";

const tags = ref<Tag[]>([]);
const selectedTagKeys = ref<string[]>([]);

onBeforeMount(() => {
  loadData();
});

const loadData = async () => {
  const response = await sendHttpRequest<undefined, undefined, Tag[]>("", "/data/tag-list");
  if (response.status !== ResponseStatus.OK) {
    return;
  }

  tags.value = response.data;
};

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

    <template #editor>
      <editor-switch
        :editors="[
          { name: 'create', label: 'Create new tag', type: 'global', default: true },
          { name: 'info', label: 'Info', type: 'single', default: true },
          { name: 'delete', label: 'Delete', type: 'single' },
        ]"
        :selected="selectedTagKeys"
        @select="onSelect"
      >
        <template #create="{ changeTo }">
          <tag-create-form
            @create="changeTo(null); loadData();"
          />
        </template>
        <template #info>
          <tag-info
            :tag-key="selectedTagKeys[0]"
          />
        </template>
        <template #delete="{ changeTo }">
          <tag-delete-form
            :tag-key="selectedTagKeys[0]"
            @delete="changeTo(null); loadData();"
            @cancel="changeTo('info')"
          />
        </template>
      </editor-switch>
    </template>
  </page-layout>
</template>

<style scoped>
</style>
