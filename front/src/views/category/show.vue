<script setup lang="ts">
import PageHeaderProp from '@/layouts/PageHeaderProp.vue';
import api from '@/api/api';
import router from '@/router';
import type { Article } from '@/store/article';
import { onMounted, reactive, ref } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute()

const list: Article[] = reactive([])
const title = ref("")

onMounted(async () => {
    try {
        const res = await api.getCategory(route.params.id)
        title.value = `分类：${res.data.Name}`
    } catch (err) {
        console.error(err)
    }
    try {
        const res = await api.getArticleListByCategory(route.params.id)
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
    <div>
        <el-card style="width: 60vw;">
            <PageHeaderProp :data="title" />
        </el-card>
        <el-card class="atc-list" v-for="atc in list" :key="atc.ID" @click="goToArticle(atc.ID)">
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
</style>