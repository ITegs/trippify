<template>
  <div id="modal">
    <h1>
      <slot name="title" />
    </h1>
    <h2>
      <slot name="subtitle" />
    </h2>
  </div>
</template>

<script setup lang="ts"></script>

<style scoped lang="scss">
#modal {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding-block: 1rem;
  border-radius: 20px 20px 0 0;
  background-color: var(--color-background);
  width: 100%;

  h1 {
    font-family: Arvo;
    font-size: 1.5rem;
    font-weight: 900;
  }

  h2 {
    font-size: 0.75rem;
  }

  ::v-deep b {
    color: var(--color-primary);
  }
}
</style>
