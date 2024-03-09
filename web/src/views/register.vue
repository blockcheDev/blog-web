<template>
    <div>
        <el-card>
            <PageHeader />
        </el-card>
        <el-card>
            <el-form :model="form" label-width="120px">
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
                    <el-button @click="goToLogin">返回登录</el-button>
                </el-form-item>
            </el-form>
        </el-card>
    </div>
</template>

<script lang="ts" setup>
import api from '@/api/api';
import router from '@/router';
import { reactive } from 'vue'
import { ElMessage } from 'element-plus'
import PageHeader from '@/layouts/PageHeader.vue';

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
        router.push("/login")
    } catch (error) {
        console.error("提交失败", error)
    }
}

const goToLogin = () => {
    router.push("/login")
}
</script>

<style>
.input {
    width: 15vw;
}
</style>