// 对应后端的 Message 结构体
export interface Message {
  uuid?: string
  sessionId?: string
  sendId: string
  receiveId: string
  content: string
  type: number
  createAt?: string
}

// 对应后端的 UserInfo 结构体
export interface User {
  uuid: string
  nickname: string
  avatar: string
  telephone?: string
  signature?: string
}

// 前端专用的会话列表类型
export interface ChatSession {
  user: User
  lastMessage: string
  unread: number
  updateTime: string
}
