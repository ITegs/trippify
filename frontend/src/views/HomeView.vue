<template>
  <div class="noscroll">
    <Map
      id="homeMap"
      :marker="marker"
      :currentSpotId="currentSpotId"
      @changedSpot="changeSpot"
      @nextSpot="toNextSpot"
      @prevSpot="toPrevSpot"
    />
    <Modal
      id="modal"
      :pretitle="`${spot.latitude} // ${spot.longitude}`"
      :title="spot.title"
      :dateFrom="dateFrom"
      :dateTo="dateTo"
      :num-pics="spot.images?.length"
    >
      <template #content>
        <TripCarousel :spot="spot" />
        <p>{{ spot.description }}</p>
      </template>
    </Modal>
  </div>
</template>

<script setup lang="ts">
import { onBeforeMount, ref, type Ref, watch } from 'vue'

import Map from '@/components/Map.vue'
import Modal from '@/components/Modal.vue'
import TripCarousel from '@/components/TripCarousel.vue'
import { useTripStore } from '@/stores/trip'
import type { Spot } from 'trippify-client/api'
import { type LatLngTuple } from 'leaflet'
import { useRoute } from 'vue-router'

const route = useRoute()

const tripStore = useTripStore()

const spot: Ref<Spot> = ref({} as Spot)
const dateFrom: Ref<Date> = ref(new Date())
const dateTo: Ref<Date | null> = ref(null)

type Marker = {
  spotId: string
  latLng: LatLngTuple
}

const marker: Ref<Marker[]> = ref([] as Marker[])

const currentSpotId = ref()

onBeforeMount(async () => {
  if (route.params.tripId) {
    await tripStore.setTrip(route.params.tripId as string)
  }
})

watch(tripStore.trip, async () => {
  const initialSpot = tripStore.trip.spots?.[tripStore.trip.spots?.length - 1]
  if (initialSpot) {
    spot.value = await tripStore.fetchSpot(initialSpot.spotId)
  }

  dateFrom.value = new Date(spot.value.date_from)
  if (spot.value.date_to) {
    dateTo.value = new Date(spot.value.date_to!)
  }

  if (tripStore.trip.spots) {
    marker.value = tripStore.trip.spots.map((spot) => {
      return { spotId: spot.spotId, latLng: [spot.latitude, spot.longitude] }
    })
  }

  currentSpotId.value = initialSpot?.spotId
})

async function changeSpot(spotId: string) {
  currentSpotId.value = spotId
  spot.value = await tripStore.fetchSpot(spotId)

  dateFrom.value = new Date(spot.value.date_from)
  if (spot.value.date_to) {
    dateTo.value = new Date(spot.value.date_to!)
  }
}

async function toNextSpot() {
  const currentSpotIndex = tripStore.trip.spots?.findIndex(
    (spot) => spot.spotId === currentSpotId.value
  )
  if (
    currentSpotIndex !== undefined &&
    tripStore.trip.spots &&
    currentSpotIndex < tripStore.trip.spots?.length - 1
  ) {
    await changeSpot(tripStore.trip.spots[currentSpotIndex + 1].spotId)
  }
}

function toPrevSpot() {
  const currentSpotIndex = tripStore.trip.spots?.findIndex(
    (spot) => spot.spotId === currentSpotId.value
  )
  if (currentSpotIndex !== undefined && currentSpotIndex > 0 && tripStore.trip.spots) {
    changeSpot(tripStore.trip.spots[currentSpotIndex - 1].spotId)
  }
}
</script>

<style scoped lang="scss">
.noscroll {
  position: fixed;
  overflow: hidden;
}

#homeMap {
  height: 100dvh;
  width: 100vw;
  z-index: 0;
}

#modal {
  width: 95%;

  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  z-index: 300;

  box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);

  p {
    //text-align: justify;"
    font-family: Arvo;
    font-size: 1rem;
    margin: 1rem;
  }
}

@media screen and (min-width: 1000px) {
  #modal {
    width: 30%;

    position: absolute;
    bottom: 0;
    left: 10%;
    transform: translateX(-20%);
    z-index: 999;

    box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
  }
}
</style>
