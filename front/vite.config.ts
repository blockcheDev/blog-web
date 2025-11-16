import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// 自动按需引入依赖
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'

// gzip压缩 打包后的资源
import viteCompression from 'vite-plugin-compression'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue({
      template: {
        compilerOptions: {
          // 在生产环境中保留 Vue 内部属性（如 __vnode）
          // 这会略微增加包大小，但能让运行时检测更可靠
          isCustomElement: () => false
        }
      }
    }),
    AutoImport({
      resolvers: [ElementPlusResolver()],
    }),
    Components({
      resolvers: [ElementPlusResolver()],
    }),
    viteCompression({
      threshold: 300000 // 对大于 300000B 的文件进行压缩
    })
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  // 在生产环境中启用开发模式的某些特性
  define: {
    __VUE_PROD_DEVTOOLS__: true, // 启用生产环境的 devtools 支持
  }
})
