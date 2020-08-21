package main

import (
	"fmt"
	"runtime"
	"testing"
)

/*

尽管 Go 编译器产生的是本地可执行代码，这些代码仍旧运行在 Go 的 runtime（这部分的代码可以在 runtime 包中找到）当中。
这个 runtime 类似 Java 和 .NET 语言所用到的虚拟机，它负责管理包括内存分配、垃圾回收（第 10.8 节）、栈处理、goroutine、channel、切片（slice）、map 和反射（reflection）等等。



runtime 调度器是个非常有用的东西，关于 runtime 包几个方法:
Gosched：让当前线程让出 cpu 以让其它线程运行,它不会挂起当前线程，因此当前线程未来会继续执行

NumCPU：返回当前系统的 CPU 核数量

GOMAXPROCS：设置最大的可同时使用的 CPU 核数(逻辑核数)

Goexit：退出当前 goroutine(但是defer语句会照常执行)

NumGoroutine：返回正在执行和排队的任务总数

GOOS：目标操作系统
*/
func TestNumCPU(t *testing.T) {
	fmt.Println("cpus:", runtime.NumCPU())
	fmt.Println("goroot:", runtime.GOROOT())
	fmt.Println("archive:", runtime.GOOS)
}

func TestGOMAXPROCS(t *testing.T) {
	runtime.GOMAXPROCS(1)
}

/*

这个函数的作用是让当前 goroutine 让出 CPU，当一个 goroutine 发生阻塞，
Go 会自动地把与该 goroutine 处于同一系统线程的其他 goroutine 转移到另一个系统线程上去，以使这些 goroutine 不阻塞。

注意关闭通道使用的是close()方法。
引申思考：close()关闭通道的时候，如果关闭一个已经关闭的通道，会报错。那我们关闭时如何确定通道此时的状态呢？
这里有一篇文章大家可以参考下
https://www.jianshu.com/p/d24dfbb33781

*/
func TestGosched(t *testing.T) {
	runtime.GOMAXPROCS(4)

	exit := make(chan int)
	go func() {
		defer close(exit)
		go func() {
			fmt.Println("b")
		}()
	}()

	for i := 0; i < 4; i++ {
		fmt.Println("a:", i)

		if i == 2 {
			runtime.Gosched() //切换任务
		}
	}
	<-exit
}
