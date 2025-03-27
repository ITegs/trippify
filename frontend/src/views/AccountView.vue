<script setup lang="ts">
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()

let username = ''
let password = ''

function login(event: Event) {
  event.preventDefault()
  userStore.login(username, password)
  userStore.getUserFromLocalStorage()
}

function logout(event: Event) {
  event.preventDefault()
  userStore.logout()
}
</script>

<template>
  <main>
    <form v-if="!userStore.user.username" @submit="login" class="login">
      <input name="username" type="text" v-model="username" placeholder="Username" />
      <input name="password" type="password" v-model="password" placeholder="Password" />
      <button>Login</button>
    </form>
    <form v-else @submit="logout">
      Logged in as {{ userStore.user.username }}
      <button>Logout</button>
    </form>
  </main>
</template>

<style scoped lang="scss">
main {
  height: 100dvh;
  width: 100%;
  padding-inline: 1rem;
  display: flex;
  align-items: center;
  justify-content: center;

  form {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;

    input {
      border: 3px solid var(--color-secondary);
      border-radius: 10px;
      padding: 0.5rem 1rem;
      font-size: 1.3rem;
      font-weight: 600;
    }

    button {
      background-color: var(--color-primary);
      color: var(--color-text);
      border: none;
      border-radius: 10px;
      padding: 0.5rem 1rem;
      font-size: 1.3rem;
      font-weight: 600;
    }
  }
}
</style>
