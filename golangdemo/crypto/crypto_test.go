package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"fmt"
	"io"
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

}

func TestDSA(t *testing.T) {

}

func TestRSA(t *testing.T) {

}

func TestRand(t *testing.T) {

}

func TestDES(t *testing.T) {

}

func TestSha1(t *testing.T) {

}

func TestSha256(t *testing.T) {

}

func TestSha512(t *testing.T) {

}
