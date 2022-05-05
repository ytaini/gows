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

func NewConsoleLogger(logLevel string) (cl *ConsoleLogger) {
	lLevel := parseLogLevel(logLevel)
	if checkLogLevel(lLevel) {
		log.Fatalf("please input legal log level!!!")
	}
	return &ConsoleLogger{
		logLevel: lLevel,
	}
}

func NewFlieLogger(filePath, fileName, logLevel string) (fl *FileLogger) {
	lLevel := parseLogLevel(logLevel)
	if checkLogLevel(lLevel) {
		log.Fatalf("please input legal log level!!!")
	}
	errFilename := fileName + ".err"
	entry := &FileLogger{
		logLevel:    lLevel,
		filePath:    filePath,
		fileName:    fileName,
		errFileName: errFilename,
		maxSize:     maxSize,
	}
	err := entry.initFile()
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	return entry
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
