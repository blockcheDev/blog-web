<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue';

// 光标位置和状态
const cursorX = ref(0);
const cursorY = ref(0);
const isVisible = ref(false);
const isHover = ref(false);
const cursorElement = ref<HTMLDivElement | null>(null);

// 尺寸配置
const CURSOR_SIZE = 8; // 光标基础尺寸
const HOVER_SCALE = 6;  // 悬停时的缩放倍率 (8px * 6 = 48px)

// 判断元素是否可交互
const isInteractive = (el: Element | null): boolean => {
  if (!el) return false;
  
  const tagName = el.tagName.toLowerCase();
  const interactiveTags = ['a', 'button', 'input', 'select', 'textarea'];
  
  // 检查当前元素是否是真正的交互元素
  const isCurrentInteractive = 
    interactiveTags.includes(tagName) ||
    window.getComputedStyle(el).cursor === 'pointer' ||
    el.hasAttribute('onclick');
  
  // 检查 Vue 事件监听器（@click 等）
  const htmlEl = el as any;
  const hasVueClick = htmlEl._vei || htmlEl.__vnode?.props;
  if (hasVueClick) {
    const props = htmlEl.__vnode?.props || {};
    if (props.onClick || props.onClickCapture) {
      // 如果有 Vue 点击事件，也算作当前元素可交互
      return true;
    }
  }
  
  // 如果当前元素本身是交互元素，直接返回 true
  if (isCurrentInteractive) return true;
  
  // 检查是否在 github-markdown-body 内
  const isInMarkdownBody = el.classList.contains('github-markdown-body') || 
                           el.closest('.github-markdown-body');
  
  // 如果在 markdown body 内但当前元素不是交互元素，不继续向上查找
  if (isInMarkdownBody) return false;
  
  // 否则检查父元素
  return el.parentElement ? isInteractive(el.parentElement) : false;
};

// 更新光标位置
const updatePosition = () => {
  if (cursorElement.value) {
    const offset = CURSOR_SIZE / 2;
    // 使用 translate 优化性能，位置不需要过渡
    cursorElement.value.style.translate = `${cursorX.value - offset}px ${cursorY.value - offset}px`;
    // scale 单独设置，会有 CSS 过渡效果
    cursorElement.value.style.scale = isHover.value ? `${HOVER_SCALE}` : '1';
  }
  requestAnimationFrame(updatePosition);
};

// 鼠标移动处理
const onMouseMove = (e: MouseEvent) => {
  cursorX.value = e.clientX;
  cursorY.value = e.clientY;
  isVisible.value = true;
  isHover.value = isInteractive(e.target as Element);
};

// 鼠标离开处理
const onMouseLeave = () => {
  isVisible.value = false;
};

// 组件挂载
onMounted(() => {
  // 设置 CSS 变量
  document.documentElement.style.setProperty('--cursor-size', `${CURSOR_SIZE}px`);
  
  document.addEventListener('mousemove', onMouseMove);
  document.addEventListener('mouseleave', onMouseLeave);
  requestAnimationFrame(updatePosition);
});

// 组件卸载
onUnmounted(() => {
  document.removeEventListener('mousemove', onMouseMove);
  document.removeEventListener('mouseleave', onMouseLeave);
});
</script>

<template>
  <div 
    ref="cursorElement"
    class="custom-cursor"
    :class="{ 'hover': isHover, 'visible': isVisible }"
  />
</template>

<style scoped>
.custom-cursor {
  position: fixed;
  top: 0;
  left: 0;
  width: var(--cursor-size, 8px);
  height: var(--cursor-size, 8px);
  border-radius: 50%;
  pointer-events: none;
  z-index: 10000;
  
  /* 默认状态：黑色边框小圆环 */
  border: 2px solid #000;
  background-color: transparent;
  
  /* 初始隐藏 */
  opacity: 0;
  
  /* 优雅的过渡效果 */
  transition: 
    opacity 0.15s cubic-bezier(0.4, 0, 0.2, 1),
    border-width 0.2s cubic-bezier(0.4, 0, 0.2, 1),
    background-color 0.2s cubic-bezier(0.4, 0, 0.2, 1),
    scale 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  
  /* 性能优化 */
  will-change: scale, translate;
}

.custom-cursor.visible {
  opacity: 1;
}

/* 悬停状态：白色圆饼 + 反色效果 */
.custom-cursor.hover {
  border-width: 0;
  background-color: #ffffff;
  mix-blend-mode: difference;
}
</style>

<style>
/* 隐藏默认光标 */
* {
  cursor: none !important;
}

/* 文本输入区域保留文本光标 */
input,
textarea,
[contenteditable="true"] {
  cursor: text !important;
}
</style>