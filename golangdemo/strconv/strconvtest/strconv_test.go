package strconvtest

import (
	"fmt"
	"strconv"
	"testing"
)

// string -> int
func TestStringToInt(t *testing.T) {
	fmt.Println(strconv.FormatInt(int64(10), 10))
	fmt.Println(strconv.Itoa(10))

	numStr := "999"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", num, num)
	}
	num2, err2 := strconv.ParseInt(numStr, 0, 0)
	if err2 != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", num2, num2)
	}
	fmt.Println("String转换int:", num2)
	/**
	另外还可以用：
	func ParseInt(s string, base int, bitSize int) (i int64, err error)
	或
	func ParseUint(s string, base int, bitSize int) (n uint64, err error)
	base指定进制（2到36），如果base为0，则会从字符串前置判断，”0x”是16进制，”0”是8进制，否则是10进制；

	bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；
	*/
}

// int -> string
func TestIntToString(t *testing.T) {
	fmt.Println(strconv.FormatInt(int64(10), 10))
	fmt.Println(strconv.Itoa(10))

	num := 200
	numStr := strconv.Itoa(num)
	fmt.Printf("type:%T value:%#v\n", numStr, numStr)
	numStr2 := strconv.FormatInt(int64(num), 10)
	fmt.Printf("type:%T value:%#v\n", numStr2, numStr2)
}

// string->bool
func TestStringToBool(t *testing.T) {
	// 使用方法：func ParseBool(str string) (bool, error)
	// 当str为：1，t，T，TRUE，true，True中的一种时为真值
	// 当str为：0，f，F，FALSE，false，False 中的一种时为假值
	fmt.Println(strconv.ParseBool("1"))
	fmt.Println(strconv.ParseBool("t"))
	fmt.Println(strconv.ParseBool("T"))
	fmt.Println(strconv.ParseBool("TRUE"))
	fmt.Println(strconv.ParseBool("true"))
	fmt.Println(strconv.ParseBool("True"))
	fmt.Println(strconv.ParseBool("0"))
	fmt.Println(strconv.ParseBool("f"))
	fmt.Println(strconv.ParseBool("F"))
	fmt.Println(strconv.ParseBool("FALSE"))
	fmt.Println(strconv.ParseBool("false"))
	fmt.Println(strconv.ParseBool("False"))

	// 其他
	fmt.Println(strconv.ParseBool("trUe"))

}

// bool -> string
func TestBoolToString(t *testing.T) {
	fmt.Printf("type:%T value:%#v\n", strconv.FormatBool(true), strconv.FormatBool(true))
	fmt.Println()
	fmt.Printf("type:%T value:%#v\n", strconv.FormatBool(false), strconv.FormatBool(false))

}

// string->float
func TestStringToFloat(t *testing.T) {
	// 使用方法：func ParseFloat(s string, bitSize int) (f float64, err error)
	// bitSize：32或64 对应系统的位数
	strF := "250.5676878"
	str, err := strconv.ParseFloat(strF, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("type:%T value:%#v\n", str, str)
	str2, err2 := strconv.ParseFloat(strF, 32)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Printf("type:%T value:%#v\n", str2, str2)
}

// float ->string
func TestFloatToString(t *testing.T) {

	/*
		使用方法：func FormatFloat(f float64, fmt byte, prec, bitSize int) string
		bitSize表示f的来源类型（32：float32、64：float64），会据此进行舍入。
		fmt表示格式：’f’（-ddd.dddd）、’b’（-ddddp±ddd，指数为二进制）、’e’（-d.dddde±dd，十进制指数）、’E’（-d.ddddE±dd，十进制指数）、’g’（指数很大时用’e’格式，否则’f’格式）、’G’（指数很大时用’E’格式，否则’f’格式）。

		prec控制精度（排除指数部分）：对’f’、’e’、’E’，它表示小数点后的数字个数；对’g’、’G’，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f。
	*/

	num := 250.562362
	str := strconv.FormatFloat(num, 'f', 4, 64)
	fmt.Printf("type:%T value:%#v\n", str, str)

	// 以上类型转string的话，可以直接用fmt.Sprintf实现。
	num2 := 250.5623232
	str2 := fmt.Sprintf("%.2f", num2)
	fmt.Printf("type:%T value:%#v\n", str2, str2)
}

// int -> flaot
func TestIntToFloat(t *testing.T) {
	num := 100
	fmt.Printf("type:%T value:%#v\n", float64(num), float64(num))
	fmt.Printf("type:%T value:%#v\n", float32(num), float32(num))

}

//  flaot-> int
func TestFloatToInt(t *testing.T) {
	num := 250.562362
	fmt.Printf("type:%T value:%#v\n", int(num), int(num))
	fmt.Printf("type:%T value:%#v\n", int32(num), int32(num))
	fmt.Printf("type:%T value:%#v\n", int64(num), int64(num))

}
