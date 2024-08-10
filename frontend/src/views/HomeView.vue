<template>
  <Map id="map" :marker="marker"/>
  <Modal id="modal" :pretitle="`${spot.latitude} // ${spot.longitude}`" :title="spot.title" :dateFrom="dateFrom"
         :dateTo="dateTo">
    <template #content>
      <TripCarousel/>
    </template>
  </Modal>
</template>

<script setup lang="ts">
import {onBeforeMount, ref, type Ref} from 'vue'

import Map from '@/components/Map.vue'
import Modal from '@/components/Modal.vue'
import TripCarousel from '@/components/TripCarousel.vue'
import {useTripStore} from '@/stores/trip'
import type {Spot} from "trippify-client";
import L from "leaflet";


const tripStore = useTripStore()

const spot: Ref<Spot> = ref({} as Spot);
const dateFrom: Ref<Date> = ref(new Date())
const dateTo: Ref<Date | null> = ref(null)

const marker: Ref<L.LatLngTuple[]> = ref([] as L.LatLngTuple[])

onBeforeMount(async () => {
  await tripStore.setTrip('66b3df242162953c2a527d20');
  const firstSpot = tripStore.trip.spots?.[0]
  if (firstSpot) {
    spot.value = await tripStore.getSpot(firstSpot.spotId)
  }

  dateFrom.value = new Date(spot.value.date_from);
  if (spot.value.date_to) {
    dateTo.value = new Date(spot.value.date_to)
  }

  marker.value = tripStore.trip.spots?.map(spot => [spot.latitude, spot.longitude]) ?? [];
});
</script>

<style scoped lang="scss">
#map {
  height: 95dvh;
  width: 100vw;
}

#modal {
  width: 95%;

  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  z-index: 999;

  box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
}

@media screen and (min-width: 1000px) {
  #modal {
    width: 40%;

    position: absolute;
    bottom: 0;
    left: 10%;
    transform: translateX(-20%);
    z-index: 999;

    box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
  }
}
</style>
