<template>
  <main>
    <div class="back" @click="back">
      <i class="fas fa-arrow-left"/>
      Back
    </div>
    <h1>Add a spot</h1>
    <hr/>
    <input type="text" v-model="spot.title" placeholder="Where are you?" class="location"/>
    <div class="geolocation">
      <button @click="getLocation">Get coordinates</button>
      <p><input type="number" v-model="spot.latitude"/>Lat <input type="number" v-model="spot.longitude"/>Lon</p>
    </div>
    <input type="file" accept="image/*" multiple @change="onFileChanged($event)"
           placeholder="Add images"/>
    <div class="selectedImages">
      <img v-for="image in spot.images" :id="image.timestamp?.toString()" :key="image.timestamp" :src="image.source"
           @dblclick="removeImage($event)">
    </div>

    <textarea v-model="spot.description" placeholder="Tell us what you did here!"/>

    <p class="status">{{ status }}</p>
    <button @click="submit" class="submit">Submit</button>

  </main>
</template>

<script setup lang="ts">
import {ref} from "vue";
import type {NewSpot} from "trippify-client";
import {useTripStore} from "@/stores/trip";

/* Lat/Lon accuracy by decimal places:
* 3 decimal places: ~111 meters
* 4 decimal places: ~11.1 meters
* 5 decimal places: ~1.11 meters
*/
const POSITIONING_ACCURACY = 1000 // -> 11.1m

const tripStore = useTripStore()

const spot = ref<NewSpot>({
  title: "",
  latitude: 0,
  longitude: 0,
  date_from: new Date().toISOString(),
  description: "",
  images: [],
})

const status = ref<string>(" ")

function back() {
  window.history.back()
}


function getLocation() {
  if (navigator.geolocation) {
    console.log("Getting location...")
    navigator.geolocation.getCurrentPosition(
        (position: GeolocationPosition) => {
          const lat = Math.round(position.coords.latitude * POSITIONING_ACCURACY) / POSITIONING_ACCURACY
          const lon = Math.round(position.coords.longitude * POSITIONING_ACCURACY) / POSITIONING_ACCURACY

          spot.value.latitude = lat
          spot.value.longitude = lon
        }
    )
  } else {
    console.log("Geolocation is not supported");
  }
}

function onFileChanged($event: Event) {
  console.log("File uploaded")
  const target = $event.target as HTMLInputElement;
  if (target && target.files) {
    for (let file of target.files) {
      const reader = new FileReader();
      reader.readAsDataURL(file);
      reader.onloadend = () => {
        if (reader.readyState === 2 && reader.result) {
          spot.value.images.push({
            source: reader.result.toString(),
            timestamp: get32BitTimestamp(),
          })
        }
      }
    }
  }
}

function get32BitTimestamp() {
  const milliseconds = Date.now();
  let seconds = Math.floor(milliseconds / 1000);

  return seconds & 0xFFFFFFFF;
}

function removeImage($event: Event) {
  const target = $event.target as HTMLImageElement;
  const timestamp = Number(target.id)
  const newArr = spot.value.images.filter((image) => image.timestamp !== timestamp)
  spot.value.images = newArr
}

function submit() {
  if (spot.value.title.length > 0 && spot.value.latitude != 0 && spot.value.longitude != 0 && spot.value.images.length > 0) {
    tripStore.addSpotToTrip("66dde729394be9748a284296", spot.value).then(
        () => window.location.href = "/"
    )
  } else {
    status.value = "Please fill out all fields!"

    setTimeout(() => {
      status.value = " "
    }, 2000)
  }
}
</script>

<style scoped lang="scss">
main {
  display: flex;
  flex-direction: column;
  align-items: center;

  .back {
    margin-top: 2rem;
    margin-left: 2rem;
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 1rem;
    align-self: flex-start;
  }


  h1 {
    font-family: Arvo;
    margin-top: 2rem;
    font-weight: bolder;
  }

  hr {
    border: 1px solid var(--color-accent);
    width: 100%;
    margin: 1rem;
  }

  .location {
    border: none;
    outline: none;
    font-family: Arvo;
    font-size: 1.5rem;
    font-weight: bold;
    text-align: center;
    margin-block: 1rem;
    background-color: var(--color-background);
  }

  button {
    outline: none;
    background-color: var(--color-background);
    border: 1px solid var(--color-primary);
    padding: 0.5rem 1rem;
    border-radius: 10px;
    font-weight: bold;
    margin: 1rem;
  }

  .geolocation {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    margin-block: 1rem;


    p {
      font-size: 0.8rem;

      input {
        border: none;
        outline: none;
        text-align: right;
        background-color: var(--color-background);
      }
    }
  }


  .selectedImages {
    margin-block: 1rem;
    display: grid;
    grid-template-columns: repeat(2, minmax(200px, 1fr));
    align-items: center;
    gap: 1rem;

    img {
      width: 100%;
      height: auto
    }
  }


  textarea {
    width: 90%;
    height: 20vh;
    margin-block: 1rem;
    border: none;
    outline: none;
    background-color: var(--color-background);
  }

  .status {
    margin-bottom: 1rem;
    color: var(--color-accent);
  }
}
</style>