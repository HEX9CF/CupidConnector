<template>
      <el-button type="success" @click="handleClick" :icon="Aim" size="small">监控</el-button>
      <el-dialog v-model="dialogFormVisible" width="400">
          <el-form :model="conf" :size="'small'">
            <el-form-item label="流量监控">
              <el-switch v-model="enable"></el-switch>
              <span style="font-size: 12px">
                &nbsp;&nbsp;&nbsp;&nbsp;以下功能仅在启用流量监控时生效
              </span>
            </el-form-item>
            <el-form-item label="监控间隔">
              <el-slider v-model="interval" show-input size="small" max="360" :disabled="!enable" />
              <span style="font-size: 12px">
              单位：分钟，若为0则关闭流量监控
              </span>
            </el-form-item>
              <el-form-item label="告警阈值">
                <el-slider v-model="alert_threshold" show-input size="small" max="100" :disabled="!enable" />
                <span style="font-size: 12px">
                剩余流量百分比，若为0则关闭流量告警
                </span>
              </el-form-item>
            <el-form-item label="注销阈值">
              <el-slider v-model="logout_threshold" show-input size="small" max="100" :disabled="!enable" />
              <span style="font-size: 12px">
              剩余流量百分比，若为0则关闭自动注销
              </span>
            </el-form-item>
            <el-form-item label="断线重连">
              <el-switch v-model="auto_reconnect" :disabled="!enable || logout_threshold != 0"></el-switch>
              <span style="font-size: 12px">
                &nbsp;&nbsp;&nbsp;&nbsp;关闭自动注销才能启用断线重连
              </span>
            </el-form-item>
          </el-form>
          <template #footer>
              <div class="dialog-footer">
                  <el-button @click="dialogFormVisible = false">取消</el-button>
                  <el-button type="primary" @click="handleConfirm" :loading="isLoading">确认</el-button>
              </div>
          </template>
      </el-dialog>
</template>

<script setup lang="ts">
import { ref } from "vue"
import { model } from '../../wailsjs/go/models'
import { ElButton, ElDialog, ElForm, ElSwitch } from "element-plus"
import { GetMonitorConf, UpdateMonitorConf } from '../../wailsjs/go/application/App'
import {Aim} from "@element-plus/icons-vue";
import {log} from "echarts/types/src/util/log";

const isLoading = ref(false);
const dialogFormVisible = ref(false)
const enable = ref<boolean>(false)
const auto_reconnect = ref<boolean>(false)
const interval = ref<number>(0)
const alert_threshold = ref<number>(0)
const logout_threshold = ref<number>(0)

const conf = ref<model.MonitorConf>({
  enable: "FALSE",
  auto_reconnect: "FALSE",
  interval: "0",
  alert_threshold: "0",
  logout_threshold: "0",
});

const handleClick = () => {
    GetMonitorConf().then(res => {
        conf.value = res.data;
        enable.value = res.data.enable === "TRUE";
        interval.value = parseInt(res.data.interval);
        alert_threshold.value = parseInt(res.data.alert_threshold);
        logout_threshold.value = parseInt(res.data.logout_threshold);
        auto_reconnect.value = res.data.auto_reconnect === "TRUE";
    })
    dialogFormVisible.value = true;
};

const handleConfirm = async () => {
  isLoading.value = true;
  conf.value.enable = enable.value ? "TRUE" : "FALSE";
  conf.value.auto_reconnect = (auto_reconnect.value && logout_threshold.value == 0) ? "TRUE" : "FALSE";
  conf.value.interval = interval.value.toString();
  conf.value.alert_threshold = alert_threshold.value.toString();
  conf.value.logout_threshold = logout_threshold.value.toString();
  await UpdateMonitorConf(conf.value).then(() => {
    isLoading.value = false;
  });
  dialogFormVisible.value = false;
};
</script>