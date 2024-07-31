<script setup lang="ts">
import { ref } from "vue";

const username = ref("");
const password = ref("");

async function loginFetch() {
  const loginUrl = "http://localhost:5173/site/user-login";

  try {
    const response = await fetch(loginUrl, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        username: username.value,
        password: password.value,
      }),
    });

    if (!response.ok) {
      throw new Error("login failed!");
    }
  } catch (error) {
    console.log(error);
  }
}

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
    <button @click="loginFetch">
      Login
    </button>
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
