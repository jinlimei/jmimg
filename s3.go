package jmimg

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (s *Service) UploadFile(name string) (string, error) {
	ctx := context.TODO()

	osf, err := os.OpenFile(name, os.O_RDONLY, 0644)

	if err != nil {
		return "", err
	}

	upName := name

	if s.nameGenerator != nil {
		upName = s.nameGenerator(name)
	}

	s3c := s3.NewFromConfig(s.cfg)

	var first512 = make([]byte, 512)
	_, err = osf.Read(first512)
	if err != nil {
		return "", err
	}

	_, err = osf.Seek(0, 0)
	if err != nil {
		return "", err
	}

	_, err = s3c.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucketName),
		Key:         aws.String(upName),
		Body:        osf,
		ContentType: aws.String(http.DetectContentType(first512)),
	})

	if err != nil {
		return "", err
	}

	err = s3.NewObjectExistsWaiter(s3c).Wait(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(upName),
	}, time.Minute)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://%s/%s", s.cdnUrl, upName), nil
}
