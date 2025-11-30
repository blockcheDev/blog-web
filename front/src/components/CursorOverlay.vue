<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue';

// DOM 引用
const cursorElement = ref<HTMLDivElement | null>(null);

// 状态变量 (非响应式，提升性能)
let cursorX = 0;
let cursorY = 0;
let currentX = 0;
let currentY = 0;
let isVisible = false;
let isClicking = false;
let isHover = false;
let isInputArea = false;

// 缓存上一次的目标元素，避免重复计算
let lastTarget: Element | null = null;

// 尺寸配置
const CURSOR_SIZE = 8;
const HOVER_SCALE = 8;
const INPUT_SCALE = 0.1;
const CLICK_SCALE = 0.1;

// 动画状态
let currentScale = 1;
let targetScale = 1;
let animationFrameId: number | null = null;
let isAnimating = false;
let lastFrameTime = 0;

// 动画配置
const SCALE_SPEED = 0.15;
const POSITION_SPEED = 0.40;
const STOP_THRESHOLD = 0.001;
const TARGET_FRAME_TIME = 1000 / 60;
const MAX_DELTA_TIME = 1000 / 30;

// 判断元素是否为文本输入区域
const isTextInput = (element: Element | null): boolean => {
  if (!element) return false;
  const tag = element.tagName.toLowerCase();
  return ['input', 'textarea'].includes(tag) || element.getAttribute('contenteditable') === 'true';
};

// 判断元素是否可交互
const isInteractive = (element: Element | null): boolean => {
  if (!element || element === document.documentElement) return false;

  // 快速检查：如果是特定容器，直接返回 false
  const classList = element.classList;
  if (classList.contains('el-overlay') ||
      classList.contains('el-drawer') ||
      classList.contains('el-drawer__wrapper') ||
      classList.contains('el-dialog__wrapper') ||
      classList.contains('el-dialog') ||
      classList.contains('el-overlay-dialog') ||
      classList.contains('el-dialog__header') ||
      classList.contains('el-dialog__body') ||
      classList.contains('el-dialog__footer')) {
    return false;
  }

  const tag = element.tagName.toLowerCase();
  // 常见交互标签
  if (['a', 'button', 'select', 'label', 'input', 'textarea'].includes(tag) ||
      element.hasAttribute('onclick')) {
    return true;
  }

  // 检查 Vue 事件监听器 (开发环境)
  const vnode = (element as any).__vnode?.props;
  if (vnode?.onClick || vnode?.onClickCapture) return true;

  // 昂贵的检查放在最后
  if (window.getComputedStyle(element).cursor === 'pointer') {
    return true;
  }

  // 特殊处理 github-markdown-body
  if (element.closest('.github-markdown-body')) {
    return tag === 'a';
  }

  return isInteractive(element.parentElement);
};

// 更新 DOM 类名
const updateCursorClass = () => {
  if (!cursorElement.value) return;
  const cl = cursorElement.value.classList;
  
  isVisible ? cl.add('visible') : cl.remove('visible');
  isHover ? cl.add('hover') : cl.remove('hover');
  isInputArea ? cl.add('input-area') : cl.remove('input-area');
  isClicking ? cl.add('clicking') : cl.remove('clicking');
};

// 动画循环
const updatePosition = (timestamp: number = 0) => {
  if (!cursorElement.value) return;

  let deltaTime = lastFrameTime ? timestamp - lastFrameTime : TARGET_FRAME_TIME;
  lastFrameTime = timestamp;
  deltaTime = Math.min(deltaTime, MAX_DELTA_TIME);
  
  const timeCompensation = Math.sqrt(deltaTime / TARGET_FRAME_TIME);
  
  // 位置插值
  const moveX = (cursorX - currentX) * POSITION_SPEED * timeCompensation;
  const moveY = (cursorY - currentY) * POSITION_SPEED * timeCompensation;
  
  currentX += moveX;
  currentY += moveY;

  // 缩放插值
  const scaleDiff = targetScale - currentScale;
  currentScale += scaleDiff * SCALE_SPEED * timeCompensation;

  const offset = CURSOR_SIZE / 2;
  
  // 应用变换
  cursorElement.value.style.transform = 
    `translate3d(${currentX - offset}px, ${currentY - offset}px, 0) scale(${currentScale})`;

  // 检查是否停止动画
  const isPositionClose = Math.abs(cursorX - currentX) < 0.1 && Math.abs(cursorY - currentY) < 0.1;
  const isScaleClose = Math.abs(targetScale - currentScale) < STOP_THRESHOLD;

  if (!isPositionClose || !isScaleClose) {
    animationFrameId = requestAnimationFrame(updatePosition);
  } else {
    // 强制对齐
    currentX = cursorX;
    currentY = cursorY;
    currentScale = targetScale;
    cursorElement.value.style.transform = 
      `translate3d(${currentX - offset}px, ${currentY - offset}px, 0) scale(${currentScale})`;
      
    isAnimating = false;
    animationFrameId = null;
    lastFrameTime = 0;
  }
};

