<template>
    <div>
        <DataDialog v-model="loginDialog.isOpen" title="登录" width="30vw">
            <div style="display: flex; flex-flow: column; align-items: center;">
                <el-form :model="form" label-width="auto" style="width: 80%;">
                    <el-form-item label="用户名">
                        <el-input class="input" v-model="form.Name" />
                    </el-form-item>
                    <el-form-item label="密码">
                        <el-input class="input" v-model="form.Password" type="password" show-password />
                    </el-form-item>
        
                    <el-form-item>
                        <el-button type="primary" @click="onSubmit">登录</el-button>
                        <el-button @click="registerRef?.open">注册</el-button>
                    </el-form-item>
                    <el-form-item>
                        <el-link href="javascript:void(0)" @click="goToGithubLogin"
                            :underline="false">
                            <svg preserveAspectRatio="xMidYMid meet" viewBox="0 0 24 24" width="1.8em" height="1.8em"
                                data-v-6c8d2bba="">
                                <path fill="currentColor"
                                    d="M12 2C6.475 2 2 6.475 2 12a9.994 9.994 0 0 0 6.838 9.488c.5.087.687-.213.687-.476c0-.237-.013-1.024-.013-1.862c-2.512.463-3.162-.612-3.362-1.175c-.113-.288-.6-1.175-1.025-1.413c-.35-.187-.85-.65-.013-.662c.788-.013 1.35.725 1.538 1.025c.9 1.512 2.338 1.087 2.912.825c.088-.65.35-1.087.638-1.337c-2.225-.25-4.55-1.113-4.55-4.938c0-1.088.387-1.987 1.025-2.688c-.1-.25-.45-1.275.1-2.65c0 0 .837-.262 2.75 1.026a9.28 9.28 0 0 1 2.5-.338c.85 0 1.7.112 2.5.337c1.912-1.3 2.75-1.024 2.75-1.024c.55 1.375.2 2.4.1 2.65c.637.7 1.025 1.587 1.025 2.687c0 3.838-2.337 4.688-4.562 4.938c.362.312.675.912.675 1.85c0 1.337-.013 2.412-.013 2.75c0 .262.188.574.688.474A10.016 10.016 0 0 0 22 12c0-5.525-4.475-10-10-10z">
                                </path>
                            </svg>
                            <span style="margin-left: 0.5em; font-size: 1.1em;">
                                Github 账号登录
                            </span>
                        </el-link>
                    </el-form-item>
                </el-form>
            </div>
        </DataDialog>
        <register ref="registerRef" />
    </div>
</template>

<script lang="ts" setup>
import api from '@/api/api';
import router from '@/router';
import { reactive, ref } from 'vue'
import DataDialog from '@/components/DataDialog.vue'
import register from '@/views/register.vue'
import { drawerMenu, loginDialog } from '@/store/store'

const goToGithubLogin = () => {
    var iHeight = 600;
    var iWidth = 800;
    var iTop = (window.screen.availHeight - 30 - iHeight) / 2;
    //获得窗口的水平位置 
    var iLeft = (window.screen.availWidth - 10 - iWidth) / 2;
    window.open('https://github.com/login/oauth/authorize?client_id=9f837cdc52f3a2edbcc9', '', 'height=' + iHeight + ',innerHeight=' + iHeight + ',width=' + iWidth + ',innerWidth=' + iWidth + ',top=' + iTop + ',left=' + iLeft + ',status=no,toolbar=no,menubar=no,location=no,resizable=no,scrollbars=0,titlebar=no');

    loginDialog.close()
    drawerMenu.close()
}

// do not use same name with ref
const form = reactive({
    Name: '',
    Password: '',
})

const onSubmit = async () => {
    try {
        const res = await api.login(form)
        console.log(res.data.msg)
        localStorage.setItem("token", res.headers.token)
        loginDialog.close()
    } catch (error) {
        console.error(error)
    }

}

const registerRef = ref<InstanceType<typeof register>>()
</script>

<style>
.input {
    width: 15vw;
}
</style>