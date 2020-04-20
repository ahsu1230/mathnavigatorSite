package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Client Configuration struct to fit incoming configuration yml file
type clientConfig struct {
	Config struct {
		Message string `yaml:"message"`
	} `yaml:"config"`
}

// Extract data from yml file
func retrieveConfigurations(fileName string) clientConfig {
	fmt.Println("Configuration File: ", fileName)

	file, errFile := os.Open(fileName)
	if errFile != nil {
		fmt.Println("Error with file ", fileName)
	}

	var cfg clientConfig
	decoder := yaml.NewDecoder(file)
	errParse := decoder.Decode(&cfg)
	if errParse != nil {
		fmt.Println("Error from parsing ", errParse)
	}
	return cfg
}

func main() {
	fmt.Println("Example client starting...")

	// Retrieve Configurations
	// We assume the configuration file will be passed when running this go app
	configFile := os.Args[1]
	config := retrieveConfigurations(configFile)
	fmt.Println("Received this message", config.Config.Message)

	// You can run this file using:
	// go run exampleCli.go exampleCli.yml

	// or via a binary:
	// go build exampleCli.go
	// ./exampleCli exampleCli.yml
}
