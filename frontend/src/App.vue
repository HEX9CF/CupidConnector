<template>
  <div class="common-layout">
      <el-container>
        <el-header>
          <Menu />
        </el-header>
        <el-main style="text-align: center">
          <Info />
          <el-row style="margin-top: 20px" :gutter="20" :justify="'center'">
            <el-col :span="4">
              <setting />
            </el-col>
            <el-col :span="4">
              <Monitor />
            </el-col>
            <el-col :span="4">
              <traffic />
            </el-col>
            <el-col :span="4">
              <Login />
            </el-col>
            <el-col :span="4">
              <Logout />
            </el-col>
          </el-row>
        </el-main>
<!--
        <el-footer style="text-align: center">
        </el-footer>
-->
      </el-container>
  </div>
</template>

<script setup lang="ts">
import Menu from "./components/Menu.vue";
import Monitor from "./components/Monitor.vue";
import Logout from "./components/Logout.vue";
import Setting from "./components/Setting.vue";
import Login from "./components/Login.vue";
import Info from "./components/Info.vue";
import Traffic from "./components/Traffic.vue";
</script>

<style scoped>

</style>