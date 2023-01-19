package init

import (
	"github.com/gin-gonic/gin"
	"github.com/he-wen-yao/my-blog/server/router"
	"net/http"
)

// RouterInit 初始化一个 gin 的路由
func RouterInit() *gin.Engine {
	app := gin.Default()
	apiGroup := app.Group("/api")
	appRouterGroup := router.Router
	{
		appRouterGroup.InitArticleRouter(apiGroup)
	}

	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 4000,
			"msg":  "没有此接口",
		})
	})
	return app
}
