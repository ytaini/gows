package mylogger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var maxSize int64 = 1 * 1024 * 1024

type FileLogger struct {
	logLevel
	filePath, fileName, errFileName string
	maxSize                         int64
	logFile, errLogFile             *os.File
}

func (fl *FileLogger) SetLogLevel(logLevel string) {
	ll := parseLogLevel(logLevel)
	if checkLogLevel(ll) {
		log.Fatalf("please input legal log level!!!")
	}
	fl.logLevel = ll
}

func (fl *FileLogger) SetFileMaxSize(size int64) {
	fl.maxSize = size
}

func (fl *FileLogger) checkFileSize(f *os.File) {

	fileInfo, err := f.Stat()
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	if fileInfo.Size() < fl.maxSize {
		return
	}

	err1 := f.Close()
	if err1 != nil {
		log.Fatalf("err: %v", err)
	}

	flag := strings.HasSuffix(fileInfo.Name(), ".err")
	now := time.Now().UnixNano()
	var fullName string

	if flag {
		fullName = fl.filePath + fl.errFileName
	} else {
		fullName = fl.filePath + fl.fileName
	}

	newName := fullName + ".bak" + fmt.Sprintf("%v", now)

	os.Rename(fullName, newName)

	file, err := os.OpenFile(fullName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("open file failed. err: %v\n", err)
	}

	if flag {
		fl.errLogFile = file
	} else {
		fl.logFile = file
	}

}

func (fl *FileLogger) initFile() error {
	file1, err1 := os.OpenFile(fl.filePath+fl.fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err1 != nil {
		return err1
	}

	file2, err2 := os.OpenFile(fl.filePath+fl.errFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err2 != nil {
		return err2
	}
	fl.logFile = file1
	fl.errLogFile = file2
	return nil
}

func (fl *FileLogger) log(ll logLevel, msg string, a ...interface{}) {
	if fl.logLevel > ll {
		return
	}
	msg = fmt.Sprintf(msg, a...)
	now := time.Now()
	timeString := now.Format("2006/01/02 15:04:05.000")
	funcName, fileName, lineNo := getInfo(3)
	fl.checkFileSize(fl.logFile)
	fmt.Fprintf(fl.logFile, "[%s] [%-7s] [%s/%s():%d] %s\n", timeString, parseLogLevelString(ll), fileName, funcName, lineNo, msg)
	if ll >= Error {
		fl.checkFileSize(fl.errLogFile)
		fmt.Fprintf(fl.errLogFile, "[%s] [%-7s] [%s/%s():%d] %s\n", timeString, parseLogLevelString(ll), fileName, funcName, lineNo, msg)
	}

}

func (fl *FileLogger) CloseLogFile() {
	fl.logFile.Close()
}

func (fl *FileLogger) CloseErrLogFile() {
	fl.errLogFile.Close()
}

func (fl *FileLogger) Debug(msg string, a ...interface{}) {
	fl.log(Debug, msg, a...)
}
func (fl *FileLogger) Info(msg string, a ...interface{}) {
	fl.log(Info, msg, a...)
}
func (fl *FileLogger) Warning(msg string, a ...interface{}) {
	fl.log(Warning, msg, a...)
}
func (fl *FileLogger) Error(msg string, a ...interface{}) {
	fl.log(Error, msg, a...)
}
func (fl *FileLogger) Fatal(msg string, a ...interface{}) {
	fl.log(Fatal, msg, a...)
}
