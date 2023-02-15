import { createApp } from 'vue'
import './style.css'
import App from './App.vue'

import router from "./router/index.js";

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import 'element-plus/theme-chalk/display.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

/* 图片预览 */
import 'viewerjs/dist/viewer.css'
import VueViewer from 'v-viewer'
import directives from "@/directives"
const app = createApp(App)
// 导入自定义指令
app.use(directives)

app.use(VueViewer)
app.use(ElementPlus)
app.use(router)
//全局注册 element-plus icon
Object.keys(ElementPlusIconsVue).forEach((key) => {
    app.component(key, ElementPlusIconsVue[key]);
});
app.mount('#app')