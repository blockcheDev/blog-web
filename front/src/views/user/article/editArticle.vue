<script setup lang="ts">
import { nextTick, onMounted, reactive, ref } from 'vue';
import api from '@/api/api'
import PageHeader from '@/layouts/PageHeader.vue';
import { allCategory, getAllCategory, allTag, getAllTag, type Article, defaultData, type Tag } from '@/store/article';
import router from '@/router';
import { useRoute } from 'vue-router';
const route = useRoute()

const article: Article = reactive(JSON.parse(JSON.stringify(defaultData))) // 深拷贝，使得data是独立于defaultData的副本
const form = reactive({
    ID: 0,
    CategoryID: 0,
    Title: '',
    Content: '',
    Type: 0,
    Tags: [] as number[],
})

onMounted(async () => {
    getAllCategory()
    getAllTag()
    try {
        const res = await api.getArticle(route.params.id)
        Object.assign(article, res.data)
    } catch (err) {
        console.error(err)
    }

    form.ID = article.ID
    form.CategoryID = article.Category.ID
    form.Title = article.Title
    form.Content = article.Content
    form.Type = article.Type
    form.Tags = article.Tags.map((tag: Tag) => tag.ID)
})

const submitForm = async (form: any) => {
    try {
        console.log(form)
        const res = await api.modifyArticle(form)
        const ID = res.data.ID
        router.push(`/article/${ID}`)
    } catch (err) {
        console.error(err)
    }
}

// 新建标签
const openDialogTag = ref(false)
const newTag = reactive({
    Name: "",
})
const submitTag = async (data: any) => {
    try {
        const res = await api.createTag(data)
        newTag.Name = ""
        getAllTag()
    } catch (err) {
        console.error(err)
    }
}

// 新建分类
const openDialogCategory = ref(false)
const newCategory = reactive({
    Name: "",
})
const submitCategory = async (data: any) => {
    try {
        const res = await api.createCategory(data)
        newCategory.Name = ""
        getAllCategory()
    } catch (err) {
        console.error(err)
    }
}

</script>

<template>
    <div>
        <el-card style="width: 80vw;">
            <PageHeader />
        </el-card>
        <el-card style="width: 80vw;">
            <el-form :model="form" label-width="80px">
                <el-form-item label="标题" prop="title">
                    <el-input v-model="form.Title"></el-input>
                </el-form-item>
                <el-form-item label="正文" prop="content">
                    <v-md-editor v-model="form.Content" height="400px"></v-md-editor>
                </el-form-item>
                <el-form-item label="文章类型">
                    <el-radio-group v-model="form.Type">
                        <el-radio :label="0">原创</el-radio>
                        <el-radio :label="1">转载</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="文章分类">
                    <el-select v-model="form.CategoryID" filterable placeholder="请选择文章的分类">
                        <el-option v-for="category in allCategory" :label="category.Name" :value="category.ID" />
                    </el-select>
                    <el-button style="margin-left: 2em;" @click="openDialogCategory = true">添加分类</el-button>
                </el-form-item>
                <el-form-item label="标签">
                    <el-select v-model="form.Tags" multiple filterable placeholder="选择标签">
                        <el-option v-for="tag in allTag" :key="tag.ID" :label="tag.Name" :value="tag.ID" />
                    </el-select>
                    <el-button style="margin-left: 2em;" @click="openDialogTag = true">添加标签</el-button>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="submitForm(form)">发布文章</el-button>
                </el-form-item>
            </el-form>
        </el-card>

        <el-dialog v-model="openDialogTag" title="添加标签">
            <el-form>
                <el-form-item label="标签名称">
                    <el-input v-model="newTag.Name" style="width: 30%;"></el-input>
                    <el-button type="primary" style="margin-left: 2em;" @click="submitTag(newTag)">添加</el-button>
                </el-form-item>
            </el-form>
        </el-dialog>
        <el-dialog v-model="openDialogCategory" title="添加分类">
            <el-form>
                <el-form-item label="分类名称">
                    <el-input v-model="newCategory.Name" style="width: 30%;"></el-input>
                    <el-button type="primary" style="margin-left: 2em;" @click="submitCategory(newCategory)">添加</el-button>
                </el-form-item>
            </el-form>
        </el-dialog>
    </div>
</template>