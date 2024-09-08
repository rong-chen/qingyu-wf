package chat

import (
	"github.com/gin-gonic/gin"
	"qingyu-wf/global"
)

func List(c *gin.Context) {
	id, _ := c.Get("id")
	list := FindList(id.(string))
	c.JSON(200, global.RespMsgData(0, "", list))
}
