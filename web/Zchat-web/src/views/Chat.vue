<script setup lang="ts">
import { ref, computed, onMounted, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'
import { useUserStore } from '@/stores/user'
import { useChatStore } from '@/stores/chat'
import { WSManager } from '@/utils/websocket'
import { getSessionList } from '@/api/chat'
import type { ChatSession, Message } from '@/models/chat'

const router = useRouter()
const userStore = useUserStore()
const chatStore = useChatStore()
const { messages: allMessages } = storeToRefs(chatStore)
const { userInfo: currentUser } = storeToRefs(userStore)

const inputContent = ref('')
const messageBoxRef = ref<HTMLElement | null>(null)

const sessionList = ref<ChatSession[]>([])

const currentSession = ref<ChatSession | null>(null)

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

const scrollToBottom = async () => {
  await nextTick()
  if (messageBoxRef.value) {
    messageBoxRef.value.scrollTop = messageBoxRef.value.scrollHeight
  }
}

const selectSession = (session: ChatSession) => {
  currentSession.value = session
  session.unread = 0
  scrollToBottom()
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
  scrollToBottom()
}

onMounted(async () => {
  if (!currentUser.value.uuid) {
    router.push('/login')
    return
  }

  WSManager.getInstance().connect(userStore.token)

  try {
    const res = await getSessionList()
    sessionList.value = res as unknown as ChatSession[]

    console.log('加载会话列表成功:', sessionList.value)
  } catch (error) {
    console.error('加载会话列表失败:', error)
  }
})
</script>

<template>
  <div
    class="h-screen w-full bg-[url('@/assets/img/chat_server_background.jpg')] bg-cover bg-center flex items-center justify-center"
  >
    <div class="w-[1000px] h-[600px] bg-white rounded-[30px] shadow-2xl flex overflow-hidden">
      <div class="w-[300px] flex border-r border-gray-200">
        <div class="w-[60px] bg-[#FCD3D3] flex flex-col items-center py-6 space-y-6">
          <div
            class="w-10 h-10 rounded-full overflow-hidden border-2 border-white cursor-pointer hover:opacity-80 transition"
          >
            <img
              :src="
                currentUser.avatar ||
                'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png'
              "
              class="w-full h-full object-cover"
            />
          </div>
          <div class="flex-1"></div>
        </div>

        <div class="flex-1 bg-[#F7F7F7] flex flex-col">
          <div class="h-16 flex items-center justify-center px-4 border-b border-gray-200">
            <input
              type="text"
              placeholder="搜索"
              class="w-full bg-gray-200 text-sm px-3 py-2 rounded-md outline-none focus:bg-white transition"
            />
          </div>

          <div class="flex-1 overflow-y-auto">
            <div
              v-if="sessionList.length === 0"
              class="flex flex-col items-center justify-center h-full text-gray-400 text-sm"
            >
              <p>暂无会话</p>
              <p class="text-xs mt-2">去搜索好友发起聊天吧</p>
            </div>

            <div
              v-for="session in sessionList"
              :key="session.user.uuid"
              @click="selectSession(session)"
              class="flex items-center p-3 cursor-pointer hover:bg-gray-200 transition"
              :class="{ 'bg-gray-200': currentSession?.user.uuid === session.user.uuid }"
            >
              <div class="relative">
                <img :src="session.user.avatar" class="w-10 h-10 rounded-full bg-white" />
                <div
                  v-if="session.unread > 0"
                  class="absolute -top-1 -right-1 bg-red-500 text-white text-[10px] w-4 h-4 rounded-full flex items-center justify-center"
                >
                  {{ session.unread }}
                </div>
              </div>
              <div class="ml-3 flex-1 overflow-hidden">
                <div class="flex justify-between items-center">
                  <span class="text-sm font-medium text-gray-800 truncate">{{
                    session.user.nickname
                  }}</span>
                  <span class="text-xs text-gray-400">{{ session.updateTime }}</span>
                </div>
                <p class="text-xs text-gray-500 truncate mt-1">{{ session.lastMessage }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="flex-1 flex flex-col bg-white">
        <div
          v-if="!currentSession"
          class="flex-1 flex items-center justify-center text-gray-300 select-none"
        >
          <div class="text-center">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-24 w-24 mx-auto mb-4 opacity-50"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="1"
                d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"
              />
            </svg>
            <p>选择一个好友开始聊天</p>
          </div>
        </div>

        <template v-else>
          <div class="h-16 border-b border-gray-200 flex items-center px-6 justify-between">
            <h2 class="text-lg font-bold text-gray-700 font-mono tracking-wide">
              {{ currentSession.user.nickname }}
            </h2>
          </div>

          <div ref="messageBoxRef" class="flex-1 overflow-y-auto p-6 space-y-4 bg-white/50">
            <div
              v-for="msg in currentMessages"
              :key="msg.uuid"
              class="flex items-start gap-3"
              :class="{ 'flex-row-reverse': msg.sendId === currentUser.uuid }"
            >
              <img
                :src="
                  msg.sendId === currentUser.uuid ? currentUser.avatar : currentSession?.user.avatar
                "
                class="w-9 h-9 rounded-full bg-gray-100"
              />
              <div
                class="max-w-[70%] px-4 py-2 rounded-lg text-sm leading-relaxed shadow-sm break-all"
                :class="
                  msg.sendId === currentUser.uuid
                    ? 'bg-[#95EC69] text-black'
                    : 'bg-white border border-gray-200 text-gray-800'
                "
              >
                {{ msg.content }}
              </div>
            </div>
          </div>

          <div class="h-[160px] border-t border-gray-200 flex flex-col">
            <div class="flex-1 p-2 bg-white flex flex-col">
              <textarea
                v-model="inputContent"
                @keydown.enter.prevent="handleSend"
                class="w-full flex-1 resize-none outline-none text-sm text-gray-700 bg-transparent placeholder-gray-300"
                placeholder="请输入内容..."
              ></textarea>
              <div class="flex justify-end pb-1 pr-2">
                <button
                  @click="handleSend"
                  class="bg-[#FCD3D3] hover:bg-red-200 text-gray-700 text-sm px-6 py-1.5 rounded-md transition shadow-sm"
                >
                  发送
                </button>
              </div>
            </div>
          </div>
        </template>
      </div>
    </div>
  </div>
</template>
