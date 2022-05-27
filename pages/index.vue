<template>
  <div class="page">
    <Header />

    <main class="content">
      Content: {{ foo }}
    </main>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  data () {
    return {
      foo: ''
    }
  },

  async fetch () {
    this.foo = await Promise.resolve('hello world')
  }
})
</script>

<style scoped>
.page {
  margin: 16px;
}
</style>
