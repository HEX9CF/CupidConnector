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
import { model } from '../../wailsjs/go/models'
import { ElButton, ElDialog, ElForm, ElSwitch } from "element-plus"
import { GetMonitorConf, UpdateMonitorConf } from '../../wailsjs/go/application/App'
import { Odometer } from "@element-plus/icons-vue";

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