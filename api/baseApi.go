package api

import (
	"ZoranYuan_blog/global"
	"ZoranYuan_blog/response"
	"ZoranYuan_blog/response/vo"
	"image/color"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type BaseApi struct {
}

func (BaseApi) GetCaptcha(ctx *gin.Context) {
	var source = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bkgColor := &color.RGBA{255, 255, 255, 255} // 红、绿、蓝、透明度（0-255）
	fontsStorage := base64Captcha.DefaultEmbeddedFonts
	// 内置字体路径（库自带的字体文件）
	fonts := []string{"wqy-microhei.ttc"} // 或其他内置字体，如 "simhei.ttf"
	// 定义验证码的配置参数
	driver := base64Captcha.NewDriverString(
		global.Config.Captcha.Height,
		global.Config.Captcha.Width,
		global.Config.Captcha.DotCount,
		0,
		global.Config.Captcha.Length,
		source,       // 字符集
		bkgColor,     // 背景颜色
		fontsStorage, // 字体存储
		fonts,        // 字体文件列表
	)
	// 创建验证码实例，第一个参数是验证码驱动，第二个参数是存储验证码数据的存储对象（这里用默认的内存存储）
	captcha := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	// 生成验证码的ID、Base64编码的图片以及验证码的值
	id, b64s, _, err := captcha.Generate()

	if err != nil {
		global.Log.Error("Failed to generate captcha", zap.Error(err))
		response.FailWithMessage("Failed to generate captcha", ctx)
	}

	response.OkWithData(vo.Captcha{
		CaptchaId: id,
		PicPath:   b64s,
	}, ctx)
}
