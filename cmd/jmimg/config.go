package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jinlimei/jmimg"
)

func getJmConfig() *jmimg.Config {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatalf("Failed to get user home dir: %v", err)
	}

	cfgName := fmt.Sprintf("%s/.jmimg.json", homeDir)

	data, err := os.ReadFile(cfgName)

	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var cfg *jmimg.Config

	err = json.Unmarshal(data, &cfg)

	if err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	return cfg
}
