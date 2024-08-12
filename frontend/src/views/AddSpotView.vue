<template>
  <main>
    <h1>Add a spot</h1>
    <button @click="getLocation">Get coordinates</button>
    <p>{{ geolocation?.coords.latitude }} {{ geolocation?.coords.latitude }}</p>
    <input type="text" v-model="location" placeholder="Enter location"/>
    <input type="file" accept="image/*" multiple @change="onFileChanged($event)"
           placeholder="Add images"/>
    <div v-for="image in images" :key="image.name">
      {{ image.name }}
    </div>
  </main>
</template>

<script setup lang="ts">
type Image = {
  name: string
  size: number
  lastModified: number
  base64: string
}

import {ref} from "vue";

const geolocation = ref<GeolocationPosition>()
const location = ref<string>("")
const images = ref<Image[]>()

function getLocation() {
  if (navigator.geolocation) {
    navigator.geolocation.getCurrentPosition((position: GeolocationPosition) => geolocation.value = position);
  }
}

function onFileChanged($event: Event) {
  const target = $event.target as HTMLInputElement;
  if (target && target.files) {
    for (let file of target.files) {
      const imgBase64 = blobToBase64(file)
      images.value?.push({
        name: file.name,
        size: file.size,
        lastModified: file.lastModified,
        base64: imgBase64
      })
    }
  }
}

function blobToBase64(file: Blob) {
  const reader = new FileReader();
  reader.readAsDataURL(file);
  reader.onload = () => {
    return reader.result;
  }
  return ""
}

</script>

<style scoped lang="scss">

</style>