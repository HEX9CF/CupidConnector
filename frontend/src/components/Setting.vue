<template>
      <el-button type="primary" @click="handleClick" :icon="Setting">设置</el-button>
      <el-dialog v-model="dialogFormVisible" title="设置" width="500">
          <el-form :model="conf">
            <el-form-item label="URL" label-width="right">
              <el-input v-model="conf.base_url" placeholder="http://a.stu.edu.cn"></el-input>
            </el-form-item>
              <el-form-item label="用户" label-width="right">
                  <el-input v-model="conf.username" placeholder="username"></el-input>
              </el-form-item>
              <el-form-item label="密码" label-width="right">
                  <el-input v-model="conf.password" placeholder="password" type="password"></el-input>
              </el-form-item>
            <el-form-item label="启动程序自动登录" label-width="right">
              <el-switch v-model="auto_login"></el-switch>
            </el-form-item>
              <el-form-item label="连接成功自动退出" label-width="right">
                  <el-switch v-model="auto_exit"></el-switch>
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
import { ElButton, ElDialog, ElForm, ElInput, ElSwitch } from "element-plus"
import { GetConf, UpdateConf } from '../../wailsjs/go/main/App'
import {Setting} from "@element-plus/icons-vue";

const isLoading = ref(false);
const dialogFormVisible = ref(false)
const auto_login = ref<boolean>(true)
const auto_exit = ref<boolean>(false)

const conf = ref<model.Conf>({
    base_url: "",
    username: "",
    password: "",
    auto_login: "TRUE",
    auto_exit: "FALSE",
});

const handleClick = () => {
    GetConf().then(res => {
        conf.value = res.data;
        auto_login.value = res.data.auto_login === "TRUE";
        auto_exit.value = res.data.auto_exit === "TRUE";
    })
    dialogFormVisible.value = true;
};

const handleConfirm = async () => {
  isLoading.value = true;
  conf.value.auto_login = auto_login.value ? "TRUE" : "FALSE";
  conf.value.auto_exit = auto_exit.value ? "TRUE" : "FALSE";
  await UpdateConf(conf.value).then(() => {
    isLoading.value = false;
  });
  dialogFormVisible.value = false;
};
</script>