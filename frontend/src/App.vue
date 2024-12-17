<template>
  <div class="common-layout">
      <el-container>
        <el-header>
          <Menu />
        </el-header>
        <el-main style="text-align: center;">
          <Info />
          <el-row style="margin-top: 20px" :gutter="20" :justify="'center'">
            <el-col :span="4">
              <setting />
            </el-col>
            <el-col :span="4">
              <Login />
            </el-col>
            <el-col :span="4">
              <Logout />
            </el-col>
          </el-row>
        </el-main>
        <el-footer style="text-align: center">
          <About />
        </el-footer>
      </el-container>
  </div>
</template>

<script setup lang="ts">
import Setting from "./components/Setting.vue";
import Info from "./components/Info.vue";
import Login from "./components/Login.vue";
import Logout from "./components/Logout.vue";
import Menu from "./components/Menu.vue";
import About from "./components/About.vue";
</script>

<style scoped>

</style>