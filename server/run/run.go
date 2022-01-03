package run

import (
	"2021/magicExcel/server/conf"
	"2021/magicExcel/server/store"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ListenOnServerRun(r *gin.Engine) {
	var addr string
	if len(conf.AppConf.AppListen) == 0 {
		zap.L().Debug("load config listen addr failed,listen on default port")
		addr = ":8090"
	} else {
		addr = conf.AppConf.AppListen
	}
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	go func() {
		zap.L().Debug("success:", zap.String("listen and server", srv.Addr))
		fmt.Printf("listen and server: %v\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Debug("run server err:", zap.Error(err))
			panic(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Debug("Shutdown server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		_ = store.RedisDb.Close()
		_ = zap.L().Sync()
		cancel()
	}()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server shutdown error:", zap.Error(err))
	}
	zap.L().Debug("server exiting")
	return
}