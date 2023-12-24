<script lang="ts" setup>
import {
    Document,
    Menu as IconMenu,
    Location,
    Setting,
    House,
} from '@element-plus/icons-vue'
import { inject } from 'vue';
import { drawerMenu } from '@/store/store';

</script>
  
<template>
    <el-menu default-active="/home" class="el-menu" router="true" @select="drawerMenu.close()" style="width: 200px;">
        <el-menu-item index="/home">
            <template #title>
                <el-icon>
                    <House />
                </el-icon>
                <span>首页</span>
            </template>
        </el-menu-item>
        <el-sub-menu index="2">
            <template #title>
                <el-icon>
                    <location />
                </el-icon>
                <span>个人中心</span>
            </template>
            <el-menu-item index="/user/push">发布文章</el-menu-item>
            <el-menu-item index="/user/article">文章管理</el-menu-item>
            <el-menu-item index="/user/info">个人资料</el-menu-item>
            <el-menu-item index="/user/setting">账户设置</el-menu-item>
        </el-sub-menu>
    </el-menu>
</template>
  