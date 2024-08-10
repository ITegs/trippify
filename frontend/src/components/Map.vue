<template>
  <div id="map"/>
</template>

<script setup lang="ts">
import {onMounted, watch} from 'vue'
import L from 'leaflet'

const props = defineProps<{
  marker: L.LatLngTuple[]
}>()

let map: L.Map


onMounted(() => {
  const initialView: L.LatLngExpression = [48.13807, 11.57523]
  map = L.map('map').setView(initialView, 13)
  L.tileLayer(import.meta.env.VITE_TILE_URL + import.meta.env.VITE_TILE_KEY, {
    maxZoom: 19,
    attribution: import.meta.env.VITE_MAP_ATTRIBUTION
  }).addTo(map)

  map.zoomControl.remove()
  map.attributionControl.setPosition('topleft')
})

watch(() => props.marker, (newMarkers) => {
  console.log(newMarkers)

  drawMarker(props.marker, map)
  drawPath(props.marker, map)
  setViewToCenter(props.marker, map)
})


function drawMarker(marker: L.LatLngTuple[], map: L.Map) {
  const markerOptions: L.CircleMarkerOptions = {
    radius: 8,
    color: '#fb7646',
    weight: 5,
    fill: true,
    fillOpacity: 1
  }

  marker.forEach((m) => {
    L.circleMarker(m, markerOptions).addTo(map)
  })
}


function drawPath(marker: L.LatLngTuple[], map: L.Map) {
  const pathOptions: L.PolylineOptions = {
    color: '#fb7646',
    weight: 3,
    dashArray: '10, 10'
  }

  L.polyline(marker, pathOptions).addTo(map)
}


function setViewToCenter(marker: L.LatLngTuple[], map: L.Map) {
  const fitBoundsOptions: L.FitBoundsOptions = {
    padding: [70, 0],
  }

  const bounds = L.latLngBounds(marker);
  map.fitBounds(bounds, fitBoundsOptions);
}
</script>

<style scoped lang="scss">
//
</style>
