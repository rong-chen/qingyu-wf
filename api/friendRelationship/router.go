package friendRelationship

import (
	"github.com/gin-gonic/gin"
	"qinyu-wf/utils"
)

type RelationshipRouter struct{}

func (f RelationshipRouter) InitRouter(router *gin.Engine) {
	r := router.Group("friendRelationship")
	{
		r.POST("apply", utils.JWTAuthMiddleware, Apply)
		r.GET("/loadingApply", utils.JWTAuthMiddleware, ApplyList)
		r.GET("/agree/:id", utils.JWTAuthMiddleware, AgreeFriend)
		r.GET("/list", utils.JWTAuthMiddleware, FriendList)
	}
}
