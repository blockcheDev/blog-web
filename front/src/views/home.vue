<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue'
import { list } from '@/store/article';
import * as store from '@/store/article';
import api from '@/api/api';
import router from '@/router';
import recentVisitors from '@/views/recentVisitors.vue'
import { Histogram } from '@element-plus/icons-vue'
import { useSeoMeta } from '@unhead/vue'

// SEO 元数据
const description = ref("欢迎来到 blockche blog，这里会随意分享一些东西。")
useSeoMeta({
    title: computed(() => `首页 - blockche blog`),
    description: computed(() => description.value),
})

// 排序状态：0-按时间，1-按阅读量
enum SortType {
    ByTime = 0,
    ByViews = 1
}

const sortStatus = ref<SortType>(SortType.ByTime)

// 计算排序后的文章列表（避免直接修改原数组）
const sortedList = computed(() => {
    const listCopy = [...list]
    if (sortStatus.value === SortType.ByViews) {
        return listCopy.sort((a, b) => b.PageViews - a.PageViews)
    }
    return listCopy.sort((a, b) => b.ID - a.ID)
})

// 排序按钮文本
const sortButtonText = computed(() =>
    sortStatus.value === SortType.ByTime ? "发表时间排序" : "阅读量排序"
)

// 网站信息
const webInfo = reactive({
    RecentVisitorsCount: 0,
    VisitorCount: 0,
    TotalPageViews: 0,
})

const recentVisitorsRef = ref<InstanceType<typeof recentVisitors>>()

// 初始化数据
onMounted(async () => {
    try {
        // 并行请求优化性能
        const [articleRes, webInfoRes] = await Promise.all([
            api.getArticle("all"),
            api.getWebInfo(),
            store.getAllCategory(),
            store.getAllTag()
        ])

        Object.assign(list, articleRes.data)
        Object.assign(webInfo, webInfoRes.data)

        // 构建 SEO 描述
        description.value += list.map(article => article.Title).join(' ')
    } catch (err) {
        console.error('加载数据失败:', err)
    }
})

// 切换排序方式
const switchSort = (): void => {
    sortStatus.value = sortStatus.value === SortType.ByTime
        ? SortType.ByViews
        : SortType.ByTime
}

// 导航函数
const goToArticle = (id: number): void => {
    router.push(`/article/${id}`)
}

const goToCategory = (): void => {
    router.push("/category")
}

const goToTag = (): void => {
    router.push("/tag")
}

</script>

