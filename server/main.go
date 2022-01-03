package main

import (
	"2021/magicExcel/server/conf"
	"2021/magicExcel/server/router"
	"2021/magicExcel/server/run"
	"2021/magicExcel/server/store"
)

func main() {
	// 1, 初始化配置
	conf.Init()
	// 2, 初始化数据库
	store.InitStore()
	r := router.NewRouter()
	run.ListenOnServerRun(r)
}
