/*
 * @Author: zwngkey
 * @Date: 2022-05-13 06:24:07
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-14 06:15:38
 * @Description:
 */
package mylogger

import (
	"log"
	"path"
	"runtime"
	"strings"
)

type logLevel int

type Logger interface {
	Debug(msg string, a ...any)
	Info(msg string, a ...any)
	Warning(msg string, a ...any)
	Error(msg string, a ...any)
	Fatal(msg string, a ...any)
}

const (
	Illegal logLevel = iota
	Debug
	Info
	Warning
	Error
	Fatal
)

func checkLogLevel(logLevel logLevel) bool {
	return logLevel == Illegal
}

func parseLogLevel(ll string) logLevel {
	ll = strings.ToLower(ll)
	switch ll {
	case "debug":
		return Debug
	case "info":
		return Info
	case "warning":
		return Warning
	case "error":
		return Error
	case "fatal":
		return Fatal
	default:
		return Illegal
	}
}

func parseLogLevelString(ll logLevel) string {
	switch ll {
	case Debug:
		return "Debug"
	case Info:
		return "Info"
	case Warning:
		return "Warning"
	case Error:
		return "Error"
	case Fatal:
		return "Fatal"
	default:
		return "Illegal"
	}
}

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, filePath, lineNo, ok := runtime.Caller(skip)
	if !ok {
		log.Fatalf("runtime.Caller() failed!!\n")
	}
	temp := strings.Split(runtime.FuncForPC(pc).Name(), "/")
	funcName = temp[len(temp)-1]
	fileName = path.Base(filePath)
	return
}
