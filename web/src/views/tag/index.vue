<script setup lang="ts">
import router from '@/router';
import * as article from '@/store/article'
import { onMounted } from 'vue';

onMounted(() => {
    article.getAllTag()
})

const goToTag = (id: any) => {
    router.push(`/tag/${id}`)
}

</script>

<template>
    <div style="display: flex;">
        <div style="display: flex; flex-flow: wrap; justify-content: left; width: 70vw; margin: 0 auto 0 auto;">
            <el-card v-for="tag in article.allTag" :key="tag.Name" @click="goToTag(tag.ID)">
                <span style="font-size:larger; font-weight: bold;">{{ tag.Name }}</span>
            </el-card>
        </div>
    </div>
</template>

<style scoped>
.el-card {
    width: 8vw;
    height: 8vw;
    margin-top: 1vw;
    margin-left: 2vw;
    display: flex;
    justify-content: center;
    align-items: center;
}

.el-card:hover {
    box-shadow: 0 0 0 8px #dadada;
    transition: all 0.2s ease-in-out;
}
</style>