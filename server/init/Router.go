package init

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/he-wen-yao/my-blog/server/middleware"
	"github.com/he-wen-yao/my-blog/server/router"
)

// RouterInit 初始化一个 gin 的路由
func RouterInit() *gin.Engine {

	app := gin.Default()
	// 解决跨域问题
	app.Use(middleware.Cors())
	// 定义路由日志的格式
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("my-blog-server %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	apiGroup := app.Group("/api")
	appRouterGroup := router.Router
	{
		appRouterGroup.InitArticleRouter(apiGroup)
	}
	// 没有对应路由
	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 4000,
			"msg":  "没有此接口",
		})
	})
	// 注册 ping 路由，检测服务是否正常
	app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return app
}
