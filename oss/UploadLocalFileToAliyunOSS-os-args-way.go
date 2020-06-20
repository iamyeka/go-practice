package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

// 通过在可执行文件后跟上参数，实现将本地文件上传到阿里云OSS的逻辑
// 使用方式：UploadLocalFileToAliyunOSS endpoint accessKeyId accessKeySecret bucketName objectKey localFilePath
func main() {
	// 创建OSSClient实例。
	endpoint := os.Args[1]
	accessKeyId := os.Args[2]
	accessKeySecret := os.Args[3]

	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	bucketName := os.Args[4]
	objectKey := os.Args[5]
	localFilePath := os.Args[6]

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
