<template>
  <div class="noscroll">
    <Header :trip="tripStore.trip as Trip" :traveler="traveler as User"/>
    <Map id="map" :marker="marker" @changedSpot="changeSpot"/>
    <Modal id="modal" :pretitle="`${spot.latitude} // ${spot.longitude}`" :title="spot.title" :dateFrom="dateFrom"
           :dateTo="dateTo">
      <template #content>
        <TripCarousel :spot="spot"/>
        <p>{{ spot.description }}</p>
      </template>
    </Modal>
  </div>
</template>

<script setup lang="ts">
import {onBeforeMount, ref, type Ref} from 'vue'

import Header from "@/components/Header.vue";
import Map from '@/components/Map.vue'
import Modal from '@/components/Modal.vue'
import TripCarousel from '@/components/TripCarousel.vue'
import {useTripStore} from '@/stores/trip'
import type {Spot, Trip, User} from "trippify-client/api";
import {type LatLngTuple} from "leaflet";
import {useUserStore} from "@/stores/user";

const tripStore = useTripStore()
const userStore = useUserStore()

const spot: Ref<Spot> = ref({} as Spot);
const dateFrom: Ref<Date> = ref(new Date())
const dateTo: Ref<Date | null> = ref(null)

type Marker = {
  spotId: string,
  latLng: LatLngTuple
}


const marker: Ref<Marker[]> = ref([] as Marker[])

const traveler: Ref<User> = ref({} as User)

onBeforeMount(async () => {
  await tripStore.setTrip('66b3df242162953c2a527d20');
  const firstSpot = tripStore.trip.spots?.[tripStore.trip.spots.length - 1]
  if (firstSpot) {
    spot.value = await tripStore.getSpot(firstSpot.spotId)
  }

  dateFrom.value = new Date(spot.value.date_from);
  if (spot.value.date_to) {
    dateTo.value = new Date(spot.value.date_to)
  }

  if (tripStore.trip.spots) {
    marker.value = tripStore.trip.spots.map(spot => {
      return {spotId: spot.spotId, latLng: [spot.latitude, spot.longitude]}
    })
  }


  traveler.value = await userStore.getUser(tripStore.trip.owner);
});

async function changeSpot(spotId: string) {
  spot.value = await tripStore.getSpot(spotId)
}
</script>

<style scoped lang="scss">
.noscroll {
  position: fixed;
  overflow: hidden;
}


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

  p {
    text-align: justify;
    font-size: 1.2rem;
    margin: 1rem;
  }
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
