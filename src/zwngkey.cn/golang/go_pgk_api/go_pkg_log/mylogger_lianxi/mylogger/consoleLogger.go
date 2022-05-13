/*
 * @Author: zwngkey
 * @Date: 2022-05-13 06:24:07
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-14 06:15:09
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

func NewConsoleLogger(logLevel string) (cl *ConsoleLogger) {
	lLevel := parseLogLevel(logLevel)
	if checkLogLevel(lLevel) {
		log.Fatalf("please input legal log level!!!")
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
	timeString := now.Format("2006/01/02 15:04:05.000")
	funcName, fileName, lineNo := getInfo(3)

	fmt.Printf("[%s] [%-7s] [%s/%s():%d] %s\n", timeString, parseLogLevelString(ll), fileName, funcName, lineNo, msg)
}

func (cl *ConsoleLogger) Debug(msg string, a ...any) {
	cl.log(Debug, msg, a...)
}
func (cl *ConsoleLogger) Info(msg string, a ...any) {
	cl.log(Info, msg, a...)
}
func (cl *ConsoleLogger) Warning(msg string, a ...any) {
	cl.log(Warning, msg, a...)
}
func (cl *ConsoleLogger) Error(msg string, a ...any) {
	cl.log(Error, msg, a...)
}
func (cl *ConsoleLogger) Fatal(msg string, a ...any) {
	cl.log(Fatal, msg, a...)
}
