<template>
  <div id="map"/>
</template>

<script setup lang="ts">
import {onMounted, watch} from 'vue'
import L, {type LatLngTuple} from 'leaflet'

type Marker = {
  spotId: string,
  latLng: LatLngTuple
}

const props = defineProps<{
  marker: Marker[]
}>()

const emit = defineEmits<{
  changedSpot: [string]
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
  drawMarker(props.marker, map)
  drawPath(props.marker, map)
  setViewToCenter(props.marker, map)
})


function drawMarker(marker: Marker[], map: L.Map) {
  const markerOptions: L.CircleMarkerOptions = {
    radius: 8,
    color: '#fb7646',
    weight: 5,
    fill: true,
    fillOpacity: 1
  }

  marker.forEach((m) => {
    const marker = L.circleMarker(m.latLng, markerOptions).addTo(map)
    marker.on("click", () => handleMarkerClick(m.spotId))
  })
}


function drawPath(marker: Marker[], map: L.Map) {
  const pathOptions: L.PolylineOptions = {
    color: '#fb7646',
    weight: 3,
    dashArray: '10, 10'
  }

  const latLngs: LatLngTuple[] = marker.map((m) => {
    return m.latLng
  })

  L.polyline(latLngs, pathOptions).addTo(map)
}


function setViewToCenter(marker: Marker[], map: L.Map) {
  const fitBoundsOptions: L.FitBoundsOptions = {
    padding: [70, 0],
  }

  const latLngs: LatLngTuple[] = marker.map((m) => {
    return m.latLng
  })

  const bounds = L.latLngBounds(latLngs);
  map.fitBounds(bounds, fitBoundsOptions);
}

function handleMarkerClick(spotId: string) {
  emit('changedSpot', spotId)
}
</script>

<style scoped lang="scss">
//
</style>
