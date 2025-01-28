package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/jinlimei/jmimg"
)

type Config struct {
	AwsProfile string `json:"aws_profile"`
	AwsRegion  string `json:"aws_region"`
	BucketName string `json:"bucket_name"`
	CDNUrl     string `json:"cdn_url"`
}

func main() {
	jmCfg := getConfig()

	awsCfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(jmCfg.AwsRegion),
		config.WithSharedConfigProfile(jmCfg.AwsProfile),
	)

	if err != nil {
		log.Fatalf("Failed to laod default config: %v", err)
	}

	svc := jmimg.New(
		awsCfg,
		jmCfg.BucketName,
		jmCfg.CDNUrl,
	)

	svc.SetUploadNameGenerator(fileGen)

	fileArgs := os.Args[1:]

	var up string

	for _, arg := range fileArgs {
		up, err = svc.UploadFile(arg)

		if err != nil {
			log.Fatalf("Failed to upload file: %v", err)
		}

		fmt.Println(up)
	}
}

func getConfig() *Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get user home dir: %v", err)
	}

	cfgName := fmt.Sprintf("%s/.jmimg.json", homeDir)

	data, err := os.ReadFile(cfgName)

	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var cfg *Config

	err = json.Unmarshal(data, &cfg)

	if err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	return cfg
}
