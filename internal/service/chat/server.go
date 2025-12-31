package chat

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"sync"	
	"Zchat/internal/model"
	"Zchat/pkg/zlog"
)

type Server struct {
	Clients  map[string]*Client
	mutex    *sync.Mutex
	Transmit chan []byte
	Login    chan *Client
	Logout   chan *Client
}

var ChatServer *Server

func InitServer() {
	if ChatServer == nil {
		ChatServer = &Server{
			Clients:  make(map[string]*Client),
			mutex:    &sync.Mutex{},
			Transmit: make(chan []byte),
			Login:    make(chan *Client),
			Logout:   make(chan *Client),
		}
	}
}

func (s *Server) Start() {
	// 记得关闭通道
	defer func() {
		close(s.Transmit)
		close(s.Login)
		close(s.Logout)
	}()
	for {
		select {
		case client := <-s.Login:
			{
				s.mutex.Lock()
				s.Clients[client.Id] = client
				s.mutex.Unlock()
				err := client.Conn.WriteMessage(websocket.BinaryMessage, []byte("welcome to chat server"))
				if err != nil {
					zlog.Info(fmt.Sprintf("Client %s logged in\n", client.Id))
				}
				zlog.Info(fmt.Sprintf("Client %s logged in\n", client.Id))
			}
		case client := <-s.Logout:
			{
				s.mutex.Lock()
				delete(s.Clients, client.Id)
				s.mutex.Unlock()
				zlog.Info(fmt.Sprintf("Client %s logged out\n", client.Id))
			}
		case data := <-s.Transmit:
			{
				// 解析消息
				var message model.Message
				if err := json.Unmarshal(data, &message); err != nil {
					zlog.Error(err.Error())
				}
				if message.ReceiveId[0] == 'U' {
					// 发送给user
					receiveClient := s.Clients[message.ReceiveId]
					receiveClient.Send <- data
				} else if message.ReceiveId[0] == 'G' {
					// 发送给Group
					var members []model.UserInfo
					for _, member := range members {
						// 跳过自己
						if member.Uuid != message.SendId {
							sendMessage := message
							sendMessage.SendId = message.ReceiveId
							sendMessage.ReceiveId = message.Uuid
							sendData, err := json.Marshal(sendMessage)
							if err != nil {
								zlog.Error(err.Error())
								break
							}
							client := s.Clients[member.Uuid]
							client.Send <- sendData
						}
					}
				}
			}
		}
	}
}