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
    <div style="display: flex; flex-flow: column;">
        <div>
            <span style="font-weight: bold;">文章作者：</span>
            <!-- <a href="http://www.hitori.cn" style="color:rgb(0, 225, 225)">{{ data.UserName }}</a> -->
            <span>{{ data.User.Name }}</span>
        </div>
        <div style="margin-top: 0.5em;">
            <span style="font-weight: bold;">文章链接：</span>
            <a :href="url" style="color:rgb(0, 225, 225)">{{ url }}</a>
        </div>
        <div style="margin-top: 0.5em;">
            <span style="font-weight: bold;">版权声明：</span>
            <span>本博客所有文章除特别声明外，均采用 </span>
            <a href="https://creativecommons.org/licenses/by-nc-sa/4.0/" style="color:rgb(0, 225, 225)">CC BY-NC-SA 4.0</a>
            <span> 许可协议。转载请注明文章出处。</span>
        </div>
    </div>
</template>