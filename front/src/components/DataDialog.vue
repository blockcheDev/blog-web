<template>
    <el-dialog v-model="visible" :title="title" width="70%" destroy-on-close>
        <!-- 内容插槽 -->
        <slot v-if="$slots.default" />
        <template v-else>
            <el-descriptions :column="2" border>
                <el-descriptions-item v-for="(value, key) in data" :key="key" :label="key">
                    {{ value }}
                </el-descriptions-item>
            </el-descriptions>
        </template>
    </el-dialog>
</template>

<script lang="ts" setup>
import { ref, watch } from 'vue'

const props = defineProps({
    modelValue: Boolean,  // 控制显示
    title: { type: String, default: '数据详情' },
    data: { type: Object, default: () => ({}) }
})

const emit = defineEmits(['update:modelValue'])

// 内部visible状态
const visible = ref(false)

// 同步外部v-model变化
watch(() => props.modelValue, (val) => {
    visible.value = val
})

// 同步内部状态到外部
watch(visible, (val) => {
    emit('update:modelValue', val)
})
</script>

<style scoped>
</style>