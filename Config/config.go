package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type DatabaseConfig struct {
}

type AppConfig struct {
	Database DatabaseConfig `yaml:"database"`
}

func LoadConfiguration(pathOfYaml string) (*AppConfig, error) {
	file, err := os.Open(pathOfYaml)
	if err != nil {
		log.Fatalf("Failed to find yaml file: %v", err)
		return nil, err
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read yaml file")
		return nil, err
	}

	var config AppConfig
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal yaml file")
		return nil, err
	}

	return &config, nil
}
