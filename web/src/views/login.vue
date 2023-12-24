<template>
    <div>
        <el-card>
            <PageHeader />
        </el-card>
        <el-card>
            <el-form :model="form" label-width="120px">
                <el-form-item label="用户名">
                    <el-input v-model="form.Name" />
                </el-form-item>
                <el-form-item label="密码">
                    <el-input v-model="form.Password" type="password" show-password />
                </el-form-item>

                <el-form-item>
                    <el-button type="primary" @click="onSubmit">登录</el-button>
                    <el-button @click="goToRegister">注册</el-button>
                </el-form-item>
            </el-form>
        </el-card>
    </div>
</template>
  
<script lang="ts" setup>
import request from '@/api/axios';
import router from '@/router';
import { reactive } from 'vue'
import { ElMessage } from 'element-plus'
import PageHeader from '@/layouts/PageHeader.vue';

// do not use same name with ref
const form = reactive({
    Name: '',
    Password: '',
})

const onSubmit = async () => {
    try {
        const res = await request.post("/login", form)
        console.log(res.data.msg)
        localStorage.setItem("token", res.headers.token)
        router.push("/home")
    } catch (error) {
        console.error(error)
    }

}

const goToRegister = () => {
    router.push("/register")
}
</script>
  