package initialize

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/he-wen-yao/my-blog/server/docs"
	"github.com/he-wen-yao/my-blog/server/router"
	"github.com/he-wen-yao/my-blog/server/util/path"
	"github.com/he-wen-yao/my-blog/server/middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RouterInit 初始化一个 gin 的路由
func RouterInit() *gin.Engine {

	app := gin.New()
	// 注册中间件
	app.Use(middleware.GinLogger())
	app.Use(gin.Recovery())
	// 解决跨域问题
	app.Use(cors.Default())
	// 定义路由日志的格式
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("my-blog-server %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	path := path.GetProjectRootPath()
	app.Static("/static", fmt.Sprintf("%s/static", path))

	// Swagger 配置
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	apiGroup := app.Group("/api")
	appRouterGroup := router.Router
	{
		appRouterGroup.InitArticleRouter(apiGroup)
		appRouterGroup.InitLoginRouter(apiGroup)
		appRouterGroup.InitFileRouter(apiGroup)
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
