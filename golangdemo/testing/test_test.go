package main

import (
	"fmt"
	"testing"
)

/*
Go 程序编写三类测试，即：功能测试（test）、基准测试（benchmark，也称性
能测试），以及示例测试（example）。

测试源码文件的主名称应该以被测源码文件的主名称为前导，并且必须以“_test”为后缀。
对于功能测试函数来说，其名称必须以Test为前缀，并且参数列表中只应有一个*testing.T类型的参数声明。
对于性能测试函数来说，其名称必须以Benchmark为前缀，并且唯一参数的类型必须是*testing.B类型的。
对于示例测试函数来说，其名称必须以Example为前缀，但对函数的参数列表没有强制规定
*/

/*
一个test文件中有多个test case时，如何控制执行顺序

1.使用t.Run来控制执行顺序和输出
*/
func Print1to20() int {
	res := 0
	for i := 0; i <= 20; i++ {
		res += i
	}
	return res
}
func TestPrint1to20(t *testing.T) {
	res := Print1to20()
	fmt.Println("Hey")
	if res != 210 {
		t.Errorf("Wrong result of Print1to20")
	}
}
func testPrint1to20(t *testing.T) {
	res := Print1to20()
	fmt.Println("Hey")
	if res != 210 {
		t.Errorf("Wrong result of Print1to20")
	}
}

func testPrint2(t *testing.T) {
	res := Print1to20()
	fmt.Println("Hey2")
	res++
	if res != 211 {
		t.Errorf("Wrong result of Print2")
	}
}
func TestAll(t *testing.T) {
	t.Run("TestPrint1to20", testPrint1to20)
	t.Run("TestPrint2", testPrint2)
}

func Test(t *testing.T) {

	fmt.Println("功能测试")
}

/*
八、benchmark函数

benchmark函数一般以Benchmark开头

benchmark的case一般会运行b.N次，每次执行都会如此

在执行过程中会根据实际case的执行时间是否稳定会增加b.N的次数以达到稳态

benchmark同样受m.Run()方法控制

*/
func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}

/*
2.使用TestMain作为初始化test，并且使用m.Run()来调用其他tests可以完成一些需要初始化操作的testing，比如数据库连接，文件打开等；

如果没有在TestMain()中调用m.Run()则除了TestMain()以外其他的tests都不会被执行

*/
func TestMain(m *testing.M) {
	fmt.Println("tests begin")
	m.Run()
}
