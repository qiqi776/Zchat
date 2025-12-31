<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const username = ref('')
const password = ref('')

const handleLogin = () => {
  console.log('Login with:', username.value, password.value)
  // 这里暂时不做真实请求，直接跳转到聊天页(虽然还没做)
  // router.push('/chat')
  alert('登录按钮点击成功！')
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="bg-white p-8 rounded-lg shadow-md w-96">
      <h1 class="text-2xl font-bold mb-6 text-center text-gray-800">Zchat 登录</h1>

      <div class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700">用户名</label>
          <input
            v-model="username"
            type="text"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            placeholder="请输入用户名"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700">密码</label>
          <input
            v-model="password"
            type="password"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            placeholder="请输入密码"
          />
        </div>

        <button
          @click="handleLogin"
          class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          登录
        </button>
      </div>
    </div>
  </div>
</template>
