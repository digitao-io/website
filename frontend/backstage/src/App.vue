<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
</script>

<template>
  <div class="test">
    <div class="sidebar">
      <RouterLink
        class="home-content"
        to="/"
      >
        Home
      </RouterLink>
      <RouterLink to="/files">
        Files
      </RouterLink>
    </div>
  </div>
  <RouterView />
</template>

<style scoped>
.sidebar {
  display: flex;
  flex-direction: column;
  padding: 0;
  position: fixed;
  width: 150px;
}

.content {
  position: fixed;
  top: 52px;
  right: 0;
  bottom:0;
  overflow: scroll;
}

</style>
