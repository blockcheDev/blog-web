<script setup lang="ts">
import router from '@/router';
import * as article from '@/store/article'
import type { Category } from '@/store/article';
import { onMounted, ref } from 'vue';
import PageHeaderProp from '@/layouts/PageHeaderProp.vue';

const title = ref("分类列表")

const list_category = ref<Category[]>([])
const sort_status = ref<number>(1)

onMounted(async () => {
    await article.getAllCategory()
    list_category.value = article.allCategory
    switchSort()
})

const switchSort = () => {
    if (sort_status.value === 0) {
        list_category.value = list_category.value.sort((a, b) => {
            return a.ID - b.ID
        })
        sort_status.value = 1
    } else {
        list_category.value = list_category.value.sort((a, b) => {
            return a.Name.localeCompare(b.Name)
        })
        sort_status.value = 0
    }
}

const goToCategory = (id: any) => {
    router.push(`/category/${id}`)
}

</script>
<template>
    <div>
        <el-card style="width: 85vw;">
            <PageHeaderProp :data="title" />
        </el-card>
        <div class="tag-container">
            <TransitionGroup name="list" tag="ul" class="card-wrapper">
                <el-card v-for="cate of list_category" :key="cate.ID" @click="goToCategory(cate.ID)"
                    class="custom-card">
                    <span class="tag-text">{{ cate.Name }}</span>
                </el-card>
            </TransitionGroup>
            <div class="card-wrapper">
                <el-card @click="switchSort()" class="custom-card" style="width: 10vw; background-color:#dadada;">
                    <span class="tag-text">{{ sort_status === 0 ? "时间排序" : "字典序排序" }}</span>
                </el-card>
            </div>
        </div>
    </div>
</template>

<style scoped>
.tag-container {
    width: 85vw;
    display: flex;
    justify-content:space-between;
    margin: auto;
}

.card-wrapper {
    display: flex;
    flex-wrap: wrap;
    padding: 0 0 0 0;
    margin: 0 0 0 0;
    gap: 2vw;
}

.custom-card {
    width: auto;
    height: 7vh;
    margin: 0 0 0 0;
    display: flex;
    justify-content: center;
    align-items: center;
}

.custom-card:hover {
    box-shadow: 0 0 0 8px #dadada;
    transition: all 0.2s ease-in-out;
}

.tag-text {
    font-size: clamp(1rem, 1.5vw, 1.2rem);
    font-weight: 600;
    text-align: center;
}

.list-move,
/* 对移动中的元素应用的过渡 */
.list-enter-active,
.list-leave-active {
    transition: all 0.5s ease;
}

.list-enter-from,
.list-leave-to {
    opacity: 0;
    transform: translateX(30px);
}

/* 确保将离开的元素从布局流中删除
  以便能够正确地计算移动的动画。 */
.list-leave-active {
    position: absolute;
}
</style>