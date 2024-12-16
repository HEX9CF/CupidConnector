<template>
  <ConfigManager />
  <Info class="info-margin" />
</template>

<script setup lang="ts">
import ConfigManager from "./components/ConfigManager.vue";
import Info from "./components/Info.vue";
</script>

<style scoped>
.info-margin {
  margin-top: 20px; /* 你可以根据需要调整这个值 */
}
</style>