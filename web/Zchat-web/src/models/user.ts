// 定义用户信息
export interface User {
  uuid: string
  nickname: string
  avatar: string
  telephone?: string
  signature?: string
}

// 登录接口返回的数据
export interface LoginResponse {
  token: string
  userInfo: User
}
