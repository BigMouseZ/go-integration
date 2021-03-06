package main

import "testing"

/*
builtin 包为Go的预声明标识符提供了文档。此处列出的条目其实并不在builtin 包中，对它们的描述只是为了让 godoc 给该语言的特殊标识符提供文档。

1、介绍builtin包
中文官方标准库给出的介绍：builtin包为Go预声明标识符提供了文档。
2、builtin内容

const (
    true = 0 == 0    // 无类型布尔值
    false = 0 != 0    // 无类型布尔值
)
true和false是两个无类型布尔值

const iota = 0  // 无类型整数值
iota是一个预定义的标识符，代表顺序按行增加的无符号整数。每个const声明单元（被括号括起来）相互独立，分别从0开始。

type bool bool
布尔类型

type byte byte
8位无符号整型。是uint8的别名，二者视为同一类型

type rune rune
32位有符号整型，int32的别名，二者视为同一类型。

type int int
至少32位的有符号整型，但和int32/rune并发同一类型。

type int8 int8
8位有符号整型，范围[-128,127]。

type int16 int16
16位有符号整型，范围[-32768,32767]。

type int32 int32
32位有符号整型，范围[-2147483648,2147483647]。

type int64 int64
64位有符号整型，范围[-9223372036854775808,9223372036854775807]。

type uint8 uint8
8位无符号整型，范围[0,65535]

type uint32 uint32
32位无符号整型，范围[0,4294967295]

type uint64 uint64
64位无符号整型，范围[0,18446744073709551615]

type float32 float32
所有IEEE-754 32位浮点数的集合，12位有效数字

type float64 float64
所有IEEE-754 64位浮点数的集合，16位有效数字

type complex64 complex64
具有float32类型实部和虚部的复数类型。

type complex128 complex128
具有float64类型实部和虚部的复数类型。

type uintptr uintptr
可以保存任意指针的位模式的整数类型

type string string
8位byte序列构成的字符串，约定但不必须是utf-8编码的文本。字符串可以为空但不能是nil，其值不可变。

type error interface {
    Error() string
}
内建error接口类型是约定用于表示错误信息，nil值表示无错误。

type Type int
在本文档中代表任意一个类型，但同一个声明里只代表同一个类型。

var nil Type  // Type必须是指针、通道、函数、接口、映射或映射
nil是预定义的标识符，标识指针、通道、函数、接口、映射或切片的零值。

type Type1 int
在本文档中代表任意一个类型，但同一个声明里只代表同一个类型。用于代表和Type不同的另一类型。

type IntegerType int
在本文档中代表一个有符号或无符合号的整数类型。

type FloatType float32
在本文档中代表一个浮点数类型。

type ComplexType complex64
在本文档中代表一个复数类型。

func real(c ComplexType) FloatType
返回复数c的实部

func imag(c ComplexType) FloatType
返回复数c的虚部。

func complex(r, i FloatType) ComplexType
使用实部r和虚部r生成一个复数。

func new(Type) *Type
内建函数new分配内存。其第一个实参为类型，而非值。其返回值为指向该类型的新分配的零值得指针。

func make(Type, size IntegerType) Type
内建函数make分配并初始化一个类型为切片、map或通道得对象。其第一个实参为类型，而非值。make得返回类型与其参数相同，而非指向它得指针。其具体结果取决于具体得类型。

切片：size指定了其长度。该切片的容量等于其长度。切片支持第二个整数实参可用来指定不同的容量；但是第二个整数值必须不小于第一个整数值。
map：初始分配的创建取决于size，但产生的映射长度为0。size可以省略，实际上，对于map来说，size不管是多少，为map分配长度都为0。
通道：通道的缓存根据指定的缓存容量初始化。若size为零或被省略，该信道即为无缓存的。
func cap(v Type) int

内建函数cap返回v的容量
func len(v Type) int

内建函数len返回v的长度
func append(slice []Type, elem ...Type) []Type

内建函数append将元素追加到切片的末尾。若它有足够的容量，其目标就会重新切片以容纳新的元素。否则，就会分配一个新的基本数组。append返回更新后的切片，因此必须存储追加后的结果。
func copy(dst, src []Type) int

内建函数copy将元素从来源切片复制到目标切片中，也能将字符串复制到字节切片中。copy返回被复制的元素数量，它会是len(src)和len(dst)中较小的那个。来源和目标的底层内存可以重叠。
func delete(m map[Type]Type1, key Type)

内建函数delete按照指定的键将元素从映射中删除。若m为nil或无此元素，delete不进行操作。
func close(c chan<- Type)

内建函数close关闭信道，该通道必须为双向的或只发送的。它应当只有发送者执行，而不应该由接收者执行，其效果是在最后发送的值被接收后停止该通道。在最后的值从已关闭的信道中被接收后，任何对其的接收操作都会无阻塞的成功。对于已关闭的信道，语句：x, ok := <- c，会将ok置为false。
func panic(v interface{})

内建函数panic停止当前Go程的正常执行。当函数F调用panic时，F的正常执行就会立刻停止。F中defer的所有函数先入后出执行后，F返回给调用者G。G如同F一样行动，层层返回，直到该GO程中所有函数都按相反的顺序停止执行。之后，程序被终止，而错误情况会被报告，包括引发该恐慌的实参值，此终止序列称为恐慌过程。
func recover() interface{}

内建函数recover允许程序管理恐慌过程中的Go程。在defer的函数中，执行recover调用会取回传至panic调用的错误值，恢复正常执行，停止恐慌过程。若recover在defer的函数之外被调用，它将不会停止恐慌过程序列。在此情况下，或当该Go程不在恐慌过程中时，或提供给panic的实参为nil时，recover就会返回nil。
func print(args ...Type)

内建函数print以特有的方法格式化参数并将结果写入标准错误，用于自举和调试。
func println(args ...Type)

println类似print，但会在参数输出之间添加空格，输出结束后换行。
*/
func Test(t *testing.T) {
	println("内建函数")

}
