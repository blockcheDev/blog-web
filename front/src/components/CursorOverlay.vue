<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue';

// 光标位置和状态
const cursorX = ref(0);
const cursorY = ref(0);
const isVisible = ref(false);
const isHover = ref(false);
const isInputArea = ref(false);
const isClicking = ref(false);
const cursorElement = ref<HTMLDivElement | null>(null);

// 尺寸配置
const CURSOR_SIZE = 8; // 光标基础尺寸
const HOVER_SCALE = 6;  // 悬停时的缩放倍率 (8px * 6 = 48px)
const INPUT_SCALE = 0.1; // 输入区域的缩放倍率（几乎不可见）
const CLICK_SCALE = 0.1; // 点击时的缩放倍率（缩小到消失）

// 调试模式（生产环境可以通过 localStorage 启用）
const DEBUG = localStorage.getItem('cursor-debug') === 'true';

// 判断元素是否为文本输入区域
const isTextInput = (element: Element | null): boolean => {
  if (!element) return false;
  
  const tag = element.tagName.toLowerCase();
  const inputTags = ['input', 'textarea'];
  
  // 检查是否为输入标签
  if (inputTags.includes(tag)) return true;
  
  // 检查是否为可编辑元素
  if (element.getAttribute('contenteditable') === 'true') return true;
  
  return false;
};

// 判断元素是否可交互
const isInteractive = (element: Element | null): boolean => {
  if (!element || element === document.documentElement) return false;

  const classList = element.classList;
  
  // 排除遮罩层本身、抽屉容器本身、对话框容器本身（但不排除其内部元素）
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
  const clickableTags = ['a', 'button', 'input', 'textarea', 'select', 'label'];

  // 检测常见可点击标签、CSS cursor、onclick 属性
  if (clickableTags.includes(tag) ||
    window.getComputedStyle(element).cursor === 'pointer' ||
    element.hasAttribute('onclick')) {
    return true;
  }

  // 检测 Vue 事件监听器（@click）- 兼容生产环境
  // 方法1: 检查 __vnode（开发环境）
  const vnode = (element as any).__vnode?.props;
  if (vnode?.onClick || vnode?.onClickCapture) return true;

  // 排除 github-markdown-body 内的非链接元素
  if (element.closest('.github-markdown-body')) {
    // 只允许链接标签触发放大
    return tag === 'a';
  }

  // 递归检查父元素
  return isInteractive(element.parentElement);
};

// 当前缩放值（用于平滑过渡）
let currentScale = 1;
let targetScale = 1;
let animationFrameId: number | null = null;
let isAnimating = false;
let lastFrameTime = 0;

// 动画配置（针对高刷新率优化）
const ANIMATION_SPEED = 0.18; // 降低插值系数，使动画更平滑
const STOP_THRESHOLD = 0.01; // 停止阈值
const TARGET_FRAME_TIME = 1000 / 60; // 基准 60fps (16.67ms)
const MAX_DELTA_TIME = 1000 / 30; // 最大帧间隔 (33.33ms)，防止异常跳跃

// 更新光标位置（支持 60Hz/120Hz/144Hz 等高刷新率）
const updatePosition = (timestamp: number = 0) => {
  if (!cursorElement.value) return;
  
  // 计算帧间隔时间（用于时间补偿）
  let deltaTime = lastFrameTime ? timestamp - lastFrameTime : TARGET_FRAME_TIME;
  lastFrameTime = timestamp;
  
  // 限制最大帧间隔，防止页面切换回来时的跳跃
  deltaTime = Math.min(deltaTime, MAX_DELTA_TIME);
  
  // 时间补偿系数（确保在不同刷新率下动画速度一致）
  // 使用平方根使补偿更平滑
  const timeCompensation = Math.sqrt(deltaTime / TARGET_FRAME_TIME);
  
  const offset = CURSOR_SIZE / 2;
  
  // 平滑插值（lerp）实现流畅过渡
  const scaleDiff = Math.abs(targetScale - currentScale);
  
  if (scaleDiff > STOP_THRESHOLD) {
    // 还在动画中，继续插值（应用时间补偿）
    // 使用缓动函数使动画更自然
    const lerpFactor = Math.min(ANIMATION_SPEED * timeCompensation, 0.5);
    currentScale += (targetScale - currentScale) * lerpFactor;
    
    // 使用 transform 一次性更新位置和缩放
    cursorElement.value.style.transform = 
      `translate(${cursorX.value - offset}px, ${cursorY.value - offset}px) scale(${currentScale})`;
    
    // 继续下一帧
    animationFrameId = requestAnimationFrame(updatePosition);
  } else {
    // 动画完成，停止 RAF
    currentScale = targetScale;
    cursorElement.value.style.transform = 
      `translate(${cursorX.value - offset}px, ${cursorY.value - offset}px) scale(${currentScale})`;
    isAnimating = false;
    animationFrameId = null;
    lastFrameTime = 0; // 重置时间戳
  }
};

