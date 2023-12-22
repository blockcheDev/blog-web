<script setup lang="ts">
import Aside from '@/layouts/aside.vue';
import Header from '@/layouts/header.vue';
import PageHeader from './PageHeader.vue';
import router from '@/router';
import { inject, ref } from 'vue';
import { Menu } from '@element-plus/icons-vue'
import { drawerMenu } from '@/store/store';

const logout = () => {
    localStorage.setItem("token", "")
    router.push("/login")
}

</script>

<template>
    <el-button type="primary" class="drawer-button" :icon="Menu" @click="drawerMenu.open()" />
    <div class="common-layout">
        <el-container>
            <el-header>
                <Header />
            </el-header>
            <!-- 占位用，高度和header一致，否则main区域会侵入header -->
            <div style="height: 60px;"></div>
            <el-container>
                <el-drawer v-model="drawerMenu.isOpen" title="I am the title" :with-header="false" direction="ltr"
                    size="15%">
                    <template #default>
                        <Aside />
                    </template>
                    <template #footer>
                        <el-button type="danger" @click="logout"
                            v-if="($route.path != '/login' && $route.path != '/register')">登出</el-button>
                    </template>
                </el-drawer>
                <el-main>
                    <el-row justify="center">
                        <el-col :span="15">
                            <router-view v-slot="{ Component }">
                                <transition mode="out-in">
                                    <component :is="Component" />
                                </transition>
                            </router-view>
                        </el-col>
                    </el-row>
                </el-main>
            </el-container>
        </el-container>
    </div>
</template>

<style scoped>
.el-header {
    /* background-color: #f6f8fa; */
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.24);
    display: flex;
    justify-content: space-between;
    align-items: center;
    /* margin-bottom: 3px; */
    backdrop-filter: saturate(50%) blur(4px);
    background: rgba(255, 255, 255, .5);
    position: fixed;
    width: 100vw;
    z-index: 9;
    height: 60px;
}

.el-main {
    align-items: center;
}

.drawer-button {
    position: fixed;
    margin-left: 1%;
    margin-top: 5%;
    z-index: 10;
}

.logoutButton {
    margin-top: auto;
    margin-left: auto;
}
</style>