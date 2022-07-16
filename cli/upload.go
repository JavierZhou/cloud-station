package cli

import (
	"fmt"

	"github.com/JavierZhou/cloud-station/store"
	"github.com/JavierZhou/cloud-station/store/aliyun"
	"github.com/spf13/cobra"
)

var (
	ossProvider  string
	ossEndpoint  string
	bucketName   string
	accessKey    string
	accessSecret string
	uploadFile   string
)

var (
	UploadCmd = &cobra.Command{
		Use:     "upload",
		Long:    "upload 云中上传",
		Short:   "upload 云中上传",
		Example: "upload -f filePath",
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				uploader store.Uploader
				err      error
			)
			switch ossProvider {
			case "aliyun":
				uploader, err = aliyun.NewAliOssStore(&aliyun.Options{
					Endpoint:     ossEndpoint,
					AccessKey:    accessKey,
					AccessSecret: accessSecret,
				})
			case "tx":
			case "aws":
			default:
				return fmt.Errorf("not support oss storage provinder")

			}
			if err != nil {
				return err
			}

			return uploader.Upload(bucketName, uploadFile, uploadFile)
		},
	}
)

func init() {
	f := UploadCmd.PersistentFlags()
	f.StringVarP(&ossProvider, "provider", "p", "aliyun", "oss storage provider")
	f.StringVarP(&ossEndpoint, "endpoint", "e", "oss-cn-shenzhen.aliyuncs.com", "oss endpoint")
	f.StringVarP(&bucketName, "bucketName", "b", "blog-javier", "oss bucketName")
	f.StringVarP(&accessKey, "accessKey", "k", "", "oss ak")
	f.StringVarP(&accessSecret, "accessSecret", "s", "", "oss sk")
	f.StringVarP(&uploadFile, "uploadFile", "f", "", "oss upload file path")
	RootCmd.AddCommand(UploadCmd)
}
