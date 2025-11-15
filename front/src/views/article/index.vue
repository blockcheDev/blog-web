<script setup lang="ts">
import api from '@/api/api';
import { useRoute } from 'vue-router'
import { reactive, onMounted, ref, computed } from 'vue'
import { marked } from 'marked'
import hljs from 'highlight.js';
import articleHeader from './articleHeader.vue'
import comment from './comment.vue';
import info from './info.vue'
import { defaultData } from '@/store/article'
import type { Article } from '@/store/article'
import { useSeoMeta } from '@unhead/vue'

const data: Article = reactive(JSON.parse(JSON.stringify(defaultData))) // 深拷贝，使得data是独立于defaultData的副本

const route = useRoute()
marked.setOptions({
    renderer: new marked.Renderer(),
    gfm: true,
    pedantic: false,
    breaks: false,
    extensions: {
        renderers: {
            ["code"]: function (code: any) {
                return "<pre class='code'><code class=language " + code.lang + ">" + hljs.highlight(code.text, { language: code.lang }).value + "</code></pre>"
            }
        },
        childTokens: {
        }
    }
});
onMounted(async () => {
    try {
        const res = await api.getArticle(route.params.id)
        Object.assign(data, res.data)
    } catch (err) {
        console.error("文章获取失败", err)
    }
})

useSeoMeta({
    title: computed(() => `${(data.Title === "" ? "为什么会没有title" : data.Title)} - blockche blog`),
    description: computed(() => data.Content),
})

</script>

<template>
    <div v-if="data.Title != ''" class="article-page">
        <el-card class="article-section">
            <articleHeader :data="data" />
        </el-card>
        <el-card class="article-section article-content">
            <v-md-editor :model-value="data.Content" mode="preview"></v-md-editor>
        </el-card>
        <el-card class="article-section">
            <info :data="data" />
        </el-card>
        <el-card class="article-section">
            <comment :articleID="route.params.id" />
        </el-card>
    </div>
</template>

<style scoped>
.article-page {
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 32px 20px;
    width: 100%;
    max-width: 60vw;
    min-height: 100vh;
    margin: 0 auto;
    box-sizing: border-box;
    overflow-x: hidden;
}

.article-section {
    width: 100%;
    margin-bottom: 24px;
    border-radius: 12px;
    box-shadow: 0 4px 16px rgba(0, 0, 0, 0.1), 0 2px 4px rgba(0, 0, 0, 0.06);
    transition: all 0.3s ease;
    border: 1px solid rgba(0, 0, 0, 0.05);
}

.article-section:last-child {
    margin-bottom: 0;
}

.article-content {
    overflow-x: auto;
}

/* 全局卡片样式优化 */
.article-section :deep(.el-card__body) {
    padding: 28px 32px;
}

/* 文章内容区域特殊处理 */
.article-content :deep(.el-card__body) {
    padding: 32px 36px;
}

/* 移动端减少卡片内边距 */
@media (max-width: 768px) {
    .article-section {
        border-radius: 10px;
        box-shadow: 0 3px 12px rgba(0, 0, 0, 0.08), 0 1px 4px rgba(0, 0, 0, 0.05);
    }
    
    .article-section :deep(.el-card__body) {
        padding: 18px 20px;
    }
    
    .article-content :deep(.el-card__body) {
        padding: 20px;
    }
    
    /* 减少 markdown 内容区域的内边距 */
    .article-content :deep(.github-markdown-body) {
        padding: 8px !important;
    }
    
    .article-content :deep(.v-md-editor-preview) {
        padding: 8px !important;
    }
}

@media (max-width: 480px) {
    .article-section {
        border-radius: 8px;
    }
    
    .article-section :deep(.el-card__body) {
        padding: 14px 16px;
    }
    
    .article-content :deep(.el-card__body) {
        padding: 16px;
    }
    
    /* 小屏幕进一步减少 markdown 内容区域的内边距 */
    .article-content :deep(.github-markdown-body) {
        padding: 4px !important;
    }
    
    .article-content :deep(.v-md-editor-preview) {
        padding: 4px !important;
    }
}

/* 桌面端 - 保持原有宽度 */
@media (min-width: 1025px) {
    .article-page {
        max-width: 60vw;
    }
}

/* 平板端 - 增加宽度利用率 */
@media (min-width: 769px) and (max-width: 1024px) {
    .article-page {
        max-width: 80vw;
        padding: 28px 20px;
    }
    
    .article-section {
        margin-bottom: 22px;
        border-radius: 11px;
        box-shadow: 0 3px 14px rgba(0, 0, 0, 0.09), 0 2px 4px rgba(0, 0, 0, 0.05);
    }
    
    .article-section :deep(.el-card__body) {
        padding: 24px 28px;
    }
}

/* 移动端 - 全宽显示 */
@media (max-width: 768px) {
    .article-page {
        max-width: 100%;
        padding: 16px 8px;
    }
    
    .article-section {
        margin-bottom: 16px;
    }
}

/* 小屏幕移动端优化 */
@media (max-width: 480px) {
    .article-page {
        padding: 12px 4px;
    }
    
    .article-section {
        margin-bottom: 12px;
    }
}
</style>