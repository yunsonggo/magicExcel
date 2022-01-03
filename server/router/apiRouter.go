package router

import (
	"2021/magicExcel/server/controller/apiController"
	"github.com/gin-gonic/gin"
)

func ApiRouter(app *gin.Engine) {
	ag := app.Group("/api")
	{
		// 验证码
		ag.GET("/captcha/img", apiController.CaptchaController)
		// 注册
		ag.POST("/user/register", apiController.RegisterController)
		// 登录
		ag.POST("/user/login", apiController.LoginController)
		ApiAuthRouter(ag)
	}
}
