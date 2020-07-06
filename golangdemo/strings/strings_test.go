package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestContain(t *testing.T) {
	// Count 计算字符串 sep 在 s 中的非重叠个数
	// 如果 sep 为空字符串，则返回 s 中的字符(非字节)个数 + 1
	// 使用 Rabin-Karp 算法实现
	s := "Hello,世界!!!!!"
	n := strings.Count(s, "!")
	fmt.Println(n) // 5
	n = strings.Count(s, "!!")
	fmt.Println(n) // 2
	n = strings.Count(s, "")
	fmt.Println(n) // 14
	// Contains 判断字符串 s 中是否包含子串 substr
	// 如果 substr 为空，则返回 true
	// func Contains(s, substr string) bool
	b := strings.Contains(s, "!!")
	fmt.Println(b) // true
	b = strings.Contains(s, "!?")
	fmt.Println(b) // false
	b = strings.Contains(s, "")
	fmt.Println(b) // true

	// ContainsAny 判断字符串 s 中是否包含 chars 中的任何一个字符
	// 如果 chars 为空，则返回 false
	// func ContainsAny(s, chars string) bool
	b = strings.ContainsAny(s, "abc")
	fmt.Println(b) // false
	b = strings.ContainsAny(s, "def")
	fmt.Println(b) // true
	b = strings.ContainsAny(s, "")
	fmt.Println(b) // false
	b = strings.Contains(s, "")
	fmt.Println(b) // true

	// ContainsRune 判断字符串 s 中是否包含字符 r
	// func ContainsRune(s string, r rune) bool
	b = strings.ContainsRune(s, '\n')
	fmt.Println(b) // false
	b = strings.ContainsRune(s, '界')
	fmt.Println(b) // true
	b = strings.ContainsRune(s, 0)
	fmt.Println(b) // false

}

func TestIndex(t *testing.T) {

	// Index 返回子串 sep 在字符串 s 中第一次出现的位置 ,中文字符3字节
	// 如果找不到，则返回 -1，如果 sep 为空，则返回 0。
	// 使用 Rabin-Karp 算法实现
	// func Index(s, sep string) int
	s := "Hello,世界! Hello!"
	i := strings.Index(s, "h")
	fmt.Println(i) // -1
	i = strings.Index(s, "!")
	fmt.Println(i) // 12
	i = strings.Index(s, "")
	fmt.Println(i) // 0
	i = strings.Index(s, "界")
	fmt.Println(i) // 9

	// LastIndex 返回子串 sep 在字符串 s 中最后一次出现的位置
	// 如果找不到，则返回 -1，如果 sep 为空，则返回字符串的长度
	// 使用朴素字符串比较算法实现
	// func LastIndex(s, sep string) int

	i = strings.LastIndex(s, "h")
	fmt.Println(i) // -1
	i = strings.LastIndex(s, "H")
	fmt.Println(i) // 14
	i = strings.LastIndex(s, "")
	fmt.Println(i) // 20
	// IndexRune 返回字符 r 在字符串 s 中第一次出现的位置
	// 如果找不到，则返回 -1
	// func IndexRune(s string, r rune) int
	i = strings.IndexRune(s, '\n')
	fmt.Println(i) // -1
	i = strings.IndexRune(s, '界')
	fmt.Println(i) // 9
	i = strings.IndexRune(s, 0)
	fmt.Println(i) // -1

	// IndexAny 返回字符串 chars 中的任何一个字符在字符串 s 中第一次出现的位置
	// 如果找不到，则返回 -1，如果 chars 为空，则返回 -1
	// func IndexAny(s, chars string) int
	i = strings.IndexAny(s, "abc")
	fmt.Println(i) // -1
	i = strings.IndexAny(s, "dof")
	fmt.Println(i) // 1
	i = strings.IndexAny(s, "")
	fmt.Println(i) // -1
}

