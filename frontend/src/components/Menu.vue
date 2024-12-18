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
    <el-menu-item>
      <el-button link :icon="Minus" @click="WindowHide"></el-button>
    </el-menu-item>
    <el-menu-item>
      <el-button link :icon="SwitchButton" @click="handleClick"></el-button>
    </el-menu-item>
  </el-menu>
</template>

<style scoped>
.el-menu--horizontal>.el-menu-item:nth-child(1) {
  margin-right: auto;
}
</style>