<script setup lang="ts">
import { onMounted, reactive } from 'vue'
import { list } from '@/store/article';
import type { Article } from '@/store/article';
import * as store from '@/store/article';
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
    store.getAllCategory()
    store.getAllTag()
})

const goToArticle = (id: any) => {
    router.push(`/article/${id}`)
}
const goToCategory = () => {
    router.push("/category")
}
</script>

<template>
    <div style="display: flex;">
        <div style="margin-left: 9vw; width: 70vw;">
            <el-card class="atc-list" v-for="atc in list" :key="atc.ID" @click="goToArticle(atc.ID)">
                <div style="display: flex; flex-flow: column;">
                    <span
                        style="font-size: 1.5em;font-weight: bold; white-space: nowrap; display: inline-block; overflow: hidden; width: 98%;">
                        {{ atc.Title }}
                    </span>
                    <span
                        style=" white-space: nowrap; display: inline-block; overflow: hidden; width: 98%; margin-top: 10px;">
                        {{ atc.Content }}
                    </span>
                    <div style="margin-top: 15px;">
                        <el-tag v-for="tag in atc.tags" style="margin-right: 10px;">{{ tag }} </el-tag>
                    </div>
                </div>
            </el-card>
        </div>
        <div style="margin-left: auto;margin-right: 5vw;">
            <el-card style="width: 15vw; margin-left: 2vw;">
                <div style="display: flex; flex-flow: column;">
                    <el-avatar style="width: 6vw; height: 6vw; margin: 1vh auto 0 auto;"
                        src="https://avatars.githubusercontent.com/u/89156012?v=4" />
                    <span style="margin: 2vh auto 0 auto; font-weight: bold; font-size: 1.6em;">blockche</span>
                    <span style="margin: 0.4vh auto 0 auto; font-size: 1.1em;">无限进步</span>
                    <div style="display: flex; justify-content: space-around; margin-top: 2vh;">
                        <div class="info" style="display: flex; flex-flow: column;">
                            <span>文章</span>
                            <span style="margin: auto; font-size: 1.2em;">{{ list.length }}</span>
                        </div>
                        <div class="info" style="display: flex; flex-flow: column;" @click="goToCategory()">
                            <span>分类</span>
                            <span style="margin: auto; font-size: 1.2em;">{{ store.allCategory.length }}</span>
                        </div>
                        <div class="info" style="display: flex; flex-flow: column;">
                            <span>标签</span>
                            <span style="margin: auto; font-size: 1.2em;">{{ store.allTag.length }}</span>
                        </div>
                    </div>
                </div>
            </el-card>
        </div>
    </div>
</template>

<style scoped>
.info:hover {
    font-weight: bold;
}

.atc-list {
    height: 140px;
    width: 60vw;
}

.atc-list:hover {
    /* border-color: #dadada;
    border-width: 8px; */
    width: 65vw;
    box-shadow: 0 0 0 8px #dadada;
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
