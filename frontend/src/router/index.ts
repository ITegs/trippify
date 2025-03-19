import {createRouter, createWebHistory} from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import AddSpotView from "@/views/AddSpotView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/:pathMatch(.*)*',
      redirect: '/'
    },
    {
      path: '/addSpot',
      name: 'addSpot',
      component: AddSpotView
    },
    {
      path: '/trip/:tripId',
      name: 'trip',
      component: HomeView
    }
  ]
})

export default router
