export interface User {
  uuid: string
  nickname: string
  avatar: string
  telephone?: string
  signature?: string
  email?: string
  gender?: number | boolean | string
  birthday?: string
  created_at?: string
}

export interface LoginResponse {
  token: string
  userInfo: User
}
