<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { sendHttpRequest, ResponseStatus } from "frontend-resources";
import type { UserLoginRequest } from "frontend-resources";
import TextInput from "@/components/Shared/TextInput.vue";
import TextButton from "@/components/Shared/TextButton.vue";

const router = useRouter();

const username = ref("");
const password = ref("");

const hasError = ref(false);

const login = async () => {
  const response = await sendHttpRequest<undefined, UserLoginRequest, undefined>("", "/site/user-login", undefined, {
    username: username.value,
    password: password.value,
  });

  if (response.status !== ResponseStatus.OK) {
    hasError.value = true;
    return;
  }

  router.push({ name: "home" });
};

</script>
<template>
  <div class="login">
    <div class="login-container">
      <h1 class="login-title">
        Login
      </h1>

      <p class="login-error-message">
        {{ hasError ? "Login failed. Please check your username and password!" : "&nbsp;" }}
      </p>

      <text-input
        v-model="username"
        class="login-username-input"
        label="Username"
        placeholder="Please input username"
        :validation="{
          type: 'string',
          minLength: 1,
        }"
        error-message="Username cannot be empty"
      />

      <text-input
        v-model="password"
        class="login-password-input"
        label="Password"
        placeholder="Please input password"
        password
        :validation="{
          type: 'string',
          minLength: 1,
        }"
        error-message="Password cannot be empty"
      />

      <text-button
        class="login-button"
        label="Login"
        type="primary"
        @click="login"
      />
    </div>
  </div>
</template>

<style scoped>
.login {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100%;
}

.login-container {
  width: 320px;
}

.login-title {
  color: var(--color-text-flavor);
  font-weight: var(--font-weight-b);
  font-size: var(--font-size-xxl);
  line-height: var(--line-height-m);
}

.login-error-message {
  margin-top: var(--margin-m);
  color: var(--color-text-flavor);
  font-size: var(--font-size-m);
  line-height: var(--line-height-m);
}

.login-username-input {
  margin-top: var(--margin-s);
}

.login-button {
  margin-top: var(--margin-m);
}
</style>
