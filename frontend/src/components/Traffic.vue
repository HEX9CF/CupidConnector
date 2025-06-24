<template>
      <el-button type="info" @click="handleClick" :icon="Odometer" size="small">流量</el-button>
      <el-dialog v-model="dialogFormVisible" width="400">
          <div id="traffic-chart" style="width: 100%; height: 300px;"></div>
          <template #footer>
              <div class="dialog-footer">
                  <el-button @click="dialogFormVisible = false">返回</el-button>
              </div>
          </template>
      </el-dialog>
</template>

<script setup lang="ts">
import { ref } from "vue"
import { ElButton, ElDialog, ElForm, ElSwitch } from "element-plus"
import { Odometer } from "@element-plus/icons-vue";

const dialogFormVisible = ref(false)

const handleClick = () => {
    dialogFormVisible.value = true;
};
</script>