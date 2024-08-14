<script setup lang="ts">
type DetailDefinition = {
  label: string;
  data: string | number | boolean | null;
};

const props = defineProps<{
  details: DetailDefinition[];
}>();
</script>

<template>
  <dl class="details-list">
    <template
      v-for="detail of props.details"
      :key="detail.label"
    >
      <dt class="details-list-title">
        {{ detail.label }}
      </dt>
      <dd class="details-list-data">
        {{ detail.data }}
      </dd>
    </template>
  </dl>
</template>

<style scoped>
.details-list-title {
  margin-top: var(--margin-s);
  color: var(--color-text-flavor);
  font-size: var(--font-size-s);
}

.details-list-data {
  margin: 0;
  font-size: var(--font-size-m);
}
</style>
