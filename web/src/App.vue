<script setup lang="ts">
import index from './layouts/index.vue';
</script>

<template>
  <index />
</template>

<style>
#app {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}

.main {
  display: flex;
  flex-flow: column;
}

.el-card {
  border-radius: 15px !important;
  margin-bottom: 30px;
  margin-left: auto;
  margin-right: auto;
  width: 60%
}

.el-dialog {
  border-radius: 15px !important;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}

/* 离开和进入过程中的样式 */
.v-enter-active,
.v-leave-active {
  /* 添加过渡动画 */
  transition: opacity 0.2s ease;
}

/* 进入之后和离开之前的样式 */
.v-enter-to,
.v-leave-from {
  opacity: 1;
}
</style>
