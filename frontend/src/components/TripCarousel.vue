<template>
  <div class="container">
    <swiper class="swiper" :autoHeight="true" :slidesPerView="1" :spaceBetween="30">
      <swiper-slide class="slide" v-for="spot in exampleData" :key="spot.id">
        <img :src="spot.img" />
        <p>{{ spot.description }}</p>
      </swiper-slide>
    </swiper>
  </div>
</template>

<script setup lang="ts">
import { Swiper, SwiperSlide } from 'swiper/vue'
import 'swiper/scss'

const exampleData = [
  {
    id: 1,
    img: 'https://trip.jo-dev.de/api/img/2022-08-20T19:04:41.899Z.jpg',
    description: 'Das Jedermannsrecht wurde ausgenutzt! 🌄🏕️'
  },
  {
    id: 2,
    img: 'https://images.unsplash.com/photo-1570641963303-92ce4845ed4c',
    description: 'yeah'
  },
  {
    id: 2,
    img: 'https://images.unsplash.com/photo-1570641963303-92ce4845ed4c',
    description: 'yeah'
  }
]
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

  p {
    text-align: justify;
    font-size: 1.2rem;
    margin: 1rem;
  }
}
</style>
