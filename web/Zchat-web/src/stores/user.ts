import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { User } from '@/models/chat'

export const useUserStore = defineStore('user', () => {
  // 1. 用户信息
  const userInfo = ref<User>({
    uuid: '',
    nickname: '',
    avatar: '',
    telephone: '',
  })

  // 尝试从本地缓存读取，这样刷新页面不会退出登录
  const token = ref(localStorage.getItem('user_token') || '')

  // 登录成功调用的动作
  const setUser = (user: User, newToken: string) => {
    userInfo.value = user
    token.value = newToken
    // 存入本地缓存
    localStorage.setItem('user_token', newToken)
    localStorage.setItem('user_info', JSON.stringify(user))
  }

  // 退出登录
  const logout = () => {
    userInfo.value = { uuid: '', nickname: '', avatar: '' }
    token.value = ''
    localStorage.removeItem('user_token')
    localStorage.removeItem('user_info')
  }

  // 初始化恢复,防止刷新丢失数据
  const init = () => {
    const storedUser = localStorage.getItem('user_info')
    if (storedUser) {
      try {
        userInfo.value = JSON.parse(storedUser)
      } catch (e) {
        console.error('解析用户信息失败', e)
      }
    }
  }

  init()
  return { userInfo, token, setUser, logout }
})
