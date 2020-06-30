package logentity

import (
	"fmt"
	"math"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"

	"go-integration/mylog/ziputil"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

type logEntity struct {
	loggerLevel      int
	loggerFilePath   string
	loggerFileName   string
	loggerFileMax    string
	loggerFileMaxDay int
	loggerFile       *os.File
}

// 结构体
func NewlogEntity(loggerLevel int, loggerFilePath, loggerFileName, loggerFileMax string, loggerFileMaxDay int) *logEntity {
	log := &logEntity{
		loggerLevel:      loggerLevel,
		loggerFilePath:   loggerFilePath,
		loggerFileName:   loggerFileName,
		loggerFileMax:    loggerFileMax,
		loggerFileMaxDay: loggerFileMaxDay,
	}
	log.initLoaggrtFile()
	return log
}

func (l *logEntity) initLoaggrtFile() {
	if len(l.loggerFileMax) < 2 {
		l.loggerFileMax = "10MB"
	} else {
		/*	unit := l.loggerFileMax[len(l.loggerFileMax)-2:]
			switch unit {
			case "KB":
				l.loggerFileMax = loggerFileMax
			case "MB":
				l.loggerFileMax = loggerFileMax
			default:
				l.loggerFileMax = "10MB"
			}*/
	}
	// os.MkdirAll(path.Dir(fn))
	err := os.MkdirAll(l.loggerFilePath, os.ModePerm)
	if err == nil {
		file, err := os.OpenFile(l.loggerFilePath+l.loggerFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
		if err != nil {
			panic("日志文件创建失败:" + err.Error())
		}
		l.loggerFile = file
	} else {
		panic("日志文件夹创建失败:" + err.Error())
	}

}

func (l *logEntity) Debug(a ...interface{}) {
	if l.loggerLevel > DEBUG {
		return
	}
	timeNow := "[" + time.Now().Format("2006-01-02 15:04:05") + "]"
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		// 判断文件是否超出大小
		fileInfo, _ := l.loggerFile.Stat()
		loggerFileMax := l.loggerFileMax
		unit := loggerFileMax[len(loggerFileMax)-2:]
		loggerFileSize, _ := strconv.ParseFloat(loggerFileMax[0:len(loggerFileMax)-2], 64)
		switch unit {
		case "KB":
			currentSize := formatFileSize(fileInfo.Size(), "KB")
			if currentSize > loggerFileSize {
				// 日志分隔
				err := ziputil.Zip(l.loggerFilePath+l.loggerFileName, l.loggerFilePath+l.loggerFileName+"."+time.Now().Format("2006-01-02")+".zip")
				if err != nil {
					fmt.Println("日志压缩失败")
				} else {
					// 清空日志文件
					fmt.Println(l.loggerFilePath + l.loggerFileName)
					if err := os.Truncate(l.loggerFilePath+l.loggerFileName, 0); err != nil {
						fmt.Println("清空日志文件 异常:", err)
					}

				}
			}
		case "MB":
			currentSize := formatFileSize(fileInfo.Size(), "MB")
			if currentSize > loggerFileSize {
				// 日志分隔
			}
		default:
			l.loggerFileMax = "10MB"
		}
		// fmt.Printf("文件大小：%v", formatFileSize(fileInfo.Size()))
		fmt.Println()
		fmt.Println(pc, file, line)
		fileName := "[" + path.Base(file) + "]"
		line := "[" + strconv.Itoa(line) + "]"
		_, err := fmt.Fprintln(l.loggerFile, timeNow, fileName, line, a)
		if err != nil {
			fmt.Println("DEBUG日志记录失败！", err)
		}
	}

}
func (l *logEntity) Info(a ...interface{}) {
	if l.loggerLevel > INFO {
		return
	}
	timeNow := "[" + time.Now().Format("2006-01-02 15:04:05") + "]"
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		fmt.Println(pc, file, line)
		_, err := fmt.Fprint(l.loggerFile, timeNow, a)
		if err != nil {
			fmt.Println("INFO日志记录失败！")
		}
	}
}

func (l *logEntity) Warn(a ...interface{}) {
	if l.loggerLevel > WARN {
		return
	}
	timeNow := "[" + time.Now().Format("2006-01-02 15:04:05") + "]"
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		fmt.Println(pc, file, line)
		_, err := fmt.Fprint(l.loggerFile, timeNow, a)
		if err != nil {
			fmt.Println("WARN日志记录失败！")
		}
	}
}

func (l *logEntity) Error(a ...interface{}) {
	if l.loggerLevel > ERROR {
		return
	}
	timeNow := "[" + time.Now().Format("2006-01-02 15:04:05") + "]"
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		fmt.Println(pc, file, line)
		_, err := fmt.Fprint(l.loggerFile, timeNow, a)
		if err != nil {
			fmt.Println("ERROR日志记录失败！")
		}
	}
}

func (l *logEntity) Fatal(a ...interface{}) {
	if l.loggerLevel > FATAL {
		return
	}
	timeNow := "[" + time.Now().Format("2006-01-02 15:04:05") + "]"
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		fmt.Println(pc, file, line)
		_, err := fmt.Fprint(l.loggerFile, timeNow, a)
		if err != nil {
			fmt.Println("FATAL日志记录失败！")
		}
	}
}

// 字节的单位转换 保留两位小数
func formatFileSize(fileSize int64, unit string) (size float64) {
	if fileSize < 1024 {
		// return strconv.FormatInt(fileSize, 10) + "B"
		size, _ := strconv.ParseFloat(fmt.Sprintf("%.2fB", float64(fileSize)/float64(1)), 64)
		return size
	} else if unit == "KB" {
		re := math.Round(float64(fileSize) / float64(1024)) // fmt.Sprintf("%.2f", float64(fileSize)/float64(1024))
		// size, _ := strconv.ParseFloat(re, 64)
		size := re
		return size

	} else if unit == "MB" {
		size, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(fileSize)/float64(1024*1024)), 64)
		return size
	} else if unit == "GB" {
		size, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(fileSize)/float64(1024*1024*1024)), 64)
		return size
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		size, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(fileSize)/float64(1024*1024*1024*1024)), 64)
		return size
	} else { // if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		size, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(fileSize)/float64(1024*1024)), 64)
		return size
	}
}
