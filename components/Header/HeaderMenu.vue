<template>
  <nav
    class="header-menu"
    :class="{
      'header-menu--show-menu': showMenu
    }"
  >
    <ul class="header-menu__list">
      <li
        v-for="menuItem in menuItems"
        :key="menuItem.key"
        class="header-menu__list-item"
      >
        {{ menuItem.label }}
      </li>
    </ul>
  </nav>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  props: {
    showMenu: { type: Boolean, required: true }
  },

  data () {
    return {
      menuItems: [
        { key: 'home', label: 'HOME' },
        { key: 'showcase', label: 'SHOWCASE' },
        { key: 'contact', label: 'CONTACT' }
      ]
    }
  }
})
</script>

<style scoped>
.header-menu__list {
  text-align: center;
  padding: 0;
  list-style-type: none;
}

.header-menu__list-item {
  line-height: 2;
  font-size: 20px;
  font-weight: 300;
  color: var(--secondary-color);
}
</style>
