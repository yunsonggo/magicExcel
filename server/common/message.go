package common

import (
	"2021/magicExcel/server/until/translation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func Success(ctx *gin.Context,data interface{},msg string) {
	ctx.JSON(http.StatusOK,gin.H{
		"code": SuccessCode,
		"msg": CodeMsg[SuccessCode] + "," +msg,
		"data":data,
	})
	return
}

func Failed(ctx *gin.Context,err error,errCode ResErrCode,errMsg interface{}) {
	var msg interface{}
	if errMsg == "" {
		msg = CodeMsg[errCode]
	} else {
		msg = errMsg
	}
	// 是否验证错误类型
	errs,ok := err.(validator.ValidationErrors)
	if !ok {
		ctx.JSON(http.StatusOK,gin.H{
			"code":errCode,
			"msg": msg,
			"error":err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK,gin.H{
			"code":errCode,
			"msg":msg,
			"error": translation.RemoveTopStruct(errs.Translate(*translation.TS)),
		})
	}
	return
}

func Res(ctx *gin.Context,err error,errCode ResErrCode,errMsg string,data interface{},msg string,end bool) bool {
	if err != nil {
		Failed(ctx,err,errCode,errMsg)
		return false
	} else {
		if end {
			Success(ctx,data,msg)
		}
		return true
	}
}