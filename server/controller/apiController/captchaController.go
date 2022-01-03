package apiController

import (
	"2021/magicExcel/server/common"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CaptchaController(ctx *gin.Context) {
	captchaBody, err := common.GenerateCaptchaCode()
	if err != nil {
		zap.L().Error("生成验证码错误", zap.Error(err))
		common.Failed(ctx, err, common.CaptchaFailed, "生成验证码错误")
		return
	}
	common.Success(ctx, captchaBody, "ok")
	return
}
