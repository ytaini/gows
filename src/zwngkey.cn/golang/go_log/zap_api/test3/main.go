package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
)

var sugarLogger *zap.SugaredLogger

func main() {
	InitLogger()
	defer sugarLogger.Sync()
	simpleHttpGet("www.baidu.com")
	simpleHttpGet("http://www.baidu.com")
}
func InitLogger() {
	encoder := getEncoder()

	writeSyncer := getLogWriter()
	core1 := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	// 有时候除了将全量日志输出到xx.log文件中之外，还希望将ERROR级别的日志单独输出到一个名为xx.err.log的日志文件中。
	// 我们可以通过以下方式实现。
	writeSyncer1 := getErrWriter()
	core2 := zapcore.NewCore(encoder, writeSyncer1, zapcore.ErrorLevel)

	//使用NewTee将core1和core2合并到core
	core := zapcore.NewTee(core1, core2)

	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	//return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)

}

func getErrWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log.err")
	return zapcore.AddSync(file)
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
	// 将日志输出到多个位置
	//return zapcore.NewMultiWriteSyncer(zapcore.AddSync(file), zapcore.AddSync(os.Stderr))
	// 利用io.MultiWriter 将日志输出到多个位置
	//mw := io.MultiWriter(file, os.Stderr)
	//return zapcore.AddSync(mw)
}

func simpleHttpGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}
