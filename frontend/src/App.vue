<template>
  <Header :trip="trip" id="header"/>
  <RouterView/>
</template>

<script setup lang="ts">
import {useTripStore} from "@/stores/trip";
import {onBeforeMount, ref} from "vue";
import Header from "@/components/Header.vue";
import type {Trip} from "trippify-client";

const tripStore = useTripStore()

const trip = ref({} as Trip)


onBeforeMount(async () => {
  await tripStore.setTrip('66dde729394be9748a284296');

  trip.value = tripStore.trip
})
</script>

<style scoped lang="scss">
#header {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 7dvh;

}
</style>
