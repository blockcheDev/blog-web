<template>
    <div>
        <DataDialog v-model="showDialog" title="注册" width="30vw">
            <div style="display: flex; flex-flow: column; align-items: center;">
                <el-form :model="form" label-width="auto" style="width: 80%;">
                    <el-form-item label="用户名">
                        <el-input class="input" v-model="form.Name" />
                    </el-form-item>
                    <el-form-item label="密码">
                        <el-input class="input" v-model="form.Password" type="password" show-password />
                    </el-form-item>
                    <el-form-item label="邮箱">
                        <el-input class="input" v-model="form.Email" />
                    </el-form-item>
                    <el-form-item label="手机号码">
                        <el-input class="input" v-model="form.Telephone" />
                        <span style="color: gray; margin-left: 1em;">*非必填</span>
                    </el-form-item>
                    <el-form-item label="性别">
                        <el-radio-group v-model="form.Gender">
                            <el-radio label="男">男</el-radio>
                            <el-radio label="女">女</el-radio>
                            <el-radio label="未知">未知</el-radio>
                        </el-radio-group>
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click="onSubmit">确认</el-button>
                        <el-button @click="close()">返回登录</el-button>
                    </el-form-item>
                </el-form>
            </div>
        </DataDialog>
    </div>
</template>

<script lang="ts" setup>
import api from '@/api/api';
import { reactive, ref } from 'vue'
import DataDialog from '@/components/DataDialog.vue'

const form = reactive({
    Name: '',
    Password: '',
    Email: '',
    Telephone: '',
    Gender: "未知",
})

const onSubmit = async () => {
    try {
        const res = await api.register(form)
        console.log(res.data.msg)
        close()
    } catch (error) {
        console.error("提交失败", error)
    }
}

const showDialog = ref(false)
const open = () => {
    showDialog.value = true
}
const close = () => {
    showDialog.value = false
}
defineExpose({
    open
})
</script>

<style>
.input {
    width: 15vw;
}
</style>