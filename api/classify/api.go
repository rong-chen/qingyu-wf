package classify

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"qingyu-wf/global"
)

func CreateApi(c *gin.Context) {
	type Params struct {
		Label string    `json:"label"`
		Cid   uuid.UUID `json:"cid"`
	}
	var p Params
	err := c.BindJSON(&p)
	if err != nil {
		c.JSON(200, global.RespMsg(7, "参数错误"))
		return
	}
	var tc TableClassify
	tc.Label = p.Label
	tc.CId = p.Cid
	err = Create(tc)
	if err != nil {
		c.JSON(200, global.RespMsg(7, "网络异常"))
		return
	}
	c.JSON(200, global.RespMsg(0, "创建成功"))
}
func List(c *gin.Context) {
	id, _ := c.Get("id")
	tc := SearchDb("c_id", id.(string))
	c.JSON(200, global.RespMsgData(0, "", tc))
}
