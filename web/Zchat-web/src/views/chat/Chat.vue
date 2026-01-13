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

const handleToContact = () => {
  router.push('/contact')
}

onMounted(async () => {
  if (!currentUser.value.uuid) {
    router.push('/login')
  } else {
    WSManager.getInstance().connect(userStore.token)
  }

  try {
    const res = await getSessionList()
    if (res) sessionList.value = res as unknown as ChatSession[]
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
      <div class="w-[30%] border-r-[3px] border-gray-300 flex">
        <div class="w-[60px] h-full bg-[#FCD3D3] flex flex-col items-center py-6">
          <div class="mb-6">
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
          </div>

          <div class="flex-1 flex flex-col space-y-6 text-gray-600">
            <button class="hover:text-gray-900 hover:scale-110 transition" title="会话聊天">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-6 w-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"
                />
              </svg>
            </button>
            <button class="hover:text-gray-900 hover:scale-110 transition" title="通讯录">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-6 w-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z"
                />
              </svg>
            </button>
            <button class="hover:text-gray-900 hover:scale-110 transition" title="朋友圈">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-6 w-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z"
                />
              </svg>
            </button>
            <button class="hover:text-gray-900 hover:scale-110 transition" title="收藏">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-6 w-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z"
                />
              </svg>
            </button>
          </div>

          <div class="flex flex-col space-y-6 text-gray-600">
            <button class="hover:text-gray-900 hover:scale-110 transition" title="设置">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-6 w-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"
                />
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
                />
              </svg>
            </button>
            <button class="hover:text-gray-900 hover:scale-110 transition" title="主页">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-6 w-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"
                />
              </svg>
            </button>
          </div>
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
          <div class="h-[8%] border-b border-gray-200 flex items-center px-6 justify-between">
            <h2 class="text-lg font-bold text-gray-700 font-['阿里妈妈数黑体_Bold']">
              {{ currentSession.user.nickname }}
            </h2>
          </div>

          <div
            ref="messageBoxRef"
            class="h-[62%] border-t-[3px] border-gray-300 overflow-y-auto p-6 space-y-4 bg-white/50"
          >
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

          <div class="h-[30%] border-t-[3px] border-gray-300 flex flex-col">
            <div class="h-10 bg-[#FCD3D3] flex items-center px-4 space-x-4">
              <button class="text-gray-700 hover:scale-110 transition" title="表情">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-5 w-5"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M14.828 14.828a4 4 0 01-5.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
                  />
                </svg>
              </button>
              <button class="text-gray-700 hover:scale-110 transition" title="文件">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-5 w-5"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M15.172 7l-6.586 6.586a2 2 0 102.828 2.828l6.414-6.586a4 4 0 00-5.656-5.656l-6.415 6.585a6 6 0 108.486 8.486L20.5 13"
                  />
                </svg>
              </button>
              <button class="text-gray-700 hover:scale-110 transition" title="记录">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-5 w-5"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"
                  />
                </svg>
              </button>
            </div>

            <div class="flex-1 p-2 bg-white flex flex-col">
              <textarea
                v-model="inputContent"
                @keydown.enter.prevent="handleSend"
                class="w-full flex-1 resize-none outline-none text-sm text-gray-700 bg-transparent placeholder-gray-300"
                placeholder="请输入内容..."
              ></textarea>

              <div class="flex justify-end pb-1 pr-1">
                <button
                  @click="handleSend"
                  class="bg-[#FCD3D3] hover:bg-red-200 text-gray-700 text-sm px-8 py-2 rounded-br-[25px] rounded-tl-[10px] transition shadow-sm font-bold"
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
