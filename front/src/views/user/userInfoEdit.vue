<template>
    <el-dialog v-model="dialogFormVisible" title="编辑用户信息">
        <el-form v-model="editForm" label-width="80px">
            <el-form-item label="用户名">
                <el-input v-model="editForm.Name" autocomplete="on" disabled="true" />
            </el-form-item>
            <el-form-item label="邮箱">
                <el-input v-model="editForm.Email" autocomplete="on" disabled="true" />
            </el-form-item>
            <el-form-item label="手机号码">
                <el-input v-model="editForm.Telephone" autocomplete="on" />
            </el-form-item>
            <el-form-item label="性别">
                <el-radio-group v-model="editForm.Gender">
                    <el-radio label="男">男</el-radio>
                    <el-radio label="女">女</el-radio>
                    <el-radio label="未知">未知</el-radio>
                </el-radio-group>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="saveEdit(editForm)">保存</el-button>
                <el-button @click="dialogFormVisible = false">取消</el-button>
            </el-form-item>
        </el-form>

    </el-dialog>
</template>

<script setup lang="ts">
import api from '@/api/api';
import { ref } from 'vue';

const props = defineProps(['user'])

const editForm = ref({
    Name: "",
    Email: "",
    Telephone: "",
    Gender: "",
});
const openDialog = () => {
    Object.assign(editForm.value, props.user)
    dialogFormVisible.value = true
}
const dialogFormVisible = ref(false)
const saveEdit = async (form: typeof editForm.value) => {
    try {
        const res = await api.modifyUser(form)
        Object.assign(props.user, form)
        console.log(form)
        dialogFormVisible.value = false
        // location.reload()
    } catch (err) {
        console.error("更新用户信息错误", err)
    }
}

defineExpose({
    openDialog
})

</script>