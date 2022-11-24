/*
 * @Author: wzmiiiiii
 * @Date: 2022-07-16 06:20:04
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-21 16:36:35
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

type level string

var maxSize int64 = 1 * 1024 * 1024

const (
	ILLEGAL = "ILLEGAL"
	DEBUG   = "DEBUG"
	INFO    = "INFO"
	WARNING = "WARNING"
	ERROR   = "ERROR"
	FATAL   = "FATAL"
)

const (
	illegal logLevel = iota
	debug
	info
	warning
	err
	fatal
)

const formatString = "2006/01/02 15:04:05.000"

const (
	errIllegal = "please input legal log level!!!"
	errCaller  = "runtime.Caller() failed!!\n"
)

type Logger interface {
	Debug(msg string, a ...any)
	Info(msg string, a ...any)
	Warning(msg string, a ...any)
	Error(msg string, a ...any)
	Fatal(msg string, a ...any)
}

func checkLogLevel(ll logLevel) bool {
	return ll == illegal
}

func parseLogLevel(ll level) logLevel {
	switch ll {
	case DEBUG:
		return debug
	case INFO:
		return info
	case WARNING:
		return warning
	case ERROR:
		return err
	case FATAL:
		return fatal
	default:
		return illegal
	}
}

func parseLogLevelString(ll logLevel) string {
	switch ll {
	case debug:
		return DEBUG
	case info:
		return INFO
	case warning:
		return WARNING
	case err:
		return ERROR
	case fatal:
		return FATAL
	default:
		return ILLEGAL
	}
}

// 获取调用者函数名,所在文件,所在行号.
func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, filePath, lineNo, ok := runtime.Caller(skip)
	if !ok {
		log.Fatalf(errCaller)
	}
	temp := strings.Split(runtime.FuncForPC(pc).Name(), "/")
	funcName = temp[len(temp)-1]
	fileName = path.Base(filePath)
	return
}
