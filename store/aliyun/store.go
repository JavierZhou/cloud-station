package aliyun

import (
	"fmt"
	"github.com/JavierZhou/cloud-station/store"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	_ store.Uploader = &AliOssStore{}
)

func NewAliOssStore(endpoint, accessKey, accessSecret string) (*AliOssStore, error) {
	c, err := oss.New(endpoint, accessKey, accessSecret)
	if err != nil {
		return nil, err
	}
	return &AliOssStore{
		client: c,
	}, nil
}

type AliOssStore struct {
	client *oss.Client
}

func (s *AliOssStore) Upload(bucketName string, objectKey string, fileName string) error {
	bucket, err := s.client.Bucket(bucketName)
	if err != nil {
		return err
	}

	err = bucket.PutObjectFromFile(objectKey, fileName)
	if err != nil {
		return err
	}

	url, err := bucket.SignURL(objectKey, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}
	fmt.Printf("文件URL: %v\n", url)
	fmt.Println("文件有效期为1天！！！")

	return nil
}
