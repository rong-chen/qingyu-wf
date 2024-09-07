package classify

import (
	"github.com/gin-gonic/gin"
	"qingyu-wf/utils"
)

type RouterClassify struct {
}

func (RouterClassify) InitRouter(router *gin.Engine) {
	r := router.Group("/classify")
	{
		r.POST("create", utils.JWTAuthMiddleware, CreateApi)
		r.GET("list", utils.JWTAuthMiddleware, List)
	}
}
