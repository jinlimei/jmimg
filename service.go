package jmimg

import (
	"github.com/aws/aws-sdk-go-v2/aws"
)

type Service struct {
	awsCfg        aws.Config
	cfg           *Config
	nameGenerator UploadNameGenerator
}

func New(awsCfg aws.Config, cfg *Config) *Service {
	return &Service{
		awsCfg: awsCfg,
		cfg:    cfg,
	}
}
