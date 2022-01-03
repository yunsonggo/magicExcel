package conf

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LoadLogger 加载日志
func LoadLogger() (err error) {
	encoder := logEncoder()
	writer := logWriter()
	level := new(zapcore.Level)
	if err = level.UnmarshalText([]byte(AppConf.LogLevel)); err != nil {
		return
	}
	var core zapcore.Core
	if AppConf.AppMode == gin.ReleaseMode {
		core = zapcore.NewCore(encoder, writer, level)
	} else {
		consoleEnCoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writer, level),
			zapcore.NewCore(consoleEnCoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	}
	caller := zap.AddCaller()
	lg := zap.New(core, caller)
	// 替换zap库中的全局logger
	zap.ReplaceGlobals(lg)
	return
}

func logEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func logWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   AppConf.LogFileName,
		MaxSize:    AppConf.LogMaxSize,    // 单位：M
		MaxBackups: AppConf.LogMaxBackups, // 备份数量
		MaxAge:     AppConf.LogMaxAge,     // 备份天数
		Compress:   false,                 // 是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

