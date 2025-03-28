<script lang="ts" setup>
import {
    Document,
    Menu as IconMenu,
    Location,
    Setting,
    House,
    Upload,
    Notebook,
    User,
    Promotion,
    CloseBold,
} from '@element-plus/icons-vue'
import { inject } from 'vue';
import { drawerMenu } from '@/store/store';
import router from '@/router';

const goTo = (addr: string) => {
    router.push(addr)
}

const isLogin = () => {
    const tokenString = localStorage.getItem("token");
    if (tokenString) {
        return true
    } else {
        return false
    }
}

const logout = () => {
    localStorage.removeItem("token")
    router.push("/home")
}

</script>
  
<template>
    <div>
        <el-drawer v-model="drawerMenu.isOpen" title="I am the title" :with-header="false" direction="ltr" size="20vw">
            <div style="display: flex; flex-flow: column; height: 100%;">
                <el-card style="margin-top: 30px;" @click="goTo('/home')">
                    <div style="display: flex;">
                        <House style="width: 1.6em; margin-right: 1em;" />
                        <span style="font-size: 1.3em; font-weight: bold;">首页</span>
                    </div>
                </el-card>
                <el-card @click="goTo('/user/push')">
                    <div style="display: flex;">
                        <Upload style="width: 1.6em; margin-right: 1em;" />
                        <span style="font-size: 1.3em; font-weight: bold;">发布文章</span>
                    </div>
                </el-card>
                <el-card @click="goTo('/user/article')">
                    <div style="display: flex;">
                        <Notebook style="width: 1.6em; margin-right: 1em;" />
                        <span style="font-size: 1.3em; font-weight: bold;">文章管理</span>
                    </div>
                </el-card>
                <el-card @click="goTo('/user/info')">
                    <div style="display: flex;">
                        <User style="width: 1.6em; margin-right: 1em;" />
                        <span style="font-size: 1.3em; font-weight: bold;">个人资料</span>
                    </div>
                </el-card>
                <el-card @click="goTo('/user/setting')">
                    <div style="display: flex;">
                        <Setting style="width: 1.6em; margin-right: 1em;" />
                        <span style="font-size: 1.3em; font-weight: bold;">账号设置</span>
                    </div>
                </el-card>
                <el-card v-if="isLogin()" style="margin-top: auto; background-color:rgb(255, 144, 144);" @click="logout()">
                    <div style="display: flex;">
                        <CloseBold style="width: 1.6em; margin-right: 1em;" />
                        <span style="font-size: 1.3em; font-weight: bold;">登出</span>
                    </div>
                </el-card>
                <el-card v-else style="margin-top: auto; background-color:aliceblue;" @click="goTo('/login')">
                    <div style="display: flex;">
                        <Promotion style="width: 1.6em; margin-right: 1em;" />
                        <span style="font-size: 1.3em; font-weight: bold;">登录</span>
                    </div>
                </el-card>
            </div>
        </el-drawer>
    </div>
</template>

<style scoped>
.el-card:hover {
    /* border-color: #dadada;
    border-width: 8px; */
    width: 90%;
    box-shadow: 0 0 0 8px #dadada;
    transition: all 0.2s ease-in-out;
}
</style>
  