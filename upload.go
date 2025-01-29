package jmimg

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	ErrNotImplementedYet = errors.New("not implemented yet")
)

func (s *Service) UploadFile(itu *ImageToUpload) (string, error) {
	var err error

	s3c := s3.NewFromConfig(s.awsCfg)

	if s.cfg.ConvertToJPEG {
		err = itu.changeToJPEG()

		if err != nil {
			return "", err
		}
	}

	if s.cfg.MaxWidth > 0 || s.cfg.MaxHeight > 0 {
		err = itu.resize(s.cfg.MaxWidth, s.cfg.MaxHeight)

		if err != nil {
			return "", err
		}
	}

	if s.cfg.StripMetadata != nil && *s.cfg.StripMetadata {
		return "", ErrNotImplementedYet
	}

	uploadFileName := itu.FileName

	if s.nameGenerator != nil {
		uploadFileName = s.nameGenerator(itu.FileName, itu.MimeType())
	}

	ctx := context.TODO()

	_, err = s3c.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.cfg.BucketName),
		Key:         aws.String(uploadFileName),
		Body:        itu.Reader(),
		ContentType: aws.String(itu.MimeType()),
	})

	if err != nil {
		return "", err
	}

	err = s3.NewObjectExistsWaiter(s3c).Wait(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(s.cfg.BucketName),
		Key:    aws.String(uploadFileName),
	}, time.Minute)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://%s/%s", s.cfg.CDNUrl, uploadFileName), nil
}
