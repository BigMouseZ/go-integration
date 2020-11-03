package main

import (
	"github.com/minio/minio-go/v6"
	"log"
)

func main() {

	var (
		endpoint        = "172.16.2.91:9000"
		accessKeyID     = "minioadmin"
		secretAccessKey = "minioadmin"
		useSSL          = false
		filePath        = "D:/d1.jpg"
	)
	// #  创建桶
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)

	// 创建一个叫mymusic的存储桶。
	bucketName := "mymusic"
	location := "us-east-1"
	objectName := "test2/test/testminio4.jpg"
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

	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: "video/jpeg"})
	if err != nil {
		log.Panic(err)
		return
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)

}
