<template>
  <button class="header-menu-toggle-button" @click="handleClick">
    <svg
      class="header-menu-toggle-button__icon"
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      fill="none"
      viewBox="0 0 24 24"
      :class="{button_withe: showMenu}"
    >
      <path v-if="showMenu" d="M1 23 23 1M1 1l22 22" />
      <path v-else d="M0 3h24M0 21h24M0 12h24" />
    </svg>
  </button>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  props: {
    showMenu: { type: Boolean, required: true }
  },

  methods: {
    handleClick () {
      this.$emit('click')
    }
  }
})
</script>

<style scoped>
.header-menu-toggle-button {
  margin: 0;
  border: none;
  padding: 0;
  background-color: transparent;
}

.header-menu-toggle-button__icon {
  stroke: var(--primary-color);
  stroke-width: 2;
}

.button_withe {
  stroke: var(--secondary-color);
  margin-right: 8px;
}
</style>
