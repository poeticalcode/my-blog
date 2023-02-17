package router

type routerGroup struct {
	articleRouter
	loginRouter
	fileRouter
}

var Router = new(routerGroup)
