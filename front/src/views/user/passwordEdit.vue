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
                <el-button type="primary" @click="saveEdit()">保存</el-button>
                <el-button @click="dialogFormVisible = false">取消</el-button>
            </el-form-item>
        </el-form>

    </el-dialog>
</template>

<script setup lang="ts">
import api from '@/api/api';
import { reactive, ref } from 'vue';
import { ElMessage } from 'element-plus'

const editForm = reactive({
    OldPassword: "",
    NewPassword: "",
    AgainPassword: "",
});
const openDialog = () => {
    dialogFormVisible.value = true
}
const dialogFormVisible = ref(false)
const saveEdit = async () => {
    if (editForm.OldPassword === "" || editForm.NewPassword === "" || editForm.AgainPassword === "") {
        ElMessage({
            message: "信息为空",
        })
    } else if (editForm.NewPassword != editForm.AgainPassword) {
        ElMessage({
            message: "新密码不一致",
        })
    } else {
        try {
            console.log(editForm)
            const res = await api.modifyUserPassword(editForm)
            dialogFormVisible.value = false
            location.reload()
        } catch (err) {
            console.error("更新密码失败", err)
        }
    }
}

defineExpose({
    openDialog
})

</script>