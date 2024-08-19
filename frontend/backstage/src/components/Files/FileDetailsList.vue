<script setup lang="ts">
import { computed, ref, watch } from "vue";
import { DateTime } from "luxon";
import { ResponseStatus, sendHttpRequest, type File, type FileIdentifier } from "frontend-resources";
import DetailsList from "@/components/Shared/DetailsList.vue";

const props = defineProps<{
  fileId: string;
}>();

const file = ref<File | null>(null);

const details = computed(() => {
  if (!file.value) {
    return [];
  }

  return [
    { label: "ID", data: file.value.id },
    { label: "Title", data: file.value.title },
    { label: "Type", data: file.value.mimeType },
    { label: "Size", data: (file.value.sizeInBytes / 1024).toFixed(2) + " KiB" },
    { label: "Created At", data: DateTime.fromISO(file.value.createdAt).toLocaleString(DateTime.DATETIME_MED_WITH_SECONDS) },
  ];
});

watch(() => props.fileId, async () => {
  const query = { id: props.fileId };
  const response = await sendHttpRequest<FileIdentifier, undefined, File>("", "/data/file-get", query);
  if (response.status !== ResponseStatus.OK) {
    return;
  }

  file.value = response.data;
}, { immediate: true });
</script>

<template>
  <div class="file-details-list">
    <img
      v-if="file?.mimeType.startsWith('image/')"
      :key="file.id"
      class="file-details-list-image"
      :src="`/data/file-download/${file.id}`"
    >

    <details-list
      :details="details"
    />
  </div>
</template>

<style scoped>
.file-details-list {
  width: 80%;
}

.file-details-list-image {
  display: block;
  width: 100%;
}
</style>
