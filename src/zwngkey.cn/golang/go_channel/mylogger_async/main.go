/*
 * @Author: zwngkey
 * @Date: 2021-12-12 15:00:25
 * @LastEditTime: 2022-11-23 20:29:09
 * @Description:
 */
package main

import (
	"time"

	"zwngkey.cn/golang/go_channel/mylogger_async/mylogger"
)

var logger mylogger.Logger

func main() {
	// logger = mylogger.NewConsoleLogger(mylogger.DEBUG)

	// logger.Debug("这是一个Debug级别的日志. err: %s", "asfdsaf")
	// logger.Info("这是一个Info级别的日志")
	// logger.Warning("这是一个Waring级别的日志")
	// logger.Error("这是一个Error级别的日志")
	// logger.Fatal("这是一个Fatal级别的日志")

	logger = mylogger.NewFileLogger("", "xx.log", mylogger.INFO)

	// logger, ok := logger.(*mylogger.FileLogger)

	// if !ok {
	// 	fmt.Println("类型错误")
	// 	return
	// }

	// logger.SetFileMaxSize(1024 * 1024)

	// defer logger.CloseErrLogFile()
	// defer logger.CloseLogFile()
	for {
		logger.Debug("这是一个Debug级别的日志 ")
		logger.Info("这是一个Info级别的日志")
		logger.Warning("这是一个Waring级别的日志")
		logger.Error("这是一个Error级别的日志")
		logger.Fatal("这是一个Fatal级别的日志")
		time.Sleep(time.Millisecond)
	}
}
