package store

import (
	redisV8 "github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GormDb *gorm.DB
	RedisDb *redisV8.Client
)


func InitStore() {
	var err error
	if err = InitMysql();err != nil {
		zap.L().Debug("Init mysql err:",zap.Error(err))
		panic(err)
	}
	if err = InitRedis(); err != nil {
		zap.L().Debug("Init redis err:",zap.Error(err))
		panic(err)
	}
	return
}
