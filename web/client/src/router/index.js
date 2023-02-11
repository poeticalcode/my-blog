import { createRouter, createWebHistory } from 'vue-router'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
import { findLast } from "lodash"
import { check, isLogin } from "../util/auth.js"
// 路由映射表
const routes = [
    {
        path: '/',
        component: () => import(/*webpackChunkName:"layout"*/ "@/layout/Layout.vue"),
        children: [
            {
                path: '/',
                component: () => import(/*webpackChunkName:"home"*/ "@/views/home/Index.vue"),
                // 任何人都可以阅读文章
                meta: {
                    title: '首页',
                }
            },
            {
                path: '/article',
                component: () => import(/*webpackChunkName:"article"*/ "@/views/article/Index.vue"),
                // 任何人都可以阅读文章
                meta: {
                    title: '文章详情',
                }
            }
        ]
    },
    {
        path: '/:pathMatch(.*)',
        component: () => import(/*webpackChunkName:"home"*/ "@/views/404/Index.vue"),
        // 404 页面
        meta: {
            title: '404',
        }
    }
]

// 创建一个路由
const router = createRouter({
    history: createWebHistory(),
    routes
})

router.beforeEach(async (to, from, next) => {
    if (to.meta.title) {
        document.title = to.meta.title
    }
    next()
})

router.afterEach(() => {
    NProgress.done()
})

// 导出一个路由
export default router;