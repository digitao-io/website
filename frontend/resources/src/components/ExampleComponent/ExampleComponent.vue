<script setup lang="ts">
import { ref } from "vue";

const message = ref("Hello world!");
</script>

<template>
  <p class="example-component">
    {{ message }}
  </p>
</template>

<style scoped>
.example-component {
  font-weight: bold;
}
</style>
