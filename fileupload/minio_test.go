package main

import (
	"fmt"
	"github.com/minio/minio-go/v6"
	"log"
	"os"
	"testing"
)

func TestPutObject(t *testing.T) {
	var (
		endpoint        = "172.16.2.91:9000"
		accessKeyID     = "minioadmin"
		secretAccessKey = "minioadmin"
		useSSL          = false
		//filePath        = "D:/d1.jpg"
	)
	// #  创建桶
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)

	// 创建一个叫mymusic的存储桶。
	bucketName := "mymusic"
	location := "us-east-1"
	//objectName := "testminio4.jpg"
	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// 检查存储桶是否已经存在。
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	file, err := os.Open("D:/d1.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	n, err := minioClient.PutObject(bucketName, "myobject.jpg", file, fileStat.Size(), minio.PutObjectOptions{ContentType: "video/jpeg"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully uploaded bytes: ", n)

}
