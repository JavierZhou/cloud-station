package aliyun

import (
	"fmt"
	"os"

	"github.com/JavierZhou/cloud-station/store"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	_ store.Uploader = &AliOssStore{}
)

type Options struct {
	Endpoint     string
	AccessKey    string
	AccessSecret string	
}

func (o *Options) varCheck() error {
	if o.Endpoint == "" || o.AccessKey == "" || o.AccessSecret == "" {
		return fmt.Errorf("endpoint or ak or sk empty")
	}

	return nil
}

func NewDefaultAliOssStore() (*AliOssStore, error) {
	return NewAliOssStore(&Options{
		AccessKey: 		os.Getenv("ALI_AK"),
		AccessSecret: 	os.Getenv("ALI_SK"),
		Endpoint: 		os.Getenv("ALI_OSS_ENDPOINT"),
	})
}

func NewAliOssStore(opts *Options) (*AliOssStore, error) {
	if err := opts.varCheck(); err != nil {
		return nil, err
	}

	c, err := oss.New(opts.Endpoint, opts.AccessKey, opts.AccessSecret)
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
