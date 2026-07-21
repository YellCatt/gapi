package config

import (
	"log"
	"path/filepath"

	"os"

	"gopkg.in/yaml.v3"
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
	configPath := "config/config.yaml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Println("config file not found, creating default config...")
		if err := createDefaultConfig(configPath); err != nil {
			log.Fatalf("failed to create default config: %v", err)
		}
	}

	file, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		log.Fatalf("failed to parse config file: %v", err)
	}
}

func createDefaultConfig(path string) error {
	defaultCfg := Config{
		Server: ServerConfig{
			Port: 8084,
		},
		Database: DatabaseConfig{
			Path: "./data.db",
		},
		Log: LogConfig{
			Path:  "./logs",
			Level: "info",
		},
	}

	data, err := yaml.Marshal(&defaultCfg)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
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
