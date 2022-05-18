<template>
  <header class="header" :class="{header_white: showMenu}">
    <div class="header__toolbar" :class="{header_toolbar_white: showMenu}">
      <div class="header__logo" :class="{header_logo_white: showMenu}">
        <HeaderLogo class="header__logo-image" :class="{headerLogoWhite: showMenu}" />

        <div class="header__logo-title" :class="{logo_title_white: showMenu}">
          DigiTAO<br>
          Software
        </div>
      </div>
      <HeaderMenuToggleButton :show-menu="showMenu" @click="toggleShowMenu" />
    </div>
    <nav v-show="showMenu" @click="toggleShowMenu">
      <ul class="nemu_font">
        <li>Home</li>
        <li>Resume</li>
        <li>Showcae</li>
        <li>Contact</li>
      </ul>
    </nav>
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

.header_toolbar_white {
  display: flex;
  justify-content: space-between;
}

.header_logo_white {
  display: flex;
  align-items: flex-end;
}

.header_white {
  position: absolute;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background-color: #242424;
}

.headerLogoWhite {
  margin-right: 8px;
  margin-top: 8px;
  margin-left: 8px;
  width: 32px;
  height: 42px;
  stroke: var(--secondary-color);
}

.logo_title_white {
  margin-bottom: -3px;
  color: white;
  line-height: 1.3;
  font-family: Raleway, sans-serif;
  font-weight: bold;
  font-size: 14px;
}

.nemu_font {
  display: flex;
  flex-direction: column;
  justify-content: content;
  align-items: center;
  color: var(--secondary-color);
  list-style: none;
}
</style>
