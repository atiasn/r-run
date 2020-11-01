package main

import (
	"flag"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

// TODO 添加日志
func ConnectOss()(*oss.Client, error){
	endpoint := os.Getenv("ALI_OSS_ENDPOINT")
	accessKeyId := os.Getenv("ALI_OSS_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("ALI_OSS_ACCESS_KEY_SECRET")
	if accessKeySecret == "" || accessKeyId == "" || endpoint == "" {
		panic("阿里云OSS信息缺失，请确认")
	}
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret, oss.UseCname(true))
	if err != nil {
		panic(err)
	}
	return client, err
}


func PutExcelFile(client *oss.Client, fileName string) (string, error){
	// 获取bucket对象
	bucketName := "trans-excel"
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		panic(err)
	}

	err = bucket.PutObjectFromFile(fileName, fileName)
	if err != nil {
		panic(err)
	}
	baseUrl := os.Getenv("ALI_OSS_ENDPOINT")
	imgUrl := baseUrl + "/" + fileName
	return imgUrl, err
}


func main()  {
	flag.Parse()
	updateFiles := flag.Args()
	client, err := ConnectOss()
	if err != nil {
		panic(err)
	}
	for index, upFile := range updateFiles {
		excelUrl, err := PutExcelFile(client, upFile)
		if err != nil {
			panic(err)
		}
		if index == 0 {
			fmt.Println("文件上传成功")
		}
		fmt.Println(excelUrl)
	}
}

