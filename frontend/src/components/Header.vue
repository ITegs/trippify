<template>
  <div class="header">
    <h1>trippify</h1>
    <div class="trip">
      <p>
        <u>{{ trip.owner.name }}</u> ist unterwegs:
      </p>
      <b>{{ trip.title }}</b>
    </div>
    <div class="profilePic">
      <img :src="trip.owner.profile_pic" alt="Profile picture"/>
    </div>
  </div>
</template>

<script setup lang="ts">
import type {Trip} from 'trippify-client/api'

const props = defineProps<{
  trip: Trip
}>()

</script>

<style scoped lang="scss">
.header {
  display: flex;
  //flex-direction: row;
  align-items: center;
  justify-content: space-between;
  background: rgba(255, 255, 255, 0.9);
  position: relative;
  z-index: 900;

  h1 {
    display: flex;
    align-items: center;
    margin-left: 1rem;

    font-family: Arvo;
    color: var(--color-primary);
    font-weight: 900;
    font-size: 1.2rem;
  }

  .trip {
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
