<template>
  <div class="header">
    <a href="/" class="logo">trippify <span class="beta">beta</span></a>

    <div v-if="trip?.owner" class="tripInfo">
      <p>
        <u>{{ trip.owner.name }}</u> ist unterwegs:
      </p>
      <b>{{ trip.title }}</b>
    </div>
    <a href="/account" class="profilePic">
      <img :src="trip.owner.profile_pic" v-if="trip?.owner" alt="Profile picture" />
    </a>
  </div>
</template>

<script setup lang="ts">
import type { Trip } from 'trippify-client/api'

const props = defineProps<{
  trip?: Trip
}>()
</script>

<style scoped lang="scss">
.header {
  display: flex;
  //flex-direction: row;
  align-items: center;
  justify-content: space-between;
  background-color: rgba(#f8fcfa, 0.85);
  position: relative;
  z-index: 900;

  .logo {
    display: flex;
    align-items: center;
    margin-left: 1rem;

    font-family: Arvo;
    color: var(--color-primary);
    font-weight: 900;
    font-size: 1.2rem;

    text-decoration: none;
  }

  .beta {
    font-size: 0.6rem;
    position: relative;
    top: 1dvh;
    font-family: monospace;
    color: var(--color-text);
  }

  .tripInfo {
    display: flex;
    flex-direction: column;
    align-items: center;
    font-size: 0.7rem;

    u {
      color: var(--color-primary);
      text-decoration: none;
    }

    b {
      color: var(--color-accent);
      font-weight: 900;
      font-size: 0.9rem;
    }
  }

  .profilePic {
    position: relative;
    top: 2.5dvh;
    margin-right: 10px;
    height: 80px;

    z-index: 901;
    border-radius: 50%;
    aspect-ratio: 1;
    background-color: var(--color-background);
    overflow: hidden;
    display: flex;
    align-items: center;
    justify-content: center;

    img {
      width: 80%;
      height: 80%;
      border-radius: 50%;
      aspect-ratio: 1;
      object-fit: cover;
    }
  }
}
</style>
