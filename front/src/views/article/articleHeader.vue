<script lang="ts" setup>
import api from '@/api/api';
import router from '@/router';
import type { Article } from '@/store/article';
import { Calendar, Edit } from '@element-plus/icons-vue'
import dayjs from 'dayjs'
import { onMounted, ref } from 'vue';
const formatDate = (date: any) => {
    return dayjs(date).format("YYYY-MM-DD")
}
const goBack = () => {
    router.back()
}
const props = defineProps<{
    data: Article
}>()

const handleTagClick = (tag_id: number) => {
    router.push(`/tag/${tag_id}`)
}
const handleCategoryClick = (category_id: number) => {
    router.push(`/category/${category_id}`)
}
const handleLikeClick = async () => {
  try {
    await api.addArticleLikeRecord(props.data.ID)
  } catch (err) {
    console.error("点赞失败", err)
    return
  }
  props.data.Likes += 1;
}

</script>

<template>
    <div class="article-header">
        <div class="title-wrapper">
            <h2 class="title">{{ data.Title }}</h2>
        </div>

        <div class="meta-container">
            <div class="meta-group">
                <div class="meta-item">
                    <span class="meta-text">发布于 {{ formatDate(data.CreatedAt) }}</span>
                </div>
                <div class="meta-item">
                    <span class="meta-text">修改于 {{ formatDate(data.ModifiedAt) }}</span>
                </div>
            </div>

            <div class="meta-group">
                <div class="meta-item">
                    <span class="meta-text">{{ data.Type === 0 ? '原创' : '转载' }}</span>
                </div>
                <div class="meta-item-clickable" @click="handleCategoryClick(data.Category.ID)">
                    <span class="meta-text">分类 {{ data.Category.Name }}</span>
                </div>
                <div class="meta-item">
                    <span class="meta-text">阅读量 {{ data.PageViews }}</span>
                </div>
                <div class="meta-item-clickable" @click="handleLikeClick()">
                    <span class="meta-text">点赞 {{ data.Likes }}</span>
                </div>
            </div>

            <div class="meta-group" style="margin-top: 0.5rem;">
                <el-tag v-for="(tag, index) in data.Tags" :key="index" class="tag-item"
                    @click="handleTagClick(tag.ID)">
                    {{ tag.Name }}
                </el-tag>
            </div>
        </div>
    </div>
</template>

<style scoped lang="scss">
.article-header {
  margin: 0 auto;          /* 水平居中 */
  padding: 1rem;         /* 内边距 */
//   transition: all 0.3s ease; /* 过渡动画 */

  .title-wrapper {
    margin-bottom: 1rem; /* 下边距 */
    
    .title {
      font-size: 2rem;      /* 字体大小 */
      text-align: center;   /* 水平居中 */
      margin: 0;            /* 清除默认边距 */
    }
  }

  .meta-container {
    display: flex;           /* 弹性布局 */
    flex-direction: column; /* 垂直排列 */
    gap: 0.8rem;              /* 子元素间距 */
  }

  .meta-group {
    display: flex;
    flex-wrap: wrap;        /* 允许换行 */
    align-items: center;    /* 垂直居中 */
    justify-content: center;/* 水平居中 */
    gap: 1rem;              /* 元素间距 */
  }

/* 通用信息项样式 */
  .meta-item, .meta-item-clickable {
    padding: 0.5rem 0.9rem;   /* 内边距 */
    background: var(--el-fill-color-light); /* 背景色 */
    border-radius: 8px;     /* 圆角半径 */
    display: inline-flex;   /* 行内弹性布局 */
    align-items: center;    /* 垂直居中 */
    transition: background-color 0.3s ease, transform 0.5s ease; /* 背景与变换过渡 */
  }

  .meta-item-clickable {
    cursor: pointer;     /* 手型光标 */
    
    &:hover {
      transform: translateY(-3px); /* 悬停上移 */
      box-shadow: 0 2px 8px var(--el-box-shadow-light); /* 阴影效果 */
    }
  }

  /* 文字样式 */
  .meta-text {
    font-size: 0.8rem;      /* 字体大小 */
    color: var(--el-text-color-secondary); /* 文字颜色 */
  }

  /* 标签样式 */
  .tag-item {
    cursor: pointer;     /* 手型光标 */
    transition: transform 0.5s ease; /* 变换动画 */
    
    &:hover {
    transform: translateY(-3px); /* 悬停上移 */
    box-shadow: 0 2px 8px var(--el-box-shadow-light); /* 阴影效果 */
    }
  }
}

/* 移动端适配 */
@media (max-width: 768px) {
  .article-header {
    padding: 1rem;         /* 缩小内边距 */
    
    .title {
      font-size: 1.5rem;  /* 标题缩小 */
    }
    
    .tag-item[effect="plain"] {
      font-size: 0.8rem;  /* 标签缩小 */
    }
  }
}
</style>