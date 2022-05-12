<template>
  <div>
    <div class="box">
      <div class="logo">logo DigiTAO Software</div>
      <div class="menu">Menu</div>
    </div>
    <div class="content">content</div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({})
</script>

<style>
  .box {
    display: flex;
    justify-content: space-between;
  }

  .menu {
    justify-content: flex-end;
  }
</style>
