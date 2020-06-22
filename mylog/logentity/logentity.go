package logentity

import (
	"fmt"
	"os"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

type logEntity struct {
	loggerLevel    int
	loggerFilePath string
	loggerFileName string
	loggerFile     *os.File
}

// 结构体
func NewlogEntity(loggerLevel int, loggerFilePath, loggerFileName string) *logEntity {
	log := &logEntity{
		loggerLevel:    loggerLevel,
		loggerFilePath: loggerFilePath,
		loggerFileName: loggerFileName,
	}
	log.initLoaggrtFile(loggerFilePath)
	return log
}

func (l *logEntity) initLoaggrtFile(loggerFilePath string) {
	file, err := os.OpenFile(loggerFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic("日志文件创建失败")
	}
	l.loggerFile = file
}

func (l *logEntity) Debug(a ...interface{}) {
	if l.loggerLevel > DEBUG {
		return
	}
	_, err := fmt.Fprint(l.loggerFile, a...)
	if err != nil {
		fmt.Println("日志记录失败！")
	}
}
func (l *logEntity) Info(a ...interface{}) {
	if l.loggerLevel > INFO {
		return
	}
	_, err := fmt.Fprint(l.loggerFile, a...)
	if err != nil {
		fmt.Println("日志记录失败！")
	}
}

func (l *logEntity) Warn(a ...interface{}) {
	if l.loggerLevel > WARN {
		return
	}
	_, err := fmt.Fprint(l.loggerFile, a...)
	if err != nil {
		fmt.Println("日志记录失败！")
	}
}

func (l *logEntity) Error(a ...interface{}) {
	if l.loggerLevel > ERROR {
		return
	}
	_, err := fmt.Fprint(l.loggerFile, a...)
	if err != nil {
		fmt.Println("日志记录失败！")
	}
}

func (l *logEntity) Fatal(a ...interface{}) {
	if l.loggerLevel > FATAL {
		return
	}
	_, err := fmt.Fprint(l.loggerFile, a...)
	if err != nil {
		fmt.Println("日志记录失败！")
	}
}
