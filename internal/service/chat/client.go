package chat

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"Zchat/pkg/zlog"
)

type Client struct {
	Conn *websocket.Conn // 每个客户端对应一个长连接
	Id   string
	Send chan []byte  	 // 只负责写数据
}

var upgrader = websocket.Upgrader{
	ReadBufferSize: 2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},		
}

// 读消息
func (c *Client) Read() {
	// 记得关闭通道
	defer func() {
		if err := c.Conn.Close(); err != nil {
			zlog.Fatal(err.Error())
		}
		close(c.Send)
	}()
	// 主循环读消息
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			zlog.Error(err.Error())
		} else {
			c.Send <- message
		}
	}
}

// 写消息
func (c *Client) Write() {
	// 记得关闭通道
	defer func() {
		if err := c.Conn.Close(); err != nil {
			zlog.Fatal(err.Error())
		}
		close(c.Send)
	}()
	// 发送消息
	for message := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			zlog.Error(err.Error())
		}
	}
}

// NewClientInit 当接受前端登录消息时就会调用,创建client
func NewClientInit(c *gin.Context, clientId string) error {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		zlog.Fatal(err.Error())
		return err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			zlog.Fatal(err.Error())
		}
	}()
	client := &Client{
		Conn: conn,
		Id: clientId,
		Send: make(chan []byte),
	}
	ChatServer.Login <- client
	go client.Read()
	go client.Write()
	return nil
}