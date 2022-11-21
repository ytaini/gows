/*
  - @Author: wzmiiiiii
  - @Date: 2022-07-16 06:20:04
  - @LastEditors: wzmiiiiii
  - @LastEditTime: 2022-11-21 11:31:07
  - @Description:
    zap 日志库
*/
package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	url := "www.baidu.com"
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)

	sugar.Infof("Failed to fetch URL: %s", url)
}
