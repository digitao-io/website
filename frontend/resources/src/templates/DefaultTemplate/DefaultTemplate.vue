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
  <div class="page">
    <header
      class="header"
      :class="{
        'header-menu-activated': menuActivated,
      }"
    >
      <div class="toolbar">
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

        <default-template-menu-button
          class="menu-button"
          :activated="menuActivated"
          @click="toggleMenuActivated"
        />
      </div>

      <nav
        class="primary-menu"
        :class="{
          'primary-menu-menu-activated': menuActivated,
        }"
      >
        <ul class="primary-menu-list">
          <li
            v-for="menuEntry of props.config.primaryMenuEntries"
            :key="menuEntry.url"
          >
            <a
              class="primary-menu-link"
              :href="menuEntry.url"
            >{{ menuEntry.label }}</a>
          </li>
        </ul>
      </nav>
    </header>

    <main class="content">
      <slot name="main" />
    </main>

    <footer class="footer">
      <hr>

      <p class="copyright">
        Except where otherwise noted, the website itself is licensed under GNU General Public License v3.0,
        and the content of the website is licensed under Creative Commons Attribution 4.0 International (CC BY 4.0).
      </p>

      <nav class="secondary-menu">
        <ul class="secondary-menu-list">
          <li
            v-for="menuEntry of props.config.secondaryMenuEntries"
            :key="menuEntry.url"
          >
            <a
              class="secondary-menu-link"
              :href="menuEntry.url"
            >{{ menuEntry.label }}</a>
          </li>
        </ul>
      </nav>
    </footer>
  </div>
</template>

<style scoped>
.page {
  display: grid;
  grid-template-columns: auto;
  grid-template-rows: repeat(3, min-content);
  grid-template-areas:
    "header"
    "content"
    "footer";
}

.header {
  grid-area: header;
  padding: 16px;
  color: var(--color-primary);
  background-color: transparent;
}
.content {
  grid-area: content;
  padding: 16px;
}
.footer {
  grid-area: footer;
  padding: 32px 16px;
  color: var(--color-primary-t1);
  font-size: var(--font-size-s);
  line-height: var(--line-height-s);
}

.header-menu-activated {
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  color: var(--color-secondary-s1);
  background-color: var(--color-primary);
}

.toolbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.logo {
  display: flex;
  align-items: end;
}
.logo-svg {
  stroke: currentColor;
  width: 32px;
  stroke-width: 2;
}
.logo-text {
  margin-left: 10px;
  font-family: var(--font-raleway);
  font-size: 14px;
  font-weight: bold;
  line-height: 15px;
}

.primary-menu {
  display: none;
}
.primary-menu-menu-activated {
  display: block;
}

.primary-menu-list {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 48px;
  padding: 0;
  list-style: none;
  font-weight: lighter;
  font-size: 20px;
  line-height: 48px;
}

.primary-menu-link {
  color: inherit;
  text-decoration: none;
  text-transform: uppercase;
  border-bottom: 0;
  transition: 0.3s;
}
.primary-menu-link:hover {
  color: var(--color-secondary-s2);
  border-bottom: 1px solid var(--color-secondary-s2);
}

.secondary-menu-list {
  list-style: none;
  padding: 0;
}
.secondary-menu-list > li {
  display: inline-block;
}
.secondary-menu-list > li:not(:last-child):after {
  display: inline-block;
  margin: 0 10px;
  content: "|";
}

.secondary-menu-link {
  color: var(--color-primary-t1);
}
.secondary-menu-link:hover {
  color: var(--color-primary-t2);
}

@media (min-width: 500px) {
  .page {
    margin: 0 auto;
    max-width: 800px;
  }
  .header {
    padding-top: 36px;
  }

  .toolbar {
    flex-direction: column;
    justify-content: start;
  }
  .logo {
    display: block;
    text-align: right;
  }
  .logo-svg {
    width: 96px;
    shape-rendering: crispEdges;
    stroke-width: 1.6;
  }
  .logo-text {
    margin: 0;
    text-align: right;
    font-size: 24px;
    line-height: 26px;
  }

  .menu-button {
    display: none;
  }

  .primary-menu {
    display: block;
  }
  .primary-menu-list {
    flex-direction: row;
    justify-content: center;
    flex-wrap: wrap;
    margin: 36px 0;
    font-size: 16px;
    line-height: 24px;
  }
  .primary-menu-list > li {
    margin: 0 16px;
  }

  .primary-menu-link:hover {
    color: var(--color-primary-t2);
    border-bottom: 1px solid var(--color-primary-t2);
  }
}

@media (min-width: 1000px) {
  .page {
    grid-template-columns: 200px 1fr;
    grid-template-rows: 1fr min-content;
    grid-template-areas:
      "header content"
      "header footer";
    margin: 0 auto;
    max-width: 1000px;
  }

  .header {
    display: flex;
    flex-direction: column;
    align-items: start;
  }
  .content {
    margin-top: 32px;
  }

  .toolbar {
    align-items: end;
  }
  .primary-menu-list {
    flex-direction: column;
    align-items: end;
    margin: 36px 0;
    font-size: 20px;
    line-height: 48px;
  }
  .primary-menu-list > li {
    margin: 0;
  }
}
</style>
