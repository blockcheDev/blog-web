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
                <el-tag v-for="(tag, index) in data.Tags" :key="index" class="article-tag"
                    size="small"
                    effect="plain"
                    @click="handleTagClick(tag.ID)">
                    {{ tag.Name }}
                </el-tag>
            </div>
        </div>
    </div>
</template>

<style scoped lang="scss">
.article-header {
  margin: 0 auto;
  padding: 0;

  .title-wrapper {
    margin-bottom: 2rem;
    padding-bottom: 1.25rem;
    border-bottom: 1px solid #e9ecef;
    
    .title {
      font-size: 2rem;
      font-weight: 700;
      text-align: center;
      margin: 0;
      color: #1a1a1a;
      line-height: 1.5;
      letter-spacing: -0.01em;
    }
  }

  .meta-container {
    display: flex;
    flex-direction: column;
    gap: 0.85rem;
  }

  .meta-group {
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    justify-content: center;
    gap: 0.65rem;
  }

/* 通用信息项样式 */
  .meta-item, .meta-item-clickable {
    padding: 0.45rem 0.9rem;
    background: #f5f6f7;
    border-radius: 18px;
    display: inline-flex;
    align-items: center;
    transition: all 0.25s ease;
    border: 1px solid transparent;
  }

  .meta-item-clickable {
    cursor: pointer;
    
    &:hover {
      transform: translateY(-2px);
      background: #e9ecef;
      border-color: #dee2e6;
      box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
    }
    
    &:active {
      transform: translateY(0);
    }
  }

  /* 文字样式 */
  .meta-text {
    font-size: 0.85rem;
    color: #6c757d;
    font-weight: 500;
  }

  /* 标签样式 */
  .article-tag {
    margin: 0;
    font-size: 0.85rem;
    cursor: pointer;
    transition: all 0.2s ease;
    border-radius: 18px;
    
    &:hover {
      opacity: 0.8;
    }
  }
}

/* 平板端适配 */
@media (min-width: 769px) and (max-width: 1024px) {
  .article-header {
    .title {
      font-size: 1.85rem;
    }
    
    .title-wrapper {
      margin-bottom: 1.75rem;
      padding-bottom: 1.1rem;
    }
    
    .meta-group {
      gap: 0.6rem;
    }
  }
}

/* 移动端适配 */
@media (max-width: 768px) {
  .article-header {
    .title-wrapper {
      margin-bottom: 1.4rem;
      padding-bottom: 0.9rem;
    }
    
    .title {
      font-size: 1.5rem;
    }
    
    .meta-container {
      gap: 0.75rem;
    }
    
    .meta-group {
      gap: 0.55rem;
    }
    
    .meta-item, .meta-item-clickable {
      padding: 0.38rem 0.75rem;
      border-radius: 14px;
    }
    
    .meta-text {
      font-size: 0.8rem;
    }
    
    .article-tag {
      font-size: 0.8rem;
      border-radius: 14px;
    }
  }
}

/* 小屏幕移动端适配 */
@media (max-width: 480px) {
  .article-header {
    .title-wrapper {
      margin-bottom: 1.2rem;
      padding-bottom: 0.8rem;
    }
    
    .title {
      font-size: 1.35rem;
    }
    
    .meta-container {
      gap: 0.7rem;
    }
    
    .meta-group {
      gap: 0.5rem;
    }
    
    .meta-item, .meta-item-clickable {
      padding: 0.35rem 0.7rem;
      border-radius: 12px;
    }
    
    .meta-text {
      font-size: 0.75rem;
    }
    
    .article-tag {
      font-size: 0.75rem;
      border-radius: 12px;
    }
  }
}
</style>