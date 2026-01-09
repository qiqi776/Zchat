import { defineStore } from 'pinia'
import type { Message } from '@/models/chat'

export const useChatStore = defineStore('chat', {
  state: () => ({
    messages: [] as Message[],
    currentSessionId: '', // 当前正在聊天的对象 UUID
  }),
  actions: {
    receiveMessage(msg: Message) {
      this.messages.push(msg)
    },
    // 发送消息的动作
    addMyMessage(msg: Message) {
      this.messages.push(msg)
      // 这里调用 WSManager 发送
      import('@/utils/websocket').then(({ WSManager }) => {
        WSManager.getInstance().send(msg)
      })
    },
  },
})
