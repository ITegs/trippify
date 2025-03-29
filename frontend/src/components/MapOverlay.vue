<template>
  <div class="container">
    <div class="leftSidebar">
      <div class="changeView" @click="nextView">
        <font-awesome-icon :icon="viewIconMap[currentView]" />
      </div>
      <div class="nextSpot" @click="() => emit('nextSpot')">
        <font-awesome-icon :icon="['fas', 'arrow-right']" />
      </div>
      <div class="previousSpot" @click="() => emit('prevSpot')">
        <font-awesome-icon :icon="['fas', 'arrow-left']" />
      </div>
    </div>
    <div
      class="addSpot"
      v-if="userStore.user.username && userStore.user.username === tripStore.trip.owner_username"
      @click="goToAddSpot"
    >
      <font-awesome-icon :icon="['fas', 'plus']" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { useUserStore } from '@/stores/user'
import { useTripStore } from '@/stores/trip'
import { type Ref, ref } from 'vue'
import type { LatLngTuple } from 'leaflet'

const userStore = useUserStore()
const tripStore = useTripStore()

const props = defineProps<{
  map: L.Map
}>()

const emit = defineEmits<{
  (e: 'nextSpot'): void
  (e: 'prevSpot'): void
}>()

enum View {
  LAST_THREE,
  CURRENT_SPOT,
  ALL_SPOTS
}

const viewIconMap: Record<View, [string, string]> = {
  [View.LAST_THREE]: ['fas', 'circle-nodes'],
  [View.CURRENT_SPOT]: ['fas', 'map-pin'],
  [View.ALL_SPOTS]: ['fas', 'map']
}

const currentView: Ref<View> = ref<View>(View.LAST_THREE)

function nextView() {
  currentView.value = (currentView.value + 1) % 3

  const spots = tripStore.trip.spots || []
  const lastSpot = spots[spots.length - 1] || { latitude: 0, longitude: 0 }
  const last3Bounds: LatLngTuple[] = spots.slice(-3).map((spot) => [spot.latitude, spot.longitude])
  const allBounds: LatLngTuple[] = spots.map((spot) => [spot.latitude, spot.longitude])

  switch (currentView.value) {
    case View.CURRENT_SPOT:
      props.map.setView([lastSpot.latitude, lastSpot.longitude], 15)
      break
    case View.LAST_THREE:
      props.map.fitBounds(last3Bounds, { paddingBottomRight: [0, 140] })
      break
    case View.ALL_SPOTS:
      props.map.fitBounds(allBounds, { paddingBottomRight: [0, 180] })
      break
  }
}

function goToAddSpot() {
  window.location.href = '/addSpot'
}
</script>

<style scoped lang="scss">
.container {
  pointer-events: none;

  * {
    pointer-events: all;
  }

  position: relative;
  height: 100%;
  width: 100%;

  .leftSidebar {
    position: absolute;
    left: 0;
    top: 7dvh;
    margin-left: 1rem;
    margin-top: 1rem;

    div {
      height: 40px;

      border-radius: 50%;
      aspect-ratio: 1;

      background-color: var(--color-text);
      color: var(--color-background);

      display: flex;
      align-items: center;
      justify-content: center;
      box-shadow:
        rgba(0, 0, 0, 0.1) 0 4px 6px -1px,
        rgba(0, 0, 0, 0.06) 0 2px 4px -1px;

      cursor: pointer;

      margin-bottom: 0.5rem;
    }
  }

  .addSpot {
    position: absolute;
    right: 0;
    top: 11dvh;
    margin-right: 32.5px;
    margin-top: 5px;
    height: 35px;

    border-radius: 50%;
    aspect-ratio: 1;

    background-color: var(--color-primary);
    color: var(--color-background);

    display: flex;
    align-items: center;
    justify-content: center;
    box-shadow:
      rgba(0, 0, 0, 0.1) 0 4px 6px -1px,
      rgba(0, 0, 0, 0.06) 0 2px 4px -1px;

    cursor: pointer;
  }
}
</style>
