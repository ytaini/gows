package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)

var sugarLogger *zap.SugaredLogger

func main() {
	InitLogger()
	defer sugarLogger.Sync()

	for {
		sugarLogger.Debugln("debug")
		sugarLogger.Infoln("info")
		sugarLogger.Warnln("warn")
		sugarLogger.Errorln("error")
		sugarLogger.DPanicln("dpanic")
		time.Sleep(time.Second / 100)
	}
}

func InitLogger() {
	encoder := getEncoder()

	writeSyncer := getLogWriter()
	core1 := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	writeSyncer1 := getErrWriter()
	core2 := zapcore.NewCore(encoder, writeSyncer1, zapcore.ErrorLevel)

	core := zapcore.NewTee(core1, core2)

	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)

}

func InitLumberjackLogger(filePath string) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
}

func getErrWriter() zapcore.WriteSyncer {
	lumberjackLogger := InitLumberjackLogger("./test.log.err")
	return zapcore.AddSync(lumberjackLogger)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberjackLogger := InitLumberjackLogger("./test.log")
	return zapcore.AddSync(lumberjackLogger)
}
