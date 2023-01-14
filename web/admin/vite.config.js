import {defineConfig, loadEnv} from 'vite'
import path from "path"

import vue from '@vitejs/plugin-vue'

import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import {ElementPlusResolver} from "unplugin-vue-components/resolvers";

// 这里属于 node 端运行

// https://vitejs.dev/config/
export default defineConfig((config) => {
    // 编程式配置
    const {command: command, mode: mode, ssrBuild: ssrBuild} = config;

    // 加载环境 环境中不以 VITE_ 开头的不会注入 import.meta.env
    const env = loadEnv(mode, __dirname)

    return {
        optimizeDeps: {},
        plugins: [
            vue(),
            AutoImport({
                resolvers: [
                    ElementPlusResolver()
                ],
            }),
            Components({
                resolvers: [
                    ElementPlusResolver()
                ],
            })
        ],
        resolve: {
            alias: {
                "@": path.resolve(__dirname, "./src"),
            }
        }
    }
})
