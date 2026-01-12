import request from '@/utils/request'
import type { LoginResponse } from '@/models/user'

// 登录 API
export const login = (data: { telephone: string; password: string }) => {
  // 注意：这里的 url '/login' 会被拼接成 '/api/login' (因为 request.ts 设置了 baseurl)
  // 如果后端的路由没有 /api 前缀，可能需要调整 request.ts 的 baseURL 或者在这里写完整路径
  return request.post<any, LoginResponse>('/login', data)
}

// 注册 API
export const register = (data: any) => {
  return request.post('/register', data)
}
