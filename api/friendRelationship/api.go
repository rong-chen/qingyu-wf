package friendRelationship

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"qinyu-wf/global"
)

func Apply(c *gin.Context) {
	var cp CreateParams
	err := c.BindJSON(&cp)
	if err != nil {
		c.JSON(200, global.RespMsg(0, "参数错误"))
		return
	}
	var fr FriendRelationship
	fr.ID, _ = uuid.NewUUID()
	fr.Status = "2"
	fr.FriendId = cp.FriendId
	fr.UserId = cp.UserId
	userId, _ := c.Get("id")
	strUserId := userId.(string)
	if fr.UserId != uuid.MustParse(strUserId) {
		c.JSON(200, global.RespMsg(7, "参数异常"))
		return
	}
	err = ApplyFriend(fr)
	if err != nil {
		c.JSON(200, global.RespMsg(7, "网络异常"))
		return
	}
	c.JSON(200, global.RespMsg(0, "添加成功"))
}
func ApplyList(c *gin.Context) {
	userId, _ := c.Get("id")
	list := FindApplyList(userId)
	c.JSON(200, global.RespMsgData(0, "", list))
}
func AgreeFriend(c *gin.Context) {
	type Agree struct {
		UserId   string `json:"userId" binding:"required"`
		Status   string `json:"status" binding:"required"`
		FriendId string `json:"friendId" binding:"required"`
		Id       string `json:"id" binding:"required"`
	}
	var a Agree
	err := c.BindJSON(&a)
	if err != nil {
		c.JSON(200, global.RespMsg(7, "参数错误"))
		return
	}
	err = UpdateFriendRelationshipStatus(a.Id, a.FriendId, a.UserId, a.Status)
	if err != nil {
		c.JSON(200, global.RespMsg(7, "网络异常"))
		return
	}
	if a.Status == "1" {
		var fr FriendRelationship
		fr.ID, _ = uuid.NewUUID()
		fr.Status = a.Status
		fr.UserId = uuid.MustParse(a.UserId)
		fr.FriendId = uuid.MustParse(a.FriendId)
		err := ApplyFriend(fr)
		if err != nil {
			return
		}
	}
	c.JSON(200, global.RespMsg(0, "添加成功"))

}
