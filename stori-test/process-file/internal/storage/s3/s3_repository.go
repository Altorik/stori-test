package s3

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	file "process-file/internal"
)

type Repository struct {
	s3Client   *s3.Client
	bucketName string
}

func NewS3Repository(s3Client *s3.Client, bucketName string) *Repository {
	return &Repository{
		s3Client:   s3Client,
		bucketName: bucketName,
	}
}

func (r *Repository) GetFileInBytes(s3File file.File) ([]byte, error) {
	downloader := manager.NewDownloader(r.s3Client, func(down *manager.Downloader) {
		down.PartSize = 10 * 1024 * 1024
	})
	bufferIn := manager.NewWriteAtBuffer([]byte{})
	_, err := downloader.Download(context.TODO(), bufferIn, &s3.GetObjectInput{
		Bucket: aws.String(r.bucketName),
		Key:    aws.String(s3File.Path + s3File.Name),
	})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return bufferIn.Bytes(), err
}
