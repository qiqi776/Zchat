import { useChatStore } from '@/stores/chat'

export class WSManager {
  private static instance: WSManager
  private ws: WebSocket | null = null
  private url: string = `ws://${location.host}/ws`

  private constructor() {}

  static getInstance() {
    if (!this.instance) {
      this.instance = new WSManager()
    }
    return this.instance
  }

  connect(token: string) {
    if (this.ws) return

    this.ws = new WebSocket(`${this.url}?token=${token}`)

    this.ws.onopen = () => {
      console.log('WS Connected')
      // 可以发送心跳
    }

    this.ws.onmessage = (event) => {
      const chatStore = useChatStore()
      // 解析后端发来的数据
      try {
        const msg = JSON.parse(event.data)
        chatStore.receiveMessage(msg)
      } catch (e) {
        console.log('Received non-JSON message:', event.data)
      }
    }

    this.ws.onclose = () => {
      console.log('WS Closed, reconnecting...')
      this.ws = null
      setTimeout(() => this.connect(token), 3000) // 断线重连
    }
  }

  send(msg: any) {
    if (this.ws && this.ws.readyState === WebSocket.OPEN) {
      this.ws.send(JSON.stringify(msg))
    } else {
      console.error('WS not ready')
    }
  }
}
