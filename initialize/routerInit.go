package initialize

import (
	"ZoranYuan_blog/global"
	"ZoranYuan_blog/middleware"
	"ZoranYuan_blog/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	// TODO 配置属于自己的 Logger() 和 Recover()
	routerApp := routers.RouterApp
	publiceRouterGroup := r.Group(global.Config.System.RouterPrefix)
	privateRouterGroup := r.Group(global.Config.System.RouterPrefix)
	privateRouterGroup.Use(middleware.JwtAuth())
	{
		routerApp.InitBaseRouter(publiceRouterGroup)
	}
	{
		routerApp.InitUserRouter(privateRouterGroup)
	}

	return r
}
