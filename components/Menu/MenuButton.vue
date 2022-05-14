<template functional>
  <button class="menu-button" @click.native="listeners['click']">
    <svg
      class="menu-button__icon"
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      fill="none"
      viewBox="0 0 24 24"
    >
      <path v-if="props.showMenu" d="M1 23 23 1M1 1l22 22" />
      <path v-else d="M0 3h24M0 21h24M0 12h24" />
    </svg>
  </button>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  props: {
    showMenu: { type: Boolean, required: true }
  }
})
</script>

<style scoped>
.menu-button {
  margin: 0;
  border: none;
  padding: 0;
  background-color: transparent;
}

.menu-button__icon {
  stroke: #242424;
  stroke-width: 2;
}
</style>
