<template>
  <div class="header">
    <h1>trippify</h1>
    <div class="trip">
      <p><u>{{ traveler.name }}</u> ist unterwegs:</p>
      <b>{{ trip.title }}</b>
    </div>
    <div class="profilePic">
      <img :src="traveler.profile_pic" alt="Profile picture"/>
    </div>
  </div>
  <div v-if="userStore.user.username === trip.owner" class="addSpot" @click="addSpot">
    <img src="../assets/icons/plus-solid.svg" alt="plus-solid"/>
  </div>
</template>

<script setup lang="ts">
import type {Trip, User} from "trippify-client";
import {useUserStore} from "@/stores/user";

const props = defineProps<{
  traveler: User
  trip: Trip
}>()

console.log(props.traveler.profile_pic)

const userStore = useUserStore()

userStore.setUser("joe", "")

const addSpot = () => {
//   TODO: go to the add spot view
}
</script>

<style scoped lang="scss">
.header {
  display: flex;
  //flex-direction: row;
  align-items: center;
  justify-content: space-between;
  background-color: transparent;
  height: 5dvh;
  margin-top: 1rem;
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
    aspect-ratio: 1;
    height: 10dvh;
    border-radius: 50%;
    overflow: hidden;
    background-color: var(--color-background);
    margin-right: 10px;
    position: relative;
    top: 2.5dvh;
    z-index: 901;
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

.addSpot {
  position: absolute;
  top: 13dvh;
  right: 2.5dvh;
  aspect-ratio: 1;
  height: 5dvh;
  border-radius: 50%;
  overflow: hidden;
  background-color: var(--color-primary);
  margin-right: 10px;
  z-index: 901;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 6px -1px, rgba(0, 0, 0, 0.06) 0px 2px 4px -1px;

  img {
    width: 60%;
    height: 60%;
    aspect-ratio: 1;
  }
}
</style>
