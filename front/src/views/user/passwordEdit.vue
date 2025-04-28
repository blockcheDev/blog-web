<template>
    <el-dialog v-model="dialogFormVisible" title="修改密码">
        <el-form v-model="editForm" label-width="80px">
            <el-form-item label="原密码">
                <el-input v-model="editForm.OldPassword" />
            </el-form-item>
            <el-form-item label="新密码">
                <el-input v-model="editForm.NewPassword" />
            </el-form-item>
            <el-form-item label="重复密码">
                <el-input v-model="editForm.AgainPassword" />
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
import { ElMessage } from 'element-plus'

const editForm = ref({
    OldPassword: "",
    NewPassword: "",
    AgainPassword: "",
});
const openDialog = () => {
    dialogFormVisible.value = true
}
const dialogFormVisible = ref(false)
const saveEdit = async (form: typeof editForm.value) => {
    if (form.OldPassword === "" || form.NewPassword === "" || form.AgainPassword === "") {
        ElMessage({
            message: "信息为空",
        })
    } else if (form.NewPassword != form.AgainPassword) {
        ElMessage({
            message: "新密码不一致",
        })
    } else if (form.NewPassword === form.OldPassword) {
        ElMessage({
            message: "新密码不能与旧密码相同",
        })
    } else {
        try {
            const res = await api.modifyUserPassword(form)
            dialogFormVisible.value = false
            // location.reload()
        } catch (err) {
            console.error("更新密码失败", err)
        }
    }
}

defineExpose({
    openDialog
})

</script>