package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log *zap.SugaredLogger

func InitLogger(dest string) *zap.SugaredLogger {
	encoder := getEncoder()

	writeSyncer := getLogWriter(dest)
	core1 := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)

	writeSyncer1 := getErrWriter(dest + ".err")
	core2 := zapcore.NewCore(encoder, writeSyncer1, zapcore.ErrorLevel)

	core := zapcore.NewTee(core1, core2)

	logger := zap.New(core, zap.AddCaller())

	return logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)

}

func getErrWriter(dest string) zapcore.WriteSyncer {
	lumberjackLogger := InitLumberjackLogger(dest)
	return zapcore.AddSync(lumberjackLogger)
}

func getLogWriter(dest string) zapcore.WriteSyncer {
	lumberjackLogger := InitLumberjackLogger(dest)
	return zapcore.AddSync(lumberjackLogger)
}

func InitLumberjackLogger(filePath string) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
}
