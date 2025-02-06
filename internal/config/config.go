package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Config struct สำหรับเก็บข้อมูลจากไฟล์ config.yaml
type Config struct {
	Port string `yaml:"port"`
	DB   struct {
		Driver   string `yaml:"driver"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"db"`
}

// LoadConfig function สำหรับโหลดค่า config จากไฟล์ config.yaml
func LoadConfig() *Config {
	// เปิดไฟล์ config.yaml
	file, err := os.Open("config.yaml")
	if err != nil {
		log.Fatalf("error opening config.yaml: %v", err)
	}
	defer file.Close()

	// อ่านข้อมูลจากไฟล์
	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("error decoding YAML: %v", err)
	}

	
	// คืนค่าคอนฟิก
	return &config
}
