import {createRouter, createWebHashHistory} from 'vue-router'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'


// 路由映射表
const routes = [
    {
        path: '/',
        component: () => import(/*webpackChunkName:"layout"*/ "@/layout/Layout.vue"),
        children: [
            {
                path: '',
                component: () => import(/*webpackChunkName:"index"*/ "@/views/Index.vue"),
                // 任何人都可以阅读文章
                meta: {
                    requiresAuth: false,
                    title: '首页'
                }
            },
            {
                path: '/article/list',
                component: () => import(/*webpackChunkName:"article"*/ "@/views/aritcle/ArticleList.vue"),
                // 任何人都可以阅读文章
                meta: {
                    requiresAuth: false,
                    title: '文章列表'
                }
            },
            {
                path: '/article/post',
                component: () => import(/*webpackChunkName:"article"*/ "@/views/aritcle/PostArticle.vue"),
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
        component: () => import(/*webpackChunkName:"article"*/ "@/views/aritcle/PostArticle.vue"),
        // 任何人都可以阅读文章
        meta: {
            requiresAuth: false,
            title: '发表文章'
        }
    },
    {path: '/index', component: () => import(/*webpackChunkName:"index"*/ "@/views/Index.vue"),},
]


// 创建一个路由
const router = createRouter({
    history: createWebHashHistory(),
    routes
})

router.beforeEach(async (to, from, next) => {
    NProgress.start()
    if (to.meta.title) {
        document.title = to.meta.title + " | 后台管理"
    }
    next()
})

router.afterEach(() => {
    NProgress.done()
})

// 导出一个路由
export default router;