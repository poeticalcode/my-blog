import { createRouter, createWebHistory } from 'vue-router'


import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
import { findLast } from "lodash"
import { check, isLogin } from "../util/auth.js"
// 路由映射表
const routes = [
    {
        path: '/console',
        component: () => import(/*webpackChunkName:"layout"*/ "@/layout/console/Layout.vue"),
        children: [{
            path: '',
            component: () => import(/*webpackChunkName:"console-index"*/ "@/views/console/home/Index.vue"),
            // 任何人都可以阅读文章
            meta: { title: '首页' }
        }, {
            path: 'article/list',
            component: () => import(/*webpackChunkName:"console-article"*/ "@/views/console/aritcle/ArticleList.vue"),
            // 任何人都可以阅读文章
            meta: { title: '文章列表' }
        }, {
            path: ':pathMatch(.*)',
            component: () => import(/*webpackChunkName:"console-home"*/ "@/views/404/Index.vue"),
            // 404 页面
            meta: { title: '404', }
        }]
    },
    {
        path: '/console/article/post',
        component: () => import(/*webpackChunkName:"console-article"*/ "@/views/console/aritcle/PostArticle.vue"),
        // 任何人都可以阅读文章
        meta: { title: '发布文章' }
    },
    {
        path: '/console/article/edit',
        component: () => import(/*webpackChunkName:"console-article"*/ "@/views/console/aritcle/EditArticle.vue"),
        // 任何人都可以阅读文章
        meta: { title: '编辑文章' }
    }, {
        path: '/',
        component: () => import(/*webpackChunkName:"layout"*/ "@/layout/Layout.vue"),
        children: [
            {
                path: '',
                component: () => import(/*webpackChunkName:"home"*/ "@/views/home/Index.vue"),
                // 任何人都可以阅读文章
                meta: { title: '首页' }
            }, {
                path: 'article/:articleId(\\d+)',
                component: () => import(/*webpackChunkName:"article"*/ "@/views/article/Index.vue"),
                // 任何人都可以阅读文章
                meta: { title: '文章详情', }
            }, {
                path: 'about',
                component: () => import(/*webpackChunkName:"about"*/ "@/views/about/Index.vue"),
                // 任何人都可以阅读文章
                meta: { title: '关于我' }
            }, {
                path: 'archive',
                component: () => import(/*webpackChunkName:"archive"*/ "@/views/archive/Index.vue"),
                // 任何人都可以阅读文章
                meta: { title: '归档' }
            }, {
                path: 'tags',
                component: () => import(/*webpackChunkName:"tags"*/ "@/views/tags/Index.vue"),
                // 任何人都可以阅读文章
                meta: { title: '标签' }
            }, {
                path: '/:pathMatch(.*)',
                component: () => import(/*webpackChunkName:"home"*/ "@/views/404/Index.vue"),
                // 404 页面
                meta: { title: '404' }
            }
        ]
    }
]

// 创建一个路由
const router = createRouter({
    history: createWebHistory(),
    routes
})

router.beforeEach(async (to, from, next) => {
    NProgress.start()
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