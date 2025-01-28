package jmimg

import (
	"github.com/aws/aws-sdk-go-v2/aws"
)

type Service struct {
	cfg           aws.Config
	bucketName    string
	cdnUrl        string
	nameGenerator UploadNameGenerator
}

func New(cfg aws.Config, bucketName, cdnUrl string) *Service {
	return &Service{
		cfg:        cfg,
		bucketName: bucketName,
		cdnUrl:     cdnUrl,
	}
}
