<template>
  <div id="map" />
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import L, { type CircleMarkerOptions, type LatLngExpression, type PolylineOptions } from 'leaflet'

const marker: LatLngExpression[] = [
  [48.13807, 11.57523],
  [48.2, 11.6],
  [48.1, 11.7],
  [48.3, 11.8],
  [48.4, 11.9]
]

const pathOptions: PolylineOptions = {
  color: '#fb7646',
  weight: 5,
  dashArray: '10, 10'
}

const markerOptions: CircleMarkerOptions = {
  color: '#fb7646',
  weight: 5,
  fill: true,
  fillOpacity: 1
}

onMounted(() => {
  const initialView: L.LatLngExpression = [48.13807, 11.57523]
  const map = L.map('map').setView(initialView, 13)
  L.tileLayer(import.meta.env.VITE_TILE_URL + import.meta.env.VITE_TILE_KEY, {
    maxZoom: 19,
    attribution: import.meta.env.VITE_MAP_ATTRIBUTION
  }).addTo(map)

  map.zoomControl.remove()
  map.attributionControl.setPosition('topleft')

  marker.forEach((m) => {
    L.circleMarker(m, markerOptions).addTo(map)
  })

  const path = L.polyline(marker, pathOptions).addTo(map)
})
</script>

<style scoped lang="scss">
//
</style>
