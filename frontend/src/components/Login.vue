<script setup lang="ts">
import {ElButton, ElNotification} from "element-plus";
import {Login} from "../../wailsjs/go/main/App";

const handleClick = () => {
  Login().then(res => {
    if (res.code === 1) {
      ElNotification({
        title: "认证成功",
        type: "success"
      });
    } else {
      ElNotification({
        title: "认证失败",
        type: "error",
        message: res.msg
      });
    }
  })
}
</script>

<template>
  <el-button type="warning" @click="handleClick">重新登录</el-button>
</template>

<style scoped>

</style>