<template>
    <div class="home-container">
        <div class="home-content">
            <!-- 文章列表区域 -->
            <div class="article-list-container">
                <TransitionGroup name="list">
                    <el-card v-for="atc in sortedList" :key="atc.ID" class="article-card" @click="goToArticle(atc.ID)">
                        <div class="article-content">
                            <!-- 标题区域 -->
                            <div class="article-header">
                                <h3 class="article-title">{{ atc.Title }}</h3>
                            </div>
                            
                            <!-- 摘要区域 -->
                            <div class="article-body">
                                <p class="article-excerpt">{{ atc.Content }}</p>
                            </div>
                            
                            <!-- 底部信息区域 -->
                            <div class="article-footer">
                                <div class="article-meta">
                                    <!-- 标签 -->
                                    <div class="article-tags" v-if="atc.Tags && atc.Tags.length > 0">
                                        <el-tag 
                                            v-for="tag in atc.Tags.slice(0, 4)" 
                                            :key="tag.Name" 
                                            class="article-tag"
                                            size="small"
                                            effect="plain">
                                            {{ tag.Name }}
                                        </el-tag>
                                        <span v-if="atc.Tags.length > 4" class="more-tags">+{{ atc.Tags.length - 4 }}</span>
                                    </div>
                                </div>
                                
                                <!-- 统计信息 -->
                                <div class="article-stats">
                                    <span class="stat-item">点赞 {{ atc.Likes }}</span>
                                    <span class="stat-item">阅读 {{ atc.PageViews }}</span>
                                </div>
                            </div>
                        </div>
                    </el-card>
                </TransitionGroup>
            </div>

            <!-- 侧边栏区域 -->
            <aside class="sidebar-container desktop-sidebar">
                <!-- 用户信息卡片 -->
                <el-card class="user-info-card">
                    <div class="user-info-content">
                        <el-avatar class="user-avatar" src="/home-avator-hitori.webp" />
                        <h2 class="user-name">blockche</h2>
                        <p class="user-motto">无限进步</p>
                        <div class="user-stats">
                            <div class="info">
                                <span>文章</span>
                                <span>{{ list.length }}</span>
                            </div>
                            <div class="info" @click="goToCategory">
                                <span>分类</span>
                                <span>{{ store.allCategory.length }}</span>
                            </div>
                            <div class="info" @click="goToTag">
                                <span>标签</span>
                                <span>{{ store.allTag.length }}</span>
                            </div>
                        </div>
                    </div>
                </el-card>

                <!-- 排序按钮 -->
                <el-card class="sort-button" @click="switchSort">
                    <span class="sort-text">{{ sortButtonText }}</span>
                </el-card>

                <!-- 网站信息卡片 -->
                <el-card class="web-info" @click="recentVisitorsRef?.open()">
                    <div class="web-info-header">
                        <el-icon :size="20">
                            <Histogram />
                        </el-icon>
                        <span class="web-info-title">网站信息</span>
                    </div>
                    <div class="web-info-item">
                        <span>总访客数</span>
                        <span>{{ webInfo.VisitorCount }}</span>
                    </div>
                    <div class="web-info-item">
                        <span>近30天访客数</span>
                        <span>{{ webInfo.RecentVisitorsCount }}</span>
                    </div>
                    <div class="web-info-item">
                        <span>总阅读量</span>
                        <span>{{ webInfo.TotalPageViews }}</span>
                    </div>
                </el-card>
            </aside>
        </div>

        <recentVisitors ref="recentVisitorsRef" />
    </div>
</template>

<style scoped>
/* ==================== 布局容器 ==================== */
.home-container {
    width: 100%;
}

.home-content {
    display: flex;
    width: 100%;
}

.article-list-container {
    margin-left: 9vw;
    width: 70vw;
}

.sidebar-container {
    display: flex;
    flex-direction: column;
    margin-left: auto;
    margin-right: 5vw;
}

/* ==================== 文章卡片 ==================== */
.article-card {
    min-height: 160px;
    width: 60vw;
    cursor: pointer;
    transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
    border-radius: 12px;
    overflow: hidden;
}

.article-card:hover {
    transform: scale(1.07);
    box-shadow: 0 0 0 8px #dadada;
}

.article-content {
    display: flex;
    flex-direction: column;
    height: 100%;
    gap: 12px;
}

/* 标题区域 */
.article-header {
    padding-bottom: 8px;
    border-bottom: 1px solid #f0f0f0;
}

.article-title {
    font-size: 1.5em;
    font-weight: 600;
    line-height: 1.4;
    color: #1a1a1a;
    margin: 0;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
    text-overflow: ellipsis;
    word-break: break-word;
}

/* 摘要区域 */
.article-body {
    flex: 1;
    min-height: 0;
}

.article-excerpt {
    font-size: 0.95em;
    line-height: 1.6;
    color: #666;
    margin: 0;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
    text-overflow: ellipsis;
    word-break: break-word;
}

/* 底部信息区域 */
.article-footer {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding-top: 12px;
    border-top: 1px solid #f0f0f0;
    gap: 12px;
}

.article-meta {
    flex: 1;
    min-width: 0;
}

.article-tags {
    display: flex;
    gap: 8px;
    flex-wrap: wrap;
    align-items: center;
}

.article-tag {
    margin: 0;
    font-size: 0.85em;
}

.more-tags {
    font-size: 0.85em;
    color: #999;
    padding: 0 4px;
}

/* 统计信息 */
.article-stats {
    display: flex;
    gap: 1rem;
}

.stat-item {
    color: #cccccc;
    font-size: 0.9em;
}

