package store

type Uploader interface {
	Upload(buketName string, objectKey string, fileName string) error
}
