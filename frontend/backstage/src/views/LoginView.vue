<script setup lang="ts">
import { ref } from "vue";

const username = ref("");
const password = ref("");
</script>
<template>
  <div class="from-table">
    <span>User</span>
    <input
      v-model="username"
      class="username-input"
      placeholder="username"
    >
    <span>Password</span>
    <input
      v-model="password"
      class="password-input"
      type="password"
      placeholder="password"
    >
    <button>Login</button>
  </div>
</template>

<style scoped>
.from-table {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

.username-input,
.password-input {
  margin-bottom: 16px;
}
</style>
