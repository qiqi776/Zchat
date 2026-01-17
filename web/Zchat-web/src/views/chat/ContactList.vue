<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'
import { useUserStore } from '@/stores/user'
import { createGroup, getContactList } from '@/api/chat' // 确保 api/chat.ts 已定义
import type { ContactItem, GroupItem } from '@/api/chat'

const router = useRouter()
const userStore = useUserStore()
const { userInfo: currentUser } = storeToRefs(userStore)

// --- 状态管理 ---
const contactSearch = ref('')
const isModalVisible = ref(false)
const isCreating = ref(false)

// 折叠面板状态
const expandContacts = ref(true)
const expandMyGroups = ref(false)
const expandJoinedGroups = ref(false)

// 数据列表
const contacts = ref<ContactItem[]>([])
const myGroups = ref<GroupItem[]>([])
const joinedGroups = ref<GroupItem[]>([])

// 创建群聊表单 (对应后端 GroupInfo 结构)
const createForm = reactive({
  name: '',
  notice: '',
  addMode: false, // false: 直接加入, true: 需要审核
  avatar: '',
})

// --- 生命周期 ---
onMounted(async () => {
  try {
    // 获取真实通讯录列表
    const res = await getContactList()
    if (res) {
      contacts.value = res
    }
  } catch (error) {
    console.error('加载通讯录失败:', error)
  }
})

// --- 交互逻辑 ---
const handleToChat = () => router.push('/chat')
const handleToOwnInfo = () => router.push('/profile')

const handleSelectContact = (item: ContactItem) => {
  console.log('点击联系人:', item)
  // TODO: 跳转到聊天页面并选中该联系人
  // router.push({ path: '/chat', query: { target: item.user_id } })
}

// --- 弹窗逻辑 ---
const openCreateModal = () => {
  isModalVisible.value = true
  // 重置表单
  createForm.name = ''
  createForm.notice = ''
  createForm.addMode = false
  createForm.avatar = ''
}

const handleCreateGroup = async () => {
  if (!createForm.name.trim()) {
    alert('请输入群名称')
    return
  }

  isCreating.value = true
  try {
    // 调用 API
    await createGroup({
      ownerId: currentUser.value.uuid, // 必须传群主ID
      name: createForm.name,
      notice: createForm.notice,
      addMode: createForm.addMode,
      avatar: createForm.avatar,
    })

    alert('群聊创建成功！')
    isModalVisible.value = false
    // TODO: 可以在这里刷新 myGroups 列表
  } catch (error) {
    console.error(error)
    alert('创建失败，请稍后重试')
  } finally {
    isCreating.value = false
  }
}
</script>

<template>
  <div
    class="h-screen w-full bg-[url('@/assets/img/chat_server_background.jpg')] bg-cover bg-center flex items-center justify-center"
  >
    <div
      class="w-[1000px] h-[600px] bg-white rounded-[30px] shadow-2xl flex overflow-hidden relative"
    >
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
                  d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z"
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
            <button
              @click="handleToOwnInfo"
              class="hover:text-gray-900 hover:scale-110 transition"
              title="主页"
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
                  d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"
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
              @click="openCreateModal"
              class="w-9 h-9 bg-[#FCD3D3] rounded-lg flex items-center justify-center hover:bg-red-200 transition text-gray-700"
              title="创建群聊"
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
                  :key="item.user_id"
                  @click="handleSelectContact(item)"
                  class="flex items-center p-2 rounded-md hover:bg-white cursor-pointer transition"
                >
                  <img
                    :src="
                      item.avatar ||
                      'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png'
                    "
                    class="w-8 h-8 rounded-full mr-3 bg-white"
                  />
                  <span class="text-sm text-gray-700">{{ item.user_name }}</span>
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
                  :key="item.uuid"
                  class="flex items-center p-2 rounded-md hover:bg-white cursor-pointer transition"
                >
                  <img
                    :src="
                      item.avatar ||
                      'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png'
                    "
                    class="w-8 h-8 rounded-full mr-3 bg-white"
                  />
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
                  :key="item.uuid"
                  class="flex items-center p-2 rounded-md hover:bg-white cursor-pointer transition"
                >
                  <img
                    :src="
                      item.avatar ||
                      'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png'
                    "
                    class="w-8 h-8 rounded-full mr-3 bg-white"
                  />
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

      <div
        v-if="isModalVisible"
        class="absolute inset-0 z-50 bg-black/50 backdrop-blur-sm flex items-center justify-center"
      >
        <div class="bg-white w-96 rounded-2xl shadow-2xl p-6 transform transition-all scale-100">
          <h3 class="text-lg font-bold text-gray-800 mb-6 text-center">创建新群聊</h3>

          <div class="space-y-4">
            <div>
              <label class="block text-xs font-medium text-gray-500 mb-1 ml-1"
                >群名称 <span class="text-red-500">*</span></label
              >
              <input
                v-model="createForm.name"
                type="text"
                placeholder="例如：周末开黑群"
                class="w-full px-4 py-2.5 rounded-lg border border-gray-300 bg-gray-50 focus:bg-white focus:ring-2 focus:ring-[#FCD3D3] focus:border-transparent outline-none transition text-sm"
              />
            </div>

            <div>
              <label class="block text-xs font-medium text-gray-500 mb-1 ml-1">群公告</label>
              <textarea
                v-model="createForm.notice"
                rows="3"
                placeholder="简单介绍一下本群..."
                class="w-full px-4 py-2.5 rounded-lg border border-gray-300 bg-gray-50 focus:bg-white focus:ring-2 focus:ring-[#FCD3D3] focus:border-transparent outline-none transition text-sm resize-none"
              ></textarea>
            </div>

            <div>
              <label class="block text-xs font-medium text-gray-500 mb-1 ml-1">加群方式</label>
              <div class="flex gap-4 mt-1">
                <label class="flex items-center cursor-pointer">
                  <input
                    type="radio"
                    :value="false"
                    v-model="createForm.addMode"
                    class="w-4 h-4 text-pink-500 focus:ring-pink-500 border-gray-300"
                  />
                  <span class="ml-2 text-sm text-gray-700">直接加入</span>
                </label>
                <label class="flex items-center cursor-pointer">
                  <input
                    type="radio"
                    :value="true"
                    v-model="createForm.addMode"
                    class="w-4 h-4 text-pink-500 focus:ring-pink-500 border-gray-300"
                  />
                  <span class="ml-2 text-sm text-gray-700">群主审核</span>
                </label>
              </div>
            </div>
          </div>

          <div class="flex gap-3 mt-8">
            <button
              @click="isModalVisible = false"
              class="flex-1 py-2.5 rounded-lg border border-gray-200 text-gray-600 hover:bg-gray-100 transition text-sm font-medium"
            >
              取消
            </button>
            <button
              @click="handleCreateGroup"
              :disabled="isCreating"
              class="flex-1 py-2.5 rounded-lg bg-[#FCD3D3] text-gray-700 font-bold hover:bg-red-200 transition disabled:opacity-50 text-sm shadow-sm flex justify-center items-center"
            >
              <svg
                v-if="isCreating"
                class="animate-spin -ml-1 mr-2 h-4 w-4 text-gray-700"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
              >
                <circle
                  class="opacity-25"
                  cx="12"
                  cy="12"
                  r="10"
                  stroke="currentColor"
                  stroke-width="4"
                ></circle>
                <path
                  class="opacity-75"
                  fill="currentColor"
                  d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                ></path>
              </svg>
              {{ isCreating ? '创建中' : '立即创建' }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
