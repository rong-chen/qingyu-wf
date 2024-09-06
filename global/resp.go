package global

import "github.com/gin-gonic/gin"

func RespMsg(code uint, msg string) *gin.H {
	return &gin.H{
		"code": code,
		"msg":  msg,
	}
}
func RespMsgData(code uint, msg string, data interface{}) *gin.H {
	return &gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}
