<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { login } from '@/api/auth'

const router = useRouter()
const userStore = useUserStore()

const form = ref({
  telephone: '',
  password: '',
})

const isLoading = ref(false)

const handleLogin = async () => {
  if (!form.value.telephone || !form.value.password) {
    alert('请输入账号和密码')
    return
  }

  isLoading.value = true

  try {
    const res = await login({
      telephone: form.value.telephone,
      password: form.value.password,
    })

    userStore.setUser(res.userInfo, res.token)
    console.log('登录成功')
    router.push('/chat')
  } catch (error: any) {
    console.error('登录失败详情:', error)
    alert('登录失败,请检查账号密码')
  } finally {
    isLoading.value = false
  }
}

const goToRegister = () => {
  router.push('/register')
}
</script>

<template>
  <div
    class="min-h-screen w-full bg-[url('@/assets/img/chat_server_background.jpg')] bg-cover bg-center flex items-center justify-center"
  >
    <div
      class="w-96 bg-white/80 backdrop-blur-md rounded-2xl shadow-xl p-8 transition-all hover:shadow-2xl"
    >
      <h2 class="text-3xl font-bold text-center text-gray-800 mb-8">Zchat 登录</h2>

      <div class="space-y-6">
        <div>
          <label class="block text-sm font-medium text-gray-600 mb-1 ml-1">账号</label>
          <input
            v-model="form.telephone"
            type="text"
            placeholder="请输入手机号"
            class="w-full px-4 py-3 rounded-lg border border-gray-300 bg-white/50 focus:bg-white focus:ring-2 focus:ring-blue-400 focus:border-transparent outline-none transition-all duration-200"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-600 mb-1 ml-1">密码</label>
          <input
            v-model="form.password"
            type="password"
            placeholder="请输入密码"
            class="w-full px-4 py-3 rounded-lg border border-gray-300 bg-white/50 focus:bg-white focus:ring-2 focus:ring-blue-400 focus:border-transparent outline-none transition-all duration-200"
            @keyup.enter="handleLogin"
          />
        </div>

        <button
          @click="handleLogin"
          :disabled="isLoading"
          class="w-full bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 rounded-lg shadow-lg transition duration-200 transform active:scale-95 disabled:opacity-50 disabled:cursor-not-allowed flex justify-center items-center"
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
          {{ isLoading ? '登录中...' : '立即登录' }}
        </button>

        <div class="flex justify-center mt-4">
          <span class="text-sm text-gray-500 mr-1">没有账号？</span>
          <button
            @click="goToRegister"
            class="text-sm text-blue-600 font-semibold hover:text-blue-800 hover:underline transition"
          >
            立即注册
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
