import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: '/login',
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/access/Login.vue'),
    },
    {
      path: '/register',
      name: 'register',
      component: () => import('../views/access/Register.vue'),
    },
    {
      path: '/chat',
      name: 'chat',
      component: () => import('../views/chat/Chat.vue'),
    },
    {
      path: '/contact',
      name: 'contact',
      component: () => import('../views/chat/ContactList.vue'),
    },
    {
      path: '/profile',
      name: 'profile',
      component: () => import('../views/chat/Profile.vue'),
    },
  ],
})

// 防止未登录用户直接访问聊天页
router.beforeEach((to, from, next) => {
  const publicPages = ['/login', '/register']
  const authRequired = !publicPages.includes(to.path)
  const loggedIn = localStorage.getItem('user_token')

  if (authRequired && !loggedIn) {
    next('/login')
  } else {
    next()
  }
})

export default router
