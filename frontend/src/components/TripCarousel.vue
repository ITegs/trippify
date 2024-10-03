<template>
  <div class="container">
    <swiper class="swiper" :autoHeight="true" :slidesPerView="1" :spaceBetween="30">
      <swiper-slide class="slide" v-for="img in spot.images" :key="img.timestamp">
        <img :src="img.source" alt="A very nice picture!"/>
      </swiper-slide>
    </swiper>
  </div>
</template>

<script setup lang="ts">
import {Swiper, SwiperSlide} from 'swiper/vue'
import 'swiper/scss'
import type {Spot} from "trippify-client";


const props = defineProps<{
  spot: Spot
}>()
</script>

<style>
.slide {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;

  img {
    height: 40dvh;
    width: 100%;
    object-fit: cover;
    border-radius: 0 0 20px 20px;
  }
}
</style>
