<template>
  <header class="header">
    <div class="header__toolbar">
      <div class="header__logo">
        <HeaderLogo class="header__logo-image" />

        <div class="header__logo-title">
          DigiTAO<br>
          Software
        </div>
      </div>

      <HeaderMenuToggleButton :show-menu="showMenu" @click="toggleShowMenu" />
    </div>

    <nav v-show="showMenu" class="header__navigation">
      <ul>
        <li>Home</li>
      </ul>
    </nav>
  </header>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  data () {
    return {
      showMenu: false,
      isShow: false
    }
  },

  methods: {
    toggleShowMenu () {
      this.showMenu = !this.showMenu
    }
  }
})
</script>

<style scoped>
.header__toolbar {
  display: flex;
  justify-content: space-between;
}

.header__logo {
  display: flex;
  align-items: flex-end;
}

.header__logo-image {
  margin-right: 8px;
  width: 32px;
  height: 42px;
  stroke: var(--primary-color);
}

.header__logo-title {
  margin-bottom: -3px;
  color: var(--primary-color);
  line-height: 1.3;
  font-family: Raleway, sans-serif;
  font-weight: bold;
  font-size: 14px;
}
</style>
