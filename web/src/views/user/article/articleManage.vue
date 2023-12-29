<script setup lang="ts">
import PageHeader from '@/layouts/PageHeader.vue';
import { list } from '@/store/article';
import api from '@/api/api';
import { onMounted, ref } from 'vue';
import dayjs from 'dayjs';
import router from '@/router';
import { computed } from '@vue/reactivity';

const formatDate = (date: any) => {
    return dayjs(date).format("YYYY-MM-DD HH:mm:ss")
}
const getCategoryName = async (id: any) => {
    try {
        const res = await api.getCategory(id)
        const name: string = res.data.Name
        return name
    } catch (err) {
        console.error(err)
        return "NULL"
    }
}

const getAllArticle = async () => {
    try {
        const res = await api.getArticle("all")
        Object.assign(list, res.data)
    } catch (err) {
        console.error(err)
    }
}

onMounted(async () => {
    await getAllArticle()
    for (var i = 0; i < list.length; i++) {
        list[i].CategoryName = await getCategoryName(list[i].CategoryID)
    }
})

const goToArticle = (id: any) => {
    router.push(`/article/${id}`)
}
const goToEditArticle = (id: any) => {
    router.push(`/user/edit/article/${id}`)
}

const openDeleteDialog = ref(false)
const deleteID = ref(0)
const openDeleteDialogFunc = (id: any) => {
    deleteID.value = id
    openDeleteDialog.value = true
}
const deleteArticle = async () => {
    try {
        const res = await api.deleteArticle(deleteID.value)
        location.reload()
        openDeleteDialog.value = false
    } catch (err) {
        console.error(err)
    }
}

const search = ref("")
const filterData = computed(() => list.filter((data) => !search.value || data.Title.toLowerCase().includes(search.value.toLowerCase())))

</script>

<template>
    <div>
        <el-card style="width: 85vw;">
            <PageHeader />
        </el-card>
        <el-card style="width: 85vw;">
            <el-table :data="filterData" stripe style="width: 100%">
                <el-table-column sortable prop="ID" label="ID" width="100" />
                <el-table-column sortable label="标题" width="300">
                    <template v-slot="scope">
                        <span style="color: blue;" @click="goToArticle(scope.row.ID)">{{ scope.row.Title }}</span>
                    </template>
                </el-table-column>
                <el-table-column sortable prop="CategoryName" label="分类" />
                <el-table-column sortable label="类型">
                    <template v-slot="scope">
                        <span>{{ scope.row.Type === 0 ? "原创" : "转载" }}</span>
                    </template>
                </el-table-column>
                <el-table-column sortable label="创建时间">
                    <template v-slot="scope">
                        <span>{{ formatDate(scope.row.CreatedAt) }}</span>
                    </template>
                </el-table-column>
                <el-table-column sortable label="修改时间">
                    <template v-slot="scope">
                        <span>{{ formatDate(scope.row.UpdatedAt) }}</span>
                    </template>
                </el-table-column>
                <el-table-column>
                    <template #header>
                        <el-input v-model="search" placeholder="搜索标题" />
                    </template>
                    <template v-slot="scope">
                        <el-button type="primary" size="small" @click="goToEditArticle(scope.row.ID)">编辑</el-button>
                        <el-button type="danger" size="small" @click="openDeleteDialogFunc(scope.row.ID)">删除</el-button>
                    </template>
                </el-table-column>

            </el-table>
        </el-card>

        <el-dialog v-model="openDeleteDialog" title="您确定要删除这篇文章吗？">
            <el-button type="danger" @click="deleteArticle()">确认删除</el-button>
            <el-button @click="openDeleteDialog = false">取消</el-button>
        </el-dialog>
    </div>
</template>