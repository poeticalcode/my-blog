import { createApp } from 'vue'
import App from './App.vue'
import '@/assets/css/main.scss'
import router from "@/router";

// ElementPlus
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import 'element-plus/theme-chalk/display.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

// directives
import directives from "@/directives"

const app = createApp(App)
// 使用指令
app.use(directives)

app.use(ElementPlus)
app.use(router)

//全局注册 element-plus icon
Object.keys(ElementPlusIconsVue).forEach((key) => {
    app.component(key, ElementPlusIconsVue[key]);
});
app.mount('#app')