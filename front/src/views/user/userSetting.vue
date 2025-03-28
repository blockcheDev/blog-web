<script setup lang="ts">
import PageHeader from '@/layouts/PageHeader.vue';
import { reactive, ref } from 'vue';
import passwordEdit from './passwordEdit.vue';
import api from '@/api/api';
import router from '@/router';
const dialog = ref();

const dialogDelete = ref(false)
const formDelete = reactive({
    Password: '',
})
const submitDelete = async (data: any) => {
    try {
        console.log(data)
        const res = await api.deleteUser(data)
        localStorage.setItem("token", "")
        dialogDelete.value = false
        router.push("/home")
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
            <el-row justify="space-between" style="margin-bottom: 30px;">
                <span>密码</span>
                <el-button type="primary" @click="dialog.openDialog()">修改密码</el-button>
            </el-row>
            <el-row justify="space-between">
                <span>注销</span>
                <el-button type="danger" @click="dialogDelete = true">账号注销</el-button>
            </el-row>
            <passwordEdit ref="dialog" />
        </el-card>

        <el-dialog v-model="dialogDelete" title="您确定要注销账号吗？">
            <el-form v-model="formDelete" label-width="80px">
                <el-form-item label="账号密码">
                    <el-input v-model="formDelete.Password" />
                </el-form-item>
                <el-form-item>
                    <el-button type="danger" @click="submitDelete(formDelete)">确认注销</el-button>
                    <el-button @click="dialogDelete = false">取消</el-button>
                </el-form-item>
            </el-form>
        </el-dialog>
    </div>
</template>

<style scoped>
.el-card {
    width: 60vw;
}
</style>