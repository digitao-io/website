<template>
  <div class="page">
    <header class="header">
      <div class="title">
        <Logo class="logo" />

        <div class="text">
          DigiTAO <br> Software
        </div>
      </div>

      <div class="menu">
        <button v-on:click="show">
          Menu
        </button>
      </div>
    </header>

    <main class="content">
      Content
      {{isenum}}
    </main>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  data () {
    return {
      isenum: true
    }
  },

  methods: {
    show () {
      if (!this.isenum) {
        this.isenum = true
      } else {
        this.isenum = false
      }
    }
  }
})
</script>

<style scoped>
@font-face {
  font-family: "Raleway";
  src: url("/fonts/Raleway.ttf") format("truetype");
}

.header {
  display: flex;
  justify-content: space-between;
}

.title {
  display: flex;
}

.logo {
  width: 26px;
  height: 34px;
  stroke: #242424;
  margin-right: 8px;
  stroke-width: 8.25;
  shape-rendering: crispEdges;
}

.text {
  display: flex;
  align-items: flex-end;
  font-family: Raleway, sans-serif;
  font-weight: bold;
  font-size: 14px;
}
</style>
