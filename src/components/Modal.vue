<template>
  <div
    id="modal"
    :class="{
      open: isOpen,
      closed: !isOpen
    }"
  >
    <div class="title">
      <h1>
        <slot name="title" />
      </h1>
      <WeatherChip :temperature="1" condition="sunny" class="weather-chip" />
    </div>
    <h2>
      <slot name="subtitle" />
    </h2>

    <i
      class="fas fa-chevron-up"
      @click="isOpen = !isOpen"
      :style="{
        transform: isOpen ? 'rotate(180deg)' : 'rotate(0deg)'
      }"
      style="transition: all 0.3s ease"
    />

    <slot v-if="isOpen" name="content" />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import WeatherChip from './WeatherChip.vue'

var isOpen = ref(false)
</script>

<style scoped lang="scss">
#modal {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-block: 1rem;
  border-radius: 20px 20px 0 0;
  background-color: var(--color-background);
  width: 100%;

  transition: height 0.5s ease-in-out;
  overflow-y: hidden;

  .title {
    width: 100%;
    display: flex;

    h1 {
      position: relative;
      left: 50%;
      transform: translateX(-50%);
      font-family: Arvo;
      font-size: 1.5rem;
      font-weight: 900;
    }

    .weather-chip {
      margin-left: auto;
      margin-right: 1rem;
    }
  }

  h2 {
    font-size: 0.75rem;
  }

  :deep(b) {
    color: var(--color-primary);
  }

  i {
    padding: 1rem;
  }
}

.open {
  height: 70vh;
}

.closed {
  height: 8rem;
}
</style>
