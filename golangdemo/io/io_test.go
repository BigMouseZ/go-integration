package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

// 定义一个 Ustr 类型（以 string 为基类型）
type Ustr string

// 实现 Ustr 类型的 Read 方法
func (s Ustr) Read(p []byte) (n int, err error) {
	i, ls, lp := 0, len(s), len(p)
	for ; i < ls && i < lp; i++ {
		// 将小写字母转换为大写字母，然后写入 p 中
		if s[i] >= 'a' && s[i] <= 'z' {
			p[i] = s[i] + 'A' - 'a'
		} else {
			p[i] = s[i]
		}
	}
	// 根据读取的字节数设置返回值
	switch i {
	case lp:
		return i, nil
	case ls:
		return i, io.EOF
	default:
		return i, errors.New("Read Fail")
	}
}
func TestReader(t *testing.T) {
	/*type Reader interface {
		Read(p []byte) (n int, err error)
	}*/
	// 从上面的定义可以看出，Reader 的标准很简单，只要某个对象实现了 Read 方法，这个对象就符合了 Reader 的标准，就可以被 ReadFull 读取。
	us := Ustr("Hello World!")     // 创建 Ustr 对象 us
	buf := make([]byte, 32)        // 创建缓冲区 buf
	n, err := io.ReadFull(us, buf) // 将 us 中的数据读取到 buf 中
	fmt.Printf("%s\n", buf)        // 显示 buf 中的内容
	// HELLO WORLD!
	fmt.Println(n, err) // 显示返回值
	// 12 unexpected EOF
}

// Reader 接口封装了基本的 Read 方法
// Read 方法用于将对象的数据流读入到 p 中
// 直到“p 被装满”或者“数据被读完”或者“数据读取出错”为止
// 返回读取的字节数 n (0 <= n <= len(p)) 和读取过程中遇到的任何错误
// 如果“p 被装满”，则 err 应该返回 nil
// 如果“数据被读完”，则 err 应该返回 EOF
// 如果“数据读取出错”，则 err 应该返回相应的错误信息
// Reader 用来输出数据
// type Reader interface {
// 	Read(p []byte) (n int, err error)
// }

// Writer 接口封装了基本的 Write 方法
// Write 方法用于将 p 中的数据写入到对象的数据流中
// 返回写入的字节数 n (0 <= n <= len(p)) 和写入时遇到的任何错误
// 如果 p 中的数据全部被写入，则 err 应该返回 nil
// 如果 p 中的数据无法被全部写入，则 err 应该返回相应的错误信息
// Writer 用来存入数据
// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }

// Closer 接口封装了基本的 Close 方法
// Close 一般用于关闭文件，关闭连接，关闭数据库等
// Close 用来关闭数据
// type Closer interface {
// 	Close() error
// }

// Seeker 接口封装了基本的 Seek 方法
// Seek 设置下一次读写操作的指针位置，每次的读写操作都是从指针位置开始的
// whence 的含义：
// 如果 whence 为 0：表示从数据的开头开始移动指针
// 如果 whence 为 1：表示从数据的当前指针位置开始移动指针
// 如果 whence 为 2：表示从数据的尾部开始移动指针
// offset 是指针移动的偏移量
// 返回移动后的指针位置和移动过程中遇到的任何错误
// Seeker 用来移动数据的读写指针

// type Seeker interface {
// 	Seek(offset int64, whence int) (ret int64, err error)
// }

// ReaderFrom 接口封装了基本的 ReadFrom 方法
// ReadFrom 从 r 中读取数据到对象的数据流中
// 直到 r 返回 EOF 或 r 出现读取错误为止
// 返回值 n 是读取的字节数
// 返回值 err 就是 r 的返回值 err
// ReadFrom 用来读出 r 中的数据
// type ReaderFrom interface {
// 	ReadFrom(r Reader) (n int64, err error)
// }

// WriterTo 接口封装了基本的 WriteTo 方法
// WriterTo 将对象的数据流写入到 w 中
// 直到对象的数据流全部写入完毕或遇到写入错误为止
// 返回值 n 是写入的字节数
// 返回值 err 就是 w 的返回值 err
// WriteTo 用来将数据写入 w 中
// type WriterTo interface {
// 	WriteTo(w Writer) (n int64, err error)
// }

