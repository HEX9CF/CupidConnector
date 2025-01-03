<script setup lang="ts">
import {ElButton} from "element-plus";
import { CircleClose } from "@element-plus/icons-vue";
import {Logout} from "../../wailsjs/go/application/App";
import {ref} from "vue";

const isLoading = ref(false);

const handleClick = async () => {
  isLoading.value = true;
  await Logout().then(() => {
    isLoading.value = false;
  });
}
</script>

<template>
  <el-button type="danger" @click="handleClick" :loading="isLoading" :icon="CircleClose" size="small">注销</el-button>
</template>

<style scoped>

</style>