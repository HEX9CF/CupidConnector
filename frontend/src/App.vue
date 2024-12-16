<template>
  <Setting />
  <Info class="info-margin" />
</template>

<script setup lang="ts">
import Setting from "./components/Setting.vue";
import Info from "./components/Info.vue";
</script>

<style scoped>
.info-margin {
  margin-top: 20px; /* 你可以根据需要调整这个值 */
}
</style>