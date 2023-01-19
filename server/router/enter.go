package router

type routerGroup struct {
	articleRouter
}

var Router = new(routerGroup)
