<script setup lang="ts">
import PageHeader from '@/layouts/PageHeader.vue';
import { ref } from 'vue';
import passwordEdit from './passwordEdit.vue';
const dialog = ref();

</script>

<template>
    <div>
        <el-card>
            <PageHeader />
        </el-card>
        <el-card>
            <el-row justify="space-between">
                <span>密码</span>
                <el-button type="primary" @click="dialog.openDialog()">修改密码</el-button>
            </el-row>
            <el-row justify="space-between">
                <span>注销</span>
                <el-button type="danger">账号注销</el-button>
            </el-row>
            <passwordEdit ref="dialog" />
        </el-card>
    </div>
</template>

<style>
.el-row {
    margin-bottom: 30px;
}
</style>