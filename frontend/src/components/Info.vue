<template>
  <el-empty v-if="isNoData" style="max-height: 250px">
    <ElButton @click="handleClick" style="margin-top: 10px;" :icon="Refresh" :loading="isLoading">刷新</ElButton>
  </el-empty>
  <div v-if="!isNoData">
    <el-row>
      <el-col :span="8">
        <el-statistic :value="info?.expiration_time">
          <template #title>
            <div style="display: inline-flex; align-items: center">
              <el-icon style="margin-right: 4px" :size="12">
                <Calendar />
              </el-icon>
              过期时间
            </div>
          </template>
        </el-statistic>
      </el-col>
      <el-col :span="8">
        <el-statistic :value="info?.user_name">
          <template #title>
            <div style="display: inline-flex; align-items: center">
              <el-icon style="margin-right: 4px" :size="12">
                <User />
              </el-icon>
              用户名称
            </div>
          </template>
        </el-statistic>
        <ElButton @click="handleClick" style="margin-top: 10px;" :icon="Refresh" :loading="isLoading" size="small">刷新
        </ElButton>
      </el-col>
      <el-col :span="8">
        <el-statistic :value="info?.account_status">
          <template #title>
            <div style="display: inline-flex; align-items: center">
              <el-icon style="margin-right: 4px" :size="12">
                <CircleCheck />
              </el-icon>
              用户状态
            </div>
          </template>
        </el-statistic>
      </el-col>
    </el-row>
      <el-row>
        <el-col :span="8">
          <Flux />
        </el-col>
        <el-col :span="16">
          <Traffic/>
        </el-col>
      </el-row>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { ElCol, ElRow } from 'element-plus'
import { GetInfo, RefreshInfo } from '../../wailsjs/go/application/App'
import { model } from '../../wailsjs/go/models'
import { Refresh, Calendar, User, CircleCheck } from "@element-plus/icons-vue";
import { EventsOn } from "../../wailsjs/runtime";
import Traffic from "./Traffic.vue";
import Flux from "./Flux.vue";

const isLoading = ref(false)
const info = ref<model.Info>()
const isNoData = ref<boolean>(true)

const getInfo = async () => {
  isLoading.value = true
  await GetInfo().then((res) => {
    console.log(res)
    if (res.code === 1) {
      info.value = res.data
      isNoData.value = false;
    } else {
      info.value = {
        user_name: "未知用户",
        overall: 0,
        used: 0,
        expiration_time: "未知",
        account_status: "未知",
      };
      isNoData.value = true;
    }
    if (info.value == null || info.value?.user_name == "") {
      info.value = {
        user_name: "未知用户",
        overall: 0,
        used: 0,
        expiration_time: "未知",
        account_status: "未知",
      };
      isNoData.value = true;
    }
    isLoading.value = false
  })
}

const handleClick = async () => {
  isLoading.value = true
  await RefreshInfo()
  isLoading.value = false
}

const updateInfo = async () => {
  isLoading.value = true
  await getInfo()
  isLoading.value = false
}

onMounted(async () => {
  EventsOn("updateInfo", async () => {
    isLoading.value = true
    await updateInfo();
    isLoading.value = false
  })
  setTimeout(async () => {
    await updateInfo()
  }, 3000)
})
</script>

<style scoped>

</style>