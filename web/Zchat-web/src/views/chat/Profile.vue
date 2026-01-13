<script setup lang="ts">
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const userStore = useUserStore()
const { userInfo } = storeToRefs(userStore)

const genderText = computed(() => {
  const g = userInfo.value.gender
  if (g === false || g === 0 || g === '0') return '男'
  if (g === true || g === 1 || g === '1') return '女'
  return g || '未知'
})

const formatDate = (dateStr?: string) => {
  if (!dateStr) return '未知'
  return dateStr.replace('T', ' ').split('.')[0]
}

const handleToChat = () => router.push('/chat')
const handleToContact = () => router.push('/contact')
const handleEdit = () => {
  console.log('点击编辑')
}
</script>

<template>
  <div
    class="h-screen w-full bg-[url('@/assets/img/chat_server_background.jpg')] bg-cover bg-center flex items-center justify-center"
  >
    <div class="w-[1000px] h-[600px] bg-white rounded-[30px] shadow-2xl flex overflow-hidden">
      <div
        class="w-[60px] h-full bg-[#FCD3D3] flex flex-col items-center py-6 border-r-[3px] border-gray-300"
      >
        <div class="mb-6">
          <div
            class="w-10 h-10 rounded-full overflow-hidden border-2 border-white cursor-pointer hover:opacity-80 transition"
          >
            <img
              :src="
                userInfo.avatar ||
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
          <button
            @click="handleToContact"
            class="hover:text-gray-900 hover:scale-110 transition"
            title="通讯录"
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
          <button class="text-gray-900 scale-110 transition cursor-default" title="主页">
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

      <div class="flex-1 bg-white flex flex-col relative">
        <div class="h-16 border-b border-gray-200 flex items-center px-8">
          <h2 class="text-xl font-bold text-gray-800 tracking-wide font-['阿里妈妈数黑体_Bold']">
            我的主页
          </h2>
        </div>

        <div class="flex-1 overflow-y-auto p-10">
          <div class="max-w-2xl mx-auto">
            <div class="flex items-center mb-10">
              <div
                class="w-24 h-24 rounded-full overflow-hidden border-4 border-gray-100 shadow-md"
              >
                <img
                  :src="
                    userInfo.avatar ||
                    'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png'
                  "
                  class="w-full h-full object-cover"
                />
              </div>
              <div class="ml-6">
                <h1 class="text-2xl font-bold text-gray-800">
                  {{ userInfo.nickname || '未设置昵称' }}
                </h1>
                <p class="text-gray-500 mt-1 text-sm">ID: {{ userInfo.uuid }}</p>
              </div>
            </div>

            <div class="grid grid-cols-1 gap-y-6 text-sm">
              <div class="flex border-b border-gray-100 pb-4">
                <span class="w-24 text-gray-500 font-medium">电话</span>
                <span class="text-gray-800 flex-1">{{ userInfo.telephone || '未绑定' }}</span>
              </div>

              <div class="flex border-b border-gray-100 pb-4">
                <span class="w-24 text-gray-500 font-medium">邮箱</span>
                <span class="text-gray-800 flex-1">{{ userInfo.email || '未绑定' }}</span>
              </div>

              <div class="flex border-b border-gray-100 pb-4">
                <span class="w-24 text-gray-500 font-medium">性别</span>
                <span class="text-gray-800 flex-1">{{ genderText }}</span>
              </div>

              <div class="flex border-b border-gray-100 pb-4">
                <span class="w-24 text-gray-500 font-medium">生日</span>
                <span class="text-gray-800 flex-1">{{ userInfo.birthday || '未设置' }}</span>
              </div>

              <div class="flex border-b border-gray-100 pb-4">
                <span class="w-24 text-gray-500 font-medium">个性签名</span>
                <span class="text-gray-800 flex-1">{{
                  userInfo.signature || '这个人很懒，什么都没写'
                }}</span>
              </div>

              <div class="flex border-b border-gray-100 pb-4">
                <span class="w-24 text-gray-500 font-medium">注册时间</span>
                <span class="text-gray-800 flex-1">{{ formatDate(userInfo.created_at) }}</span>
              </div>
            </div>

            <div class="mt-10 flex justify-end">
              <button
                @click="handleEdit"
                class="bg-[#FCD3D3] hover:bg-red-200 text-gray-700 font-medium px-8 py-2.5 rounded-lg shadow-sm transition flex items-center"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-4 w-4 mr-2"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"
                  />
                </svg>
                编辑资料
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
