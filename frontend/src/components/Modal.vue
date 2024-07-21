<template>
  <div
    id="modal"
    :class="{
      open: isOpen,
      closed: !isOpen
    }"
  >
    <div
      id="head"
      v-drag="dragHandler"
      :style="{
        boxShadow: isOpen ? '0 0 10px 0 rgba(0, 0, 0, 0.1)' : 'none'
      }"
    >
      <span class="pull-line" />

      <div class="title">
        <p>{{ pretitle }}</p>

        <h1>{{ title }}</h1>

        <div v-if="duration !== 0" style="justify-self: right">
          <p v-if="!dateTo">{{ dateFrom.toLocaleDateString('de-DE') }} - Heute</p>
          <p v-else>
            {{ dateFrom.toLocaleDateString('de-DE') }} - {{ dateTo.toLocaleDateString('de-DE') }}
          </p>
        </div>
        <div v-else style="justify-self: right">
          <p>Am</p>
        </div>

        <div v-if="duration !== 0" style="justify-self: right">
          <h2 v-if="duration === 1">1 Tag</h2>
          <h2 v-else>{{ duration }} Tage</h2>
        </div>
        <div v-else style="justify-self: right">
          <h2>{{ dateFrom.toLocaleDateString('de-DE') }}</h2>
        </div>
      </div>

      <h2 id="isCurrent" v-if="!dateTo">Aktueller Ort</h2>
    </div>

    <div
      class="content"
      :style="{
        paddingTop: !isOpen ? '10vh' : '0',
        opacity: isOpen ? '1' : '0'
      }"
    >
      <slot name="content" />
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  pretitle: String
  title: String
  dateFrom: Date
  dateTo: Date | null
}>()

import { ref } from 'vue'
import WeatherChip from './WeatherChip.vue'
import { useDrag } from '@vueuse/gesture'

const today = new Date()
let duration = 0
if (props.dateTo) {
  duration = Math.floor((props.dateTo.getTime() - props.dateFrom.getTime()) / (1000 * 60 * 60 * 24))
} else {
  duration = Math.floor((today.getTime() - props.dateFrom.getTime()) / (1000 * 60 * 60 * 24))
}

var isOpen = ref(false)

function openModal() {
  isOpen.value = true
}

function closeModal() {
  isOpen.value = false
}

const dragHandler = ({ movement: [, y], dragging }: any) => {
  if (dragging) {
    if (y < -50) {
      openModal()
    } else if (y > 50) {
      closeModal()
    }
  }
}

useDrag(dragHandler, {
  filterTaps: true,
  preventWindowScrollY: true,
  useTouch: true,
  axis: 'y',
  domTarget: document.getElementById('head')
})
</script>

<style scoped lang="scss">
.open {
  height: calc(100dvh - 5dvh - 200px);
}

.closed {
  height: 18dvh;
}

#modal {
  background-color: var(--color-background);
  transition: height 0.5s ease-in-out;
  border-radius: 20px 20px 0 0;
  overflow-y: hidden;

  #head {
    width: 100%;
    // height: 16vh;
    display: flex;
    flex-direction: column;
    align-items: center;
    padding-top: 1rem;
    padding-bottom: 0.5rem;

    .pull-line {
      height: 4px;
      width: 60px;
      border-radius: 5px;
      background-color: #dbdbdb;
    }

    .title {
      width: 100%;
      display: grid;
      grid-template-columns: 1.5fr 1fr;
      grid-template-rows: 1fr 1fr;
      grid-auto-flow: column;
      align-items: end;
      padding-inline: 1rem;

      p {
        color: var(--color-text-light);
        font-size: 0.8rem;
        white-space: nowrap;
      }

      h1 {
        color: var(--color-text);
        font-family: Arvo;
        font-weight: 600;
        white-space: nowrap;
        overflow: scroll;
        font-size: 1.7rem;
      }

      h2 {
        color: var(--color-primary);
        font-family: Arvo;
        font-size: 1.5rem;
      }
    }

    #isCurrent {
      color: var(--color-primary);
      font-size: 0.75rem;
      align-self: flex-start;
      padding-left: 1rem;
      position: relative;
      top: -0.2rem;
    }
  }

  .content {
    height: 53vh;
    padding-block: 2rem;
    overflow-y: auto;
    transition: all 0.5s ease-in-out;
  }
}
</style>
