package conf

import "go.uber.org/zap"

func Init() {
	var err error
	// 1, 加载配置
	if err = LoadConf();err != nil {
		panic(err)
	}
	// 2, 加载日志
	if err = LoadLogger(); err != nil {
		panic(err)
	}
	zap.L().Debug("init logger success")
	return
}
