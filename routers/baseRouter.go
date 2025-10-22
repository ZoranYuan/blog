package routers

import (
	"ZoranYuan_blog/api"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct {
}

func (BaseRouter) InitBaseRouter(r *gin.RouterGroup) {
	baseGroup := r.Group("base")
	baseApi := api.ApiApp.BaseApi
	// TODO 加入中间件
	{
		baseGroup.GET("captcha", baseApi.GetCaptcha)
	}
}