func TestSplit(t *testing.T) {
	// SplitN 以 sep 为分隔符，将 s 切分成多个子串，结果中不包含 sep 本身
	// 如果 sep 为空，则将 s 切分成 Unicode 字符列表。
	// 如果 s 中没有 sep 子串，则将整个 s 作为 []string 的第一个元素返回
	// 参数 n 表示最多切分出几个子串，超出的部分将不再切分。
	// 如果 n 为 0，则返回 nil，如果 n 小于 0，则不限制切分个数，全部切分
	// func SplitN(s, sep string, n int) []string
	s := "Hello, 世界! Hello!"
	ss := strings.SplitN(s, " ", 2)
	fmt.Printf("%q\n", ss) // ["Hello," "世界! Hello!"]
	ss = strings.SplitN(s, " ", -1)
	fmt.Printf("%q\n", ss) // ["Hello," "世界!" "Hello!"]
	ss = strings.SplitN(s, "", 3)
	fmt.Printf("%q\n", ss) // ["H" "e" "llo, 世界! Hello!"]

	// SplitAfterN 以 sep 为分隔符，将 s 切分成多个子串，结果中包含 sep 本身
	// 如果 sep 为空，则将 s 切分成 Unicode 字符列表。
	// 如果 s 中没有 sep 子串，则将整个 s 作为 []string 的第一个元素返回
	// 参数 n 表示最多切分出几个子串，超出的部分将不再切分。
	// 如果 n 为 0，则返回 nil，如果 n 小于 0，则不限制切分个数，全部切分
	// func SplitAfterN(s, sep string, n int) []string

	ss = strings.SplitAfterN(s, " ", 2)
	fmt.Printf("%q\n", ss) // ["Hello, " "世界! Hello!"]
	ss = strings.SplitAfterN(s, " ", -1)
	fmt.Printf("%q\n", ss) // ["Hello, " "世界! " "Hello!"]
	ss = strings.SplitAfterN(s, "", 3)
	fmt.Printf("%q\n", ss) // ["H" "e" "llo, 世界! Hello!"]

	// Split 以 sep 为分隔符，将 s 切分成多个子切片，结果中不包含 sep 本身
	// 如果 sep 为空，则将 s 切分成 Unicode 字符列表。
	// 如果 s 中没有 sep 子串，则将整个 s 作为 []string 的第一个元素返回
	// func Split(s, sep string) []string
	ss = strings.Split(s, " ")
	fmt.Printf("%q\n", ss) // ["Hello," "世界!" "Hello!"]
	ss = strings.Split(s, ", ")
	fmt.Printf("%q\n", ss) // ["Hello" "世界! Hello!"]
	ss = strings.Split(s, "")
	fmt.Printf("%q\n", ss) // 单个字符列表

	// SplitAfter 以 sep 为分隔符，将 s 切分成多个子切片，结果中包含 sep 本身
	// 如果 sep 为空，则将 s 切分成 Unicode 字符列表。
	// 如果 s 中没有 sep 子串，则将整个 s 作为 []string 的第一个元素返回
	// func SplitAfter(s, sep string) []string
	ss = strings.SplitAfter(s, " ")
	fmt.Printf("%q\n", ss) // ["Hello, " "世界! " "Hello!"]
	ss = strings.SplitAfter(s, ", ")
	fmt.Printf("%q\n", ss) // ["Hello, " "世界! Hello!"]
	ss = strings.SplitAfter(s, "")
	fmt.Printf("%q\n", ss) // 单个字符列表

	// Fields 以连续的空白字符为分隔符，将 s 切分成多个子串，结果中不包含空白字符本身
	// 空白字符有：\t, \n, \v, \f, \r, ' ', U+0085 (NEL), U+00A0 (NBSP)
	// 如果 s 中只包含空白字符，则返回一个空列表
	// func Fields(s string) []string
	ss = strings.Fields(s)
	fmt.Printf("%q\n", ss) // ["Hello," "世界!" "Hello!"]

	// FieldsFunc 以一个或多个满足 f(rune) 的字符为分隔符，
	// 将 s 切分成多个子串，结果中不包含分隔符本身。
	// 如果 s 中没有满足 f(rune) 的字符，则返回一个空列表。
	// func FieldsFunc(s string, f func(rune) bool) []string
	ss = strings.FieldsFunc(s, isSlash)
	fmt.Printf("%q\n", ss) // ["C:" "Windows" "System32" "FileName"]
}
func isSlash(r rune) bool {
	return r == '\\' || r == '/'
}

func TestJion(t *testing.T) {

	// Join 将 a 中的子串连接成一个单独的字符串，子串之间用 sep 分隔
	// func Join(a []string, sep string) string
	ss := []string{"Monday", "Tuesday", "Wednesday"}
	s := strings.Join(ss, "|")
	fmt.Println(s)

}
func TestFix(t *testing.T) {
	// HasPrefix 判断字符串 s 是否以 prefix 开头
	// func HasPrefix(s, prefix string) bool
	s := "Hello 世界!"
	b := strings.HasPrefix(s, "hello")
	fmt.Println(b) // false
	b = strings.HasPrefix(s, "Hello")
	fmt.Println(b) // true
	// HasSuffix 判断字符串 s 是否以 prefix 结尾
	// func HasSuffix(s, suffix string) bool
	b = strings.HasSuffix(s, "世界")
	fmt.Println(b) // false
	b = strings.HasSuffix(s, "世界!")
	fmt.Println(b) // true
}

func TestMap(t *testing.T) {
	// Map 将 s 中满足 mapping(rune) 的字符替换为 mapping(rune) 的返回值。
	// 如果 mapping(rune) 返回负数，则相应的字符将被删除。
	// func Map(mapping func(rune) rune, s string) string
	s := "C:\\Windows\\System32\\FileName"
	ms := strings.Map(Slash, s)
	fmt.Printf("%q\n", ms) // "C:/Windows/System32/FileName"
}
func Slash(r rune) rune {
	if r == '\\' {
		return '/'
	}
	return r
}

func TestRepeat(t *testing.T) {
	// Repeat 将 count 个字符串 s 连接成一个新的字符串
	// func Repeat(s string, count int) string
	s := "Hello!"
	rs := strings.Repeat(s, 3)
	fmt.Printf("%q\n", rs) // "Hello!Hello!Hello!"
}

