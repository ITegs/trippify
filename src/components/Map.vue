<template>
  <div id="map" />
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import L from 'leaflet'

onMounted(() => {
  const initialView: L.LatLngExpression = [48.13807, 11.57523]
  const map = L.map('map').setView(initialView, 13)
  L.tileLayer('https://tile.openstreetmap.de/{z}/{x}/{y}.png', {
    maxZoom: 19,
    attribution: '© OpenStreetMap contributors'
  }).addTo(map)

  map.zoomControl.remove()
  map.attributionControl.setPosition('topleft')
})
</script>

<style scoped lang="scss">
//
</style>
