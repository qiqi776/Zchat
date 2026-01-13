import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: '/Login',
    },
    {
      path: '/login',
      name: 'Login',
      component: () => import('../views/access/Login.vue'),
    },
    {
      path: '/register',
      name: 'Register',
      component: () => import('../views/access/Register.vue'),
    },
    {
      path: '/chat',
      name: 'Chat',
      component: () => import('../views/chat/Chat.vue'),
    },
    {
      path: '/chat/contactlist',
      name: 'ContactList',
      component: () => import('../views/chat/ContactList.vue'),
    },
    {
      path: '/profile',
      name: 'Profile',
      component: () => import('../views/chat/Profile.vue'),
    },
  ],
})

router.beforeEach((to, from, next) => {
  next()
})

export default router
