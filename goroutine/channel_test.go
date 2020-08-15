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
	//相当于跑步的接力
	ch := make(chan bool)
	go rece(ch)
	ch <- true
	println("main函数结束")
}

func send(ch chan int) {
	defer close(ch) //
	for i := 0; i < 10; i++ {
		ch <- i
	}

}

//缓冲通道
func TestChannelCache(t *testing.T) {
	//缓冲实现异步，用的比较多
	ch := make(chan int, 100)
	go send(ch)
	//len:获取数据量，cap:获取容量
	println(len(ch), cap(ch))
	//利用for循环取值
	/*for {
		//使用value,ok:=<-ch取值方式，当通道关闭的时候 ok= false
		ret, ok := <-ch
		if !ok {
			break
		}
		println(ret)
	}*/
	//利用for range循环取值========此方法用的较多
	for ret := range ch {
		println(ret)
	}
	println(len(ch), cap(ch))
	println("main函数结束")
}
