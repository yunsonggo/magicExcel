package router

import (
	"2021/magicExcel/server/conf"
	"2021/magicExcel/server/middleware"
	"2021/magicExcel/server/until/translation"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func NewRouter() *gin.Engine {
	if conf.AppConf.AppMode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	err := translation.InitTrans("zh")
	if err != nil {
		zap.L().Error("init translation err:", zap.Error(err))
	}
	r := gin.New()
	// 启用全局中间件
	// r.Use(conf.GinLogger(),conf.GinRecovery(true),middleware.CorsMiddleware(),middleware.RateLimitMiddleware(60*time.Second,10))
	r.Use(conf.GinLogger(), middleware.CorsMiddleware(), middleware.RateLimitMiddleware(60*time.Second, 10))
	// 静态文件
	r.StaticFS("/api/static", http.Dir("./public"))
	pprof.Register(r) // 注册pprof 性能分析
	ApiRouter(r)
	return r
}
