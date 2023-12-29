<template>
    <el-page-header :icon="ArrowLeft" @back="goBack">
        <template #content>
            <span class="text-large font-600 mr-3"> {{ data }} </span>
        </template>
    </el-page-header>
</template>
  
<script lang="ts" setup>
import router from '@/router';
import { ArrowLeft } from '@element-plus/icons-vue'
const goBack = () => {
    router.back()
}

const prop = defineProps(['data'])
</script>