<template>
    <main>
        <ElButton @click="handleClick">配置信息</ElButton>
        <el-dialog v-model="dialogFormVisible" title="配置信息" width="500">
            <el-form :model="conf">
                <el-form-item label="用户" label-width="right">
                    <el-input v-model="conf.username" placeholder="username"></el-input>
                </el-form-item>
                <el-form-item label="密码" label-width="right">
                    <el-input v-model="conf.password" placeholder="password" type="password"></el-input>
                </el-form-item>
                <el-form-item label="连接后自动关闭" label-width="right">
                    <el-switch v-model="auto_exit"></el-switch>
                </el-form-item>
            </el-form>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="dialogFormVisible = false">Cancel</el-button>
                    <el-button type="primary" @click="handleConfirm">
                        Confirm
                    </el-button>
                </div>
            </template>
        </el-dialog>
    </main>
</template>

<script setup lang="ts">
import { ref } from "vue"
import { model } from '../../wailsjs/go/models'
import { ElButton, ElDialog, ElForm, ElInput, ElNotification, ElSwitch } from "element-plus"
import { GetConf, UpdateConf } from '../../wailsjs/go/main/App'

const dialogFormVisible = ref(false)
const auto_exit = ref<boolean>(false)

const conf = ref<model.Conf>({
    username: "",
    password: "",
    auto_exit: "FALSE",
    url: ""
});

const handleClick = () => {
    GetConf().then(res => {
        conf.value = res.data;
        auto_exit.value = res.data.auto_exit === "TRUE";
    })
    dialogFormVisible.value = true;
};
const handleConfirm = () => {
    conf.value.auto_exit = auto_exit.value ? "TRUE" : "FALSE";
    UpdateConf(conf.value).then(res => {
        if (res.code === 1)
            ElNotification({
                title: "更新成功",
                type: "success"
            });
        else
            ElNotification({
                title: "更新失败",
                type: "error",
                message: res.msg
            });
    })
    dialogFormVisible.value = false;
};
</script>