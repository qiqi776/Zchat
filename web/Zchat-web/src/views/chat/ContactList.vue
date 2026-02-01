<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'
import { useUserStore } from '@/stores/user'
import { createGroup, getUserList, getMyGroupList } from '@/api/chat'
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

// 创建群聊表单
const createForm = reactive({
  name: '',
  notice: '',
  add_mode: false,
  avatar: '',
})

// 获取联系人
const fetchContacts = async () => {
  if (!currentUser.value.uuid) return
  try {
    const res = await getUserList({ owner_id: currentUser.value.uuid })
    if (res) contacts.value = res
  } catch (error) {
    console.error('加载通讯录失败:', error)
  }
}

// 获取创建的群组
const fetchMyGroups = async () => {
  if (!currentUser.value.uuid) return
  try {
    const res = await getMyGroupList({ owner_id: currentUser.value.uuid })
    if (res) myGroups.value = res
  } catch (error) {
    console.error('加载群组失败:', error)
  }
}

// 生命周期
onMounted(() => {
  fetchContacts()
  fetchMyGroups()
})

// 交互逻辑
const handleToChat = () => router.push('/chat')
const handleToOwnInfo = () => router.push('/profile')

const handleSelectContact = (item: ContactItem) => {
  console.log('点击联系人:', item)
  // TODO: 跳转到聊天页面并选中该联系人
  // router.push({ path: '/chat', query: { target: item.user_id } })
}

const toggleMyGroups = () => {
  expandMyGroups.value = !expandMyGroups.value
  // 懒加载优化
  if (expandMyGroups.value && myGroups.value.length === 0) {
    fetchMyGroups()
  }
}

// 弹窗逻辑
const openCreateModal = () => {
  isModalVisible.value = true
  // 重置表单
  createForm.name = ''
  createForm.notice = ''
  createForm.add_mode = false
  createForm.avatar = ''
}

const handleCreateGroup = async () => {
  if (!createForm.name.trim()) {
    alert('请输入群名称')
    return
  }

  isCreating.value = true
  try {
    await createGroup({
      owner_id: currentUser.value.uuid,
      name: createForm.name,
      notice: createForm.notice,
      add_mode: createForm.add_mode,
      avatar: createForm.avatar,
    })

    alert('群聊创建成功！')
    isModalVisible.value = false

    await fetchMyGroups()
    expandMyGroups.value = true
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
          </div>
          <div class="flex flex-col space-y-6 text-gray-600">
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
                <div v-if="contacts.length === 0" class="text-xs text-gray-400 pl-4">
                  暂无联系人
                </div>
              </div>
            </div>

            <div class="select-none">
              <div
                @click="toggleMyGroups"
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
                  :key="item.group_id"
                  class="flex items-center p-2 rounded-md hover:bg-white cursor-pointer transition"
                >
                  <img
                    :src="
                      item.avatar ||
                      'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png'
                    "
                    class="w-8 h-8 rounded-full mr-3 bg-white"
                  />
                  <span class="text-sm text-gray-700">{{ item.group_name }}</span>
                </div>
                <div v-if="myGroups.length === 0" class="text-xs text-gray-400 pl-4">
                  暂无创建的群聊
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
                <div class="text-xs text-gray-400 pl-4">暂无加入的群聊</div>
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
                    v-model="createForm.add_mode"
                    class="w-4 h-4 text-pink-500 focus:ring-pink-500 border-gray-300"
                  />
                  <span class="ml-2 text-sm text-gray-700">直接加入</span>
                </label>
                <label class="flex items-center cursor-pointer">
                  <input
                    type="radio"
                    :value="true"
                    v-model="createForm.add_mode"
                    class="w-4 h-4 text-pink-500 focus:ring-pink-500 border-gray-300"
                  />
                  <span class="ml-2 text-sm text-gray-700">群主审核</span>
                </label>
              </div>
            </div>

            <div>
              <label class="block text-xs font-medium text-gray-500 mb-1 ml-1">群头像</label>
              <input
                v-model="createForm.avatar"
                type="text"
                placeholder="URL (选填)"
                class="w-full px-4 py-2.5 rounded-lg border border-gray-300 bg-gray-50 focus:bg-white focus:ring-2 focus:ring-[#FCD3D3] focus:border-transparent outline-none transition text-sm"
              />
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
              <span v-if="!isCreating">立即创建</span>
              <span v-else>创建中...</span>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
