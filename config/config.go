package config

import (
	"log"
	"path/filepath"

	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Log      LogConfig      `yaml:"log"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type DatabaseConfig struct {
	Path string `yaml:"path"`
}

type LogConfig struct {
	Path  string `yaml:"path"`
	Level string `yaml:"level"`
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

func GetLogPath() string {
	return cfg.Log.Path
}

func GetLogLevel() string {
	return cfg.Log.Level
}

func InitDirectories() error {
	if err := os.MkdirAll(cfg.Log.Path, 0755); err != nil {
		return err
	}

	dbDir := filepath.Dir(cfg.Database.Path)
	if dbDir != "." && dbDir != "" {
		if err := os.MkdirAll(dbDir, 0755); err != nil {
			return err
		}
	}

	return nil
}