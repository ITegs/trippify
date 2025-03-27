<template>
  <Header :trip="trip" id="header" />
  <RouterView />
</template>

<script setup lang="ts">
import { useTripStore } from '@/stores/trip'
import Header from '@/components/Header.vue'
import { storeToRefs } from 'pinia'
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()
userStore.getUserFromLocalStorage()

const tripStore = useTripStore()
tripStore.init()
const { trip } = storeToRefs(tripStore)
</script>

<style scoped lang="scss">
#header {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 7dvh;
}
</style>
