package router

import "github.com/he-wen-yao/my-blog/server/router/client"

type appRouterGroup struct {
	ClientRouter client.RouterGroup
}

var AppRouterGroup = new(appRouterGroup)
