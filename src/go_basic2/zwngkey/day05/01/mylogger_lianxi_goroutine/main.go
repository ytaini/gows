/*
 * @Author: zwngkey
 * @Date: 2021-12-26 23:24:12
 * @LastEditTime: 2022-04-30 20:52:40
 * @Description:
 */
package main

import (
	"go_basic2/zwngkey/day05/01/mylogger_lianxi_goroutine/mylogger"
)

var logger mylogger.Logger

func main() {
	// log := mylogger.NewConsoleLogger("Debug")

	// log.Debug("这是一个Debug级别的日志. err: %s", "asfdsaf")
	// log.Info("这是一个Info级别的日志")
	// log.Warning("这是一个Waring级别的日志")
	// log.Error("这是一个Error级别的日志")
	// log.Fatal("这是一个Fatal级别的日志")

	logger = mylogger.NewFlieLogger("./", "xx.log", "Info")

	// logger, ok := logger.(*mylogger.FileLogger)

	// if !ok {
	// 	fmt.Println("类型错误")
	// 	return
	// }

	// logger.SetFileMaxSize(1024 * 1024)

	// defer logger.CloseErrLogFile()
	// defer logger.CloseLogFile()

	for i := 0; i < 10000; i++ {
		logger.Debug("这是一个Debug级别的日志 ")
		logger.Info("这是一个Info级别的日志")
		logger.Warning("这是一个Waring级别的日志")
		logger.Error("这是一个Error级别的日志")
		logger.Fatal("这是一个Fatal级别的日志")
	}

}
