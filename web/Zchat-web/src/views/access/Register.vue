<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { register } from '@/api/auth'

const router = useRouter()

const form = ref({
  nickname: '',
  telephone: '',
  password: '',
})

const isLoading = ref(false)
const errorMessage = ref('')
const successMessage = ref('')

// 校验
const checkTelephoneValid = () => {
  const regex = /^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$/
  return regex.test(form.value.telephone)
}

// 关闭提示框
const closeAlert = () => {
  errorMessage.value = ''
  successMessage.value = ''
}

// 注册
const handleRegister = async () => {
  closeAlert()

  if (!form.value.nickname || !form.value.telephone || !form.value.password) {
    alert('请填写完整注册信息')
    return
  }

  if (!checkTelephoneValid()) {
    errorMessage.value = '请输入有效的手机号码'
    return
  }

  isLoading.value = true

  try {
    await register({
      nickname: form.value.nickname,
      telephone: form.value.telephone,
      password: form.value.password,
    })

    successMessage.value = '注册成功！即将跳转登录...'

    setTimeout(() => {
      router.push('/login')
    }, 1500)
  } catch (error: any) {
    console.error('注册失败:', error)
    errorMessage.value = error.response?.data?.msg || '注册失败，请稍后重试'
  } finally {
    isLoading.value = false
  }
}

const goToLogin = () => {
  router.push('/login')
}
</script>

<template>
  <div
    class="min-h-screen w-full bg-[url('@/assets/img/chat_server_background.jpg')] bg-cover bg-center flex items-center justify-center"
  >
    <div
      class="w-96 bg-white/80 backdrop-blur-md rounded-2xl shadow-xl p-8 transition-all hover:shadow-2xl"
    >
      <h2 class="text-3xl font-bold text-center text-gray-800 mb-8">注册账号</h2>

      <div
        v-if="errorMessage"
        class="mb-4 p-3 bg-red-100 border border-red-200 text-red-700 rounded-lg flex items-center justify-between text-sm animate-pulse"
      >
        <div class="flex items-center">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-5 w-5 mr-2"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
            />
          </svg>
          {{ errorMessage }}
        </div>
        <button @click="errorMessage = ''" class="hover:text-red-900">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-4 w-4"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M6 18L18 6M6 6l12 12"
            />
          </svg>
        </button>
      </div>

      <div
        v-if="successMessage"
        class="mb-4 p-3 bg-green-100 border border-green-200 text-green-700 rounded-lg flex items-center justify-between text-sm"
      >
        <div class="flex items-center">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-5 w-5 mr-2"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
            />
          </svg>
          {{ successMessage }}
        </div>
      </div>

      <div class="space-y-5">
        <div>
          <label class="block text-sm font-medium text-gray-600 mb-1 ml-1">昵称</label>
          <input
            v-model="form.nickname"
            type="text"
            placeholder="3-10个字符"
            class="w-full px-4 py-3 rounded-lg border border-gray-300 bg-white/50 focus:bg-white focus:ring-2 focus:ring-green-400 focus:border-transparent outline-none transition-all duration-200"
            @input="closeAlert"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-600 mb-1 ml-1">账号 (手机号)</label>
          <input
            v-model="form.telephone"
            type="text"
            placeholder="请输入手机号"
            class="w-full px-4 py-3 rounded-lg border border-gray-300 bg-white/50 focus:bg-white focus:ring-2 focus:ring-green-400 focus:border-transparent outline-none transition-all duration-200"
            @input="closeAlert"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-600 mb-1 ml-1">密码</label>
          <input
            v-model="form.password"
            type="password"
            placeholder="请设置密码"
            class="w-full px-4 py-3 rounded-lg border border-gray-300 bg-white/50 focus:bg-white focus:ring-2 focus:ring-green-400 focus:border-transparent outline-none transition-all duration-200"
            @keyup.enter="handleRegister"
            @input="closeAlert"
          />
        </div>

        <button
          @click="handleRegister"
          :disabled="isLoading"
          class="w-full bg-green-600 hover:bg-green-700 text-white font-bold py-3 rounded-lg shadow-lg transition duration-200 transform active:scale-95 disabled:opacity-50 disabled:cursor-not-allowed flex justify-center items-center mt-6"
        >
          <svg
            v-if="isLoading"
            class="animate-spin -ml-1 mr-3 h-5 w-5 text-white"
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
          {{ isLoading ? '注册中...' : '立即注册' }}
        </button>

        <div class="flex justify-center mt-4">
          <span class="text-sm text-gray-500 mr-1">已有账号？</span>
          <button
            @click="goToLogin"
            class="text-sm text-green-600 font-semibold hover:text-green-800 hover:underline transition"
          >
            去登录
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
