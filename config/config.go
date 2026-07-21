package config

import (
	"log"

	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type DatabaseConfig struct {
	Path string `yaml:"path"`
}

var cfg Config

func LoadConfig() {
	file, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		log.Fatalf("failed to parse config file: %v", err)
	}
}

func GetServerPort() int {
	return cfg.Server.Port
}

func GetDatabasePath() string {
	return cfg.Database.Path
}