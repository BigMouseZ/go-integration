package main

import (
	"flag"
	"fmt"
	"testing"
)

func TestFlag(t *testing.T) {
	// go flag 包用来解析命令行参数，通过一个简单的例子来了解下

	// 参数1：命令行传递参数的名称  参数2：默认值   参数3：参数的说明
	username := flag.String("name", "", "Input your name")

	// 注意 username 是string 的指针 type: *string
	//  在例如：
	// var ip = flag.Int("flagname", 1234, "help message for flagname")
	// ip 的类型是： *int
	flag.Parse()
	fmt.Println("Hello,", *username)
}
