<script setup lang="ts">
import api from '@/api/api'
import dayjs from 'dayjs'
import { reactive, ref } from 'vue'
import DataDialog from '@/components/DataDialog.vue'

interface Visitor {
    IP: string,
    Ts: number,
    Region: string,

    Date: string,
}
const RecentVisitors = reactive<Visitor[]>([])

const showDialog = ref(false)

const open = async () => {
    try {
        const res = await api.getRecentVisitors()
        Object.assign(RecentVisitors, res.data)
    } catch (err) {
        console.error(err)
        return
    }
    for (var visitor of RecentVisitors) {
        visitor.Date = dayjs.unix(visitor.Ts).format("YYYY-MM-DD HH:mm:ss")
    }
    showDialog.value = true
}
const close = () => {
    showDialog.value = false
}
defineExpose({
    open
})
</script>

<template>
    <DataDialog v-model="showDialog" title="最近访客">
        <el-table :data="RecentVisitors" style="width: 100%">
            <el-table-column prop="IP" label="IP" width="180" />
            <el-table-column prop="Date" label="Date" width="180" />
            <el-table-column prop="Region" label="Region" />
        </el-table>
    </DataDialog>
</template>