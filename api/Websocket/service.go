package Websocket

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"qingyu-wf/api/chat"
	"strings"
	"sync"
)

var (
	connections = sync.Map{} // 线程安全的连接映射
	upgrade     = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // 允许所有跨域请求
		},
	}
	messageChannel = make(chan MessageParams)
)

func HandleWebSocket(c *gin.Context) {
	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	userId := c.Param("id")

	connections.Store(userId, conn)
	go Receiver()
	//go sender()
	go func() {
		for {
			messageType, msg, err := conn.ReadMessage()
			if err != nil {
				break
			}
			switch messageType {
			case websocket.TextMessage:
				params := handleMessage(msg)
				if params.Type == "" {
					conn.WriteMessage(messageType, []byte("数据异常"))
					break
				}
				messageChannel <- params
			}
		}
		defer func() {
			connections.Delete(userId)
			conn.Close()
		}()
	}()
}

func Receiver() {
	//	接收事件
	for {
		msg := <-messageChannel
		switch msg.Type {
		case "text":
			var table chat.ContentChatTable
			table.Receiver = msg.Receiver
			table.Sender = msg.Sender
			table.Content = msg.Message
			table.ContentType = msg.Type
			err := chat.Create(table)
			if err != nil {
				return
			}
			conn, ok := GetSyncMapConn(table.Receiver)
			if ok {
				data, err := json.Marshal(msg)
				if err != nil {
					return
				}
				conn.WriteMessage(websocket.TextMessage, []byte(data))
			}
		}
	}
}
func sender() {
	//	发送事件
}

// GetSyncMapConn 取出conn
func GetSyncMapConn(id string) (*websocket.Conn, bool) {
	conn, ok := connections.Load(id)
	if !ok {
		return nil, ok
	}
	return conn.(*websocket.Conn), ok
}

// 解析 JSON 消息
func handleMessage(msg []byte) (message MessageParams) {
	err := json.Unmarshal(msg, &message)

	if strings.TrimSpace(message.Sender) == "" {
		return
	}

	if strings.TrimSpace(message.Receiver) == "" {
		return
	}

	if strings.TrimSpace(message.Type) == "" {
		return
	}

	if strings.TrimSpace(message.Message) == "" {
		return
	}

	if err != nil {
		return
	}
	return
}
