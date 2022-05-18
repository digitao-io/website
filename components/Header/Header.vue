<template>
  <header class="header">
    <div
      class="header__content"
      :class="{
        'header__content--show-menu': showMenu
      }"
    >
      <div class="header__toolbar">
        <div class="header__logo">
          <HeaderLogo class="header__logo-image" />

          <div class="header__logo-title">
            DigiTAO<br>
            Software
          </div>
        </div>

        <HeaderMenuToggleButton
          class="header__menu-toggle-button"
          :show-menu="showMenu"
          @click="toggleShowMenu"
        />
      </div>

      <HeaderMenu
        class="header__menu"
        :show-menu="showMenu"
      />
    </div>

    <div v-show="showMenu" class="header__spacer" />
  </header>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  data () {
    return {
      showMenu: false
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

.header__menu {
  display: none;
  margin-top: 48px;
}

.header__spacer {
  height: 42px;
}

.header__content.header__content--show-menu {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  padding: 16px;
  background-color: var(--primary-color);
}

.header__content.header__content--show-menu .header__logo-image {
  stroke: var(--secondary-color);
}

.header__content.header__content--show-menu .header__logo-title {
  color: var(--secondary-color);
}

.header__content.header__content--show-menu .header__menu {
  display: block;
}
</style>
