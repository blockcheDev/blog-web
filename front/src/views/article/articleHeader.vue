<script lang="ts" setup>
import api from '@/api/api';
import router from '@/router';
import { ArrowLeft } from '@element-plus/icons-vue'
import { Calendar } from '@element-plus/icons-vue'
import dayjs from 'dayjs'
import { onMounted, ref } from 'vue';
const formatDate = (date: any) => {
    return dayjs(date).format("YYYY-MM-DD")
}
const goBack = () => {
    router.back()
}
const props = defineProps(['data'])

</script>
<template>
    <div style="display: flex; flex-flow: column;">
        <div style="margin: auto;">
            <h2>{{ data.Title }}</h2>
        </div>
        <div style="display: flex; flex-flow: column; margin: auto;">
            <p style="margin: 0.3em auto;">
                发布于 {{ formatDate(data.CreatedAt) }} | 修改于 {{ formatDate(data.UpdatedAt) }}
            </p>
            <p style="margin: 0.3em auto;">
                {{ data.CategoryName }} | {{ "data.Type===0" ? "原创" : "转载" }}
            </p>
            <p style="margin: auto auto;">
                <el-tag v-for="tag in data.tags" style="margin-left: 10px;">{{ tag }} </el-tag>
            </p>
        </div>
    </div>
</template>
  