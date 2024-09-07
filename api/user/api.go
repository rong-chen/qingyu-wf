package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"qingyu-wf/global"
	"qingyu-wf/utils"
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
	var lp LoginParams
	err := c.BindJSON(&lp)
	if err != nil {
		c.JSON(200, global.RespMsg(7, "参数异常"))
		return
	}
	user := SearchDb("username", lp.Username)
	if user.ID == uuid.Nil {
		c.JSON(200, global.RespMsg(7, "暂无该用户"))
		return
	}
	if ok := utils.Vaild(lp.Password, user.Password); !ok {
		c.JSON(200, global.RespMsg(7, "密码错误"))
		return
	}
	strUsr := user.ID.String()
	jwt, err := utils.GenerateJWT(strUsr)
	if err != nil {
		return
	}
	c.JSON(200, global.RespMsgData(0, "登陆成功", jwt))
}

func GetUserInfo(c *gin.Context) {
	id, _ := c.Get("id")
	user := SearchDb("id", id.(string))
	c.JSON(200, global.RespMsgData(0, "", user))
}

func GetUserInfoById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(200, global.RespMsg(7, "参数不正确"))
	}
	user := SearchDb("id", id)
	c.JSON(200, global.RespMsgData(0, "", user))
}
