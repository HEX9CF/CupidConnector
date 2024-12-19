<script setup lang="ts">
import { SwitchButton, Minus } from "@element-plus/icons-vue";
import { Exit } from "../../wailsjs/go/main/App";
import { WindowHide } from "../../wailsjs/runtime/runtime";

const handleClick = () => {
  Exit();
}
</script>

<template>
  <el-menu class="el-menu-demo" mode="horizontal" :ellipsis="false" text-color="#fff" active-text-color="#fff">
    <el-menu-item>
      <h1>Cupid Connector</h1>
    </el-menu-item>
    <div class="button-container">
      <el-button class="minus-button" text :icon="Minus" @click="WindowHide"></el-button>
      <el-button class="exit-button" text :icon="SwitchButton" @click="handleClick"></el-button>
    </div>
  </el-menu>
</template>

<style scoped>
.el-menu--horizontal{
  --wails-draggable:drag
}
.el-menu--horizontal>.el-menu-item:nth-child(1) {
  margin-right: auto;
}

.button-container {
  display: flex;
  align-items: center;
}

.minus-button,
.exit-button {
  width: 60px;
  height: 100%;
  padding: 0; /* 移除默认的内边距 */
  margin: 0; /* 移除默认的外边距 */
}
</style>