<script setup lang="ts">
const props = defineProps<{
  label: string;
  type: "primary" | "secondary";
}>();
</script>

<template>
  <button
    class="text-button"
    :class="{
      'text-button--primary': props.type === 'primary',
      'text-button--secondary': props.type === 'secondary',
    }"
  >
    {{ props.label }}
  </button>
</template>

<style scoped>
.text-button {
  position: relative;
  display: inline-block;
  border: var(--border-xs) solid var(--color-bg-flavor);
  padding: var(--padding-s) var(--padding-m);
  font-weight: var(--font-weight-b);
  font-size: var(--font-size-m);
  line-height: var(--line-height-m);
  cursor: pointer;
}
.text-button:active {
  top: 2px;
}

.text-button.text-button--primary {
  color: var(--color-text-light);
  background-color: var(--color-bg-flavor);
}
.text-button.text-button--primary:hover {
  border: var(--border-xs) solid var(--color-bg-flavor-s);
  background-color: var(--color-bg-flavor-s);
}

.text-button.text-button--secondary {
  color: var(--color-text-flavor);
  background-color: var(--color-bg);
}
.text-button.text-button--secondary:hover {
  background-color: var(--color-bg-flavor-t);
}
</style>
