package router

import (
	"github.com/gin-gonic/gin"
	"github.com/he-wen-yao/my-blog/server/api"
)

type fileRouter struct {
}

// InitLoginRouter 初始化文件操作相关路由
func (a fileRouter) InitFileRouter(router *gin.RouterGroup) {
	group := router.Group("/file")
	fileApi := api.V1.FileApi
	{
		// 上传文件
		group.POST("/image", fileApi.UploadImage)
	}

}
