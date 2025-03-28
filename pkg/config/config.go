package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DBConfig struct {
		MongoURI string `yaml:"MongoURI"`
	} `yaml:"DBConfig"`
	SearchEngineConfig struct {
		RussianStopWordsPath string `yaml:"RussianStopWordsPath"`
		EnglishStopWordsPath string `yaml:"EnglishStopWordsPath"`
		Port                 int    `yaml:"Port"`
	} `yaml:"SearchEngineConfig"`
}

func GetConfig() (*Config, error) {
	config_path := os.Getenv("CONFIG_PATH")
	data, err := os.ReadFile(config_path)
	if err != nil {
		return nil, err
	}

	fmt.Println("Содержимое файла:", string(data))

	cnfg := Config{}
	if err = yaml.Unmarshal(data, &cnfg); err != nil {
		return nil, err
	}

	fmt.Printf("--- t:\n%v\n\n", cnfg)

	return &cnfg, nil
}
