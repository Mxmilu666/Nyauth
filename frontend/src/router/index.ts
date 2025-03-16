import { createRouter, createWebHistory } from 'vue-router'

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
                    path: 'security',
                    name: 'SecurityPage',
                    component: () => import('@/pages/Console/Security/security.vue')
                }
            ]
        },
        {
            path: '/authorize',
            name: 'Authorize',
            component: () => import('@/pages/Authorize/Authorize.vue')
        }
    ]
})

export default router