// ReaderAt 接口封装了基本的 ReadAt 方法
// ReadAt 从对象数据流的 off 处读出数据到 p 中
// 忽略数据的读写指针，从数据的起始位置偏移 off 处开始读取
// 如果对象的数据流只有部分可用，不足以填满 p
// 则 ReadAt 将等待所有数据可用之后，继续向 p 中写入
// 直到将 p 填满后再返回
// 在这点上 ReadAt 要比 Read 更严格
// 返回读取的字节数 n 和读取时遇到的错误
// 如果 n < len(p)，则需要返回一个 err 值来说明
// 为什么没有将 p 填满（比如 EOF）
// 如果 n = len(p)，而且对象的数据没有全部读完，则
// err 将返回 nil
// 如果 n = len(p)，而且对象的数据刚好全部读完，则
// err 将返回 EOF 或者 nil（不确定）
// type ReaderAt interface {
// 	ReadAt(p []byte, off int64) (n int, err error)
// }

// WriterAt 接口封装了基本的 WriteAt 方法
// WriteAt 将 p 中的数据写入到对象数据流的 off 处
// 忽略数据的读写指针，从数据的起始位置偏移 off 处开始写入
// 返回写入的字节数和写入时遇到的错误
// 如果 n < len(p)，则必须返回一个 err 值来说明
// 为什么没有将 p 完全写入
// type WriterAt interface {
// 	WriteAt(p []byte, off int64) (n int, err error)
// }

// ByteReader 接口封装了基本的 ReadByte 方法
// ReadByte 从对象的数据流中读取一个字节到 c 中
// 如果对象的数据流中没有可读数据，则返回一个错误信息
//
// type ByteReader interface {
// 	ReadByte() (c byte, err error)
// }

// ByteScanner 在 ByteReader 的基础上增加了一个 UnreadByte 方法
// UnreadByte 用于撤消最后一次的 ReadByte 操作
// 即将对象的读写指针移到上次 ReadByte 之前的位置
// 如果上一次的操作不是 ReadByte，则 UnreadByte 返回一个错误信息
// type ByteScanner interface {
// 	ByteReader
// 	UnreadByte() error
// }

// ByteWriter 接口封装了基本的 WriteByte 方法
// WriteByte 将一个字节 c 写入到对象的数据流中
// 返回写入过程中遇到的任何错误

// type ByteWriter interface {
// 	WriteByte(c byte) error
// }

// RuneReader 接口封装了基本的 ReadRune 方法
// ReadRune 从对象的数据流中读取一个字符到 r 中
// 如果对象的数据流中没有可读数据，则返回一个错误信息
// type RuneReader interface {
// 	ReadRune() (r rune, size int, err error)
// }

// RuneScanner 在 RuneReader 的基础上增加了一个 UnreadRune 方法
// UnreadRune 用于撤消最后一次的 ReadRune 操作
// 即将对象的读写指针移到上次 ReadRune 之前的位置
// 如果上一次的操作不是 ReadRune，则 UnreadRune 返回一个错误信息
// type RuneScanner interface {
// 	RuneReader
// 	UnreadRune() error
// }

func TestNewBuffer(t *testing.T) {
	// bytes.NewBuffer 实现了很多基本的接口，可以通过 bytes 包学习接口的实现
	bb := bytes.NewBuffer([]byte("Hello World!"))
	b := make([]byte, 32)

	bb.Read(b)
	fmt.Printf("%s\n", b) // Hello World!

	bb.WriteString("New Data!\n")
	bb.WriteTo(os.Stdout) // New Data!

	bb.WriteString("Third Data!")
	bb.ReadByte()
	fmt.Println(bb.String()) // hird Data!
	bb.UnreadByte()
	fmt.Println(bb.String()) // Third Data!
}
func TestWriteString(t *testing.T) {
	// WriteString 将字符串 s 写入到 w 中
	// 返回写入的字节数和写入过程中遇到的任何错误
	// 如果 w 实现了 WriteString 方法
	// 则调用 w 的 WriteString 方法将 s 写入 w 中
	// 否则，将 s 转换为 []byte
	// 然后调用 w.Write 方法将数据写入 w 中
	// func WriteString(w Writer, s string) (n int, err error)
	// os.Stdout 实现了 Writer 接口
	io.WriteString(os.Stdout, "Hello World!")
	// Hello World!
}

