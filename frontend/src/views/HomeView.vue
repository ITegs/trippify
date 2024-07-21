<template>
  <Map id="map" />
  <Modal id="modal" :pretitle="pretitle" :title="title" :dateFrom="dateFrom" :dateTo="dateTo">
    <template #content>
      <TripCarousel />
    </template>
  </Modal>
</template>

<script setup lang="ts">
import { ref, type Ref } from 'vue'

import Map from '@/components/Map.vue'
import Modal from '@/components/Modal.vue'
import TripCarousel from '@/components/TripCarousel.vue'

const pretitle: Ref<String> = ref('55°29′N 8°27′E')
const title: Ref<String> = ref('Malmö')
const dateFrom: Ref<Date> = ref(new Date(Date.parse('2024-07-18')))
const dateTo: Ref<Date | null> = ref(null)
// const dateTo: Ref<Date | null> = ref(new Date(Date.parse('2024-07-24')))
</script>

<style scoped lang="scss">
#map {
  height: 95dvh;
  width: 100vw;
}

#modal {
  width: 95%;

  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  z-index: 999;

  box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
}

@media screen and (min-width: 1000px) {
  #modal {
    width: 40%;

    position: absolute;
    bottom: 0;
    left: 10%;
    transform: translateX(-20%);
    z-index: 999;

    box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
  }
}
</style>
