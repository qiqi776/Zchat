import request from '@/utils/request'
import type { ChatSession } from '@/models/chat'

export interface GroupItem {
  uuid: string // 注意：后端 json tag 是 "uuid"
  name: string
  avatar: string
  notice?: string
  ownerId?: string // 后端是 ownerId
  memberCnt?: number // 后端是 memberCnt
}

// 对应后端的 MyUserListRespond
export interface ContactItem {
  user_id: string
  user_name: string
  avatar: string
}

// 对应后端的 GroupInfo
export interface CreateGroupParams {
  ownerId: string
  name: string
  notice?: string
  avatar?: string
  addMode: boolean
}

// 获取会话列表
export const getSessionList = () => {
  return request.get<ChatSession[]>('/chat/sessions')
}

// 获取某个会话的历史消息
export const getHistoryMessages = (sessionId: string) => {
  return request.get(`/chat/history?sessionId=${sessionId}`)
}

// 获取通讯录列表
export const getContactList = () => {
  return request.get<any, ContactItem[]>('/contact/list')
}

// 创建群聊
export const createGroup = (data: CreateGroupParams) => {
  return request.post('/group/createGroup', data)
}
