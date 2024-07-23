<script setup lang="ts">
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
</script>

<template>
  <header class="page-header">
    <div class="logo-area">
      <div class="page-header__logo">
        logo
      </div>

      <div class="page-header__logo-text">
        digiTAO <br> Software
      </div>
    </div>

    <div class="page-header__menu-container">
      [加上菜单按钮与菜单等内容]
      <nav class="page-header__main-menu">
        <p>该部分内容仅作示例，请删除：</p>
        <ul>
          <li
            v-for="menuEntry of props.config.primaryMenuEntries"
            :key="menuEntry.url"
          >
            {{ menuEntry.label }}
          </li>
        </ul>
      </nav>
    </div>
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
  justify-content: space-between;
}
.logo-area {
  display: flex;
}
.page-header__logo {
  width: 50px;
  height: 60px;
}
.page-header__logo-text {
  text-align: left;
  font-family: var(--font-raleway);
  font-size: 14px;
  font-weight: bold;
  color: var(--color-primary);
  line-height: 15px;
}
</style>