func TestNewReader(t *testing.T) {
	// ReadAtLeast 从 r 中读取数据到 buf 中，要求至少读取 min 个字节
	// 返回读取的字节数 n 和读取过程中遇到的任何错误
	// 如果 n < min，则 err 返回 ErrUnexpectedEOF
	// 如果 r 中无可读数据，则 err 返回 EOF
	// 如果 min 大于 len(buf)，则 err 返回 ErrShortBuffer
	// 只有当 n >= min 时 err 才返回 nil
	// func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
	r := strings.NewReader("Hello World!")
	b := make([]byte, 32)
	n, err := io.ReadAtLeast(r, b, 20)
	fmt.Printf("%s\n%d, %v", b, n, err)
	// Hello World!
	// 12, unexpected EOF
}
func TestCopyN(t *testing.T) {
	// CopyN 从 src 中复制 n 个字节的数据到 dst 中
	// 它返回复制的字节数 written 和复制过程中遇到的任何错误
	// 只有当 written = n 时，err 才返回 nil
	// 如果 dst 实现了 ReadFrom 方法，则调用 ReadFrom 来执行复制操作
	// func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
	r := strings.NewReader("Hello World!")
	n, err := io.CopyN(os.Stdout, r, 20)
	fmt.Printf("\n%d, %v", n, err)
	// Hello World!
	// 12, EOF
}
func TestCopy(t *testing.T) {
	// Copy 从 src 中复制数据到 dst 中，直到所有数据复制完毕
	// 返回复制过程中遇到的任何错误
	// 如果数据复制完毕，则 err 返回 nil，而不是 EOF
	// 如果 dst 实现了 ReadeFrom 方法，则调用 dst.ReadeFrom(src) 复制数据
	// 如果 src 实现了 WriteTo 方法，则调用 src.WriteTo(dst) 复制数据
	// func Copy(dst Writer, src Reader) (written int64, err error)
	r := strings.NewReader("Hello World!")
	n, err := io.Copy(os.Stdout, r)
	fmt.Printf("\n%d, %v", n, err)
	// Hello World!
	// 12, <nil>
}
func TestLimitReader(t *testing.T) {
	// LimitReader 覆盖了 r 的 Read 方法
	// 使 r 只能读取 n 个字节的数据，读取完毕后返回 EOF
	// func LimitReader(r Reader, n int64) Reader

	// LimitedReader 结构用来实现 LimitReader 的功能
	// type LimitedReader struct

	r := strings.NewReader("Hello World!")
	lr := io.LimitReader(r, 5)
	n, err := io.CopyN(os.Stdout, lr, 20)
	fmt.Printf("\n%d, %v", n, err)
	// Hello
	// 5, EOF
}
func TestNewSectionReader(t *testing.T) {
	// NewSectionReader 封装 r，并返回 SectionReader 类型的对象
	// 使 r 只能从 off 的位置读取 n 个字节的数据，读取完毕后返回 EOF
	// func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader

	// SectionReader 结构用来实现 NewSectionReader 的功能
	// SectionReader 实现了 Read、Seek、ReadAt、Size 方法
	// type SectionReader struct
	// 	// Size 返回 s 中被限制读取的字节数
	// 	func (s *SectionReader) Size()
	r := strings.NewReader("Hello World!")
	sr := io.NewSectionReader(r, 0, 5)
	n, err := io.CopyN(os.Stdout, sr, 20)
	fmt.Printf("\n%d, %v", n, err)
	fmt.Printf("\n%d", sr.Size())
	// World
	// 5, EOF
	// 5
}
func TestTeeReader(t *testing.T) {
	// TeeReader 覆盖了 r 的 Read 方法
	// 使 r 在读取数据的同时，自动向 w 中写入数据
	// 所有写入时遇到的错误都被作为 err 返回值
	// func TeeReader(r Reader, w Writer) Reader

	r := strings.NewReader("Hello World!")
	tr := io.TeeReader(r, os.Stdout)
	b := make([]byte, 32)
	tr.Read(b)
	// World World!
}
