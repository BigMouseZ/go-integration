package main

import (
	"fmt"
	"os"
	"testing"
)

/*

1 %v    值的默认格式表示。当输出结构体时，扩展标志（%+v）会添加字段名
2 %#v    值的Go语法表示
3 %T    值的类型的Go语法表示
4 %%    百分号

%b    表示为二进制
%c    该值对应的unicode码值
%d    表示为十进制
%o    表示为八进制
%q    该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
%x    表示为十六进制，使用a-f
%X    表示为十六进制，使用A-F
%U    表示为Unicode格式：U+1234，等价于"U+%04X"

1 %s    直接输出字符串或者[]byte %q    该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
2 %x    每个字节用两字符十六进制数表示（使用a-f）
3 %X    每个字节用两字符十六进制数表示（使用A-F）


1 %f:    默认宽度，默认精度
2 %9f    宽度9，默认精度
3 %.2f   默认宽度，精度2 %9.2f  宽度9，精度2 %9.f   宽度9，精度0

+    总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义）；
-    在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）；
#    切换格式：
      八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）；
     对%q（%#q），如果strconv.CanBackquote返回真会输出反引号括起来的未转义字符串；
     对%U（%#U），如果字符是可打印的，会在输出Unicode格式、空格、单引号括起来的go字面值；
' '    对数值，正数前加空格而负数前加负号；
      对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格；
0    使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面；
*/

func TestPrint(t *testing.T) {

	/*
	   func Print(a ...interface{}) (n int, err error)
	   Print采用默认格式将其参数格式化并写入标准输出。如果两个相邻的参数都不是字符串，会在它们的输出之间添加空格。返回写入的字节数和遇到的任何错误。
	   func Println(a ...interface{})
	   与Print的区别
	   1.会在每个参数间隔中增加空格
	   2.在结尾处自动增加了一个\n参数
	*/
	var len = 0
	// 拼接字符串
	len, _ = fmt.Print("aa", "bb")
	// 验证返回字节数
	fmt.Println("返回字节数", len)

	// 验证有连续两个参数非字符串会增加空格
	len, _ = fmt.Print("aa", "bb", 1, 2, 3, 4)
	fmt.Println("返回字节数", len)

	// 验证Println的输出
	len, _ = fmt.Println("aa", "bb", 1, 2)
	// aabb12 6个字节  3个间隔增加空格 1个字节 最后的\n一个字节 len=10
	fmt.Println("返回字节数", len)

	// Printf根据format参数生成格式化的字符串并写入标准输出。返回写入的字节数和遇到的任何错误。

	var name interface{} = "yinzhengjie"
	fmt.Printf("My name is %v !\n", name)
	var age interface{} = 18
	fmt.Printf("I am [%d] years old。", age)
}
func TestSprint(t *testing.T) {

	// Sprint采用默认格式将其参数格式化，串联所有输出生成并返回一个字符串。如果两个相邻的参数都不是字符串，会在它们的输出之间添加空格。
	s := fmt.Sprint("aa", "bb", 1, 2, 3, 4)
	fmt.Println("返回字符串：", s)

	// Sprintf根据format参数生成格式化的字符串并返回该字符串。
	var name interface{} = "yinzhengjie"
	s = fmt.Sprintf("My name is %v !\n", name)
	fmt.Println("返回字符串：", s)

}

func TestFprint(t *testing.T) {
	logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s", "timeNow", "fileName", 2, "file", "Debug", "msg")

	// Fprint采用默认格式将其参数格式化并写入w。如果两个相邻的参数都不是字符串，会在它们的输出之间添加空格。返回写入的字节数和遇到的任何错误。
	_, _ = fmt.Fprint(os.Stdout, logMsg, "测试", 1, 2, 3)
	// Fprintf 将参数列表 a 填写到格式字符串 format 的占位符中
	// 并将填写后的结果写入 w 中，返回写入的字节数
	var name interface{} = "yinzhengjie"
	_, _ = fmt.Fprintf(os.Stdout, "My name is %v !\n", name)

	// Fprintln采用默认格式将其参数格式化并写入w。总是会在相邻参数的输出之间添加空格并在输出结束后添加换行符。返回写入的字节数和遇到的任何错误。

	_, _ = fmt.Fprintln(os.Stdout, logMsg, "测试", 1, 2, 3)

}

func TestScan(t *testing.T) {
	// func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
	// 作用:Fscanf从r(即实现io.Reader接口的类型)扫描文本，根据format 参数指定的格式将成功读取的空白分隔的值保存进成功传递给本函数的参数(注意  换行也是认为空白分割的)。
	// 返回成功扫描的条目个数和遇到的任何错误。
	var (
		isOk bool
		str  string
	)
	len, err := fmt.Fscanf(os.Stdin, "%t %s", &isOk, &str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("读取到的长度", len) // 读取到的长度2
	// 假设终端输入 t string\n
	fmt.Println(isOk, str) // true string

	// fmt.Scan从标准输入中扫描用户输入的数据，将以空白符分隔的数据分别存入指定的参数中。
	// Scan 从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符。
	// 本函数返回成功扫描的数据个数和遇到的任何错误。如果读取的数据个数比提供的参数少，会返回一个错误报告原因
	var (
		name    string
		age     int
		married bool
	)
	len, err = fmt.Scan(&name, &age, &married)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("扫描结果 name:%s age:%d married:%t\t", name, age, married)

	// Scanf从标准输入扫描文本，根据format参数指定的格式去读取由空白符分隔的值保存到传递给本函数的参数中。
	// 1:name 2:10 3:t
	_, err = fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("扫描结果 name:%s age:%d married:%t", name, age, married)

	// fmt.Scanf和fmt.Scan不同于以简单的空格符作为输入数据的分隔符，fmt.Scanf为输入数据指定了具体的输入内容格式，只有按照格式输入数据才会被扫描并存入对应变量。

	// func Scanln(a ...interface{}) (n int, err error)
	// Scanln类似于Scan，它遇到换行才停止扫描。最后一个数据后面必须有换行或者到达结束位置。
	fmt.Scanln(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t", name, age, married)

	// Fscan系列
	/*	func Fscan(r io.Reader, a ...interface{}) (n int, err error)
		func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
		func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)*/
	// 这几个函数功能类似于fmt.Scan 、fmt.Scanf、fmt.Scanln三个函数，只不过它们不是从标准输入中读取数据而是从io.Reader中读取数据。

	// Sscan系列
	/*	func Sscan(str string, a ...interface{}) (n int, err error)
		func Sscanln(str string, a ...interface{}) (n int, err error)
		func Sscanf(str string, format string, a ...interface{}) (n int, err error)*/

	// 这几个函数功能类似于fmt.Scan 、fmt.Scanf、fmt.Scanln三个函数，只不过它们不是从标准输入中读取数据而是从指定字符串中读取数据。

}
