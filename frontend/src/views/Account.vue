<script setup lang="ts">
import Logout from "../components/Logout.vue";
import Monitor from "../components/Monitor.vue";
import Info from "../components/Info.vue";
import Setting from "../components/Setting.vue";
import Login from "../components/Login.vue";
</script>

<template>
  <Info />
  <el-row style="margin-top: 20px" :gutter="20" :justify="'center'">
    <el-col :span="4">
      <setting />
    </el-col>
    <el-col :span="4">
      <Monitor />
    </el-col>
    <el-col :span="4">
      <Login />
    </el-col>
    <el-col :span="4">
      <Logout />
    </el-col>
  </el-row>
</template>

<style scoped>

</style>