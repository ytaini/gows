package zlog

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"wzmiiiiii.cn/zinx/utils"
)

var sugaredLogger *zap.SugaredLogger

// 初始化zap logger
func initLogger() {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	writerSyncer1 := zapcore.AddSync(os.Stderr)

	lumberjackLogger := &lumberjack.Logger{
		Filename:   utils.Config.LogPath,
		MaxSize:    5,
		MaxAge:     1,
		MaxBackups: 5,
		Compress:   false,
	}

	writerSyncer2 := zapcore.AddSync(lumberjackLogger)

	writerSyncer := zapcore.NewMultiWriteSyncer(writerSyncer1, writerSyncer2)

	core := zapcore.NewCore(encoder, writerSyncer, zap.InfoLevel)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(0))
	sugaredLogger = logger.Sugar()
}

func GetSugaredLogger() *zap.SugaredLogger {
	return sugaredLogger
}
