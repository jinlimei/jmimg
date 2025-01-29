package main

import (
	"context"
	"fmt"
	"log"
	"os"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/jinlimei/jmimg"
)

var (
	BuildTime  string
	CommitHash string
	GoVersion  string
	GitTag     string
	Version    string
)

func main() {
	jmCfg := getJmConfig()

	awsCfg, err := awsConfig.LoadDefaultConfig(
		context.TODO(),
		awsConfig.WithRegion(jmCfg.AwsRegion),
		awsConfig.WithSharedConfigProfile(jmCfg.AwsProfile),
	)

	if err != nil {
		log.Fatalf("Failed to laod default config: %v", err)
	}

	svc := jmimg.New(
		awsCfg,
		jmCfg,
	)

	svc.SetUploadNameGenerator(fileGen)

	fileArgs := os.Args[1:]

	if len(fileArgs) == 0 {
		fmt.Printf("Usage: %s [file1] [file2] [file3] ... [fileN]\n", os.Args[0])
		fmt.Printf("\nBuild Time: %s\n", BuildTime)
		fmt.Printf("Commit Hash: %s\n", CommitHash)
		fmt.Printf("Go Version: %s\n", GoVersion)
		fmt.Printf("Git Tag: %s\n", GitTag)
		fmt.Printf("Version: %s\n", Version)

		os.Exit(1)
		return
	}

	var (
		uri  string
		file *os.File
		itu  *jmimg.ImageToUpload
	)

	for _, fileName := range fileArgs {
		file, err = os.OpenFile(fileName, os.O_RDONLY, 0640)

		if err != nil {
			log.Fatalf("Failed to open %s: %v", fileName, err)
		}

		itu, err = jmimg.NewImageUpload(fileName, file)

		if err != nil {
			log.Fatalf("Failed to build image upload for %s: %v", fileName, err)
		}

		uri, err = svc.UploadFile(itu)

		if err != nil {
			log.Fatalf("Failed to upload file: %v", err)
		}

		fmt.Println(uri)
	}

	fmt.Printf("Uploaded %d files\n", len(fileArgs))
}
