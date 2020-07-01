package main

import (
	"fmt"
	"time"
)

// time包提供了时间的显示和测量用的函数。日历的计算采用的是公历。
func main() {
	timeTemplate1 := "2006-01-02 15:04:05" // 常规类型
	timeTemplate2 := "20060102150405"      // 常规类型
	fmt.Println(time.Now().Format(timeTemplate1))
	fmt.Println(time.Now().Format(timeTemplate2))
	// 获取毫秒
	fmt.Println(time.Microsecond)

	// 获取月
	fmt.Println(time.Month(12))
	// 获星期
	fmt.Println(time.Weekday(1))
	// 当前时间
	fmt.Println(time.Now())
	fmt.Println(time.Now().String())
	fmt.Println(time.Now().Day())
	// 当前时间-小时
	fmt.Println(time.Now().Hour())
	fmt.Println(time.Now().Minute())
	fmt.Println(time.Now().Second())
	fmt.Println(time.Now().Nanosecond())
	// 当前时间unix时间戳since 1970 -1- 1
	fmt.Println(time.Now().Unix())

	// 当前时间unix时间戳(nanoseconds),since 1970 -1- 1,
	fmt.Println(time.Now().UnixNano())

	// 当前时间加三个小时
	fmt.Println(time.Now().Add(time.Hour * 3))

	// 时间戳转化成时间
	currentTime := time.Now().Unix()
	tm := time.Unix(currentTime, 0)
	fmt.Println(tm)
}
