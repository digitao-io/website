<script setup lang="ts">
const slotName = "main";
</script>

<template>
  <default-template :config="{ primaryMenuEntries: [], secondaryMenuEntries: [] }">
    <template #[slotName]>
      <markdown-component :config="{ markdown: '# Hello World\nThis is a hello world from somewhere ...' }" />
    </template>
  </default-template>
</template>

<style scoped>
.message {
  color: green;
}
</style>
