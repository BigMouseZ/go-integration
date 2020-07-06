package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestReplace(t *testing.T) {
	// Replace 返回 s 的副本，并将副本中的 old 字符串替换为 new 字符串
	// 替换次数为 n 次，如果 n 为 -1，则全部替换
	// 如果 old 为空，则在副本的每个字符之间都插入一个 new
	// func Replace(s, old, new string, n int) string
	s := "Hello 世界！"
	s = strings.Replace(s, " ", ",", -1)
	fmt.Println(s)
	s = strings.Replace(s, "", "|", -1)
	fmt.Println(s)
}

func TestEqualFold(t *testing.T) {
	// EqualFold 判断 s 和 t 是否相等。忽略大小写，同时它还会对特殊字符进行转换
	// 比如将“ϕ”转换为“Φ”、将“Ǆ”转换为“ǅ”等，然后再进行比较
	// func EqualFold(s, t string) bool
	s1 := "Hello 世界! ϕ Ǆ"
	s2 := "hello 世界! Φ ǅ"
	b := strings.EqualFold(s1, s2)
	fmt.Printf("%v\n", b) // true

}
func TestReader(t *testing.T) {

	// Reader 结构通过读取字符串，实现了 io.Reader，io.ReaderAt，
	// io.Seeker，io.WriterTo，io.ByteScanner，io.RuneScanner 接口
	/*type Reader struct {
		s string // 要读取的字符串
		i int // 当前读取的索引位置，从 i 处开始读取数据
		prevRune int // 读取的前一个字符的索引位置，小于 0 表示之前未读取字符
	}*/

	// 通过字符串 s 创建 strings.Reader 对象
	// 这个函数类似于 bytes.NewBufferString
	// 但比 bytes.NewBufferString 更有效率，而且只读
	// func NewReader(s string) *Reader { return &Reader{s, 0, -1} }

	// Len 返回 r.i 之后的所有数据的字节长度
	// func (r *Reader) Len() int
	s := "Hello 世界!"
	// 创建 Reader
	r := strings.NewReader(s)
	// 获取字符串的编码长度
	fmt.Println(r.Len()) // 13

	// Read 将 r.i 之后的所有数据写入到 b 中（如果 b 的容量足够大）
	// 返回读取的字节数和读取过程中遇到的错误
	// 如果无可读数据，则返回 io.EOF
	// func (r *Reader) Read(b []byte) (n int, err error)
	s = "Hello World!"
	// 创建 Reader
	r = strings.NewReader(s)
	// 创建长度为 5 个字节的缓冲区
	b := make([]byte, 5)
	// 循环读取 r 中的字符串
	for n, _ := r.Read(b); n > 0; n, _ = r.Read(b) {
		fmt.Printf("%q, ", b[:n]) // "Hello", " Worl", "d!"
	}

	// ReadAt 将 off 之后的所有数据写入到 b 中（如果 b 的容量足够大）
	// 返回读取的字节数和读取过程中遇到的错误
	// 如果无可读数据，则返回 io.EOF
	// 如果数据被一次性读取完毕，则返回 io.EOF
	// func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)

	s = "Hello World!"
	// 创建 Reader
	r = strings.NewReader(s)
	// 创建长度为 5 个字节的缓冲区
	b = make([]byte, 5)
	// 读取 r 中指定位置的字符串
	n, _ := r.ReadAt(b, 0)
	fmt.Printf("%q\n", b[:n]) // "Hello"
	// 读取 r 中指定位置的字符串
	n, _ = r.ReadAt(b, 6)
	fmt.Printf("%q\n", b[:n]) // "World"

	// ReadByte 将 r.i 之后的一个字节写入到返回值 b 中
	// 返回读取的字节和读取过程中遇到的错误
	// 如果无可读数据，则返回 io.EOF
	// func (r *Reader) ReadByte() (b byte, err error)

	s = "Hello World!"
	// 创建 Reader
	r = strings.NewReader(s)
	// 读取 r 中的一个字节
	for i := 0; i < 3; i++ {
		b, _ := r.ReadByte()
		fmt.Printf("%q, ", b) // 'H', 'e', 'l',
	}

	// UnreadByte 撤消前一次的 ReadByte 操作，即 r.i--
	// func (r *Reader) UnreadByte() error

	s = "Hello World!"
	// 创建 Reader
	r = strings.NewReader(s)
	// 读取 r 中的一个字节
	for i := 0; i < 3; i++ {
		b, _ := r.ReadByte()
		fmt.Printf("%q, ", b) // 'H', 'H', 'H',
		r.UnreadByte()        // 撤消前一次的字节读取操作
	}

	// ReadRune 将 r.i 之后的一个字符写入到返回值 ch 中
	// ch： 读取的字符
	// size：ch 的编码长度
	// err： 读取过程中遇到的错误
	// 如果无可读数据，则返回 io.EOF
	// 如果 r.i 之后不是一个合法的 UTF-8 字符编码，则返回 utf8.RuneError 字符
	// func (r *Reader) ReadRune() (ch rune, size int, err error)

	s = "你好 世界！"
	// 创建 Reader
	r = strings.NewReader(s)
	// 读取 r 中的一个字符
	for i := 0; i < 5; i++ {
		b, n, _ := r.ReadRune()
		fmt.Printf(`"%c:%v", `, b, n)
		// "你:3", "好:3", " :1", "世:3", "界:3",
	}

	// 撤消前一次的 ReadRune 操作
	// func (r *Reader) UnreadRune() error
	s = "你好 世界！"
	// 创建 Reader
	r = strings.NewReader(s)
	// 读取 r 中的一个字符
	for i := 0; i < 5; i++ {
		b, _, _ := r.ReadRune()
		fmt.Printf("%q, ", b)
		// '你', '你', '你', '你', '你',
		r.UnreadRune() // 撤消前一次的字符读取操作
	}

}
func TestSeek(t *testing.T) {
	// Seek 用来移动 r 中的索引位置
	// offset：要移动的偏移量，负数表示反向移动
	// whence：从那里开始移动，0：起始位置，1：当前位置，2：结尾位置
	// 如果 whence 不是 0、1、2，则返回错误信息
	// 如果目标索引位置超出字符串范围，则返回错误信息
	// 目标索引位置不能超出 1 << 31，否则返回错误信息
	// func (r *Reader) Seek(offset int64, whence int) (int64, error)
	s := "Hello World!"
	// 创建 Reader
	r := strings.NewReader(s)
	// 创建读取缓冲区
	b := make([]byte, 5)
	// 读取 r 中指定位置的内容
	r.Seek(6, 0) // 移动索引位置到第 7 个字节
	r.Read(b)    // 开始读取
	fmt.Printf("%q\n", b)
	r.Seek(-5, 1) // 将索引位置移回去
	r.Read(b)     // 继续读取
	fmt.Printf("%q\n", b)
}
func TestWriteTo(t *testing.T) {
	// WriteTo 将 r.i 之后的数据写入接口 w 中
	// func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
	s := "Hello World!"
	// 创建 Reader
	r := strings.NewReader(s)
	// 创建 bytes.Buffer 对象，它实现了 io.Writer 接口
	buf := bytes.NewBuffer(nil)
	// 将 r 中的数据写入 buf 中
	r.WriteTo(buf)
	fmt.Printf("%q\n", buf) // "Hello World!"
}
func TestReplacer(t *testing.T) {
	// Replacer 根据一个替换列表执行替换操作
	/*type Replacer struct {
		Replace(s string) string
		WriteString(w io.Writer, s string) (n int, err error)
	}*/

	// NewReplacer 通过“替换列表”创建一个 Replacer 对象。
	// 按照“替换列表”中的顺序进行替换，只替换非重叠部分。
	// 如果参数的个数不是偶数，则抛出异常。
	// 如果在“替换列表”中有相同的“查找项”，则后面重复的“查找项”会被忽略
	// func NewReplacer(oldnew ...string) *Replacer

	// Replace 返回对 s 进行“查找和替换”后的结果
	// Replace 使用的是 Boyer-Moore 算法，速度很快
	// func (r *Replacer) Replace(s string) string
	srp := strings.NewReplacer("Hello", "你好", "World", "世界", "!", "！")
	s := "Hello World!Hello World!hello world!"
	rst := srp.Replace(s)
	fmt.Print(rst) // 你好 世界！你好 世界！hello world！

	wl := []string{"Hello", "Hi", "Hello", "你好"}
	srp = strings.NewReplacer(wl...)
	s = "Hello World! Hello World! hello world!"
	rst = srp.Replace(s)
	fmt.Print(rst) // Hi World! Hi World! hello world!

}
func TestWriteString(t *testing.T) {
	// WriteString 对 s 进行“查找和替换”，然后将结果写入 w 中
	// func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)
	wl := []string{"Hello", "你好", "World", "世界", "!", "！"}
	srp := strings.NewReplacer(wl...)
	s := "Hello World!Hello World!hello world!"
	srp.WriteString(os.Stdout, s)
	// 你好 世界！你好 世界！hello world！
}
