<template>
  <main>
    <div class="head">
      <div class="back" @click="goBack">
        <font-awesome-icon :icon="['fas', 'chevron-left']"/>
        Zurück
      </div>
      <p>Neuer Spot</p>
    </div>

    <input type="text" v-model="spot.title" placeholder="Wo bist du?" class="location"/>

    <div class="map">
      <SmallMap :position="[spot.latitude, spot.longitude]" :marker="[[spot.latitude, spot.longitude]]"
                :draggable="false" class="mapC"
                v-if="spot.latitude !== 0 && spot.longitude !== 0"/>
      <p v-else>Standort läd...</p>
      <div class="refresh" @click="getLocation">
        <font-awesome-icon :icon="['fas', 'sync']" inverse/>
      </div>
    </div>

    <div class="images">
      <div v-for="image in spot.images" :id="image.timestamp?.toString()" :key="image.timestamp" class="image">
        <span class="removeImg" @click="removeImage($event)">
          <font-awesome-icon :icon="['fas', 'times']"/>
        </span>
        <img :src="image.source" alt="A very nice picture"/>
      </div>
      <span class="addImg" @click="triggerFileInput">
      <font-awesome-icon :icon="['fas', 'plus']" size="2x"/>
      </span>
    </div>
    <input type="file" accept="image/*" multiple @change="onFileChanged($event)" ref="imageUpload" v-show="false"/>

    <div class="description">
      <p>Beschreibung</p>
      <textarea v-model="spot.description" placeholder="Tell us what you did here!"/>
    </div>

    <div class="submit">
      <p class="status">{{ status }}</p>
      <button @click="submit">Hochladen</button>
    </div>

  </main>
</template>

<script setup lang="ts">
import {onMounted, ref} from "vue";
import type {NewSpot} from "trippify-client";
import {useTripStore} from "@/stores/trip";
import router from "@/router";
import SmallMap from "@/components/SmallMap.vue";
import Compressor from 'compressorjs';

/* Lat/Lon accuracy by decimal places:
* 3 decimal places: ~111 meters
* 4 decimal places: ~11.1 meters
* 5 decimal places: ~1.11 meters
*/
const POSITIONING_ACCURACY = 1000 // -> 11.1m

const COMPRESSION_FACTOR = 0.2

const tripStore = useTripStore()

const spot = ref<NewSpot>({
  title: "",
  latitude: 0,
  longitude: 0,
  date_from: new Date().toISOString(),
  description: "",
  images: [],
})

const imageUpload = ref<HTMLInputElement | null>(null)

function triggerFileInput() {
  if (imageUpload.value) {
    imageUpload.value.click()
  }
}

const status = ref<string>(" ")


function goBack() {
  router.back()
}

onMounted(() =>
    getLocation()
)


function getLocation() {
  if (window.navigator.geolocation) {
    console.log("Getting location...")
    window.navigator.geolocation.getCurrentPosition(
        (position: GeolocationPosition) => {
          const lat = Math.round(position.coords.latitude * POSITIONING_ACCURACY) / POSITIONING_ACCURACY
          const lon = Math.round(position.coords.longitude * POSITIONING_ACCURACY) / POSITIONING_ACCURACY

          spot.value.latitude = lat
          spot.value.longitude = lon
        }
    )
  } else {
    alert("Geolocation is not supported")
    console.log("Geolocation is not supported");
  }
}

