package user

import "github.com/gin-gonic/gin"

type Router struct {
}

func (Router) InitRouter(router *gin.Engine) {
	r := router.Group("user")
	{
		r.POST("create", CreateApi)
	}
}
