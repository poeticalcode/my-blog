package router

import (
	"github.com/gin-gonic/gin"
	"github.com/he-wen-yao/my-blog/server/api"
)

type articleRouter struct {
}

// InitArticleRouter 初始化客户端文章相关路由
func (a articleRouter) InitArticleRouter(router *gin.RouterGroup) {
	group := router.Group("/article")
	articleApi := api.V1.ArticleApi
	{
		// 注册获取文章详情的路由
		group.GET("/:id", articleApi.FetchArticleDetail)
		// 注册删除文章的路由
		group.DELETE("/:id", articleApi.DeleteArticleById)
		// 注册分页获取文章的路由
		group.GET("/list", articleApi.FetchArticleListByPaging)
		// 注册创建文章的路由
		group.POST("", articleApi.CreateArticle)
		// 注册更新文章的路由
		group.PUT("", articleApi.UpdateArticle)
	}
}
