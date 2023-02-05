import {createApp} from 'vue'
import './style.css'
import App from './App.vue'

import router from "./router/index.js";

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'


import Auth from "./directives/auth.js";


const app = createApp(App);
app.use(ElementPlus)
app.use(router)
app.use(Auth)

//全局注册 element-plus icon
Object.keys(ElementPlusIconsVue).forEach((key) => {
    app.component(key, ElementPlusIconsVue[key]);
});
app.mount('#app')
