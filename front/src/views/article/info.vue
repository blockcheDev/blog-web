<script setup lang="ts">
import type { Article } from '@/store/article';
import { onMounted, ref } from 'vue';

const props = defineProps<{
    data: Article
}>()
const url = ref("")
onMounted(() => {
    url.value = location.href
})
</script>

<template>
    <div class="info-container">
        <div class="info-item">
            <span class="info-label">文章作者：</span>
            <!-- <a href="http://www.hitori.cn" class="info-link">{{ data.UserName }}</a> -->
            <span class="info-value">{{ data.User.Name }}</span>
        </div>
        <div class="info-item">
            <span class="info-label">文章链接：</span>
            <a :href="url" class="info-link">{{ url }}</a>
        </div>
        <div class="info-item">
            <span class="info-label">版权声明：</span>
            <span class="info-value">本博客所有文章除特别声明外，均采用 </span>
            <a href="https://creativecommons.org/licenses/by-nc-sa/4.0/" class="info-link">CC BY-NC-SA 4.0</a>
            <span class="info-value"> 许可协议。转载请注明文章出处。</span>
        </div>
    </div>
</template>

<style scoped>
.info-container {
    display: flex;
    flex-direction: column;
    gap: 0.85rem;
    padding: 0.25rem 0;
}

.info-item {
    display: flex;
    flex-wrap: wrap;
    align-items: baseline;
    line-height: 1.75;
    padding: 0.6rem 0;
}

.info-item:not(:last-child) {
    border-bottom: 1px solid #f5f5f5;
}

.info-label {
    font-weight: 600;
    flex-shrink: 0;
    color: #2c3e50;
    font-size: 0.9em;
    min-width: 85px;
}

.info-value {
    word-break: break-word;
    color: #5a6c7d;
    font-size: 0.88em;
}

.info-link {
    color: #0066cc;
    word-break: break-all;
    text-decoration: none;
    transition: all 0.2s ease;
    font-weight: 500;
    border-bottom: 1px solid transparent;
    font-size: 0.88em;
}

.info-link:hover {
    color: #0052a3;
    border-bottom-color: #0052a3;
}

/* 移动端适配 */
@media (max-width: 768px) {
    .info-container {
        gap: 0.75rem;
        padding: 0.2rem 0;
    }
    
    .info-item {
        font-size: 0.9em;
        padding: 0.5rem 0;
        flex-direction: column;
        align-items: flex-start;
    }
    
    .info-label {
        margin-bottom: 0.35em;
        min-width: auto;
    }
    
    .info-value,
    .info-link {
        font-size: 0.93em;
    }
}

/* 小屏幕移动端 */
@media (max-width: 480px) {
    .info-container {
        gap: 0.65rem;
        padding: 0.15rem 0;
    }
    
    .info-item {
        font-size: 0.85em;
        padding: 0.45rem 0;
    }
    
    .info-label {
        font-size: 0.92em;
        margin-bottom: 0.3em;
    }
    
    .info-value,
    .info-link {
        font-size: 0.9em;
    }
}
</style>