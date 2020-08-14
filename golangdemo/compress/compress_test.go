package main

import (
	"compress/gzip"
	"log"
	"os"
	"testing"
)

/*
 const (
        NoCompression = flate.NoCompression // 不压缩
        BestSpeed = flate.BestSpeed // 最快速度
        BestCompression = flate.BestCompression // 最佳压缩比
        DefaultCompression = flate.DefaultCompression   // 默认压缩比
    )
    var (
        // 当读取gzip数据时发现无效的校验和时将返回该错误
        ErrChecksum = errors.New("gzip: invalid checksum")
        // 当读取gzip数据时发现无效的数据头时将返回该错误
        ErrHeader = errors.New("gzip: invalid header")
    )
    //数据头结构
    type Header struct {
        Comment string    // 文件注释
        Extra   []byte    // 附加数据
        ModTime time.Time // 文件修改时间
        Name    string    // 文件名
        OS      byte      // 操作系统类型
    }

*/
func TestZip(t *testing.T) {
	fw, err := os.Create("demo.gzip") // 创建gzip包文件，返回*io.Writer
	if err != nil {
		log.Fatalln(err)
	}
	defer fw.Close()

	// 实例化新的gzip.Writer
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	// 获取要打包的文件信息
	fr, err := os.Open("demo.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer fr.Close()

	// 获取文件头信息
	fi, err := fr.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	// 创建gzip.Header
	gw.Header.Name = fi.Name()

	// 读取文件数据
	buf := make([]byte, fi.Size())
	_, err = fr.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}

	// 写入数据到zip包
	_, err = gw.Write(buf)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestUnZip(t *testing.T) {
	// 打开gzip文件
	fr, err := os.Open("demo.gzip")
	if err != nil {
		log.Fatalln(err)
	}
	defer fr.Close()

	// 创建gzip.Reader
	gr, err := gzip.NewReader(fr)
	if err != nil {
		log.Fatalln(err)
	}
	defer gr.Close()

	// 读取文件内容
	buf := make([]byte, 1024*1024*10) // 如果单独使用，需自己决定要读多少内容，根据官方文档的说法，你读出的内容可能超出你的所需（当你压缩gzip文件中有多个文件时，强烈建议直接和tar组合使用）
	n, err := gr.Read(buf)

	// 将包中的文件数据写入
	fw, err := os.Create(gr.Header.Name)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = fw.Write(buf[:n])
	if err != nil {
		log.Fatalln(err)
	}
}