/* ==================== 用户信息卡片 ==================== */
.user-info-card {
    width: 15vw;
    margin-left: 2vw;
}

.user-info-content {
    display: flex;
    flex-direction: column;
    align-items: center;
}

.user-avatar {
    width: 6vw;
    height: 6vw;
    margin-top: 1vh;
}

.user-name {
    margin: 2vh 0 0;
    font-weight: bold;
    font-size: 1.6em;
}

.user-motto {
    margin: 0.4vh 0 0;
    font-size: 1.1em;
    color: #666;
}

.user-stats {
    display: flex;
    flex-direction: column;
    width: 100%;
    margin-top: 2vh;
}

/* ==================== 信息项 ==================== */
.info {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-top: 1.5vh;
    padding: 6px 0.5vw;
    border-radius: 8px;
    border: 1px solid #e0e0e0;
    cursor: default;
    transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.info:hover {
    cursor: pointer;
    background: #ffffff;
    border-color: #2196f3;
    box-shadow: 0 2px 8px rgba(33, 150, 243, 0.2);
    transform: translateY(-2px);
}

.info span:first-child {
    font-size: 0.9rem;
    color: #666;
}

.info span:last-child {
    font-weight: 500;
    color: #666;
}

/* ==================== 排序按钮 ==================== */
.sort-button {
    width: 15vw;
    margin-left: 2vw;
    background-color: #eeeeee;
    cursor: pointer;
    transition: all 0.2s ease-in-out;
}

.sort-button:hover {
    box-shadow: 0 0 0 8px #dadada;
}

.sort-text {
    font-size: 1.1rem;
    font-weight: bold;
}

/* ==================== 网站信息卡片 ==================== */
.web-info {
    display: flex;
    flex-direction: column;
    width: 15vw;
    margin-left: 2vw;
    cursor: pointer;
    transition: all 0.2s ease-in-out;
}

.web-info:hover {
    box-shadow: 0 0 0 8px #dadada;
}

.web-info-header {
    display: flex;
    align-items: center;
    gap: 0.3rem;
    margin-bottom: 0.5rem;
}

.web-info-title {
    font-size: 1.2rem;
    font-weight: bold;
}

.web-info-item {
    display: flex;
    justify-content: space-between;
    margin-top: 0.5rem;
    color: #666;
}

/* ==================== 列表过渡动画 ==================== */
.list-move,
.list-enter-active,
.list-leave-active {
    transition: all 0.5s ease;
}

.list-enter-from,
.list-leave-to {
    opacity: 0;
    transform: translateX(30px);
}

.list-leave-active {
    position: absolute;
}

/* ==================== 响应式布局 ==================== */
@media screen and (max-width: 768px) {
    .home-content {
        flex-direction: column;
        align-items: center;
    }

    .article-list-container {
        margin-left: 0;
        width: 95vw;
    }

    .sidebar-container {
        display: none;
    }

    .article-card {
        width: 90vw !important;
        min-height: 140px;
    }

    .article-title {
        font-size: 1.2em !important;
    }

    .article-card:hover {
        transform: translateY(-2px);
    }

    .article-tag {
        font-size: 0.75em;
    }

    .article-stats {
        font-size: 0.85em;
    }
}

@media screen and (max-width: 480px) {
    .article-list-container {
        width: 98vw;
    }

    .article-card {
        width: 95vw !important;
    }

    .article-title {
        font-size: 1.1em !important;
    }

    .article-excerpt {
        font-size: 0.9em;
    }
}

/* ==================== 触摸设备优化 ==================== */
@media (hover: none) and (pointer: coarse) {
    .article-card:hover {
        transform: none;
        box-shadow: none;
    }

    .article-card:active {
        transform: scale(0.98);
        box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
    }

    .info:hover,
    .sort-button:hover,
    .web-info:hover {
        box-shadow: none;
        transform: none;
    }

    .info:active,
    .sort-button:active,
    .web-info:active {
        box-shadow: 0 0 0 4px #dadada;
        transform: scale(0.98);
    }
}
</style>
