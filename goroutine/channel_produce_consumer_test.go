package main

import (
	"math/rand"
	"testing"
	"time"
)

//生产消费者模型
//使用goroutine和channel实现一个简单的生产者消费者模型

//生产者：产生随机数 math/rand

//消费者：计算每个随机数的每个位的数字的和  123232=？

type ch1 chan *item
type ch2 chan *result

// 1个生产者 20消费者
type item struct {
	id  int64
	num int64
}

type result struct {
	id  int64
	num int64
}

//生产者
func producer() {
	//1、生成随机数
	var id int64
	for {
		rand.Seed(time.Now().UnixNano())
		ret := rand.Int63()
		id++
		tem := &item{
			id:  id,
			num: ret,
		}
		//2、把随机数发送到通道里面
		ch1 <- tem
	}

}

//消费者
func consumer() {
	tem := <-ch1
}

//打印结果

func printResult() {

}

func TestPC(t *testing.T) {
	//go producer()
	//go consumer()
	//
	////打印结果
	//printResult()

	//给rand加随机数种子实现每次执行都能产生真正的随机数
	rand.Seed(time.Now().UnixNano())
	ret := rand.Intn(101) //[1,101]
	println(ret)
}
