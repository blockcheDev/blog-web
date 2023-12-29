<script setup lang="ts">
import api from '@/api/api';
import { onMounted, reactive, ref } from 'vue';
import type { Comment } from '@/store/article';
import dayjs from 'dayjs';
const formatDate = (date: any) => {
    return dayjs(date).format("YYYY-MM-DD hh:mm")
}

const prop = defineProps(['articleID']);

const comment: Comment = reactive({})
const pushComment = async () => {
    comment.ArticleID = Number(prop.articleID);
    try {
        const res = await api.pushComment(comment)
        location.reload()
    } catch (err) {
        console.error(err)
    }
}

const list: Comment[] = reactive([])

onMounted(async () => {
    try {
        const res = await api.getCommentListByArticle(prop.articleID)
        Object.assign(list, res.data)
    } catch (err) {
        console.error(err)
    }
    for (var i = 0; i < list.length; i++) {
        try {
            const res = await api.getUser(list[i].UserID)
            list[i].UserName = res.data.Name
        } catch (err) {
            console.error(err)
        }
        list[i].Floor = i + 1
    }
})

</script>

<template>
    <div>
        <div>
            <span style="font-size: 1.5em; font-weight: bold;">评论</span>
            <div style="margin-top: 1.5vh;" />
            <el-input v-model="comment.Content" :autosize="{ minRows: 2, maxRows: 10 }" type="textarea"
                placeholder="输入评论" />
            <div style="margin-top: 1.5vh;" />
            <div style="display: flex;">
                <el-button style="margin-left: auto;" type="primary" @click="pushComment()">发送</el-button>
            </div>
        </div>
        <div>
            <div v-for="cmt in list" style="margin-top: 1em; display: flex; flex-flow: column;">
                <div style="display: flex; align-items: center;">
                    <el-tag size="small">#{{ cmt.Floor }}楼</el-tag>
                    <span style="font-size: 1.1em; margin-left: 1em;">{{ cmt.UserName }}</span>
                    <span style="font-size: 1.1em; color: gray; margin-left: 0.5em;">{{ formatDate(cmt.CreatedAt) }}</span>
                </div>
                <el-card style="margin: 0.5em auto 0 0">
                    <span>{{ cmt.Content }}</span>
                </el-card>
            </div>
        </div>
    </div>
</template>