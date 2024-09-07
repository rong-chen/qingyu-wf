package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"qingyu-wf/api/Websocket"
	"qingyu-wf/api/classify"
	"qingyu-wf/api/friendRelationship"
	"qingyu-wf/api/user"
)

// routerInterface 定义接口，要求实现 InitRouter 方法
type routerInterface interface {
	InitRouter(engine *gin.Engine)
}

var routerList = []routerInterface{
	new(user.Router),
	new(friendRelationship.RelationshipRouter),
	new(classify.RouterClassify),
	new(Websocket.RouterWebsocket),
}

func Init() {
	router := gin.Default()
	for i := range routerList {
		routerList[i].InitRouter(router)
	}
	router.Use(LoadTls())
	err := router.RunTLS(":8082", "localhost.pem", "localhost-key.pem")
	if err != nil {
		panic(err)
	}
}

func LoadTls() gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:8000",
		})
		err := middleware.Process(c.Writer, c.Request)
		if err != nil {
			//如果出现错误，请不要继续。
			fmt.Println(err)
			return
		}
		// 继续往下处理
		c.Next()
	}
}
