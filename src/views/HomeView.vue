<template>
  <div id="map" />
  <Modal id="modal">
    <template #title>Johannes Trip</template>
    <template #subtitle>Letzter Stop: <b>München Maxvorstadt</b></template>
    <template #content>
      <img
        src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4
  //8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg=="
        alt="Red dot"
      />
    </template>
  </Modal>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import L from 'leaflet'
import Modal from '@/components/Modal.vue'

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

#modal {
  position: absolute;
  bottom: 0;
  z-index: 999;
}
</style>
