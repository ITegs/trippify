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
    <h1>
      <span v-for="(letter, i) in 'trippify'" :key="i" :style="`animation-delay: -${1 - i * 0.1}s`">
        {{ letter }}
      </span>
    </h1>
    <form v-if="!userStore.user.username" @submit="login" class="login">
      <input name="username" type="text" v-model="username" placeholder="Username" />
      <input name="password" type="password" v-model="password" placeholder="Password" />
      <button>Login</button>
    </form>
    <form v-else @submit="logout">
      Logged in as {{ userStore.user.username }}
      <button>Logout</button>
    </form>
    <p>
      Hier log ich mich ein :)
      <br />
      <br />
      Profilerstellung kommt erst in der v3. Sorry!
      <br />
      Zum Trost eine lustige Katze:
    </p>
    <img src="/cat.gif" alt="Funny cat" />
  </main>
</template>

<style scoped lang="scss">
main {
  height: 100dvh;
  width: 100%;
  padding-inline: 1rem;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;

  h1 {
    font-family: Arvo;
    color: var(--color-primary);
    font-weight: 900;
    font-size: 1.5rem;
    margin-bottom: 2rem;
    border: 2px solid var(--color-secondary);
    padding: 0.5rem 1rem;
    border-radius: 0.2em;
    transform: skewX(10deg);

    span {
      display: inline-block;
      animation: wave 2s ease-in-out infinite;
    }
  }

  p {
    font-family: Arvo;
    text-align: center;
    margin-bottom: 1rem;
  }

  form {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
    margin-bottom: 3rem;

    input {
      border: 2px solid var(--color-secondary);
      border-radius: 10px;
      padding: 0.5rem 1rem;
      font-size: 1rem;
      font-weight: 600;
    }

    button {
      background-color: var(--color-primary);
      color: var(--color-text);
      border: none;
      border-radius: 10px;
      padding: 0.5rem 1rem;
      font-size: 1rem;
      font-weight: 600;
    }
  }

  @keyframes wave {
    0%,
    100% {
      transform: translateY(-3px);
    }
    50% {
      transform: translateY(3px);
    }
  }
}
</style>
