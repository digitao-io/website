<script setup lang="ts">
/* eslint-disable @typescript-eslint/no-explicit-any */
type DataExtractor = (rowData: any) => string | number | boolean | null;
type KeyExtractor = (rowData: any) => string;

type ColumnDefinition = {
  label: string;
  dataExtractor: DataExtractor;
  align: "left" | "right";
  multiline?: true;
  isKey?: true;
};

const props = defineProps<{
  columns: ColumnDefinition[];
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  data: any[];
  selected: string[];
}>();

const emit = defineEmits<{
  select: [string[]];
}>();

const keyExtractor = props.columns.find((column) => column.isKey)!.dataExtractor as KeyExtractor;
const onRowClick = (rowData: any) => {
  const key = keyExtractor(rowData);
  emit("select", [key]);
};

/* eslint-enable @typescript-eslint/no-explicit-any */
</script>

<template>
  <div class="data-table">
    <table class="data-table-table">
      <thead>
        <tr class="data-table-head-row">
          <th
            v-for="column of props.columns"
            :key="column.label"
            class="data-table-head-cell"
            :class="{
              'data-table-head-cell__align-left': column.align === 'left',
              'data-table-head-cell__align-right': column.align === 'right',
            }"
          >
            {{ column.label }}
          </th>
        </tr>
      </thead>

      <tbody>
        <tr
          v-for="rowData of data"
          :key="keyExtractor(rowData)"
          class="data-table-body-row"
          :class="{
            'data-table-body-row__selected': props.selected.includes(keyExtractor(rowData))
          }"
          @click="onRowClick(rowData)"
        >
          <td
            v-for="column of props.columns"
            :key="column.label"
            class="data-table-body-cell"
            :class="{
              'data-table-body-cell__multiline': column.multiline,
              'data-table-body-cell__align-left': column.align === 'left',
              'data-table-body-cell__align-right': column.align === 'right',
            }"
          >
            {{ column.dataExtractor(rowData) }}
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
.data-table-table {
  border-collapse: collapse;
  width: 100%;
}

.data-table-table,
.data-table-head-cell,
.data-table-body-cell {
  border: var(--border-xs) solid var(--color-bg-flavor-s);
}

.data-table-head-cell,
.data-table-body-cell {
  padding: var(--padding-s) var(--padding-m);
}

.data-table-head-cell.data-table-head-cell__align-left,
.data-table-body-cell.data-table-body-cell__align-left {
  text-align: left;
}

.data-table-head-cell.data-table-head-cell__align-right,
.data-table-body-cell.data-table-body-cell__align-right {
  text-align: right;
}

.data-table-body-cell {
  white-space: nowrap;
}

.data-table-body-cell__multiline {
  white-space: normal;
}

.data-table-head-row {
  color: var(--color-text-light);
  background-color: var(--color-bg-flavor);
}

.data-table-body-row {
  color: var(--color-text);
  cursor: pointer;
}

.data-table-body-row.data-table-body-row__selected {
  background-color: var(--color-bg-flavor-t);
}
</style>
