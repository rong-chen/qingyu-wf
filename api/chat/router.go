package chat

import (
	"github.com/gin-gonic/gin"
	"qingyu-wf/utils"
)

type RouterChat struct {
}

func (receiver RouterChat) InitRouter(router *gin.Engine) {
	r := router.Group("/chat")
	{
		r.GET("list", utils.JWTAuthMiddleware, List)
	}

}
