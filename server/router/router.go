package router

type routerGroup struct {
	articleRouter
	loginRouter
}

var Router = new(routerGroup)
