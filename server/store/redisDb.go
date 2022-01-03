package store

import (
	"2021/magicExcel/server/conf"
	"context"

	redisV8 "github.com/go-redis/redis/v8"
)

func InitRedis() (err error) {
	RedisDb = redisV8.NewClient(&redisV8.Options{
		Addr:     conf.AppConf.RedisAddr,
		Password: conf.AppConf.RedisPass,
		DB:       conf.AppConf.RedisDb,
	})
	_, err = RedisDb.Ping(context.Background()).Result()
	return
}
