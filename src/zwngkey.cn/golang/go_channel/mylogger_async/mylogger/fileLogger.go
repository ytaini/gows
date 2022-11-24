package mylogger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type FileLogger struct {
	logLevel
	filePath, fileName, errFileName string
	maxSize                         int64
	logFile, errLogFile             *os.File
	logCh                           chan *logMsg
}
type logMsg struct {
	logLevel
	funcName, fileName, timestamp, msg string
	lineNo                             int
}

func NewFileLogger(filePath, fileName string, logLevel level) (f *FileLogger) {
	lLevel := parseLogLevel(logLevel)
	if checkLogLevel(lLevel) {
		log.Fatalf(errIllegal)
	}
	errFilename := fileName + ".err"
	f = &FileLogger{
		logLevel:    lLevel,
		filePath:    filePath,
		fileName:    fileName,
		errFileName: errFilename,
		maxSize:     maxSize,
		logCh:       make(chan *logMsg, 50000),
	}
	err := f.initFile()
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	return f
}

func (f *FileLogger) initFile() error {
	file1, err1 := os.OpenFile(f.filePath+f.fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err1 != nil {
		return err1
	}

	file2, err2 := os.OpenFile(f.filePath+f.errFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err2 != nil {
		return err2
	}
	f.logFile = file1
	f.errLogFile = file2

	go f.writeLogBackground()
	return nil
}

func (f *FileLogger) SetLogLevel(logLevel level) {
	ll := parseLogLevel(logLevel)
	if checkLogLevel(ll) {
		log.Fatalf(errIllegal)
	}
	f.logLevel = ll
}

func (f *FileLogger) SetFileMaxSize(size int64) {
	f.maxSize = size
}

func (f *FileLogger) checkFileSize(fl *os.File) {
	fileInfo, err := fl.Stat()
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	if fileInfo.Size() < f.maxSize {
		return
	}
	err1 := fl.Close()
	if err1 != nil {
		log.Fatalf("err: %v", err)
	}

	flag := strings.HasSuffix(fileInfo.Name(), ".err")
	now := time.Now().UnixNano()
	var fullName string
	if flag {
		fullName = f.filePath + f.errFileName
	} else {
		fullName = f.filePath + f.fileName
	}

	newName := fullName + ".bak" + fmt.Sprintf("%v", now)

	os.Rename(fullName, newName)

	file, err := os.OpenFile(fullName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("open file failed. err: %v\n", err)
	}

	if flag {
		f.errLogFile = file
	} else {
		f.logFile = file
	}
}

func (f *FileLogger) writeLogBackground() {
	for {
		select {
		case t := <-f.logCh:
			f.checkFileSize(f.logFile)
			fmtStr := "[%s] [%-7s] [%s/%s():%d] %s\n"
			tmp := fmt.Sprintf(fmtStr, t.timestamp, parseLogLevelString(f.logLevel), t.fileName, t.funcName, t.lineNo, t.msg)
			fmt.Fprint(f.logFile, tmp)
			if t.logLevel >= err {
				f.checkFileSize(f.errLogFile)
				fmt.Fprint(f.errLogFile, tmp)
			}
		default:
			time.Sleep(time.Second / 10)
		}
	}
}

func (f *FileLogger) log(ll logLevel, msg string, a ...any) {
	if f.logLevel > ll {
		return
	}
	funcName, fileName, lineNo := getInfo(3)
	tmp := &logMsg{
		logLevel:  ll,
		msg:       fmt.Sprintf(msg, a...),
		timestamp: time.Now().Format(formatString),
		funcName:  funcName,
		fileName:  fileName,
		lineNo:    lineNo,
	}
	select {
	case f.logCh <- tmp:
	default:
	}
}

func (f *FileLogger) CloseLogFile() {
	f.logFile.Close()
}

func (f *FileLogger) CloseErrLogFile() {
	f.errLogFile.Close()
}

func (f *FileLogger) Debug(msg string, a ...any) {
	f.log(debug, msg, a...)
}
func (f *FileLogger) Info(msg string, a ...any) {
	f.log(info, msg, a...)
}
func (f *FileLogger) Warning(msg string, a ...any) {
	f.log(warning, msg, a...)
}
func (f *FileLogger) Error(msg string, a ...any) {
	f.log(err, msg, a...)
}
func (f *FileLogger) Fatal(msg string, a ...any) {
	f.log(fatal, msg, a...)
}
