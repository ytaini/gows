package main

import (
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.Logger

func main() {
	InitLogger()
	defer logger.Sync()
	simpleHttpGet("www.baidu.com")
	simpleHttpGet("http://www.baidu.com")
}

func InitLogger() {
	logger, _ = zap.NewProduction()
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("Error fetching url..", zap.String("url", url), zap.Error(err))
	} else {
		logger.Info("success...", zap.String("statusCode", resp.Status), zap.String("url", url))
		_ = resp.Body.Close()
	}
}
