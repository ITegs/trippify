<template>
  <div>
    <p>{{ temperature }}°</p>
    <i :class="icon" />
  </div>
</template>

<script setup lang="ts">
import { computed, type PropType } from 'vue'

type Condition = 'sunny' | 'cloudy' | 'rain'

const props = defineProps({
  temperature: {
    type: Number,
    required: true
  },
  condition: {
    type: String as PropType<Condition>,
    required: true
  }
})

const icon = computed((): string => {
  switch (props.condition) {
    case 'sunny':
      return 'fas fa-sun'
    case 'cloudy':
      return 'fas fa-cloud'
    case 'rain':
      return 'fas fa-cloud-rain'
    default:
      return 'fas fa-question-square'
  }
})
</script>

<style scoped lang="scss">
div {
  display: flex;
  align-items: center;
  width: fit-content;
  padding: 0.2rem 0.4rem;

  background-color: var(--color-secondary);
  border: 2px solid var(--color-primary);
  border-radius: 20px;

  i {
    padding-left: 0.2rem;
  }
}
</style>
