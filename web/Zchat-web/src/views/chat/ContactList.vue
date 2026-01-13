<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const userStore = useUserStore()
const { userInfo: currentUser } = storeToRefs(userStore)

// 模拟通讯录数据
const contactSearch = ref('')
const expandContacts = ref(true)
const expandMyGroups = ref(false)
const expandJoinedGroups = ref(false)

const contacts = ref([
  { id: 'c1', name: '张三', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Felix' },
  { id: 'c2', name: '李四', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Aneka' },
])
const myGroups = ref([
  {
    id: 'g1',
    name: '前端交流群',
    avatar: 'https://api.dicebear.com/7.x/identicon/svg?seed=group1',
  },
])
const joinedGroups = ref([
  {
    id: 'g2',
    name: '公司全员群',
    avatar: 'https://api.dicebear.com/7.x/identicon/svg?seed=group2',
  },
])

const handleToChat = () => {
  router.push('/chat')
}

const handleSelectContact = (item: any) => {
  console.log('选择了联系人:', item)
  // 这里可以做跳转到聊天，或者在右侧显示资料
}
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
            <button
              @click="handleToChat"
              class="hover:text-gray-900 hover:scale-110 transition"
              title="会话聊天"
            >
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

            <button class="text-gray-900 scale-110 transition cursor-default" title="通讯录">
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
          </div>
        </div>

        <div class="flex-1 bg-[#F7F7F7] flex flex-col">
          <div class="h-16 flex items-center px-4 border-b border-gray-200 gap-2">
            <input
              v-model="contactSearch"
              type="text"
              placeholder="搜索联系人/群聊"
              class="flex-1 bg-gray-200 text-sm px-3 py-2 rounded-md outline-none focus:bg-white transition"
            />
            <button
              class="w-9 h-9 bg-[#FCD3D3] rounded-lg flex items-center justify-center hover:bg-red-200 transition text-gray-700"
            >
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
                  d="M12 4v16m8-8H4"
                />
              </svg>
            </button>
          </div>

          <div class="flex-1 overflow-y-auto p-2 space-y-2">
            <div class="select-none">
              <div
                @click="expandContacts = !expandContacts"
                class="flex items-center px-2 py-2 text-gray-600 cursor-pointer hover:bg-gray-200 rounded-md transition"
              >
                <svg
                  :class="{ 'rotate-90': expandContacts }"
                  class="w-4 h-4 mr-2 transition-transform duration-200"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M9 5l7 7-7 7"
                  />
                </svg>
                <span class="font-medium text-sm">联系人</span>
              </div>

              <div v-show="expandContacts" class="pl-4 space-y-1 mt-1">
                <div
                  v-for="item in contacts"
                  :key="item.id"
                  @click="handleSelectContact(item)"
                  class="flex items-center p-2 rounded-md hover:bg-white cursor-pointer transition"
                >
                  <img :src="item.avatar" class="w-8 h-8 rounded-full mr-3 bg-white" />
                  <span class="text-sm text-gray-700">{{ item.name }}</span>
                </div>
              </div>
            </div>

            <div class="select-none">
              <div
                @click="expandMyGroups = !expandMyGroups"
                class="flex items-center px-2 py-2 text-gray-600 cursor-pointer hover:bg-gray-200 rounded-md transition"
              >
                <svg
                  :class="{ 'rotate-90': expandMyGroups }"
                  class="w-4 h-4 mr-2 transition-transform duration-200"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M9 5l7 7-7 7"
                  />
                </svg>
                <span class="font-medium text-sm">我创建的群聊</span>
              </div>
              <div v-show="expandMyGroups" class="pl-4 space-y-1 mt-1">
                <div
                  v-for="item in myGroups"
                  :key="item.id"
                  class="flex items-center p-2 rounded-md hover:bg-white cursor-pointer transition"
                >
                  <img :src="item.avatar" class="w-8 h-8 rounded-full mr-3 bg-white" />
                  <span class="text-sm text-gray-700">{{ item.name }}</span>
                </div>
              </div>
            </div>

            <div class="select-none">
              <div
                @click="expandJoinedGroups = !expandJoinedGroups"
                class="flex items-center px-2 py-2 text-gray-600 cursor-pointer hover:bg-gray-200 rounded-md transition"
              >
                <svg
                  :class="{ 'rotate-90': expandJoinedGroups }"
                  class="w-4 h-4 mr-2 transition-transform duration-200"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M9 5l7 7-7 7"
                  />
                </svg>
                <span class="font-medium text-sm">我加入的群聊</span>
              </div>
              <div v-show="expandJoinedGroups" class="pl-4 space-y-1 mt-1">
                <div
                  v-for="item in joinedGroups"
                  :key="item.id"
                  class="flex items-center p-2 rounded-md hover:bg-white cursor-pointer transition"
                >
                  <img :src="item.avatar" class="w-8 h-8 rounded-full mr-3 bg-white" />
                  <span class="text-sm text-gray-700">{{ item.name }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div
        class="flex-1 flex flex-col bg-white items-center justify-center text-gray-300 select-none"
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
              d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0z"
            />
          </svg>
          <p>选择一个联系人查看详细资料</p>
        </div>
      </div>
    </div>
  </div>
</template>
