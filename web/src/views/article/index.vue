<script setup lang="ts">
import api from '@/api/api';
import { useRoute } from 'vue-router'
import { reactive, onMounted, ref } from 'vue'
import { marked } from 'marked'
import hljs from 'highlight.js';
import articleHeader from './articleHeader.vue'
import { defaultData } from '@/store/article'
import type { Article } from '@/store/article'

const data: Article = reactive(JSON.parse(JSON.stringify(defaultData))) // 深拷贝，使得data是独立于defaultData的副本

const route = useRoute()
const html = ref()
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
        html.value = marked(data.Content)
    } catch (err) {
        console.error("文章获取失败", err)
    }
    try {
        const res = await api.getTagByArticle(data.ID)
        Object.assign(data.tags, res.data)
    } catch (err) {
        console.error(err)
    }
    try {
        const res = await api.getCategory(data.CategoryID)
        data.CategoryName = res.data.Name
    } catch (err) {
        console.error(err)
    }
})

</script>

<template>
    <div v-if="data.Title != ''" style="display: flex; justify-content: center;">
        <!-- <div>
            <el-card style="width: 15vw; margin-left: 5vw; margin-right: 2vw;">
                <div style="display: flex;">
                    <el-avatar style="width: 5vw; height: 5vw; margin: auto;"
                        src="https://avatars.githubusercontent.com/u/89156012?v=4" />
                </div>
            </el-card>
        </div> -->
        <div>
            <el-card>
                <articleHeader :data="data" />
            </el-card>
            <el-card>
                <!-- <div v-html="html"></div> -->
                <v-md-editor :model-value="data.Content" mode="preview"></v-md-editor>
            </el-card>
        </div>
    </div>
</template>

<style scoped>
.el-card {
    width: 60vw;
}
</style>