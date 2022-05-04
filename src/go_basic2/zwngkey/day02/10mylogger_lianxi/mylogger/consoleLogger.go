package mylogger

import (
	"fmt"
	"time"
)

type ConsoleLogger struct {
	logLevel
}

func (cl *ConsoleLogger) log(ll logLevel, msg string, a ...interface{}) {
	if cl.logLevel > ll {
		return
	}
	msg = fmt.Sprintf(msg, a...)
	now := time.Now()
	timeString := now.Format("2006/01/02 15:04:05.000")
	funcName, fileName, lineNo := getInfo(3)

	fmt.Printf("[%s] [%-7s] [%s/%s():%d] %s\n", timeString, parseLogLevelString(ll), fileName, funcName, lineNo, msg)
}

func (cl *ConsoleLogger) Debug(msg string, a ...interface{}) {
	cl.log(Debug, msg, a...)
}
func (cl *ConsoleLogger) Info(msg string, a ...interface{}) {
	cl.log(Info, msg, a...)
}
func (cl *ConsoleLogger) Warning(msg string, a ...interface{}) {
	cl.log(Warning, msg, a...)
}
func (cl *ConsoleLogger) Error(msg string, a ...interface{}) {
	cl.log(Error, msg, a...)
}
func (cl *ConsoleLogger) Fatal(msg string, a ...interface{}) {
	cl.log(Fatal, msg, a...)
}
