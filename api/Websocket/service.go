package Websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/http"
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
	uid, err := uuid.Parse(userId)
	if err != nil {
		conn.Close()
		return
	}
	connections.Store(uid, conn)
	go Receiver()
	go sender()
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
		defer conn.Close()
	}()
}

func Receiver() {
	//	接收事件
	for {
		msg := <-messageChannel
		fmt.Println(msg)
	}
}
func sender() {
	//	发送事件
}

// GetSyncMapConn 取出conn
func GetSyncMapConn(id uuid.UUID) (*websocket.Conn, bool) {
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
