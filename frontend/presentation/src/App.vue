<script setup lang="ts">
import { useSSRContext } from "vue";

const ctx = import.meta.env.SSR
  ? useSSRContext()
  : window.context;
const pageDetails = ctx.pageDetails;
const slots = Object.keys(pageDetails.slots);
const componentDetailsArrayBySlot = ctx.pageDetails.slots;
</script>

<template>
  <component
    :is="pageDetails.template"
    :config="pageDetails.config"
  >
    <template
      v-for="slot of slots"
      #[slot]
    >
      <template
        v-for="componentDetails, i of componentDetailsArrayBySlot[slot]"
        :key="`${slot}-${componentDetails.component}-${i}`"
      >
        <component
          :is="componentDetails.component"
          :config="componentDetails.config"
        />
      </template>
    </template>
  </component>
</template>

<style scoped>
</style>
