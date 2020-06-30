package main

import (
	"fmt"
	"log"
	"time"

	"go-integration/golangdemo/time/timetest"
)

// time包提供了时间的显示和测量用的函数。日历的计算采用的是公历。
func main() {
	var Sunday int
	Sunday = int(time.Sunday)
	fmt.Println(Sunday)
	log.Println(time.Sunday)

	fmt.Println(timetest.Sunday)
}
