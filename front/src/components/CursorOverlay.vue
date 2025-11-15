<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed } from 'vue';

// Props: 光标大小、颜色、混合模式
const props = withDefaults(defineProps<{
  size?: number;
  color?: string;
  blendMode?: 'difference' | 'exclusion' | 'screen';
}>(), {
  size: 40,
  color: '#eeeeee',
  blendMode: 'difference'
});

// 状态：光标位置、可见性、鼠标按下状态、文本选择状态
const x = ref(0);
const y = ref(0);
const visible = ref(false);
const isMouseDown = ref(false);
const isSelecting = ref(false);
const dotRef = ref<HTMLDivElement | null>(null);
let rafId: number | null = null;

// 计算属性：光标尺寸、透明度、缩放
const cursorSize = computed(() => Math.max(10, props.size));
const cursorOpacity = computed(() => visible.value && !isMouseDown.value && !isSelecting.value ? 1 : 0);
const cursorScale = computed(() => isMouseDown.value ? 0 : 1);
// 鼠标移动：更新位置并显示光标
const handleMove = (e: MouseEvent) => {
  x.value = e.clientX;
  y.value = e.clientY;
  visible.value = true;
};

// 鼠标按下：标记按下状态
const handleDown = (e: MouseEvent) => {
  if (e.button === 0) {
    isMouseDown.value = true;
  }
};

// 鼠标释放：清除按下状态
const handleUp = () => {
  isMouseDown.value = false;
};

// 检测文本选择状态
const checkSelection = () => {
  const selection = window.getSelection();
  const hasSelection = !!(selection && !selection.isCollapsed);
  if (hasSelection !== isSelecting.value) {
    isSelecting.value = hasSelection;
    document.documentElement.classList.toggle('is-selecting-text', hasSelection);
  }
};

// RAF 循环：同步光标位置到 DOM
const loop = () => {
  if (dotRef.value) {
    const half = cursorSize.value / 2;
    dotRef.value.style.translate = `${x.value - half}px ${y.value - half}px`;
    dotRef.value.style.scale = String(cursorScale.value);
  }
  rafId = requestAnimationFrame(loop);
};

// 挂载：绑定事件监听器
onMounted(() => {
  window.addEventListener('mousemove', handleMove, { passive: true });
  window.addEventListener('mouseleave', () => visible.value = false, { passive: true });
  window.addEventListener('mousedown', handleDown, true);
  window.addEventListener('mouseup', handleUp, true);
  document.addEventListener('selectionchange', checkSelection);
  rafId = requestAnimationFrame(loop);
});

// 卸载：清理事件监听器
onUnmounted(() => {
  window.removeEventListener('mousemove', handleMove);
  window.removeEventListener('mouseleave', () => visible.value = false);
  window.removeEventListener('mousedown', handleDown, true);
  window.removeEventListener('mouseup', handleUp, true);
  document.removeEventListener('selectionchange', checkSelection);
  document.documentElement.classList.remove('is-selecting-text');
  if (rafId) cancelAnimationFrame(rafId);
});
</script>

<template>
  <!-- 自定义光标：圆形，跟随鼠标，使用 mix-blend-mode 实现反色 -->
  <div ref="dotRef" class="cursor-dot" :style="{
    opacity: cursorOpacity,
    width: cursorSize + 'px',
    height: cursorSize + 'px',
    background: color,
    mixBlendMode: blendMode
  }" />
</template>

<style scoped>
/* 光标样式：固定定位、圆形、GPU 加速 */
.cursor-dot {
  position: fixed;
  left: 0;
  top: 0;
  z-index: 9999;
  border-radius: 50%;
  box-shadow: 0 0 0 1px rgba(0, 0, 0, 0.15);
  pointer-events: none;
  user-select: none;
  transition: opacity 0.2s ease-out, scale 0.2s ease-out;
  will-change: translate, scale;
  backface-visibility: hidden;
  contain: layout style paint;
}
</style>

<style>
/* 全局隐藏系统光标 */
html,
body,
#app,
*,
*:hover,
*:active,
*:focus {
  cursor: none !important;
}

/* 输入框显示文本光标 */
input,
textarea,
[contenteditable="true"],
.allow-text-cursor {
  cursor: text !important;
}

/* 选择文本时显示文本光标 */
.is-selecting-text,
.is-selecting-text * {
  cursor: text !important;
}
</style>