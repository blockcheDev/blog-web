<template>
    <div style="height: 80vh; display: flex; justify-content: center; align-items: center;">
        <el-card style="height: 30vh; width: 50vw; display: flex; justify-content: center; align-items: center;">
            <span style="font-size: 2em; display: block;">{{ text }}</span>
        </el-card>
    </div>
</template>

<script lang="ts" setup>
import api from '@/api/api';
import { ref } from 'vue'
import { onMounted } from 'vue'
import router from '@/router';
import { useRoute } from 'vue-router'
const route = useRoute()
const text = ref("请稍等 OvO")

onMounted(async () => {
    const code = <string>route.query.code
    try {
        const res = await api.loginWithGithub(code)
        localStorage.setItem("token", res.headers.token)
        text.value = res.data.msg + " ♪(´▽｀)"
        setTimeout(() => {
            router.push("/home")
        }, 1500)
    } catch (err) {
        text.value = "对不起，你能重试一下吗 >_<"
        setTimeout(() => {
            router.push("/login")
        }, 1500)
        console.error(err)
    }
})
</script>