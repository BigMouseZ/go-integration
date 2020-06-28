package ziputil

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// srcFile could be a single file or a directory
func Zip(srcFile string, destZip string) error {
	zipfile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	worlkerr := filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 将文件或者目录信息转换为zip格式的文件信息
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(path, filepath.Dir(srcFile)+"/")
		// header.Name = path
		if info.IsDir() {
			header.Name += "/"
		} else {
			// 确定采用的压缩算法（这个是内建注册的deflate）
			header.Method = zip.Deflate
		}
		//

		//创建在zip内的文件或者目录
		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if ! info.IsDir() {
			//打开需要压缩的文件
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			//将待压缩文件拷贝给zip内文件
			_, err = io.Copy(writer, file)
		}
		return err
	})
	if worlkerr != nil {
		fmt.Println(worlkerr)
	}
	return err
}
