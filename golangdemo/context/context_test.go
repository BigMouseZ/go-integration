package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

/*

一、介绍

go中有Context 包，专门用来简化 对于处理单个请求的多个 goroutine 之间与请求域的数据、取消信号、截止时间等相关操作，这些操作可能涉及多个 API 调用。
你可以通过 go get golang.org/x/net/context 命令获取这个包。

例如：在 Go http包的Server中，每一个请求在都有一个对应的 goroutine 去处理。请求处理函数通常会启动额外的 goroutine 用来访问后端服务，比如数据库和RPC服务。
用来处理一个请求的 goroutine 通常需要访问一些与请求特定的数据，比如终端用户的身份认证信息、验证相关的token、请求的截止时间。
当一个请求被取消或超时时，所有用来处理该请求的 goroutine 都应该迅速退出，然后系统才能释放这些 goroutine 占用的资源。

用context会方便很多，context是一个可继承的树状的结构。

二、使用 Context 的程序包需要遵循如下的原则：

Context 变量需要作为第一个参数使用，一般命名为ctx。不要把 Context 存在一个结构体当中
即使方法允许，也不要传入一个 nil 的 Context ，如果你不确定你要用什么 Context 的时候传一个 context.TODO
使用 context 的 Value 相关方法只应该用于在程序和接口中传递的和请求相关的元数据，不要用它来传递一些可选的参数
同样的 Context 可以用来传递到不同的 goroutine 中，Context 在多个goroutine 中是安全的
三、主要方法

func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)

func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

func WithValue(parent Context, key interface{}, val interface{}) Context

*/

func process(ctx context.Context) {
	ret, ok := ctx.Value("trace_id").(int)
	if !ok {
		ret = 21342423
	}

	fmt.Printf("ret:%d\n", ret)

	s, _ := ctx.Value("session").(string)
	fmt.Printf("session:%s\n", s)

}
func TestWithValue(t *testing.T) {
	// 1、WithValue:可以把需要的信息放到context中，需要时把变量取出来
	ctx := context.WithValue(context.Background(), "trace_id", 2222222)
	ctx = context.WithValue(ctx, "session", "sdlkfjkaslfsalfsafjalskfj")
	process(ctx)

}

type Result struct {
	r   *http.Response
	err error
}

func TestWithTimeout(t *testing.T) {
	// 2、 WithTimeout: 可以用来控制goroutine超时，context包中提供的WithTimeout(本质上调用的是WithDeadline) 方法
	// context包中提供的WithTimeout(本质上调用的是WithDeadline) 方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan Result, 1)
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		fmt.Println("http request failed, err:", err)
		return
	}
	go func() {
		resp, err := client.Do(req)
		pack := Result{r: resp, err: err}
		c <- pack
	}()
	select {
	case <-ctx.Done(): // ctx到时，这个channel里面就会有数据。
		tr.CancelRequest(req)
		res := <-c
		fmt.Println("Timeout! err:", res.err)
	case res := <-c:
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("Server Response: %s", out)
	}
	return

}

func worker(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(n, "i exited")
				return // returning not to leak the goroutine
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func test() {
	ctx, cancel := context.WithCancel(context.Background())
	//cancel() 通知子goroutine结束
	defer cancel() // cancel when we are finished consuming integers
	intChan := worker(ctx)
	for n := range intChan {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func TestWithCancel(t *testing.T) {
	// 3、withCancel:WithCancel返回一个继承的Context,这个Context在父Context的Done被关闭时关闭自己的Done通道，或者在自己被Cancel的时候关闭自己的Done。
	// WithCancel同时还返回一个取消函数cancel，这个cancel用于取消当前的Context。
	test()
	time.Sleep(time.Second * 5)
}

func TestWithDeadline(t *testing.T) {
	// 4、withDeadline:deadline保存了超时的时间，当超过这个时间，会触发cancel, 如果超过了过期时间，会自动撤销它的子context

	d := time.Now().Add(50 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}
