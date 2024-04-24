<template>
  <Map id="map" />
  <Modal id="modal">
    <template #title>{{ title }}</template>
    <template #subtitle>
      <span v-html="subtitle" />
    </template>
    <template #content>
      <TripCarousel />
    </template>
  </Modal>
</template>

<script setup lang="ts">
import { ref } from 'vue'

import Map from '@/components/Map.vue'
import Modal from '@/components/Modal.vue'
import TripCarousel from '@/components/TripCarousel.vue'

const title = ref('Johannes Trip')
const subtitle = ref('to the <b>moon</b>! &#128640;')
</script>

<style scoped lang="scss">
#map {
  height: 95dvh;
  width: 100vw;
}

#modal {
  position: absolute;
  bottom: 0;
  z-index: 999;
}
</style>
