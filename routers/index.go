package routers

type Router struct {
	BaseRouter
	UserRouter
}

var RouterApp = new(Router)
