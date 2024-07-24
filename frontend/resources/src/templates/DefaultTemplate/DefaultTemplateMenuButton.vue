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
  background-color: transparent;
  font-size: 7px;
  cursor: pointer;
  padding: 3em 1.5em;
}

.menu-button > div {
  width: 4em;
  height: 0.3em;
  background-color: var(--color-secondary);
}

.menu-button > div::before,
.menu-button > div::after {
  transition: font-size 0s;
  content: "";
  display: block;
  position: absolute;
  width: 4em;
  height: 0.3em;
  background-color: var(--color-primary);
}
.menu-button > div::before {
  margin-top: -1.5em;
}
.menu-button > div::after {
  margin-top: 1.5em;
}

.menu-button > div {
  transition: 0.4s;
}

.menu-button > div::before,
.menu-button > div::after {
  transition: 0.4s;
  transform-origin: center center;
}

.menu-button-activated > div::before {
  animation: stripe-top-activated 0.4s linear forwards;
}
.menu-button-activated > div {
  background-color: transparent;
}
.menu-button-activated > div::after {
  animation: stripe-bottom-activated 0.4s linear forwards;
}

.menu-button-deactivated > div::before {
  animation: stripe-top-deactivated 0.4s linear forwards;
}
.menu-button-deactivated > div {
  background-color: var(--color-primary);
}
.menu-button-deactivated > div::after {
  animation: stripe-bottom-deactivated 0.4s linear forwards;
}

@keyframes stripe-top-activated {
  0% {
  }
  20% {
    margin-top: 0em;
    transform: rotate(0deg);
  }
  60% {
    margin-top: 0em;
    transform: rotate(55deg);
  }
  100% {
    margin-top: 0em;
    transform: rotate(45deg);
  }
}

@keyframes stripe-bottom-activated {
  0% {
  }
  20% {
    margin-top: 0em;
    transform: rotate(0deg);
  }
  60% {
    margin-top: 0em;
    transform: rotate(-55deg);
  }
  100% {
    margin-top: 0em;
    transform: rotate(-45deg);
  }
}

@keyframes stripe-top-deactivated {
  0% {
    margin-top: 0em;
    transform: rotate(45deg);
  }
  20% {
    transform: rotate(0deg);
  }
  60% {
    margin-top: 1.7em;
    transform: rotate(0deg);
  }
  100% {
    margin-top: 1.5em;
    transform: rotate(0deg);
  }
}

@keyframes stripe-bottom-deactivated {
  0% {
    margin-top: 0em;
    transform: rotate(-45deg);
  }
  20% {
    transform: rotate(0deg);
  }
  60% {
    margin-top: -1.7em;
    transform: rotate(0deg);
  }
  100% {
    margin-top: -1.5em;
    transform: rotate(0deg);
  }
}
</style>
