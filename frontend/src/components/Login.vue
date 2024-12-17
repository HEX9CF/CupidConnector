<script setup lang="ts">
import {ElButton} from "element-plus";
import { RefreshRight } from "@element-plus/icons-vue";
import {Login} from "../../wailsjs/go/main/App";
import {ref} from "vue";

const isLoading = ref(false);

const handleClick = async () => {
  isLoading.value = true;
  await Login().then(() => {
    isLoading.value = false;
  });
}
</script>

<template>
  <el-button type="warning" @click="handleClick" :loading="isLoading" :icon="RefreshRight" size="small">重登</el-button>
</template>

<style scoped>

</style>