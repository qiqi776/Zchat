import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/Login.vue' // 引入页面

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'login',
      component: Login,
    },
    {
      path: '/chat',
      name: 'chat',
      // 路由懒加载 (推荐写法)
      component: () => import('../views/Chat.vue'),
    },
  ],
})

export default router
