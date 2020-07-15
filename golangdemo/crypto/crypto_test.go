package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
	"math/big"
	"testing"
)

func TestMd5(t *testing.T) {
	data := []byte("These pretzels are making me thirsty.")
	fmt.Printf("%x", md5.Sum(data))
	fmt.Println()
	h := md5.New()
	io.WriteString(h, "The fog is getting thicker!")
	io.WriteString(h, "And Leon's getting laaarger!")
	fmt.Printf("%x", h.Sum(nil))
	fmt.Println()

	str := "1"
	// 方法一
	data = []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) // 将[]byte转成16进制
	fmt.Println(md5str1)
	fmt.Println()
	// 方法二
	w := md5.New()
	io.WriteString(w, str)
	// 将str写入到w中
	md5str2 := fmt.Sprintf("%x", w.Sum(nil))

	fmt.Println(md5str2)

}

func TestAES(t *testing.T) {
	x := []byte("世界上最邪恶最专制的现代奴隶制国家--朝鲜")
	key := []byte("hgfedcba87654321")
	x1 := encryptAES(x, key)
	fmt.Print(string(x1))
	x2 := decryptAES(x1, key)
	fmt.Print(string(x2))
}
func padding(src []byte, blocksize int) []byte {
	padnum := blocksize - len(src)%blocksize
	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)
	return append(src, pad...)
}

func unpadding(src []byte) []byte {
	n := len(src)
	unpadnum := int(src[n-1])
	return src[:n-unpadnum]
}

func encryptAES(src []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	src = padding(src, block.BlockSize())
	blockmode := cipher.NewCBCEncrypter(block, key)
	blockmode.CryptBlocks(src, src)
	return src
}

func decryptAES(src []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	blockmode := cipher.NewCBCDecrypter(block, key)
	blockmode.CryptBlocks(src, src)
	src = unpadding(src)
	return src
}

func TestCipher(t *testing.T) {
	// 这个包实现了标准的块加密模式。我们可以看一下cipher.Block
	//
	// type Block interface {
	// 	BlockSize() int
	// 	Encrypt(dst, src []byte)
	// 	Decrypt(dst, src []byte)
	// }
	// Go中定义了一个接口BlockMode代表各种模式
	//
	// type BlockMode interface {
	// 	BlockSize() int
	// 	CryptBlocks(dst, src []byte)
	// }
	// 该包提供了获取BlockMode实例的两个方法
	//
	// func NewCBCDecrypter(b Block, iv []byte) BlockMode
	// func NewCBCEncrypter(b Block, iv []byte) BlockMode
	// 即一个CBC加密，一个CBC解密

}

func TestDSA(t *testing.T) {

}

func TestRSA(t *testing.T) {

}

func TestDES(t *testing.T) {

}

func TestRand(t *testing.T) {
	// rand包实现了用于加解密的更安全的随机数生成器。
	// 1、Int
	n, err := rand.Int(rand.Reader, big.NewInt(128))
	if err == nil {
		fmt.Println("rand.Int：", n, n.BitLen())
	}
	// 2、Prime
	p, err := rand.Prime(rand.Reader, 5)
	if err == nil {
		fmt.Println("rand.Prime：", p)
	}
	// 3、Read
	b := make([]byte, 32)
	m, err := rand.Read(b)
	if err == nil {
		fmt.Println("rand.Read：", b[:m])
		fmt.Println("rand.Read：", base64.URLEncoding.EncodeToString(b))
	}

	// rand.Int： 92 7
	// rand.Prime： 29
	// rand.Read： [207 47 241 208 190 84 109 134 86 106 87 223 111 113 203 155 44 118 71 20 186 62 66 130 244 98 97 184 8 179 6 230]
	// rand.Read： zy_x0L5UbYZWalffb3HLmyx2RxS6PkKC9GJhuAizBuY=
}

func TestSha1(t *testing.T) {
	h := sha1.New()
	h.Write([]byte("hello,world"))
	l := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(l)

	s := "sha1 this string"
	// 产生一个散列值得方式是 sha1.New()，sha1.Write(bytes)，然后 sha1.Sum([]byte{})。这里我们从一个新的散列开始。
	h = sha1.New()
	// 写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	h.Write([]byte(s))
	// 这个用来得到最终的散列值的字符切片。Sum 的参数可以用来都现有的字符切片追加额外的字节切片：一般不需要要。
	bs := h.Sum(nil)
	// SHA1 值经常以 16 进制输出，例如在 git commit 中。使用%x 来将散列结果格式化为 16 进制字符串。
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}

func TestSha256(t *testing.T) {
	// 第一种调用方法
	sum := sha256.Sum256([]byte("hello world\n"))
	fmt.Printf("%x\n", sum)

	// 第二种调用方法
	h := sha256.New()
	h.Write([]byte("hello world\n"))
	fmt.Printf("%x\n", h.Sum(nil))
}

func TestSha512(t *testing.T) {
	// 第一种调用方法
	sum := sha512.Sum512([]byte("hello world\n"))
	fmt.Printf("%x\n", sum)

	// 第二种调用方法
	h := sha512.New()
	h.Write([]byte("hello world\n"))
	fmt.Printf("%x\n", h.Sum(nil))
}
