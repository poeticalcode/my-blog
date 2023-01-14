import {createRouter, createWebHashHistory} from 'vue-router'

import Layout from '@/layout/Layout.vue';
import Index from '@/views/Index.vue'

import ArticleList from "@/views/aritcle/ArticleList.vue";
import PostArticle from "@/views/aritcle/PostArticle.vue";


// 路由映射表
const routes = [
    {
        path: '/',
        component: Layout,
        children: [
            {
                path: '',
                component: Index,
                // 任何人都可以阅读文章
                meta: {
                    requiresAuth: false,
                    title: '首页'
                }
            },
            {
                path: '/article/list',
                component: ArticleList,
                // 任何人都可以阅读文章
                meta: {
                    requiresAuth: false,
                    title: '文章列表'
                }
            },
            {
                path: '/article/post',
                component: PostArticle,
                // 任何人都可以阅读文章
                meta: {
                    requiresAuth: false,
                    title: '图片列表'
                }
            }
        ]
    },
    {
        path: '/post',
        component: PostArticle,
        // 任何人都可以阅读文章
        meta: {
            requiresAuth: false,
            title: '发表文章'
        }
    },
    {path: '/index', component: Index},
]


// 创建一个路由
const router = createRouter({
    history: createWebHashHistory(),
    routes
})

router.beforeEach(async (to, from, next) => {
    if (to.meta.title) {
        document.title = to.meta.title + " | 后台管理"
    }
    next()
})

// 导出一个路由
export default router;