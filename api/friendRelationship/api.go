package friendRelationship

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"qingyu-wf/api/user"
	"qingyu-wf/global"
)

func Apply(c *gin.Context) {
	var cp CreateParams
	err := c.BindJSON(&cp)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, global.RespMsg(7, "参数错误"))
		return
	}
	obj := FindAwaitingAgreeTable(cp.UserId, cp.FriendId)
	if obj.ID != uuid.Nil {
		c.JSON(200, global.RespMsg(7, "申请已存在"))
		return
	}
	var aat AwaitingAgreeTable
	aat.ID, _ = uuid.NewUUID()
	aat.FriendId = cp.FriendId
	aat.UserId = cp.UserId
	aat.Status = "2"
	userId, _ := c.Get("id")
	strUserId := userId.(string)
	if aat.UserId != uuid.MustParse(strUserId) {
		c.JSON(200, global.RespMsg(7, "参数异常"))
		return
	}
	err = ApplyFriend(aat)
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
	strId := c.Param("id")
	if strId == "" {
		c.JSON(200, global.RespMsg(7, "参数错误"))
		return
	}
	err := UpdateAwaitAgreeTableStatus(strId, "1")
	if err != nil {
		c.JSON(200, global.RespMsg(7, "网络异常"))
		return
	}
	table := FindAwaitAgreeTable(strId)
	var fr FriendRelationship
	fr.Status = "1"
	fr.ID, _ = uuid.NewUUID()
	fr.UserId = table.UserId
	fr.FriendId = table.FriendId
	err = CreateRelationshipList(fr)
	fr.ID, _ = uuid.NewUUID()
	fr.UserId = table.FriendId
	fr.FriendId = table.UserId
	err = CreateRelationshipList(fr)
	if err != nil {
		c.JSON(200, global.RespMsg(7, "网络异常"))
		return
	}
	c.JSON(200, global.RespMsg(0, "添加成功"))
}

func FriendList(c *gin.Context) {
	id, _ := c.Get("id")
	list := SearchFriendList(id.(string))
	var frl []FriendsList
	for i := range list {
		frl = append(frl, FriendsList{
			FriendId:   list[i].FriendId,
			FriendInfo: user.SearchDb("id", list[i].FriendId.String()),
			Status:     list[i].Status,
			ClassifyId: list[i].ClassifyId,
		})
	}
	c.JSON(200, global.RespMsgData(0, "", frl))
}
