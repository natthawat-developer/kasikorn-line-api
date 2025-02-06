package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port string `yaml:"port"`

	DB struct {
		Driver   string `yaml:"driver"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"db"`

	CORS struct {
		AllowOrigins string `yaml:"allow_origins"`
		AllowMethods string `yaml:"allow_methods"`
		AllowHeaders string `yaml:"allow_headers"`
	} `yaml:"cors"`
}

func LoadConfig() *Config {

	file, err := os.Open("config.yaml")
	if err != nil {
		log.Fatalf("error opening config.yaml: %v", err)
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("error decoding YAML: %v", err)
	}

	return &config
}
