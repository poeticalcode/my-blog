package router

import "github.com/gin-gonic/gin"

type Router interface {
	RegisterRouter(router *gin.RouterGroup)
}

var RouterArray = []Router{}

func InitRouter() {
	app := gin.Default()
	
	router := app.Group("/api") // API 组
	for _, item := range RouterArray {
		item.RegisterRouter(router)
	}
}
