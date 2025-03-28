import './assets/main.scss'
import 'leaflet/dist/leaflet.css'

import { createApp } from 'vue'
import { GesturePlugin } from '@vueuse/gesture'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

//icons
import {
  faChevronDown,
  faChevronLeft,
  faCircleNodes,
  faMap,
  faMapPin,
  faPlus,
  faSync,
  faTimes
} from '@fortawesome/free-solid-svg-icons'

library.add(faPlus, faChevronDown, faChevronLeft, faSync, faTimes, faMap, faMapPin, faCircleNodes)

const app = createApp(App).component('font-awesome-icon', FontAwesomeIcon)

app.use(GesturePlugin)

app.use(createPinia())
app.use(router)

app.mount('#app')
