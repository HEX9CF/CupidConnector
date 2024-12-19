<template>
  <el-button type="primary" @click="handleClick" :icon="Setting" size="small">设置</el-button>
  <el-dialog v-model="dialogFormVisible" width="400">
    <el-form :model="conf">
      <el-form-item label="URL" required>
        <el-input v-model="conf.base_url" placeholder="https://a.stu.edu.cn"></el-input>
      </el-form-item>
      <el-form-item label="用户" required>
        <el-input v-model="conf.username" placeholder="username"></el-input>
      </el-form-item>
      <el-form-item label="密码" required>
        <el-input v-model="conf.password" placeholder="password" type="password"></el-input>
      </el-form-item>
      <el-form-item label="启动程序自动登录" class="form-btn">
        <el-switch v-model="auto_login"></el-switch>
      </el-form-item>
      <el-form-item label="认证成功自动退出" class="form-btn">
        <el-switch v-model="auto_exit"></el-switch>
      </el-form-item>
      <el-form-item label="自动隐藏程序窗口" class="form-btn">
        <el-switch v-model="auto_hide"></el-switch>
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
import { GetBasicConf, UpdateBasicConf } from '../../wailsjs/go/main/App'
import { Setting } from "@element-plus/icons-vue";

const isLoading = ref(false);
const dialogFormVisible = ref(false)
const auto_login = ref<boolean>(true)
const auto_exit = ref<boolean>(false)
const auto_hide = ref<boolean>(false)

const conf = ref<model.BasicConf>({
  base_url: "",
  username: "",
  password: "",
  auto_login: "TRUE",
  auto_exit: "FALSE",
  auto_hide: "FALSE"
});

const handleClick = () => {
  GetBasicConf().then(res => {
    conf.value = res.data;
    auto_login.value = res.data.auto_login === "TRUE";
    auto_exit.value = res.data.auto_exit === "TRUE";
    auto_hide.value = res.data.auto_hide === "TRUE";
  })
  dialogFormVisible.value = true;
};

const handleConfirm = async () => {
  isLoading.value = true;
  conf.value.auto_login = auto_login.value ? "TRUE" : "FALSE";
  conf.value.auto_exit = auto_exit.value ? "TRUE" : "FALSE";
  conf.value.auto_hide = auto_hide.value ? "TRUE" : "FALSE";
  await UpdateBasicConf(conf.value).then(() => {
    isLoading.value = false;
  });
  dialogFormVisible.value = false;
};
</script>

<style>

</style>