package api

import "github.com/he-wen-yao/my-blog/server/api/client"

type appApiGroup struct {
	// 客户端 API
	ClientApi client.ApiGroup
}

// AppApiGroup 博客系统的 APi
var AppApiGroup = new(appApiGroup)
