<script setup lang="ts">
import {ElButton} from "element-plus";
import {Logout} from "../../wailsjs/go/main/App";

const handleClick = () => {
  Logout();
}
</script>

<template>
  <el-button type="danger" @click="handleClick">注销账号</el-button>
</template>

<style scoped>

</style>