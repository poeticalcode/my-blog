package client

import (
	"github.com/gin-gonic/gin"
	"github.com/he-wen-yao/my-blog/server/api"
)

type articleRouter struct {
}

// InitArticleRouter 初始化客户端文章相关路由
func (a articleRouter) InitArticleRouter(router *gin.RouterGroup) {
	group := router.Group("/article")
	clientApi := api.AppApiGroup.ClientApi
	{
		// 注册获取文章详情的路由
		group.GET("/:id", clientApi.FetchArticleDetail)
		// 注册分页获取文章的路由
		group.POST("/list", clientApi.FetchArticleListByPaging)
		// 注册创建文章的路由
		group.POST("", clientApi.CreateArticle)
	}
}
