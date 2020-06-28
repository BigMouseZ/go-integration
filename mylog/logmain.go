package main

import (
	"fmt"

	"go-integration/mylog/logentity"
	"go-integration/mylog/ziputil"
)

func main() {
	// 日志测试
	str := "100KB"
	fmt.Print(str[len(str)-2:])
	logger := logentity.NewlogEntity(logentity.DEBUG, "./log/", "test.log","10KB")

	ziputil.Zip("./log/test.log","./log/test.zip")

	logger.Debug("debug日志测试dafdvv asd水电费v啊啊所大无是的v求稳上多 阿萨德大是房东 房东")
}
