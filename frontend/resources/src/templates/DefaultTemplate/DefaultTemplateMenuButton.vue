<script setup lang="ts">
import { ref } from "vue";

const props = defineProps<{
  activated: boolean;
}>();

const initialState = ref(true);
</script>

<template>
  <button
    class="menu-button"
    :class="{
      'menu-button-activated': props.activated && !initialState,
      'menu-button-deactivated': !props.activated && !initialState,
    }"
    @click="initialState = false"
  >
    <div />
  </button>
</template>

<style scoped>
.menu-button {
  display: block;
  margin: 0 -1.2em;
  border: 0;
  color: inherit;
  background-color: transparent;
  font-size: 7px;
  cursor: pointer;
  padding: 3em 1.5em;
}

.menu-button > div {
  width: 4em;
  height: 0.3em;
  background-color: currentColor;
}

.menu-button > div::before,
.menu-button > div::after {
  content: "";
  display: block;
  position: absolute;
  width: 4em;
  height: 0.3em;
  background-color: currentColor;
}
.menu-button > div::before {
  margin-top: -1.5em;
}
.menu-button > div::after {
  margin-top: 1.5em;
}

.menu-button > div::before,
.menu-button > div::after {
  transform-origin: center center;
}

.menu-button-activated > div::before {
  margin-top: 0em;
  transform: rotate(45deg);
}
.menu-button-activated > div {
  background-color: transparent;
}
.menu-button-activated > div::after {
  margin-top: 0em;
  transform: rotate(-45deg);
}
</style>
