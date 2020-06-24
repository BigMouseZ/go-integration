package logentity

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"
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
	loggerFileMax  string
	loggerFile     *os.File
}

// 结构体
func NewlogEntity(loggerLevel int, loggerFilePath, loggerFileName, loggerFileMax string) *logEntity {
	log := &logEntity{
		loggerLevel:    loggerLevel,
		loggerFilePath: loggerFilePath,
		loggerFileName: loggerFileName,
		loggerFileMax:  loggerFileMax,
	}
	log.initLoaggrtFile(loggerFilePath, loggerFileName, loggerFileMax)
	return log
}

func (l *logEntity) initLoaggrtFile(loggerFilePath, loggerFileName, loggerFileMax string) {
	if len(loggerFileMax) < 2 {
		l.loggerFileMax = "10MB"
	} else {
		unit := loggerFileMax[len(loggerFileMax)-2:]
		switch unit {
		case "KB":
			l.loggerFileMax = loggerFileMax
		case "MB":
			l.loggerFileMax = loggerFileMax
		default:
			l.loggerFileMax = "10MB"
		}
	}
	// os.MkdirAll(path.Dir(fn))
	err := os.MkdirAll(loggerFilePath, os.ModePerm)
	if err == nil {
		file, err := os.OpenFile(loggerFilePath+loggerFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
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
		switch unit {
		case "KB":
			formatFileSize(fileInfo.Size())
			// loggerFileMax
		case "MB":
			l.loggerFileMax = loggerFileMax
		default:
			l.loggerFileMax = "10MB"
		}
		fmt.Printf("文件大小：%v", formatFileSize(fileInfo.Size()))
		fmt.Println()
		fmt.Println(pc, file, line)
		fileName := "[" + path.Base(file) + "]"
		line := "[" + strconv.Itoa(line) + "]"
		_, err := fmt.Fprintln(l.loggerFile, timeNow, fileName, line, a)
		if err != nil {
			fmt.Println("DEBUG日志记录失败！")
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
func formatFileSize(fileSize int64) (size string) {
	if fileSize < 1024 {
		// return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.2fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { // if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}