// 启动动画（仅在状态变化时调用）
const startAnimation = () => {
  if (!isAnimating) {
    isAnimating = true;
    updatePosition();
  }
};

// 更新光标位置（仅更新坐标，不触发动画）
const updateCursorPosition = () => {
  if (cursorElement.value) {
    const offset = CURSOR_SIZE / 2;
    cursorElement.value.style.transform = 
      `translate(${cursorX.value - offset}px, ${cursorY.value - offset}px) scale(${currentScale})`;
  }
};

// 鼠标移动处理
const onMouseMove = (e: MouseEvent) => {
  cursorX.value = e.clientX;
  cursorY.value = e.clientY;
  isVisible.value = true;
  
  const target = e.target as Element;
  const wasInputArea = isInputArea.value;
  const wasHover = isHover.value;
  
  isInputArea.value = isTextInput(target);
  isHover.value = !isInputArea.value && isInteractive(target);
  
  // 计算新的目标缩放值（点击状态优先级最高）
  const newTargetScale = isClicking.value ? CLICK_SCALE : 
                         (isInputArea.value ? INPUT_SCALE : 
                         (isHover.value ? HOVER_SCALE : 1));
  
  // 只在状态变化时才启动动画
  if (newTargetScale !== targetScale) {
    targetScale = newTargetScale;
    startAnimation();
  } else {
    // 状态没变，只更新位置
    updateCursorPosition();
  }
};

// 鼠标离开处理
const onMouseLeave = () => {
  isVisible.value = false;
};

// 鼠标按下处理
const onMouseDown = () => {
  isClicking.value = true;
  targetScale = CLICK_SCALE;
  startAnimation();
};

// 鼠标松开处理
const onMouseUp = () => {
  isClicking.value = false;
  // 恢复到当前应有的状态
  targetScale = isInputArea.value ? INPUT_SCALE : (isHover.value ? HOVER_SCALE : 1);
  startAnimation();
};

// 组件挂载
onMounted(() => {
  // 设置 CSS 变量
  document.documentElement.style.setProperty('--cursor-size', `${CURSOR_SIZE}px`);
  
  document.addEventListener('mousemove', onMouseMove);
  document.addEventListener('mouseleave', onMouseLeave);
  document.addEventListener('mousedown', onMouseDown);
  document.addEventListener('mouseup', onMouseUp);
  // 不再启动无限 RAF 循环，只在需要时才启动动画
});

// 组件卸载
onUnmounted(() => {
  document.removeEventListener('mousemove', onMouseMove);
  document.removeEventListener('mouseleave', onMouseLeave);
  document.removeEventListener('mousedown', onMouseDown);
  document.removeEventListener('mouseup', onMouseUp);
  
  // 清理可能存在的动画帧
  if (animationFrameId !== null) {
    cancelAnimationFrame(animationFrameId);
  }
});
</script>

<template>
  <div 
    ref="cursorElement"
    class="custom-cursor"
    :class="{ 
      'hover': isHover, 
      'visible': isVisible, 
      'input-area': isInputArea,
      'clicking': isClicking 
    }"
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
  
  /* 优雅的过渡效果（移除 scale，由 JS RAF 控制） */
  transition: 
    opacity 0.15s cubic-bezier(0.4, 0, 0.2, 1),
    border-width 0.03s cubic-bezier(1, 0, 1, 1),
    background-color 0.03s cubic-bezier(1, 0, 1, 1);  
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