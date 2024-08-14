<script setup lang="ts">
const props = defineProps<{
  title: string;
}>();
</script>

<template>
  <div class="page-layout">
    <h1 class="page-layout-title">
      {{ props.title }}
    </h1>

    <div class="page-layout-container">
      <div class="page-layout-data-area">
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
        <slot name="editor" />
      </div>
    </div>
  </div>
</template>

<style scoped>
.page-layout {
  margin: var(--margin-m);
}

.page-layout-container {
  display: flex;
  flex-wrap: wrap;
}

.page-layout-data-area {
  flex: 3 3;
  margin-right: var(--margin-l);
}

.page-layout-editor-area {
  flex: 2 2;
}

.page-layout-title {
  color: var(--color-text-flavor);
  font-weight: var(--font-weight-b);
  font-size: var(--font-size-xxl);
  line-height: var(--line-height-m);
}
</style>
