<template>
  <div id="map" />
</template>

<script setup lang="ts">
import * as L from 'leaflet'
import { type LatLngTuple } from 'leaflet'
import { onMounted, watch } from 'vue'

const {
  position,
  zoom = 16,
  marker,
  draggable = true
} = defineProps<{
  position: LatLngTuple
  zoom?: number
  marker?: LatLngTuple[]
  draggable?: boolean
}>()

let map: L.Map

onMounted(() => {
  map = L.map('map').setView(position, zoom)
  L.tileLayer(import.meta.env.VITE_TILE_URL + import.meta.env.VITE_TILE_KEY, {
    maxZoom: 19,
    attribution: import.meta.env.VITE_MAP_ATTRIBUTION
  }).addTo(map)

  if (!draggable) {
    map.dragging.disable()
  }

  map.zoomControl.remove()
  map.attributionControl.remove()

  if (marker) {
    drawMarker(marker)
  }
})

function drawMarker(mark: LatLngTuple[]) {
  const markerOptions: L.CircleMarkerOptions = {
    radius: 8,
    color: '#fb7646',
    weight: 5,
    fill: true,
    fillOpacity: 1
  }

  mark.forEach((m) => {
    const marker = L.circleMarker(m, markerOptions).addTo(map)
  })
}

// Watch the position prop to update map position
watch(
  () => position,
  (newPosition) => {
    if (map) {
      map.flyTo(newPosition) // Fly to the new position when it changes
    }
  },
  { immediate: true }
) // Ensure it runs immediately on mount as well
</script>

<style scoped lang="scss">
#map {
  //height: 100%;
}
</style>
