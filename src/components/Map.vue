<template>
  <div id="map" />
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import L from 'leaflet'

onMounted(() => {
  const map = L.map('map').setView([51.505, -0.09], 13)
  L.tileLayer('https://{s}.basemaps.cartocdn.com/light_all/{z}/{x}/{y}{r}.png', {
    maxZoom: 19,
    attribution: '© OpenStreetMap contributors'
  }).addTo(map)

  map.zoomControl.remove()
  map.attributionControl.setPosition('topleft')
})
</script>

<style scoped lang="scss">
#map {
  height: 95vh;
}
</style>
