<script setup lang="ts">
import { computed, ref } from "vue";
import Ajv from "ajv";

const props = defineProps<{
  label: string;
  placeholder: string;
  validation: object;
  errorMessage: string;
  password?: boolean;
  disabled?: boolean;
}>();

const model = defineModel<string>();

const interacted = ref(false);
const hasError = computed(() => !(new Ajv()).validate(props.validation, model.value));
</script>

<template>
  <div class="text-input">
    <label class="text-input-label">
      <span class="text-input-label-text">
        {{ props.label }}
      </span>
      <input
        v-model="model"
        class="text-input-input"
        size="1"
        :type="props.password ? 'password' : 'text'"
        :placeholder="placeholder"
        :disabled="!!props.disabled"
        @input="interacted = true"
      >
    </label>
    <span class="text-input-error-message">
      {{ hasError && interacted ? props.errorMessage : "&nbsp;" }}
    </span>
  </div>
</template>

<style scoped>
.text-input {
  display: block;
  width: 100%;
}

.text-input-label {
  display: block;
}

.text-input-label-text {
  display: block;
  color: var(--color-text);
  font-size: var(--font-size-m);
  line-height: var(--line-height-m);
}

.text-input-input {
  display: block;
  border: var(--border-xs) solid var(--color-text);
  padding: var(--padding-s) var(--padding-s);
  width: 100%;
  box-sizing: border-box;
  color: var(--color-text);
  font-size: var(--font-size-m);
  line-height: var(--line-height-m);
  outline: none;
}
.text-input-input:focus {
  border: var(--border-xs) solid var(--color-text-flavor);
}
.text-input-input:disabled {
  color: var(--color-text-t);
  background-color: var(--color-bg-s);
  cursor: not-allowed;
}

.text-input-error-message {
  display: block;
  padding: 0 var(--padding-s);
  height: var(--line-height-m);
  color: var(--color-text-flavor);
  font-size: var(--font-size-m);
  line-height: var(--line-height-m);
}
</style>
