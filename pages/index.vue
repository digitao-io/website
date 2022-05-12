<template>
  <div class="page">
    <header class="header">
      <div class="title">
        <div class="logo">
          Logo
        </div>

        <div class="text">
          DigiTAO Software
        </div>
      </div>

      <div class="menu">
        Menu
      </div>
    </header>

    <main class="content">
      Content
    </main>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({})
</script>

<style scoped>
  .header {
    display: flex;
    justify-content: space-between;
  }

  .title {
    display: flex;
  }
</style>
