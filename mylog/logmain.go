package main

import (
	"go-integration/mylog/logentity"
)

func main() {
	// 日志测试
	logger := logentity.NewlogEntity(logentity.DEBUG, "./log/", "test.log", "1KB", 1)

	// ziputil.Zip("./log","./mylog/test.zip")

	logger.Debug("debug日志测试dafdvv asd水电费v啊啊所大无是的v求稳上多 阿萨德大是房东 房东")
}
