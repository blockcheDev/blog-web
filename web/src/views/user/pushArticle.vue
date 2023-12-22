<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue';
import api from '@/api/api'
import PageHeader from '@/layouts/PageHeader.vue';
import router from '@/router';

const form = reactive({
    CategoryID: undefined,
    Title: '',
    Content: '',
    Type: 0,
    Tags: [],
})

const allCategory = ref()
const allTag = ref()

onMounted(async () => {
    try {
        const res = await api.getCategory("all")
        allCategory.value = res.data
    } catch (err) {
        console.error(err)
    }
    getAllTag()
})
const getAllTag = async () => {
    try {
        const res = await api.getTag("all")
        allTag.value = res.data
    } catch (err) {
        console.error(err)
    }
}

const submitForm = async (form: any) => {
    try {
        console.log(form)
        const res = await api.pushArticle(form)
        const ID = res.data.ID
        router.push(`/article/${ID}`)
    } catch (err) {
        console.error(err)
    }
}

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

</script>

<template>
    <div>
        <el-card>
            <PageHeader />
        </el-card>
        <el-card>
            <el-form :model="form" label-width="80px">
                <el-form-item label="标题" prop="title">
                    <el-input v-model="form.Title"></el-input>
                </el-form-item>
                <el-form-item label="正文" prop="content">
                    <el-input type="textarea" v-model="form.Content"></el-input>
                </el-form-item>
                <el-form-item label="文章类型">
                    <el-radio-group v-model="form.Type">
                        <el-radio :label="0">原创</el-radio>
                        <el-radio :label="1">转载</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="文章分类">
                    <el-select v-model="form.CategoryID" placeholder="请选择文章的分类">
                        <el-option v-for="category in allCategory" :label="category.Name" :value="category.ID" />
                    </el-select>
                </el-form-item>
                <el-form-item label="标签">
                    <el-select v-model="form.Tags" multiple placeholder="选择标签">
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
    </div>
</template>