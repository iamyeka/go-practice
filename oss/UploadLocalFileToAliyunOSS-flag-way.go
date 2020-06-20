package main

import (
	"flag"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

var endpoint, accessKeyId, accessKeySecret, bucketName, objectKey, localFilePath string

func init() {
	flag.StringVar(&endpoint, "endpoint", "", "aliyun enpoint")
	flag.StringVar(&accessKeyId, "accessKeyId", "", "aliyun accessKeyId")
	flag.StringVar(&accessKeySecret, "accessKeySecret", "", "aliyun accessKeySecret")
	flag.StringVar(&bucketName, "bucketName", "", "aliyun bucketName")
	flag.StringVar(&objectKey, "objectKey", "", "bucket objectKey")
	flag.StringVar(&localFilePath, "localFilePath", "", "localFilePath")
}

/*
用flag模块传递参数，实现将本地文件上传到阿里云OSS的逻辑
使用方式：
1. 打包成对应操作系统的可执行文件。如linux系统：GOOS=linux go build UploadLocalFileToAliyunOSS-flag-way.go
2. UploadLocalFileToAliyunOSS -endpoint *** -accessKeyId *** -accessKeySecret *** -bucketName *** -objectKey *** -localFilePath ***
*/
func main() {
	flag.Parse()

	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 上传本地文件。
	err = bucket.PutObjectFromFile(objectKey, localFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}
