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

	str := "These pretzels are making me thirsty."
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
	has, err := AesEncryptSimple([]byte("guoshuaijie"), "signalwayKey", "signalwayiv")
	if err != nil {
		panic(err)
	}
	md5str1 := fmt.Sprintf("%x", has) // 将[]byte转成16进制
	fmt.Println(md5str1)
	raw := []byte("guoshuaijie")
	key := []byte("huyanyan87654321")
	//key2:=[]byte("guoshuai87654321")
	encryptByte := EncryptAES(raw, key)
	if encryptByte != nil {
		decryptByte := DecryptAES(encryptByte, key)
		if decryptByte != nil {
			fmt.Println(string(decryptByte))
		}
	}
}

// 加密
func AesEncryptSimple(origData []byte, key string, iv string) ([]byte, error) {
	return AesDecryptPkcs5(origData, []byte(key), []byte(iv))
}

func AesEncryptPkcs5(origData []byte, key []byte, iv []byte) ([]byte, error) {
	return AesEncrypt(origData, key, iv, PKCS5Padding)
}

func AesEncrypt(origData []byte, key []byte, iv []byte, paddingFunc func([]byte, int) []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = paddingFunc(origData, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// 解密
func AesDecryptSimple(crypted []byte, key string, iv string) ([]byte, error) {
	return AesDecryptPkcs5(crypted, []byte(key), []byte(iv))
}

func AesDecryptPkcs5(crypted []byte, key []byte, iv []byte) ([]byte, error) {
	return AesDecrypt(crypted, key, iv, PKCS5UnPadding)
}

func AesDecrypt(crypted, key []byte, iv []byte, unPaddingFunc func([]byte) []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = unPaddingFunc(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	if length < unpadding {
		return []byte("unpadding error")
	}
	return origData[:(length - unpadding)]
}

func padding(src []byte, blocksize int) []byte {
	padnum := blocksize - len(src)%blocksize
	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)
	return append(src, pad...)
}

func unPadding(src []byte) []byte {
	n := len(src)
	unPadNum := int(src[n-1])
	if unPadNum > n {
		return nil
	}
	return src[:n-unPadNum]
}

func EncryptAES(src []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}
	src = padding(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	blockMode.CryptBlocks(src, src)
	return src
}

func DecryptAES(src []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	blockMode.CryptBlocks(src, src)
	src = unPadding(src)
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
