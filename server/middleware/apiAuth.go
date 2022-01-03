package middleware

import (
	"2021/magicExcel/server/common"
	"2021/magicExcel/server/until/jwt"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

func ApiAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err error
		var token string
		authHeader := ctx.Request.Header.Get("Authorization")
		if len(authHeader) == 0 {
			err = errors.New("验证失败")
			common.Failed(ctx, err, common.TokenFailed, "验证失败")
			ctx.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			err = errors.New("token格式错误")
			common.Failed(ctx, err, common.TokenFailed, "token格式错误")
			ctx.Abort()
			return
		}
		token = parts[1]
		_, claims, parseErr := jwt.ParseToken(token)
		if parseErr != nil {
			err = errors.New("解析token失败")
			common.Failed(ctx, err, common.TokenFailed, "解析token失败")
			ctx.Abort()
			return
		}
		userId := claims.UserId
		ctx.Set("ctxUserId", userId)
		ctx.Next()
	}
}
