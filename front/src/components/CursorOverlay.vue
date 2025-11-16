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

// 调试模式（生产环境可以通过 localStorage 启用）
const DEBUG = localStorage.getItem('cursor-debug') === 'true';

// 判断元素是否可交互
const isInteractive = (element: Element | null): boolean => {
  if (!element || element === document.documentElement) return false;

  const tag = element.tagName.toLowerCase();
  const clickableTags = ['a', 'button', 'input', 'textarea', 'select', 'label'];

  // 检测常见可点击标签、CSS cursor、onclick 属性
  if (clickableTags.includes(tag) ||
    window.getComputedStyle(element).cursor === 'pointer' ||
    element.hasAttribute('onclick')) {
    return true;
  }

  // // 检测 Element Plus 组件（通过类名）- 作为备用方案
  // const classList = element.className;
  // if (typeof classList === 'string') {
  //   const clickableClasses = [
  //     'el-button',
  //     'el-card',
  //     'el-link',
  //     'el-menu-item',
  //     'el-dropdown-item',
  //     'el-tab',
  //     'el-radio',
  //     'el-checkbox',
  //     'el-switch',
  //     'el-tag',
  //     'el-badge',
  //     'el-avatar',
  //     'el-image',
  //     'el-pagination',
  //     'el-breadcrumb-item'
  //   ];
    
  //   if (clickableClasses.some(cls => classList.includes(cls))) {
  //     return true;
  //   }
  // }

  // 检测 Vue 事件监听器（@click）- 兼容生产环境
  // 方法1: 检查 __vnode（开发环境）
  const vnode = (element as any).__vnode?.props;
  if (vnode?.onClick || vnode?.onClickCapture) return true;

  // 方法2: 检查 _vei（Vue Event Invoker - 生产环境也可用）
  // _vei 结构: { click: invoker, ... }
  const vei = (element as any)._vei;
  if (vei && (vei.click || vei.onClick)) return true;

  // 方法3: 检查实际的点击事件监听器（仅开发环境可用）
  const listeners = (window as any).getEventListeners?.(element);
  if (listeners?.click?.length > 0) return true;
  
  // // 方法4: 检查原生事件监听器（通过检查是否有事件属性）
  // try {
  //   const hasClickListener = element.onclick !== null || 
  //                           (element as any)._events?.click ||
  //                           Object.keys(vei || {}).some(key => key.toLowerCase().includes('click'));
  //   if (hasClickListener) return true;
  // } catch (e) {
  //   // 忽略错误
  // }

  // 调试：输出元素信息
  if (DEBUG) {
    const el = element as any;
    const debugInfo: any = {
      tag,
      classList: element.className,
      hasVnode: !!el.__vnode,
      hasVei: !!el._vei
    };
    
    if (el._vei) {
      debugInfo.veiKeys = Object.keys(el._vei);
    }
    if (el.__vnode?.props) {
      debugInfo.vnodeProps = Object.keys(el.__vnode.props);
    }
    
    console.log('Element debug:', debugInfo);
  }

  // 方法5: 检查 role 属性
  const role = element.getAttribute('role');
  if (role && ['button', 'link', 'menuitem', 'tab'].includes(role)) {
    return true;
  }

  // 方法6: 检查 data-clickable 属性（可以手动添加到需要交互的元素上）
  if (element.hasAttribute('data-clickable')) {
    return true;
  }

  // 排除 github-markdown-body 内的非链接元素
  if (element.closest('.github-markdown-body')) {
    // 只允许链接标签触发放大
    return tag === 'a';
  }

  // 递归检查父元素
  return isInteractive(element.parentElement);
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