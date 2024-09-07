package user

import (
	"github.com/gin-gonic/gin"
	"qingyu-wf/utils"
)

type Router struct {
}

func (Router) InitRouter(router *gin.Engine) {
	r := router.Group("user")
	{
		r.POST("create", CreateApi)
		r.POST("login", Login)
		r.GET("info", utils.JWTAuthMiddleware, GetUserInfo)
		r.GET("info/:id", utils.JWTAuthMiddleware, GetUserInfoById)
	}
}
