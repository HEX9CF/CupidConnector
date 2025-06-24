<template>
  <div class="common-layout">
      <el-container>
        <el-header>
          <Menu />
        </el-header>
        <el-main style="text-align: center">
          <el-tabs type="border-card" tab-position="top">
            <el-tab-pane label="账号">
              <Account/>
            </el-tab-pane>
            <el-tab-pane label="监控">
              <Traffic/>
            </el-tab-pane>
            <el-tab-pane label="关于">
              <About />
            </el-tab-pane>
          </el-tabs>
        </el-main>
        <el-footer style="text-align: center">
        </el-footer>
      </el-container>
  </div>
</template>

<script setup lang="ts">
import Menu from "./components/Menu.vue";
import About from "./views/About.vue";
import Account from "./views/Account.vue";
import Traffic from "./views/Traffic.vue";
</script>

<style scoped>

</style>