<template>
    <div>
        <el-card style="width: 60vw;">
            <PageHeader />
        </el-card>
        <el-card style="width: 60vw;">
            <div style="display: flex;">
                <div style="margin: 7vh; display: flex; flex-flow: column;">
                    <!-- <el-avatar :size="250" src="https://avatars.githubusercontent.com/u/89156012?v=4" /> -->
                    <el-avatar :size="250" :src=user.AvatarUrl />
                    <div style="margin-top: 1em;">
                        <span style=" font-size: 1.5em; font-weight: bold;">{{ user.Name }}</span>
                        <span style=" font-size: 1.3em; color: gray;"> ID:{{ user.ID }}</span>
                    </div>
                    <el-button color="#f3f4f6" style="margin-top: 0.5em; border-radius: 10px;"
                        @click="dialog.openDialog()">编辑信息</el-button>
                </div>
                <div style="margin: auto 7vh auto 7vh;">
                    <el-card style="width: 25vw;">
                        <div style="display: flex; ">
                            <User style="width: 1.5em;" />
                            <span style="margin-left: 1em;">账号</span>
                            <span style="margin-left: auto;">{{ user.Name }}</span>
                        </div>
                    </el-card>
                    <el-card style="width: 25vw;">
                        <div style="display: flex;">
                            <Message style="width: 1.5em;" />
                            <span style="margin-left: 1em;">邮箱</span>
                            <span style="margin-left: auto;">{{ user.Email }}</span>
                        </div>
                    </el-card>
                    <el-card style="width: 25vw;">
                        <div style="display: flex;">
                            <iphone style="width: 1.5em;" />
                            <span style="margin-left: 1em;">手机号码</span>
                            <span style="margin-left: auto;">{{ user.Telephone }}</span>
                        </div>
                    </el-card>
                    <el-card style="width: 25vw; margin-bottom: 0;">
                        <div style="display: flex;">
                            <Male v-if="user.Gender === '男'" style="width: 1.5em;" />
                            <Female v-else-if="user.Gender === '女'" style="width: 1.5em;" />
                            <Aim v-else style="width: 1.5em;" />
                            <span style="margin-left: 1em;">性别</span>
                            <span style="margin-left: auto;">{{ user.Gender }}</span>
                        </div>
                    </el-card>
                </div>
            </div>
        </el-card>

        <userInfoEdit :user="user" ref="dialog" />
    </div>
</template>

<script setup lang="ts">
import api from '@/api/api';
import {
    Iphone,
    User,
    Female,
    Male,
    Aim,
    Message,
} from '@element-plus/icons-vue'
import { onMounted, reactive, ref } from 'vue';
import userInfoEdit from './userInfoEdit.vue';
import PageHeader from '@/layouts/PageHeader.vue';

const dialog = ref()

const user = reactive({
    ID: 0,
    Name: "",
    Email: "",
    Telephone: "",
    Gender: "",
    AvatarUrl: "",
})

onMounted(async () => {
    try {
        const res = await api.getUser()
        Object.assign(user, res.data)
        if (user.AvatarUrl === "") {
            user.AvatarUrl = "/avatar.png"
        }
    } catch (err) {
        console.error("获取用户信息失败", err)
    }
})

</script>

<style scoped></style>