package middlewares

import (
	"fmt"
	"os"
	"gopkg.in/yaml.v2"
)

type Config struct {
	App struct {
		Build      string `yaml:"build"`
		CorsOrigin string `yaml:"corsOrigin"`
	} `yaml:"app"`
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
		DbName	 string `yaml:"dbName"`
	} `yaml:"database"`
}

func RetrieveConfigurations(fileName string) Config {
	fmt.Println("Configuration File: ", fileName)

	file, errFile := os.Open(fileName)
	if errFile != nil {
		fmt.Println("Error with file ", fileName)
	}

	var cfg Config
	decoder := yaml.NewDecoder(file)
	errParse := decoder.Decode(&cfg)
	if errParse != nil {
		fmt.Println("Error from parsing ", errParse)
	}
	return cfg
}
