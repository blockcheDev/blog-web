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
                <v-md-editor :model-value="data.Content" mode="preview"></v-md-editor>
            </el-card>
            <el-card>
                <info :data="data" />
            </el-card>
            <el-card>
                <comment :articleID="route.params.id" />
            </el-card>
        </div>
    </div>
</template>

<style scoped>
.el-card {
    width: 60vw;
}
</style>