package user

import (
	"github.com/gin-gonic/gin"
	"qinyu-wf/global"
)

func CreateApi(c *gin.Context) {
	c.JSON(200, global.RespMsg(200, "成功"))
}
