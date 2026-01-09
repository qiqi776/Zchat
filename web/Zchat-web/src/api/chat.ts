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
