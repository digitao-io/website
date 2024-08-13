<script setup lang="ts">
import { RouterView } from "vue-router";
import NavigationMenu from "@/components/App/NavigationMenu.vue";
</script>

<template>
  <div class="app">
    <header class="app-header">
      <div class="app-title">
        Backstage
      </div>
      <navigation-menu />
    </header>

    <main class="app-content">
      <router-view />
    </main>
  </div>
</template>

<style scoped>
.app {
  display: grid;
  grid-template-columns: 150px 1fr;
  grid-template-rows: 1fr;
  grid-template-areas:
    "header content";

  min-height: 100vh;
}

.app-header {
  grid-area: header;
  color: var(--color-text-light);
  background-color: var(--color-bg-dark);
}

.app-title {
  padding: var(--padding-m);
  font-weight: var(--font-weight-b);
  font-size: var(--font-size-l);
  line-height: var(--line-height-m);
}

.app-content {
  grid-area: content;
}
</style>
