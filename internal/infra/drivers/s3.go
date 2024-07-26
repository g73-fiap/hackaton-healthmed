package drivers

import "mime/multipart"

type S3Driver interface {
	PutObject(name string, file *multipart.FileHeader) (string, error)
	RemoveObject(name string) error
}

type s3Driver struct {
}

func NewS3Driver() S3Driver {
	return &s3Driver{}
}

func (s *s3Driver) PutObject(name string, file *multipart.FileHeader) (string, error) {
	return "location", nil
}
func (s *s3Driver) RemoveObject(name string) error {
	return nil
}
