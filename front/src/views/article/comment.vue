<script setup lang="ts">
import api from '@/api/api';
import { onMounted, reactive, ref } from 'vue';
import type { Comment } from '@/store/article';
import { userInfo } from '@/store/user'
import dayjs from 'dayjs';
const formatDate = (date: any) => {
    return dayjs(date).format("YYYY-MM-DD hh:mm")
}

const prop = defineProps(['articleID']);

const comment: Partial<Comment> = reactive({})
const pushComment = async () => {
    comment.ArticleID = Number(prop.articleID);
    try {
        const res = await api.pushComment(comment)
        // 等待700ms再执行内部方法
        setTimeout(() => {
            location.reload()
        }, 700)
    } catch (err) {
        console.error(err)
    }
}

const list: Comment[] = reactive([])

onMounted(async () => {
    try {
        const res = await api.getCommentListByArticle(prop.articleID)
        console.log(res.data)
        Object.assign(list, res.data)
    } catch (err) {
        console.error(err)
    }
    for (var i = 0; i < list.length; i++) {
        list[i].Floor = i + 1
    }
    userInfo.getInfo()
})

const deleteComment = async (id: any) => {
    try {
        const res = await api.deleteComment(id)
        setTimeout(() => {
            location.reload()
        }, 700)
    } catch (err) {
        console.error(err)
    }
}

</script>

<template>
    <div class="comment-container">
        <div class="comment-input-section">
            <span class="comment-title">评论</span>
            <div class="input-spacer" />
            <el-input v-model="comment.Content" :autosize="{ minRows: 2, maxRows: 10 }" type="textarea"
                placeholder="输入评论" class="comment-textarea" />
            <div class="input-spacer" />
            <div class="button-wrapper">
                <el-button type="primary" @click="pushComment()">发送</el-button>
            </div>
        </div>
        <div class="comment-list">
            <div v-for="cmt in list" :key="cmt.ID" class="comment-item">
                <div class="comment-header">
                    <el-tag size="small" class="floor-tag">#{{ cmt.Floor }}楼</el-tag>
                    <span class="user-name">{{ cmt.User.Name === "" ? "账号已注销" : cmt.User.Name }}</span>
                    <span class="comment-time">{{ formatDate(cmt.CreatedAt) }}</span>
                    <el-button v-if="cmt.User.Name === userInfo.Name" link class="delete-btn"
                        @click="deleteComment(cmt.ID)">删除</el-button>
                </div>
                <el-card class="comment-content-card">
                    <span class="comment-text">{{ cmt.Content }}</span>
                </el-card>
            </div>
        </div>
    </div>
</template>

<style scoped>
.comment-container {
    width: 100%;
}

.comment-input-section {
    margin-bottom: 1.5rem;
    padding: 1rem 1.25rem;
    border-radius: 10px;
    border: 1px solid #e9ecef;
}

.comment-title {
    font-size: 1.35em;
    font-weight: 700;
    color: #1a1a1a;
    margin-bottom: 0.5rem;
}

.input-spacer {
    margin-top: 0.65rem;
}

.comment-textarea {
    width: 100%;
}

.comment-textarea :deep(.el-textarea__inner) {
    border-radius: 8px;
    border: 1px solid #dee2e6;
    transition: all 0.3s ease;
}

.comment-textarea :deep(.el-textarea__inner):focus {
    border-color: #409eff;
    box-shadow: 0 0 0 3px rgba(64, 158, 255, 0.1);
}

.button-wrapper {
    display: flex;
    justify-content: flex-end;
}

.button-wrapper :deep(.el-button) {
    border-radius: 8px;
    padding: 10px 24px;
    font-weight: 500;
}

.comment-list {
    margin-top: 0;
}

.comment-item {
    margin-top: 0.85rem;
    display: flex;
    flex-direction: column;
    padding: 0.75rem 0.85rem;
    background: #ffffff;
    border-radius: 10px;
    border: 1px solid #e9ecef;
    transition: all 0.3s ease;
}

.comment-item:first-child {
    margin-top: 0;
}

.comment-header {
    display: flex;
    align-items: center;
    flex-wrap: wrap;
    gap: 0.65em;
    margin-bottom: 0.5rem;
    padding-bottom: 0.5rem;
    border-bottom: 1px solid #f0f0f0;
}

.floor-tag {
    flex-shrink: 0;
    border-radius: 12px;
    font-weight: 600;
}

.user-name {
    font-size: 0.95em;
    font-weight: 600;
    color: #1a1a1a;
}

.comment-time {
    font-size: 0.85em;
    color: #6c757d;
}

.delete-btn {
    margin-left: auto;
    color: #dc3545;
    font-weight: 500;
}

.delete-btn:hover {
    color: #c82333;
}

.comment-content-card {
    width: 100%;
    background: transparent;
    border: none;
    box-shadow: none;
    margin: 0;
    border-radius: 0;
    overflow: visible;
}

.comment-content-card :deep(.el-card__body) {
    padding: 0 !important;
    margin: 0;
    border-radius: 0;
    overflow: visible;
}

.comment-text {
    word-break: break-word;
    white-space: pre-wrap;
    color: #495057;
    line-height: 1.5;
    font-size: 0.93em;
    display: block;
    margin: 0;
    padding: 0;
}

/* 移动端适配 */
@media (max-width: 768px) {
    .comment-input-section {
        padding: 0.85rem 1rem;
        border-radius: 9px;
    }
    
    .comment-title {
        font-size: 1.25em;
    }
    
    .input-spacer {
        margin-top: 0.75rem;
    }
    
    .comment-item {
        padding: 0.65rem 0.75rem;
        border-radius: 9px;
        margin-top: 0.75rem;
    }
    
    .user-name {
        font-size: 0.95em;
    }
    
    .comment-time {
        font-size: 0.85em;
    }
    
    .comment-header {
        gap: 0.55em;
        margin-bottom: 0.45rem;
        padding-bottom: 0.45rem;
    }
    
    .delete-btn {
        margin-left: 0;
        font-size: 0.9em;
    }
    
    .comment-text {
        font-size: 0.9em;
    }
}

/* 小屏幕移动端 */
@media (max-width: 480px) {
    .comment-input-section {
        padding: 0.75rem 0.85rem;
        border-radius: 8px;
    }
    
    .comment-title {
        font-size: 1.2em;
    }
    
    .input-spacer {
        margin-top: 0.7rem;
    }
    
    .comment-item {
        padding: 0.55rem 0.65rem;
        border-radius: 8px;
        margin-top: 0.7rem;
    }
    
    .user-name {
        font-size: 0.9em;
    }
    
    .comment-time {
        font-size: 0.8em;
    }
    
    .input-spacer {
        margin-top: 0.8rem;
    }
    
    .comment-text {
        font-size: 0.85em;
    }
    
    .button-wrapper :deep(.el-button) {
        padding: 8px 20px;
        font-size: 0.9em;
    }
}
</style>