func TestUpper(t *testing.T) {
	// ToUpper 将 s 中的所有字符修改为其大写格式
	// 对于非 ASCII 字符，它的大写格式需要查表转换
	// func ToUpper(s string) string

	// ToLower 将 s 中的所有字符修改为其小写格式
	// 对于非 ASCII 字符，它的小写格式需要查表转换
	// func ToLower(s string) string

	// ToTitle 将 s 中的所有字符修改为其 Title 格式
	// 大部分字符的 Title 格式就是其 Upper 格式
	// 只有少数字符的 Title 格式是特殊字符
	// 这里的 ToTitle 主要给 Title 函数调用
	// func ToTitle(s string) string
	s := "heLLo worLd Ａｂｃ"
	us := strings.ToUpper(s)
	ls := strings.ToLower(s)
	ts := strings.ToTitle(s)
	fmt.Printf("%q\n", us) // "HELLO WORLD ＡＢＣ"
	fmt.Printf("%q\n", ls) // "hello world ａｂｃ"
	fmt.Printf("%q\n", ts) // "HELLO WORLD ＡＢＣ"

	// Title 将 s 中的所有单词的首字母修改为其 Title 格式
	// BUG: Title 规则不能正确处理 Unicode 标点符号
	// func Title(s string) string
	ts = strings.Title(s)
	fmt.Printf("%q\n", ts) // "HeLLo WorLd"
}

func TestTrimFunc(t *testing.T) {

	// TrimLeftFunc 将删除 s 头部连续的满足 f(rune) 的字符
	// func TrimLeftFunc(s string, f func(rune) bool) string

	// TrimRightFunc 将删除 s 尾部连续的满足 f(rune) 的字符
	// func TrimRightFunc(s string, f func(rune) bool) string
	s := "\\\\HostName\\C\\Windows\\"
	ts := strings.TrimLeftFunc(s, isSlash)
	fmt.Printf("%q\n", ts) // "HostName\\C\\Windows\\"
	ts = strings.TrimRightFunc(s, isSlash)
	fmt.Printf("%q\n", ts) // "\\\\HostName\\C\\Windows"

	// TrimFunc 将删除 s 首尾连续的满足 f(rune) 的字符
	// func TrimFunc(s string, f func(rune) bool) string
	ts = strings.TrimFunc(s, isSlash)
	fmt.Printf("%q\n", ts) // "HostName\\C\\Windows"
}

func TestIndexFunc(t *testing.T) {
	// 返回 s 中第一个满足 f(rune) 的字符的字节位置。
	// 如果没有满足 f(rune) 的字符，则返回 -1
	// func IndexFunc(s string, f func(rune) bool) int
	s := "C:\\Windows\\System32"
	i := strings.IndexFunc(s, isSlash)
	fmt.Printf("%v\n", i) // 2

	// 返回 s 中最后一个满足 f(rune) 的字符的字节位置。
	// 如果没有满足 f(rune) 的字符，则返回 -1
	// func LastIndexFunc(s string, f func(rune) bool) int
	i = strings.LastIndexFunc(s, isSlash)
	fmt.Printf("%v\n", i) // 10
}

func TestTrim(t *testing.T) {
	// Trim 将删除 s 首尾连续的包含在 cutset 中的字符
	// func Trim(s string, cutset string) string
	s := " Hello 世界! "
	ts := strings.Trim(s, " Helo!")
	fmt.Printf("%q\n", ts) // "世界"

	// TrimLeft 将删除 s 头部连续的包含在 cutset 中的字符
	// func TrimLeft(s string, cutset string) string
	ts = strings.TrimLeft(s, " Helo")
	fmt.Printf("%q\n", ts) // "世界! "

	// TrimRight 将删除 s 尾部连续的包含在 cutset 中的字符
	// func TrimRight(s string, cutset string) string
	ts = strings.TrimRight(s, " 世界!")
	fmt.Printf("%q\n", ts) // " Hello"

	// TrimSpace 将删除 s 首尾连续的的空白字符
	// func TrimSpace(s string) string
	ts = strings.TrimSpace(s)
	fmt.Printf("%q\n", ts) // "Hello 世界!"

	// TrimPrefix 删除 s 头部的 prefix 字符串
	// 如果 s 不是以 prefix 开头，则返回原始 s
	// func TrimPrefix(s, prefix string) string
	ts = strings.TrimPrefix(s, "Hello")
	fmt.Printf("%q\n", ts) // " 世界"

	// TrimSuffix 删除 s 尾部的 suffix 字符串
	// 如果 s 不是以 suffix 结尾，则返回原始 s
	// func TrimSuffix(s, suffix string) string
	s = "Hello 世界!!!!!"
	ts = strings.TrimSuffix(s, "!!!!")
	fmt.Printf("%q\n", ts) // " 世界"

	/*	注：TrimSuffix只是去掉s字符串结尾的suffix字符串，只是去掉１次，
		而TrimRight是一直去掉s字符串右边的字符串，只要有响应的字符串就去掉，是一个多次的过程
		，这也是二者的本质区别．*/
}
