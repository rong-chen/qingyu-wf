package Websocket

import "github.com/gin-gonic/gin"

type RouterWebsocket struct{}

func (w RouterWebsocket) InitRouter(router *gin.Engine) {
	r := router.Group("/conn")
	{
		r.GET("ws/:id", HandleWebSocket)
	}
}
