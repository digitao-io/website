<script setup lang="ts">
import { ref } from "vue";
import DefaultTemplateMenuButton from "./DefaultTemplateMenuButton.vue";

export type DefaultTemplateConfig = {
  primaryMenuEntries: DefaultTemplateConfigMenuEntry[];
  secondaryMenuEntries: DefaultTemplateConfigMenuEntry[];
};

export type DefaultTemplateConfigMenuEntry = {
  label: string;
  url: string;
};

const props = defineProps<{
  config: DefaultTemplateConfig;
}>();

const menuActivated = ref(false);

const toggleMenuActivated = () => {
  menuActivated.value = !menuActivated.value;
};
</script>

<template>
  <header class="page-header">
    <div class="logo">
      <svg
        class="logo-svg"
        viewBox="0 0 32 42"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path d="M21 21H31V41H1V11H21V21ZM21 21V31H11V21H21Z" />
        <path d="M31 1H21V11H31V1Z" />
      </svg>

      <div class="logo-text">
        digiTAO<br>Software
      </div>
    </div>

    <div class="page-header__menu-container">
      <default-template-menu-button
        :activated="menuActivated"
        @click="toggleMenuActivated"
      />
    </div>

    <!--nav class="page-header__main-menu">
      <p>该部分内容仅作示例，请删除：</p>
      <ul>
        <li
          v-for="menuEntry of props.config.primaryMenuEntries"
          :key="menuEntry.url"
        >
          {{ menuEntry.label }}
        </li>
      </ul>
    </nav-->
  </header>
  <main class="page-content">
    <slot name="main" />
  </main>
  <footer class="page-footer">
    [这里添加版权信息等]
    <nav class="page-header__secondary-menu">
      <p>该部分内容仅作示例，请删除：</p>
      <ul>
        <li
          v-for="menuEntry of props.config.secondaryMenuEntries"
          :key="menuEntry.url"
        >
          {{ menuEntry.label }}
        </li>
      </ul>
    </nav>
  </footer>
</template>

<style scoped>
.page-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.logo {
  display: flex;
  align-items: end;
}
.logo-svg {
  stroke: var(--color-primary);
  width: 32px;
  stroke-width: 2;
}
.logo-text {
  margin-left: 10px;
  font-family: var(--font-raleway);
  font-size: 14px;
  font-weight: bold;
  color: var(--color-primary);
  line-height: 15px;
}
</style>
