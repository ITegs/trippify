<template>
  <div
    id="modal"
    :style="heightMode === 'absolute' ? { top: absYPos + 'px' } : { top: relYPos + 'dvh' }"
    :class="{
      closed: modalState === ModalState.closed,
      half: modalState === ModalState.half,
      full: modalState === ModalState.full,
      animate: heightMode === 'relative'
    }"
  >
    <div
      id="head"
      :style="{
        boxShadow: modalState !== ModalState.closed ? '0 0 10px 0 rgba(0, 0, 0, 0.1)' : 'none'
      }"
    >
      <div id="drag-section" v-drag="dragHandler">
        <span class="pull-line" />
      </div>

      <div class="title">
        <p>{{ pretitle }}</p>

        <h1>{{ title }}</h1>

        <div v-if="duration !== 0" style="justify-self: right" class="dateLabel">
          <p v-if="!dateTo">{{ dateFrom.toLocaleDateString('de-DE') }} - Heute</p>
          <p v-else>
            {{ dateFrom.toLocaleDateString('de-DE') }} - {{ dateTo.toLocaleDateString('de-DE') }}
          </p>
        </div>
        <div v-else style="justify-self: right" class="dateLabel">
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

      <div
        class="numPics"
        :class="modalState !== ModalState.closed ? 'hidden' : ''"
        v-show="numPics"
      >
        <p v-if="numPics !== 1">{{ numPics }} Fotos</p>
        <p v-else>{{ numPics }} Foto</p>
        <font-awesome-icon :icon="['fas', 'chevron-down']" />
      </div>
    </div>

    <div
      class="content"
      :style="{
        paddingTop: !modalState ? '10vh' : '0',
        opacity: modalState ? '1' : '0',
        height: windowHeight * FULL_MODAL_PERCENT_HEIGHT - 100 + 'px'
      }"
    >
      <slot name="content" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { onUpdated, ref, type Ref } from 'vue'
import { useDrag } from '@vueuse/gesture'

const CLOSED_MODAL_PERCENT_HEIGHT = 0.2
const HALF_MODAL_PERCENT_HEIGHT = 0.55
const FULL_MODAL_PERCENT_HEIGHT = 0.9

const props = defineProps<{
  pretitle: String
  title: String
  dateFrom: Date
  dateTo: Date | null
  numPics: Number
}>()

const today = new Date()
const duration = ref(0)

onUpdated(() => {
  if (props.dateTo) {
    duration.value = Math.floor(
      (props.dateTo.getTime() - props.dateFrom.getTime()) / (1000 * 60 * 60 * 24)
    )
  } else {
    duration.value = Math.floor(
      (today.getTime() - props.dateFrom.getTime()) / (1000 * 60 * 60 * 24)
    )
  }
})

enum ModalState {
  'closed',
  'half',
  'full'
}

const windowHeight = window.screen.height
const heightMode = ref<'absolute' | 'relative'>('relative')
const relYPos = ref(100 - CLOSED_MODAL_PERCENT_HEIGHT * 100)
const absYPos = ref(windowHeight * (1 - CLOSED_MODAL_PERCENT_HEIGHT))
let lastY = absYPos.value

const modalState: Ref<ModalState> = ref(ModalState.closed)

const dragHandler = ({ movement: [, y], dragging }: any) => {
  if (dragging) {
    heightMode.value = 'absolute'
    if (absYPos.value + y > 40 && absYPos.value + y < windowHeight - 60) {
      absYPos.value = y + lastY
    }
  } else {
    heightMode.value = 'relative'
    lastY = absYPos.value

    switch (modalState.value) {
      case ModalState.closed:
        if (lastY < windowHeight / 2) {
          fullModal()
        } else if (lastY > windowHeight - 250) {
          closeModal()
        } else {
          halfModal()
        }
        break

      case ModalState.half:
        if (lastY < windowHeight / 2 - 150) {
          fullModal()
        } else if (lastY > windowHeight / 2 - 50) {
          closeModal()
        } else {
          halfModal()
        }
        break

      case ModalState.full:
        if (lastY > windowHeight / 2 - 150) {
          closeModal()
        } else if (lastY > windowHeight / 2 - 250) {
          halfModal()
        } else {
          fullModal()
        }
        break
    }
  }
}

function closeModal() {
  modalState.value = ModalState.closed
  relYPos.value = 100 - CLOSED_MODAL_PERCENT_HEIGHT * 100
  absYPos.value = windowHeight * (1 - CLOSED_MODAL_PERCENT_HEIGHT)
  lastY = absYPos.value
}

function halfModal() {
  modalState.value = ModalState.half
  relYPos.value = 100 - HALF_MODAL_PERCENT_HEIGHT * 100
  absYPos.value = windowHeight * (1 - HALF_MODAL_PERCENT_HEIGHT)
  lastY = absYPos.value
}

function fullModal() {
  modalState.value = ModalState.full
  relYPos.value = 100 - FULL_MODAL_PERCENT_HEIGHT * 100
  absYPos.value = windowHeight * (1 - FULL_MODAL_PERCENT_HEIGHT)
  lastY = absYPos.value
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
.closed {
}

.half {
  .content {
    :deep p {
      display: none;
    }
  }
}

.full,
.half {
  #head {
    .title {
      grid-template-rows: 1fr !important;
      //padding-bottom: 0.5rem;

      p {
        display: none;
      }

      .dateLabel {
        display: none;
      }
    }
  }
}

.full {
  .content {
    :deep(img) {
      height: 100% !important;
    }

    :deep(.swiper-wrapper) {
      height: 100% !important;
    }
  }
}

#modal {
  background-color: var(--color-background);
  border-radius: 20px 20px 0 0;
  overflow-y: hidden;
  position: relative;

  &.animate {
    transition: top 0.2s ease-out;
  }

  #head {
    padding-bottom: 1rem;

    #drag-section {
      width: 100%;
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
        overflow: hidden;
        text-overflow: ellipsis;
      }

      h1 {
        color: var(--color-text);
        font-family: Arvo;
        font-weight: 600;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
        font-size: 1.5rem;
      }

      h2 {
        color: var(--color-primary);
        font-family: Arvo;
        font-size: 1.2rem;
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

    .numPics {
      position: relative;
      display: flex;
      flex-direction: column;
      align-items: center;
      opacity: 1;

      p {
        font-style: italic;
      }

      &.hidden {
        height: 0;
        opacity: 0;
        //transform: translateY(-30px);
        transition: all 0.2s ease-in-out;
      }
    }
  }

  .content {
    padding-block: 2rem;
    overflow-y: auto;
    transition: all 0.5s ease-in-out;
  }
}
</style>
