<template>
  <div class="container">
    <div
      class="addSpot"
      v-if="userStore.user.username && userStore.user.username === tripStore.trip.owner_username"
      @click="goToAddSpot"
    >
      <font-awesome-icon :icon="['fas', 'plus']" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { useUserStore } from '@/stores/user'
import { useTripStore } from '@/stores/trip'

function goToAddSpot() {
  window.location.href = '/addSpot'
}

const userStore = useUserStore()
const tripStore = useTripStore()
</script>

<style scoped lang="scss">
.container {
  pointer-events: none;

  * {
    pointer-events: all;
  }

  position: relative;
  height: 100%;
  width: 100%;

  .addSpot {
    position: absolute;
    right: 0;
    top: 11dvh;
    margin-right: 32.5px;
    margin-top: 5px;
    height: 35px;

    border-radius: 50%;
    aspect-ratio: 1;

    background-color: var(--color-primary);
    color: var(--color-background);

    display: flex;
    align-items: center;
    justify-content: center;
    box-shadow:
      rgba(0, 0, 0, 0.1) 0 4px 6px -1px,
      rgba(0, 0, 0, 0.06) 0 2px 4px -1px;
  }
}
</style>
