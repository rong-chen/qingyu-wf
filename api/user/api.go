package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"qinyu-wf/global"
	"qinyu-wf/utils"
)

func CreateApi(c *gin.Context) {
	var cp CreateParams
	err := c.BindJSON(&cp)
	if err != nil {
		c.JSON(200, global.RespMsg(7, "参数错误"))
		return
	}

	if SearchDb("username", cp.Username).ID != uuid.Nil {
		c.JSON(200, global.RespMsg(7, "用户名已被占用"))
		return
	}
	var user User
	user.ID, _ = uuid.NewUUID()
	user.Username = cp.Username
	user.Gender = "3"
	user.NickName = uuid.NewString()
	user.Password, _ = utils.Encryption(cp.Password)
	err = Create(user)
	if err != nil {
		c.JSON(200, global.RespMsg(7, "创建失败"))
		return
	}
	c.JSON(200, global.RespMsg(200, "创建成功"))
}
func Login(c *gin.Context) {

}
