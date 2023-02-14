package router

import (
	"github.com/gin-gonic/gin"
	"github.com/he-wen-yao/my-blog/server/api"
	"github.com/he-wen-yao/my-blog/server/middleware"
)

type loginRouter struct {
}

// InitLoginRouter 初始化登录相关路由
func (a loginRouter) InitLoginRouter(router *gin.RouterGroup) {
	group := router.Group("")
	loginApi := api.V1.LoginApi
	{
		// 注册获取文章详情的路由
		group.POST("/login", loginApi.Login)
		// 获取登录用户信息
		group.GET("/login-user-info", middleware.JWTAuth(), loginApi.FetchLoginUserInfo)
	}

}
