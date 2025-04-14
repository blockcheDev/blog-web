<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'
import { list } from '@/store/article';
import * as store from '@/store/article';
import api from '@/api/api';
import router from '@/router';

const RecentVisitorsCount = reactive({
    RecentVisitorsCount: 0,
})
onMounted(async () => {
    try {
        const res = await api.getArticle("all")
        Object.assign(list, res.data)
    } catch (err) {
        console.error(err)
    }
    store.getAllCategory()
    store.getAllTag()
    try {
        const res = await api.getRecentVisitorsCount()
        Object.assign(RecentVisitorsCount, res.data)
    } catch (err) {
        console.error(err)
    }
    
})

const goToArticle = (id: any) => {
    router.push(`/article/${id}`)
}
const goToCategory = () => {
    router.push("/category")
}
const goToTag = () => {
    router.push("/tag")
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
                        <el-tag v-for="tag in atc.Tags" style="margin-right: 10px;">{{ tag.Name }} </el-tag>
                    </div>
                </div>
            </el-card>
        </div>
        <div style="display: flex; flex-flow: column; margin-left: auto;margin-right: 5vw;">
            <el-card style="width: 15vw; margin-left: 2vw;">
                <div style="display: flex; flex-flow: column;">
                    <el-avatar style="width: 6vw; height: 6vw; margin: 1vh auto 0 auto;"
                        src="https://avatars.githubusercontent.com/u/89156012?v=4" />
                    <span style="margin: 2vh auto 0 auto; font-weight: bold; font-size: 1.6em;">blockche</span>
                    <span style="margin: 0.4vh auto 0 auto; font-size: 1.1em;">无限进步</span>
                    <div style="display: flex; flex-flow: column; margin-top: 2vh;">
                        <div class="info">
                            <span>文章</span>
                            <span>{{ list.length }}</span>
                        </div>
                        <div class="info" @click="goToCategory()">
                            <span>分类</span>
                            <span>{{ store.allCategory.length }}</span>
                        </div>
                        <div class="info" @click="goToTag()">
                            <span>标签</span>
                            <span>{{ store.allTag.length }}</span>
                        </div>
                    </div>
                </div>
            </el-card>
            <el-card style="width: 15vw; margin-left: 2vw;">
                <span>最近访客 {{ RecentVisitorsCount.RecentVisitorsCount }}</span>
            </el-card>
        </div>
    </div>
</template>

<style scoped>
.info {
    display: flex;
    margin-top: 1.5vh;

    padding: 6px;
    border-radius: 8px;
    border: 1px solid #e0e0e0;
    /* background: #f8f9fa; */
    cursor: default;
    
    /* 过渡动画 */
    transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
    
    /* 内容对齐 */
    align-items: center;
}

.info:hover {
    /* 悬停效果 */
    cursor: pointer;
    background: #ffffff;
    border-color: #2196f3;
    box-shadow: 0 2px 8px rgba(33, 150, 243, 0.2);
    transform: translateY(-2px);
}

.info span:first-child {
    font-size: 0.9rem;
    color: #666;

    margin-left: 0.5vw;
}

.info span:last-child {
    font-weight: 500;
    color: #666;

    margin-left: 0.5vw;
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
