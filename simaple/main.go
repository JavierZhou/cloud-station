package main

import (
	"flag"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

var (
	AccessKey    = ""
	AccessSecret = ""
	OssEndpoint  = ""
	BucketName   = ""
	UploadFile   = ""
	Help         = false
)

func upload(filePath string) error {
	client, err := oss.New(OssEndpoint, AccessKey, AccessSecret)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(BucketName)
	if err != nil {
		return err
	}

	err = bucket.PutObjectFromFile(filePath, filePath)
	if err != nil {
		return err
	}

	url, err := bucket.SignURL(filePath, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}
	fmt.Printf("文件URL: %v\n", url)
	fmt.Println("文件有效期为1天！！！")

	return nil
}

func varCheck() error {
	if AccessKey == "" && AccessSecret == "" && OssEndpoint == "" {
		return fmt.Errorf("endpoint or ak or sk is empty")
	}

	if UploadFile == "" {
		return fmt.Errorf("upload file path is empty")
	}

	return nil
}

func loadParams() {
	flag.BoolVar(&Help, "h", false, "this is help doc")
	flag.StringVar(&UploadFile, "f", "", "上传文件的名称")
	flag.Parse()

	if Help {
		usage()
		os.Exit(0)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `cloud-station version 0.0.1
Usage: cloud-station -f <file_path>
`)
	flag.PrintDefaults()
}

func main() {
	loadParams()

	if err := varCheck(); err != nil {
		fmt.Printf("参数校验异常: %v\n", err)
		usage()
		os.Exit(1)
	}

	err := upload(UploadFile)
	if err != nil {
		fmt.Printf("上传文件异常：%v\n", err)
		usage()
		os.Exit(1)
	}

	fmt.Printf("文件 [%v] 上传成功\n", UploadFile)
}