const startAnimation = () => {
  if (!isAnimating) {
    isAnimating = true;
    lastFrameTime = 0;
    animationFrameId = requestAnimationFrame(updatePosition);
  }
};

const onMouseMove = (e: MouseEvent) => {
  cursorX = e.clientX;
  cursorY = e.clientY;
  
  if (!isVisible) {
    isVisible = true;
    currentX = cursorX;
    currentY = cursorY;
    updateCursorClass();
  }

  startAnimation();

  const target = e.target as Element;
  
  // 核心优化：如果目标元素没有变化，跳过昂贵的检查
  if (target !== lastTarget) {
    lastTarget = target;
    
    const wasInputArea = isInputArea;
    const wasHover = isHover;
    
    isInputArea = isTextInput(target);
    // 只有非输入区域才检查是否可交互
    isHover = !isInputArea && isInteractive(target);
    
    if (wasInputArea !== isInputArea || wasHover !== isHover) {
      updateCursorClass();
    }
  }

  // 计算目标缩放
  const newTargetScale = isClicking ? CLICK_SCALE : 
                         (isInputArea ? INPUT_SCALE : 
                         (isHover ? HOVER_SCALE : 1));

  if (newTargetScale !== targetScale) {
    targetScale = newTargetScale;
  }
};

const onMouseLeave = () => {
  isVisible = false;
  updateCursorClass();
};

const onMouseDown = () => {
  isClicking = true;
  updateCursorClass();
  targetScale = CLICK_SCALE;
  startAnimation();
};

const onMouseUp = () => {
  isClicking = false;
  updateCursorClass();
  targetScale = isInputArea ? INPUT_SCALE : (isHover ? HOVER_SCALE : 1);
  startAnimation();
};

onMounted(() => {
  document.documentElement.style.setProperty('--cursor-size', `${CURSOR_SIZE}px`);
  document.addEventListener('mousemove', onMouseMove, { passive: true });
  document.addEventListener('mouseleave', onMouseLeave, { passive: true });
  document.addEventListener('mousedown', onMouseDown, { passive: true });
  document.addEventListener('mouseup', onMouseUp, { passive: true });
});

onUnmounted(() => {
  document.removeEventListener('mousemove', onMouseMove);
  document.removeEventListener('mouseleave', onMouseLeave);
  document.removeEventListener('mousedown', onMouseDown);
  document.removeEventListener('mouseup', onMouseUp);
  if (animationFrameId !== null) {
    cancelAnimationFrame(animationFrameId);
  }
});
</script>

<template>
  <div ref="cursorElement" class="custom-cursor" />
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
  
  /* 优雅的过渡效果（移除 scale，由 JS RAF 控制） */
  transition: 
    opacity 0.3s cubic-bezier(0.25, 0.8, 0.25, 1),
    border-width 0.1s cubic-bezier(0.25, 0.8, 0.25, 1),
    background-color 0.1s cubic-bezier(0.25, 0.8, 0.25, 1);  
  /* 性能优化：启用 GPU 加速 */
  will-change: transform;
  transform: translateZ(0);
  backface-visibility: hidden;
  perspective: 1000px;
}

.custom-cursor.visible {
  opacity: 1;
}

/* 悬停状态：白色圆饼 + 反色效果 */
.custom-cursor.hover {
  border-width: 0;
  border-color: transparent;
  background-color: #ffffff;
  mix-blend-mode: difference;
}

/* 输入区域状态和点击状态：统一缩放到几乎不可见 */
.custom-cursor.input-area,
.custom-cursor.clicking {
  opacity: 0.3;
  border-color: transparent;
  background-color: #000000;
}
</style>

<style>
/* 隐藏默认光标 */
* {
  cursor: none !important;
}

/* 文本输入区域恢复原生光标 */
input,
textarea,
[contenteditable="true"] {
  cursor: text !important;
}
</style>