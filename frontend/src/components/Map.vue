<template>
  <div>
    <MapOverlay id="mapOverlay"/>
    <div id="map"/>
  </div>
</template>

<script setup lang="ts">
import {onMounted, watch} from 'vue'
import * as L from 'leaflet'
import type {LatLngTuple} from "leaflet";
import 'leaflet-doubletapdrag';
import 'leaflet-doubletapdragzoom';
import MapOverlay from "@/components/MapOverlay.vue";

type Marker = {
  spotId: string,
  latLng: LatLngTuple
}

const props = defineProps<{
  marker: Marker[]
}>()

const emit = defineEmits<{
  (e: 'changedSpot', spotId: string): void
}>()

let map: L.Map

onMounted(() => {
  const initialView: L.LatLngExpression = [48.13807, 11.57523]
  map = L.map('map', {
    // @ts-ignore
    doubleTapDragZoom: 'center',
    doubleTapDragZoomOptions: {
      reverse: false,
    },
  }).setView(initialView, 5)
  L.tileLayer(import.meta.env.VITE_TILE_URL + import.meta.env.VITE_TILE_KEY, {
    maxZoom: 19,
    attribution: import.meta.env.VITE_MAP_ATTRIBUTION
  }).addTo(map)

  map.zoomControl.remove()
  // map.attributionControl.setPosition('topleft') //TODO
})

watch(() => props.marker, (newMarkers) => {
  drawPath(props.marker, map)
  drawMarker(props.marker, map)
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
    paddingBottomRight: [0, 120],
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
#mapOverlay {
  position: fixed;
  background-color: rgba(0, 0, 0, 0);
  z-index: 100;
}


#map {
  z-index: 0;
  width: 100%;
  height: 100%;
}
</style>
