<script setup lang="ts">
import { useSSRContext } from "vue";

import { StaticConfigResolver } from "./resolving/static-config-resolver";

const ctx = useSSRContext();

const resolver = new StaticConfigResolver(ctx!.page);
await resolver.resolve();

const template = resolver.getTempalte();
const slots = resolver.getSlots();
const getComponents = (slot: string) => resolver.getComponents(slot);
</script>

<template>
  <component
    :is="template.template"
    :config="template.staticConfig"
  >
    <template
      v-for="slot of slots"
      #[slot]
    >
      <template
        v-for="component, i of getComponents(slot)"
        :key="`${slot}-${component.component}-${i}`"
      >
        <component
          :is="component.component"
          :config="component.staticConfig"
        />
      </template>
    </template>
  </component>
</template>

<style scoped>
</style>
