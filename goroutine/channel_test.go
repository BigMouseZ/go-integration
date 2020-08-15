package main

import "testing"

func TestChannel(t *testing.T) {
	//定义一个ch1变量
	// 是一个channel类型
	// 这个channel内部传递的数据是int类型
	var ch1 chan int
	var ch2 chan string
	println("ch1", ch1)
	println("ch2", ch2)
	//channel 是引用类型
	//make 函数初始化：slice、 map channel

	ch3 := make(chan int, 1)
	//通道的操作：发送、接收、关闭
	//发送和接收都用一个符号：<-
	ch3 <- 10
	ret := <-ch3
	println(ret)
	//g关闭：
	close(ch3)
	//1.关闭的通道，能取到对应的数据类型零值
	//2、往关闭的通道中传值，会引发panic
	//3。关闭一个已经关闭的通道，会引发panic

}
func rece(ch chan bool) {
	ret := <-ch
	println(ret)
}

//无缓冲通道
func TestChannelNoCache(t *testing.T) {
	ch := make(chan bool)
	go rece(ch)
	ch <- true
	println("main函数结束")
}

//缓冲通道
func TestChannelCache(t *testing.T) {
	ch := make(chan bool, 1)
	ch <- false
	//len:获取数据量，cap:获取容量
	println(len(ch), cap(ch))
	go rece(ch)
	ch <- true
	println("main函数结束")
}
