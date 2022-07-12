package example_test

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"testing"
)

var (
	// 全局client，init里初始化
	client *oss.Client
)

var (
	AccessKey    = os.Getenv("ALI_AK")
	AccessSecret = os.Getenv("ALI_SK")
	OssEndpoint  = os.Getenv("ALI_OSS_ENDPOINT")
	BucketName   = os.Getenv("ALI_BUCKET_NAME")
)

// 测试阿里云OSS SDK ListBuckets 接口
func TestListBuckets(t *testing.T) {
	lsRes, err := client.ListBuckets()
	if err != nil {
		// HandleError(err)
		t.Log(err)
	}

	for _, bucket := range lsRes.Buckets {
		fmt.Println("Buckets:", bucket.Name)
	}
}

// 测试阿里云OSS SDK PutObjectFromFile 接口
func TestUploadFile(t *testing.T) {
	bucket, err := client.Bucket(BucketName)
	if err != nil {
		// HandleError(err)
		t.Log(err)
	}

	err = bucket.PutObjectFromFile("mydir/test.go", "oss_test.go")
	if err != nil {
		// HandleError(err)
		t.Log(err)
	}
}

//初始化一个OSS client
func init() {
	c, err := oss.New(OssEndpoint, AccessKey, AccessSecret)
	if err != nil {
		// HandleError(err)
		panic(err)
	}
	client = c
}
