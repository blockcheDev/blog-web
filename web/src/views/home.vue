<script setup lang="ts">
import { onMounted, reactive } from 'vue'
import { list } from '@/store/article';
import type { Article } from '@/store/article';
import api from '@/api/api';
import router from '@/router';

onMounted(async () => {
    try {
        const res = await api.getArticle("all")
        console.log(res.data)
        Object.assign(list, res.data)
    } catch (err) {
        console.error(err)
    }
    for (var i = 0; i < list.length; i++) {
        const atc = list[i]
        try {
            const res = await api.getTagByArticle(atc.ID)
            list[i].tags = res.data
        } catch (err) {
            console.error(err)
        }
    }
})

const goToArticle = (id: any) => {
    router.push(`/article/${id}`)
}
</script>

<template>
    <div class="main">
        <el-card class="atc-list" v-for="atc in list" :key="atc.ID" style="height: 140px; width: 60vw;"
            @click="goToArticle(atc.ID)">
            <div style="display: flex; flex-flow: column;">
                <span
                    style="font-size: 1.5em;font-weight: bold; white-space: nowrap; display: inline-block; overflow: hidden; width: 98%;">
                    {{ atc.Title }}
                </span>
                <span style=" white-space: nowrap; display: inline-block; overflow: hidden; width: 98%; margin-top: 10px;">
                    {{ atc.Content }}
                </span>
                <div style="margin-top: 15px;">
                    <el-tag v-for="tag in atc.tags" style="margin-right: 10px;">{{ tag }} </el-tag>
                </div>
            </div>
        </el-card>
    </div>
</template>

<style scoped>
.atc-list:hover {
    border-color: #dadada;
    border-width: 8px;
    transition: all 0.2s ease-in-out;
}

.list-anmi-move,
/* 对移动中的元素应用的过渡 */
.list-anmi-enter-active,
.list-anmi-leave-active {
    transition: all 0.5s ease;
}

.list-anmi-enter-from,
.list-anmi-leave-to {
    opacity: 0;
    transform: translateX(30px);
}

/* 确保将离开的元素从布局流中删除
  以便能够正确地计算移动的动画。 */
.list-anmi-leave-active {
    position: absolute;
}
</style>
