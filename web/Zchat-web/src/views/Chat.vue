<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { storeToRefs } from 'pinia'
import { useRouter } from 'vue-router'
import type { ChatSession, Message } from '@/models/chat'
import { useChatStore } from '@/stores/chat'
import { useUserStore } from '@/stores/user'
import { WSManager } from '@/utils/websocket'
// import { getSessionList } from '@/api/chat'

const router = useRouter()
const chatStore = useChatStore()
const userStore = useUserStore()

const { messages: allMessages } = storeToRefs(chatStore)
const { userInfo: currentUser } = storeToRefs(userStore)

const sessionList = ref<ChatSession[]>([])

// 当前选中的会话
const currentSession = ref<ChatSession | null>(null)

// 过滤消息
const currentMessages = computed(() => {
  if (!currentSession.value) return []
  const targetId = currentSession.value.user.uuid
  const myId = currentUser.value.uuid

  return allMessages.value.filter(
    (msg) =>
      (msg.sendId === myId && msg.receiveId === targetId) ||
      (msg.sendId === targetId && msg.receiveId === myId),
  )
})

const inputContent = ref('')

// 页面加载时，拉取数据
onMounted(async () => {
  // 安全检查：如果没ID，踢回登录页
  if (!currentUser.value.uuid) {
    alert('请先登录')
    router.push('/')
    return
  }

  // 连接 WebSocket
  WSManager.getInstance().connect(userStore.token)

  try {
    // TODO:
    // const res = await getSessionList()
    // sessionList.value = res.data

    console.log('正在加载会话列表...')
  } catch (error) {
    console.error('加载会话失败', error)
  }
})

const selectSession = (session: ChatSession) => {
  currentSession.value = session
}

const handleSend = () => {
  if (!inputContent.value.trim() || !currentSession.value) return

  const msg: Message = {
    uuid: Date.now().toString(),
    sendId: currentUser.value.uuid,
    receiveId: currentSession.value.user.uuid,
    content: inputContent.value,
    type: 0,
  }

  chatStore.addMyMessage(msg)
  WSManager.getInstance().send(msg)
  inputContent.value = ''
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
          v-for="(msg, index) in allMessages"
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
