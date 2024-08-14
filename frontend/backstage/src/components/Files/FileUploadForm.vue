<script setup lang="ts">
import axios from "axios";
import { computed, ref, shallowRef } from "vue";
import { ResponseStatus, sendHttpRequest } from "frontend-resources";
import type { FileCreateResponse, FileData } from "frontend-resources";
import TextButton from "@/components/Shared/TextButton.vue";
import TextInput from "@/components/Shared/TextInput.vue";
import DetailsList from "@/components/Shared/DetailsList.vue";

const emit = defineEmits<{
  create: [];
}>();

const file = shallowRef<File | null>(null);
const title = ref<string>("");

const mimeType = computed(() => file.value?.type ?? null);
const sizeInBytes = computed(() => file.value?.size ?? null);
const fileDetails = computed(() => [
  { label: "Type", data: mimeType.value ?? "N/A" },
  { label: "Size", data: ((sizeInBytes.value ?? 0) / 1024)?.toFixed(2) + " KiB" },
]);

const onFileSelectClick = () => {
  if (file.value !== null) {
    file.value = null;
    return;
  }

  const inputElement = document.createElement("input");
  inputElement.type = "file";
  inputElement.onchange = () => {
    file.value = inputElement.files?.[0] ?? null;
  };
  inputElement.click();
};

const onFileSelectDrop = (event: DragEvent) => {
  event.preventDefault();

  const droppedFile = event.dataTransfer?.files?.[0];
  if (!droppedFile) {
    return;
  }

  file.value = droppedFile;
};

const onFileSelectDragover = (event: DragEvent) => {
  event.preventDefault();
};

const onUploadClick = async () => {
  const body = {
    title: title.value,
    mimeType: mimeType.value!,
    sizeInBytes: sizeInBytes.value!,
  };

  const fileCreateRes = await sendHttpRequest<undefined, FileData, FileCreateResponse>("", "/data/file-create", undefined, body);
  if (fileCreateRes.status !== ResponseStatus.OK) {
    return;
  }

  const form = new FormData();
  form.append("file", file.value!);

  await axios({
    method: "POST",
    url: fileCreateRes.data.fileUploadUrl,
    data: form,
  });

  emit("create");
};
</script>

<template>
  <div class="file-upload-form">
    <text-input
      v-model="title"
      label="Title"
      placeholder="A human-readable title for the file"
      :validation="{
        type: 'string',
        minLength: 1,
      }"
      error-message="Title cannot be empty"
    />

    <button
      class="file-upload-form-drop-area"
      :class="{
        'file-upload-form-drop-area--has-file': file !== null,
      }"
      @dragover="onFileSelectDragover"
      @drop="onFileSelectDrop"
      @click="onFileSelectClick"
    >
      {{ file === null ? "Drop file here, or click to select a file." : "Drop a new file, or click to remove file." }}
    </button>

    <details-list
      :details="fileDetails"
    />

    <text-button
      class="file-upload-button"
      label="Upload"
      type="primary"
      @click="onUploadClick"
    />
  </div>
</template>

<style scoped>
.file-upload-form {
  width: 80%;
}

.file-upload-form-drop-area {
  display: block;
  border: var(--border-xs) dashed var(--color-text);
  width: 100%;
  height: 240px;
  color: var(--color-text);
  background-color: var(--color-bg);
  font-size: var(--font-size-l);
  cursor: pointer;
}

.file-upload-form-drop-area.file-upload-form-drop-area--has-file {
  border: var(--border-xs) dashed var(--color-text-flavor);
  color: var(--color-text-flavor);
  background-color: var(--color-bg-flavor-t);
}

.file-upload-button {
  margin-top: var(--margin-m);
}
</style>