async function onFileChanged($event: Event) {
  console.log("File uploaded")
  const target = $event.target as HTMLInputElement;
  if (target && target.files) {
    for (let file of target.files) {
      const compressedImage = await compressImage(file, COMPRESSION_FACTOR)
      const reader = new FileReader();
      reader.readAsDataURL(compressedImage);
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

async function compressImage(image: File, quality: number): Promise<File | Blob> {
  return await new Promise((resolve, reject) => {
    new Compressor(image, {
      quality,
      success: resolve,
      error: reject,
    });
  });

}


function get32BitTimestamp() {
  const milliseconds = Date.now();
  let seconds = Math.floor(milliseconds / 1000);

  return seconds & 0xFFFFFFFF;
}

function removeImage($event: Event) { // TODO: fix
  const target = $event.target as HTMLImageElement;
  const timestamp = Number(target.id)
  const newArr = spot.value.images.filter((image) => image.timestamp !== timestamp)
  spot.value.images = newArr
}

function submit() {
  const tripId = tripStore.trip._id

  if (spot.value.title.length > 0 && spot.value.latitude != 0 && spot.value.longitude != 0 && spot.value.images.length > 0) {
    tripStore.addSpotToTrip(tripId, spot.value).then(
        () => window.location.href = "/"
    )
  } else {
    status.value = "Please fill out all fields!"

    setTimeout(() => {
      status.value = " "
    }, 3000)
  }
}
</script>

<style scoped lang="scss">
main {
  width: 100%;
  padding-inline: 1rem;

  .head {
    background-color: var(--color-background);
    width: 100%;
    padding-top: 4rem;
    padding-bottom: 1rem;

    display: grid;
    grid-template-columns: 100px 1fr 100px;
    align-items: center;
    justify-items: center;

    position: fixed;
    z-index: 2;

    .back {
      display: flex;
      flex-direction: row;
      align-items: center;
      gap: 0.5rem;
      justify-self: flex-start;
    }


    p {
      font-family: Arvo;
      font-weight: 700;
    }
  }

  .location {
    margin-top: 5rem;

    position: relative;
    top: 1rem;
    z-index: 1;

    font-family: Arvo;
    font-size: 1.5rem;
    font-weight: 900;

    width: 100%;
    background-color: transparent;
    color: var(--color-accent);
    border: none;
    outline: none;

    &::placeholder {
      color: var(--color-accent);
      opacity: 0.5;
    }

    &::-ms-input-placeholder {
      color: var(--color-accent);
      opacity: 0.5;
    }
  }

  .map {
    height: 25vh;
    background-color: var(--color-primary);
    border-radius: 20px;

    position: relative;

    .mapC {
      //height: 25vh;
      height: 100%;
      z-index: 0;
      border-radius: 20px;
    }

    p {
      position: absolute;
      text-align: center;
      width: 100%;
      top: 50%;
      transform: translateY(-50%);

      font-family: Arvo;
      font-size: 1rem;
      font-weight: 900;
      color: #D9D9D9;

      animation: loading 1.5s infinite linear;
    }

    .refresh {
      position: absolute;
      left: calc(100% - 45px);
      top: calc(100% - 45px);
      z-index: 1;

      display: flex;
      align-items: center;
      justify-content: center;
      background-color: var(--color-text);

      height: 35px;
      border-radius: 50%;
      aspect-ratio: 1;
    }
  }

  .images {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 0.5rem;

    height: 25vh;
    overflow-x: scroll;
    overflow-y: hidden;
    margin-inline: -1rem;
    padding-inline: 1rem;

    margin-block: 1rem;

    .image {
      position: relative;
      height: 100%;

      img {
        height: 100%;
        object-fit: cover;
        border-radius: 20px;
      }

      .removeImg {
        position: absolute;
        top: 0;
        right: 0;
        z-index: 100;

        height: 20px;
        width: 20px;
        background-color: var(--color-primary);
        color: var(--color-background);
        border-radius: 50%;

        display: flex;
        align-items: center;
        justify-content: center;
      }
    }

    .addImg {
      height: 100%;
      aspect-ratio: 2/3;
      border-radius: 20px;
      background-color: #D9D9D9;

      display: flex;
      align-items: center;
      justify-content: center;
    }

    &::-webkit-scrollbar {
      display: none;
    }

    -ms-overflow-style: none; /* IE and Edge */
    scrollbar-width: none; /* Firefox */
  }

  .description {
    width: 100%;
    height: 20vh;
    padding: 1rem;

    background-color: #D9D9D9;
    color: var(--color-text);
    font-family: Lato;
    font-size: 1rem;
    border-radius: 20px;

    p {
      font-weight: 700;
      margin-bottom: 0.5rem;
    }

    textarea {
      border: none;
      outline: none;
      width: 100%;
      height: 80%;

      background-color: #D9D9D9;
      font-family: Lato;

      font-size: 1rem;
      margin: 0;
      padding: 0;

      &::placeholder {
        color: var(--color-text);
        opacity: 0.5;
      }

      &::-ms-input-placeholder {
        color: var(--color-text);
        opacity: 0.5;
      }
    }
  }

  .submit {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;

    margin-block: 1rem;

    .status {
      color: var(--color-accent);
    }

    button {
      background-color: var(--color-primary);
      color: var(--color-background);
      outline: none;
      border: none;
      padding: 0.5rem 1rem;
      border-radius: 10px;
      font-size: 1rem;
      font-family: Lato;
      font-weight: 700;
    }
  }
}

@keyframes loading {
  0% {
    opacity: 1;
  }
  50% {
    opacity: 0.5;
  }
  100% {
    opacity: 1;
  }
}
</style>