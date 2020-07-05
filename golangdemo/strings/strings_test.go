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
