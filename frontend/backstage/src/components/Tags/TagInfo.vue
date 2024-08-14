<script setup lang="ts">
const props = defineProps<{
  tagKey: string;
}>();
</script>

<template>
  <p class="tag-info">
    You have selected tag <strong>{{ props.tagKey }}</strong>.
  </p>
</template>

<style scoped>
.tag-info {
  font-size: var(--font-size-l);
  line-height: var(--line-height-m);
}

.tag-info strong {
  color: var(--color-text-flavor);
}
</style>
