<template>
  <div
    id="modal"
    :class="{
      open: isOpen,
      closed: !isOpen
    }"
  >
    <div class="head">
      <i
        class="fas fa-chevron-up"
        @click="isOpen = !isOpen"
        :style="{
          transform: isOpen ? 'rotate(180deg)' : 'rotate(0deg)'
        }"
        style="transition: all 0.3s ease"
      />
      <div class="title">
        <h1>
          <slot name="title" />
        </h1>
        <WeatherChip :temperature="1" condition="sunny" class="weather-chip" />
      </div>

      <h2>
        <slot name="subtitle" />
      </h2>
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
import { ref } from 'vue'
import WeatherChip from './WeatherChip.vue'

var isOpen = ref(false)
</script>

<style scoped lang="scss">
.open {
  height: 65vh;
}

.closed {
  height: 15vh;
}

#modal {
  background-color: var(--color-background);
  width: 100%;
  transition: height 0.5s ease-in-out;
  border-radius: 20px 20px 0 0;
  overflow-y: hidden;

  .head {
    width: 100%;
    height: 12vh;
    display: flex;
    flex-direction: column;
    align-items: center;
    padding-top: 1rem;

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
  }

  .content {
    height: 53vh;
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 2rem 1rem;
    overflow-y: auto;
    transition: all 0.5s ease-in-out;

    :deep(img) {
      width: 80%;
      border-radius: 20px;
    }
  }
}
</style>
