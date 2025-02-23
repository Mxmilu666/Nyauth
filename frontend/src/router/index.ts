import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      name: 'Home',
      path: '/',
      component: () => import('@/pages/Home/Home.vue'),
    },
    {
      name: 'Login',
      path: '/login',
      component: () => import('@/pages/Login/Login.vue'),
    },
  ],
})

export default router
