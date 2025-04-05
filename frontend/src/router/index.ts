import { createRouter, createWebHistory } from 'vue-router'
import { Cookie } from '@/utils/cookie'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            name: 'Home',
            path: '/',
            component: () => import('@/pages/Home/Home.vue')
        },
        {
            name: 'Login',
            path: '/login',
            component: () => import('@/pages/Userauth/Auth.vue')
        },
        {
            name: 'ResetPassword',
            path: '/reset-password',
            component: () => import('@/pages/Userauth/Reset.vue')
        },
        // console
        {
            path: '/console',
            name: 'Console',
            component: () => import('@/pages/Console/Console.vue'),
            children: [
                {
                    path: '',
                    name: 'HomePage',
                    component: () => import('@/pages/Console/Home.vue')
                },
                {
                    path: 'info',
                    name: 'InfoPage',
                    component: () => import('@/pages/Console/Profile/Info.vue')
                },
                {
                    path: 'info/username',
                    name: 'UsernameInfoPage',
                    component: () => import('@/pages/Console/Profile/Detail/UserName.vue')
                },
                {
                    path: 'security',
                    name: 'SecurityPage',
                    component: () => import('@/pages/Console/Security/security.vue')
                }
            ]
        },
        // oauth
        {
            path: '/oauth',
            name: 'Oauth',
            children: [
                {
                    path: 'authorize',
                    name: 'Authorize',
                    component: () => import('@/pages/Authorize/Authorize.vue')
                }
            ]
        }
    ]
})

// 全局前置守卫
router.beforeEach((to, from, next) => {
    // 检查 path 是否为 /console 或者 /oauth
    if (to.path.startsWith('/console') || to.path.startsWith('/oauth')) {
        // 验证是否存在 token
        const token = Cookie.get('token')
        if (!token) {
            // 无token，重定向到登录页
            next({
                name: 'Login',
                query: { redirect: to.fullPath }
            })
            return
        }
    }
    // 其他情况正常放行
    next()
})

export default router
