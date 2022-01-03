package middleware

import (
	"2021/magicExcel/server/conf"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie", "Authorization"}
	corsExample := conf.AppConf.Cors
	// 运行Release模式下才会进行跨域保护 保证开发过程中不会被跨域困扰
	if gin.Mode() == gin.ReleaseMode {
		//corsConfig.AllowOrigins = []string{corsExample}
		corsConfig.AllowOrigins = corsExample
	} else {
		corsConfig.AllowAllOrigins = true
	}
	corsConfig.AllowCredentials = true
	return cors.New(corsConfig)
}
