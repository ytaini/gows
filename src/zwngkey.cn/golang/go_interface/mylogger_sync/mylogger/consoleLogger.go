/*
 * @Author: wzmiiiiii
 * @Date: 2022-07-16 06:20:04
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-21 16:35:52
 * @Description:
 */
package mylogger

import (
	"fmt"
	"log"
	"time"
)

type ConsoleLogger struct {
	logLevel
}

func NewConsoleLogger(logLevel level) (cl *ConsoleLogger) {
	lLevel := parseLogLevel(logLevel)
	if checkLogLevel(lLevel) {
		log.Fatalf(errIllegal)
	}
	return &ConsoleLogger{
		logLevel: lLevel,
	}
}

func (cl *ConsoleLogger) log(ll logLevel, msg string, a ...any) {
	if cl.logLevel > ll {
		return
	}
	msg = fmt.Sprintf(msg, a...)
	now := time.Now()

	timeString := now.Format(formatString)

	funcName, fileName, lineNo := getInfo(3)
	fmtStr := "[%s] [%-7s] [%s/%s():%d] %s\n"
	fmt.Printf(fmtStr, timeString, parseLogLevelString(ll), fileName, funcName, lineNo, msg)
}

func (cl *ConsoleLogger) Debug(msg string, a ...any) {
	cl.log(debug, msg, a...)
}
func (cl *ConsoleLogger) Info(msg string, a ...any) {
	cl.log(info, msg, a...)
}
func (cl *ConsoleLogger) Warning(msg string, a ...any) {
	cl.log(warning, msg, a...)
}
func (cl *ConsoleLogger) Error(msg string, a ...any) {
	cl.log(err, msg, a...)
}
func (cl *ConsoleLogger) Fatal(msg string, a ...any) {
	cl.log(fatal, msg, a...)
}
