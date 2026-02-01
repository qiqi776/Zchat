import request from '@/utils/request'
import type { ChatSession } from '@/models/chat'

// 群组信息
export interface GroupItem {
  group_id: string
  group_name: string
  avatar: string
  notice?: string
  owner_id?: string
  member_cnt?: number
}

// 联系人信息 MyUserListRespond
export interface ContactItem {
  user_id: string
  user_name: string
  avatar: string
}

// 创建群组参数 CreateGroupRequest
export interface CreateGroupParams {
  owner_id: string
  name: string
  notice?: string
  avatar?: string
  add_mode: boolean
}

// 获取会话列表
export const getSessionList = () => {
  return request.get<ChatSession[]>('/chat/sessions')
}

// 获取某个会话的历史消息
export const getHistoryMessages = (sessionId: string) => {
  return request.get(`/chat/history?sessionId=${sessionId}`)
}

// 通讯录-获取联系人列表
export const getUserList = (data: { owner_id: string }) => {
  return request.post<any, ContactItem[]>('/user/getUserList', data)
}

// 通讯录-获取创建的群组
export const getMyGroupList = (data: { owner_id: string }) => {
  return request.post<any, GroupItem[]>('/group/loadMyGroup', data)
}

// 通讯录-创建群聊
export const createGroup = (data: CreateGroupParams) => {
  return request.post('/group/createGroup', data)
}
