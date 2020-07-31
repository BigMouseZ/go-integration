package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // 最后执行，防止程序出错
	fmt.Println("hello world", i)
	if i == 8 {
		panic("报错误 ")
	}

}

func TestWaitGroup(t *testing.T) {
	defer fmt.Println("结束")
	wg.Add(10) // 计数器+10
	for i := 0; i < 10; i++ {
		go hello(i) // 1 创建一个goroutine 2.在新的goroutine中执行hello函数
	}
	fmt.Println("main func")
	time.Sleep(time.Second * 2)
	// 等hello执行完（执行hello函数的那个goroutine执行完）
	wg.Wait()

}
