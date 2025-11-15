<script setup lang="ts">
import Aside from '@/layouts/aside.vue';
import Header from '@/layouts/header.vue';
import Footer from '@/layouts/footer.vue';
import router from '@/router';
import { inject, ref } from 'vue';
import { Menu } from '@element-plus/icons-vue'
import { drawerMenu } from '@/store/store';

</script>

<template>
    <button class="drawer-button" @click="drawerMenu.open()">
        <Menu class="menu-icon" />
    </button>
    <div class="common-layout">
        <el-container>
            <el-header>
                <Header />
            </el-header>
            <!-- 占位用，高度和header一致，否则main区域会侵入header -->
            <div style="height: 60px;"></div>
            <el-container>
                <Aside v-if="drawerMenu.isOpen" />
                <el-main>
                    <router-view v-slot="{ Component }">
                        <transition mode="out-in">
                            <component :is="Component" />
                        </transition>
                    </router-view>
                </el-main>
            </el-container>
            <el-footer>
                <Footer />
            </el-footer>
        </el-container>
    </div>
</template>

<style scoped>
.el-header {
    /* background-color: #f6f8fa; */
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.24);
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
    left: 20px;
    bottom: 40px;
    z-index: 10;
    width: 48px;
    height: 48px;
    border-radius: 50%;
    border: none;
    background: #ababab;
    color: white;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    box-shadow: 0 4px 12px rgba(57, 57, 57, 0.3);
    transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
    outline: none;
}

.drawer-button:hover {
    transform: translateY(-3px) scale(1.05);
    box-shadow: 0 6px 20px rgba(57, 57, 57, 0.4);
}

.drawer-button:active {
    transform: translateY(-1px) scale(1.02);
    box-shadow: 0 3px 10px rgba(57, 57, 57, 0.3);
}

.drawer-button .menu-icon {
    width: 24px;
    height: 24px;
    transition: transform 0.3s ease;
}

.drawer-button:hover .menu-icon {
    transform: rotate(90deg);
}

.logoutButton {
    margin-top: auto;
    margin-left: auto;
}
</style>
