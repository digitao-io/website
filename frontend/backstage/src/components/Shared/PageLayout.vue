<script setup lang="ts">
const props = defineProps<{
  title: string;
}>();
</script>

<template>
  <div class="page-layout">
    <div class="page-layout-data-area">
      <h1 class="page-layout-title">
        {{ props.title }}
      </h1>

      <div class="page-layout-filter-area">
        <slot name="filter" />
      </div>

      <div class="page-layout-table-area">
        <slot name="table" />
      </div>

      <div class="page-layout-pagination-area">
        <slot name="pagination" />
      </div>
    </div>

    <div class="page-layout-editor-area">
      Hello world!
    </div>
  </div>
</template>

<style scoped>
.page-layout {
  display: flex;
  flex-wrap: wrap;
  margin: var(--margin-m);
}

.page-layout-data-area {
  margin-right: var(--margin-l);
  min-width: 50%;
}

.page-layout-title {
  color: var(--color-text-flavor);
  font-weight: var(--font-weight-b);
  font-size: var(--font-size-xxl);
  line-height: var(--line-height-m);
}
</style>
