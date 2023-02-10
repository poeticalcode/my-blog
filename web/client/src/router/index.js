import {createRouter, createWebHashHistory} from 'vue-router'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
import {findLast} from "lodash"
import {check, isLogin} from "../util/auth.js"
// 路由映射表
const routes = [
    {
        path: '/',
        component: () => import(/*webpackChunkName:"layout"*/ "@/layout/Layout.vue"),
        children: [
            {
                path: '/',
                component: () => import(/*webpackChunkName:"article"*/ "@/views/home/Index.vue"),
                // 任何人都可以阅读文章
                meta: {
                    requiresAuth: false,
                    title: '首页',
                    authority: ["admin"]
                }
            }
        ]
    },
    
]


// 创建一个路由
const router = createRouter({
    history: createWebHashHistory(),
    routes
})


router.beforeEach(async (to, from, next) => {
    if (to.path !== from.path) { // 同页面变化就不显示加载进度条
        NProgress.start()
    }
    const record = findLast(to.matched, record => record.meta.authority)
    if (record && !check(record.meta.authority)) {
        if (!isLogin() && to.path != "/login") {
            next({path: "/login"})
        } else {
            next({path: "/403"})
        }
        NProgress.done()
    }
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