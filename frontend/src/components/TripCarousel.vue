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
    description:
      'Die Wanderung beginnt am späten Nachmittag, als wir uns auf den Weg machen, um den Gipfel des Berges zu erreichen, rechtzeitig zum Sonnenuntergang. Mit jedem Schritt auf dem steinigen Pfad steigt die Vorfreude, das Naturschauspiel aus der besten Perspektive zu erleben. Als wir den Gipfel erreichen, breitet sich vor uns ein atemberaubender Anblick aus: Der Himmel färbt sich in leuchtenden Orangetönen, die sich im Kontrast zum tiefblauen Horizont abzeichnen. Die letzten Sonnenstrahlen tauchen die umliegende Landschaft in ein goldenes Licht, und die Stille auf dem Gipfel lässt uns den Moment in völliger Ruhe genießen. Während die Sonne langsam hinter den fernen Bergen versinkt, fühlen wir uns erfüllt und dankbar für die Schönheit der Natur, die uns für einen kurzen Moment alles um uns herum vergessen lässt.'
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
    margin: 20px;
    text-align: justify;
    font-size: 1.2rem;
    font-family: 'Arvo';
  }
}
</style>
