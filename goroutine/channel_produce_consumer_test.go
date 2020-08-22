package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//生产消费者模型
//使用goroutine和channel实现一个简单的生产者消费者模型

//生产者：产生随机数 math/rand

//消费者：计算每个随机数的每个位的数字的和  123232=？

var itemChan chan *item
var resultChan chan *result

// 1个生产者 20消费者
type item struct {
	id  int64
	num int64
}

type result struct {
	item *item
	sum  int64
}

//生产者
func producer(ch chan *item) {
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
		//2、把随机数发送到通道里面,测试
		ch <- tem
	}

}

//消费者
func consumer(ch chan *item, resChan chan *result) {
	for tem := range ch {
		//
		ret := calc(tem.num)
		//构建result结构体
		resObj := &result{
			item: tem,
			sum:  ret,
		}
		//结构体指针
		resChan <- resObj
	}

}

//计算一个数每个位的和

func calc(num int64) int64 {
	var sum int64
	for num > 0 {
		sum = sum + num%10
		num = num / 10
	}
	return sum
}

func startWorker(n int, ch chan *item, resChan chan *result) {
	for i := 0; i < n; i++ {
		go consumer(ch, resChan)
	}

}

//打印结果
func printResult(resChan chan *result) {
	for ret := range resChan {
		fmt.Printf("id:%v,num:%v,sum:%v \n", ret.item.id, ret.item.num, ret.sum)
		time.Sleep(time.Second)
		//fmt.Println(ret.sum, ret.item)
	}

}

func TestPC(t *testing.T) {
	itemChan = make(chan *item, 100)
	resultChan = make(chan *result, 100)
	go producer(itemChan)
	startWorker(20, itemChan, resultChan)
	//
	////打印结果
	printResult(resultChan)

	//给rand加随机数种子实现每次执行都能产生真正的随机数
	/*rand.Seed(time.Now().UnixNano())
	ret := rand.Intn(101) //[1,101]
	println(ret)*/
}
