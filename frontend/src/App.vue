<template>
  <div class="common-layout">
    <el-container>
      <el-header>
        <el-menu
            class="el-menu-demo"
            mode="horizontal"
            :ellipsis="false"
            text-color="#fff"
            active-text-color="#fff"
        >
          <el-menu-item>
            <h1>Cupid Connector</h1>
          </el-menu-item>
          <el-menu-item>
            <setting />
          </el-menu-item>
          <el-menu-item>
            <Login />
          </el-menu-item>
          <el-menu-item>
            <Logout />
          </el-menu-item>
        </el-menu>
      </el-header>
      <el-main style="text-align: center;">
        <Info />
      </el-main>
      <el-footer>
        <div style="text-align: center">
          本项目采用 GPL-3.0 license 进行许可<br/>
          仓库地址：https://github.com/HEX9CF/CupidConnector
        </div>
      </el-footer>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import Setting from "./components/Setting.vue";
import Info from "./components/Info.vue";
import Login from "./components/Login.vue";
import Logout from "./components/Logout.vue";
</script>

<style scoped>
 .el-menu--horizontal > .el-menu-item:nth-child(1) {
   margin-right: auto;
 }
</style>