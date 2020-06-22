package timeutils

import (
	"fmt"
	"time"
)

func main() {
	// 时间转换的模板，golang里面只能是 "2006-01-02 15:04:05" （go的诞生时间）
	timeTemplate1 := "2006-01-02 15:04:05" // 常规类型
	timeTemplate2 := "20060102150405"      // 常规类型
	t1 := "2019-01-08 13:50:30"            // 外部传入的时间字符串

	t2 := "20190520040930"
	// ======= 将时间字符串转换为时间戳 =======
	stamp, _ := time.ParseInLocation(timeTemplate1, t1, time.Local)  // 使用parseInLocation将字符串格式化返回本地时区时间
	stamp2, _ := time.ParseInLocation(timeTemplate2, t2, time.Local) // 使用parseInLocation将字符串格式化返回本地时区时间
	// time.Now().Format()
	fmt.Println(stamp)
	fmt.Println(stamp2)
}

func GetCurrentTimeByString(t2 string) time.Time {
	timeTemplate2 := "20060102150405" // 常规类型
	stamp2, _ := time.ParseInLocation(timeTemplate2, t2, time.Local) // 使用parseInLocation将字符串格式化返回本地时区时间
	return stamp2
}
