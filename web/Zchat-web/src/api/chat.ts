import request from '@/utils/request'
import type { ChatSession } from '@/models/chat'

// 获取会话列表
export const getSessionList = () => {
  return request.get<ChatSession[]>('/chat/sessions')
}

// 获取某个会话的历史消息
export const getHistoryMessages = (sessionId: string) => {
  return request.get(`/chat/history?sessionId=${sessionId}`)
}

// 创建群聊
export const createGroup = (data: { owner_id: string; name: string; notice?: string }) => {
  // 注意：这里假设后端接口是 /group/create，如果是 /createGroup 请自行修改字符串
  return request.post('/group/createGroup', data)
}
