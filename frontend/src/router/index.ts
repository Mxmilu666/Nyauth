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
            component: () => import('@/pages/Userauth/Login.vue')
        },
        {
            name: 'Register',
            path: '/regsiter',
            component: () => import('@/pages/Userauth/Register.vue')
        },
        {
            name: 'ResetPassword',
            path: '/reset-password',
            component: () => import('@/pages/Userauth/Reset.vue')
        }
    ]
})

export default router
