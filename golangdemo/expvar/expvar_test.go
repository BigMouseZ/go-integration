package main

import (
	"testing"
)

func Test(t *testing.T) {

	// expvar包提供了公共变量的标准接口，如服务的操作计数器。本包通过HTTP在/debug/vars位置以JSON格式导出了这些变量。
	// 对这些公共变量的读写操作都是原子级的。
	/**

	"公共变量"即Var是一个实现了String()函数的接口，定义如下

	type Var interface {
	      String() string
	}
	实际类型的Var包括：Int、Float、String和Map，每个具体的类型都包含这几个函数：

	1）New*()  //  新建一个变量
	2）Set(*)   //  设置这个变量
	3）Add(*)  //  在原有变量上加上另一个变量
	4）String()  // 实现Var接口
	除此之外，Map还有几个特有的函数：
	1）Init()                  // 初始化Map
	2）Get(key string)  // 根据key获取value
	3）Do(f func(Key Value))  // 对Map中的每对key/value执行函数f
	*/

}
