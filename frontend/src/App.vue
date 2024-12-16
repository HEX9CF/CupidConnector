<template>
  <div class="common-layout">
    <el-container>
      <el-header>
        <Login />
        <Setting />
      </el-header>
      <el-main style="text-align: center;">
        <Info class="info-margin" />
      </el-main>
      <el-footer></el-footer>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import Setting from "./components/Setting.vue";
import Info from "./components/Info.vue";
import Login from "./components/Login.vue";
</script>

<style scoped>
.info-margin {
  margin-top: 20px; /* 你可以根据需要调整这个值 */
}
</style>