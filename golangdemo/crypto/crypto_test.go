package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"testing"
)

func TestMd5(t *testing.T) {
	data := []byte("These pretzels are making me thirsty.")
	fmt.Printf("%x", md5.Sum(data))

	h := md5.New()
	io.WriteString(h, "The fog is getting thicker!")
	io.WriteString(h, "And Leon's getting laaarger!")
	fmt.Printf("%x", h.Sum(nil))
}

func TestAES(t *testing.T) {

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
