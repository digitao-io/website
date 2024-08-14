<script setup lang="ts">
import { ref } from "vue";
import { RouterLink } from "vue-router";
import { routes } from "@/router";

type MenuEntry = {
  label: string;
  name: string;
};

const menuEntries = ref<MenuEntry[]>(routes
  .filter((route) => route.label)
  .map((route) => ({ label: route.label!, name: route.name as string })));
</script>

<template>
  <nav class="navigation-menu">
    <ul class="navigation-menu-list">
      <li
        v-for="menuEntry of menuEntries"
        :key="menuEntry.name"
        class="navigation-menu-item"
      >
        <router-link
          class="navigation-menu-item-link"
          active-class="navigation-menu-item-link__active"
          :to="{ name: menuEntry.name }"
        >
          {{ menuEntry.label }}
        </router-link>
      </li>
    </ul>
  </nav>
</template>

<style scoped>
.navigation-menu-list {
  padding: 0;
}

.navigation-menu-item-link {
  display: block;
  padding: var(--padding-s) var(--padding-m);
  font-weight: var(--font-weight-m);
  text-decoration: none;
  color: var(--color-text-light);
  background-color: var(--color-bg-flavor);
}
.navigation-menu-item-link:hover {
  border-right: var(--border-l) solid var(--color-bg-flavor-s);
}

.navigation-menu-item-link__active {
  background-color: var(--color-bg-flavor-s);
}
</style>
