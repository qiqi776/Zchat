<script setup lang="ts">
import { ref } from 'vue'
import type { ChatSession, Message } from '@/models/chat'

// --- 模拟数据 (等后端接口好了，再换成真实请求) ---
const currentUser = ref({
  uuid: 'my-uuid',
  avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Felix',
})

// 模拟左侧会话列表
const sessionList = ref<ChatSession[]>([
  {
    user: {
      uuid: 'u1',
      nickname: 'Apylee后继',
      avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Annie',
    },
    lastMessage: '项目进度怎么样了？',
    unread: 2,
    updateTime: '10:42',
  },
  {
    user: {
      uuid: 'u2',
      nickname: '产品经理',
      avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Bob',
    },
    lastMessage: '这个需求还得改改...',
    unread: 0,
    updateTime: '昨天',
  },
])

// 模拟当前选中的会话
const currentSession = ref<ChatSession | null>(sessionList.value[0] || null)
// 模拟右侧聊天记录
const messageList = ref<Message[]>([
  { sendId: 'u1', receiveId: 'my-uuid', content: '嗨，你好！', type: 0 },
  { sendId: 'my-uuid', receiveId: 'u1', content: '你好呀，在写代码呢。', type: 0 },
  { sendId: 'u1', receiveId: 'my-uuid', content: '基于Go的分布式IM项目进度如何？', type: 0 },
])

const inputContent = ref('')

// --- 交互逻辑 ---
const selectSession = (session: ChatSession) => {
  currentSession.value = session
  // TODO: 这里以后要调用加载历史消息的接口
}

const handleSend = () => {
  if (!inputContent.value.trim() || !currentSession.value) return

  // 1. 构造消息对象
  const msg: Message = {
    sendId: currentUser.value.uuid,
    receiveId: currentSession.value.user.uuid,
    content: inputContent.value,
    type: 0,
  }

  // 2. 推入列表 (前端先展示，假装发成功了)
  messageList.value.push(msg)

  // 3. 清空输入框
  inputContent.value = ''

  // TODO: 这里以后要调用 WebSocket 发送给后端
  console.log('Sending to backend:', msg)
}
</script>

<template>
  <div class="flex h-screen w-full bg-gray-100 overflow-hidden">
    <div class="w-16 bg-gray-800 flex flex-col items-center py-6 space-y-6">
      <img
        :src="currentUser.avatar"
        class="w-10 h-10 rounded-full border-2 border-gray-600 cursor-pointer hover:border-white transition"
      />

      <div class="text-gray-400 hover:text-white cursor-pointer"><i class="iconfont">💬</i></div>
      <div class="text-gray-400 hover:text-white cursor-pointer"><i class="iconfont">👥</i></div>
      <div class="mt-auto text-gray-400 hover:text-white cursor-pointer">⚙️</div>
    </div>

    <div class="w-72 bg-white border-r border-gray-200 flex flex-col">
      <div class="p-4 bg-gray-50 border-b">
        <input
          type="text"
          placeholder="搜索联系人..."
          class="w-full px-3 py-2 bg-gray-200 rounded text-sm focus:outline-none focus:ring-1 focus:ring-blue-500"
        />
      </div>

      <div class="flex-1 overflow-y-auto">
        <div
          v-for="session in sessionList"
          :key="session.user.uuid"
          @click="selectSession(session)"
          :class="[
            'flex items-center p-3 cursor-pointer hover:bg-gray-100 transition',
            currentSession?.user.uuid === session.user.uuid
              ? 'bg-blue-50 border-l-4 border-blue-500'
              : '',
          ]"
        >
          <img :src="session.user.avatar" class="w-10 h-10 rounded-full bg-gray-300" />
          <div class="ml-3 flex-1 overflow-hidden">
            <div class="flex justify-between items-center">
              <span class="font-medium text-gray-800 truncate">{{ session.user.nickname }}</span>
              <span class="text-xs text-gray-400">{{ session.updateTime }}</span>
            </div>
            <p class="text-sm text-gray-500 truncate mt-1">{{ session.lastMessage }}</p>
          </div>
        </div>
      </div>
    </div>

    <div v-if="currentSession" class="flex-1 flex flex-col bg-[#F5F5F5]">
      <div class="h-14 bg-white border-b flex items-center px-6 justify-between">
        <h2 class="text-lg font-medium text-gray-800">{{ currentSession.user.nickname }}</h2>
        <span class="text-gray-400 cursor-pointer">...</span>
      </div>

      <div class="flex-1 overflow-y-auto p-6 space-y-4">
        <div
          v-for="(msg, index) in messageList"
          :key="index"
          :class="['flex', msg.sendId === currentUser.uuid ? 'justify-end' : 'justify-start']"
        >
          <img
            v-if="msg.sendId !== currentUser.uuid"
            :src="currentSession.user.avatar"
            class="w-9 h-9 rounded-full mr-2"
          />

          <div
            :class="[
              'max-w-[70%] px-4 py-2 rounded-lg text-sm leading-relaxed shadow-sm',
              msg.sendId === currentUser.uuid
                ? 'bg-blue-500 text-white rounded-tr-none'
                : 'bg-white text-gray-800 rounded-tl-none',
            ]"
          >
            {{ msg.content }}
          </div>

          <img
            v-if="msg.sendId === currentUser.uuid"
            :src="currentUser.avatar"
            class="w-9 h-9 rounded-full ml-2"
          />
        </div>
      </div>

      <div class="h-40 bg-white border-t p-4 flex flex-col">
        <div class="flex gap-4 text-gray-500 mb-2">
          <span class="cursor-pointer hover:text-gray-700">😊</span>
          <span class="cursor-pointer hover:text-gray-700">📁</span>
        </div>

        <textarea
          v-model="inputContent"
          @keydown.enter.prevent="handleSend"
          class="flex-1 w-full resize-none outline-none text-gray-700"
          placeholder="输入消息..."
        ></textarea>

        <div class="flex justify-end mt-2">
          <button
            @click="handleSend"
            class="px-6 py-1.5 bg-blue-500 text-white rounded hover:bg-blue-600 text-sm transition"
          >
            发送
          </button>
        </div>
      </div>
    </div>

    <div v-else class="flex-1 flex items-center justify-center bg-gray-50 text-gray-400">
      <div class="text-center">
        <div class="text-6xl mb-4">👋</div>
        <p>选择一个联系人开始聊天</p>
      </div>
    </div>
  </div>
</template>
