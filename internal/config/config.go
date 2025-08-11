package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string `yaml:"address" env-required:"true"`
}
type PG struct {
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	DBName   string `yaml:"dbname" env-required:"true"`
	Port     string `yaml:"port" env-required:"true"`
	Host     string `yaml:"host" env-required:"true"`
}
type Database struct {
	PG PG
}
type Key struct {
	HMACKey string `yaml:"hmac_key" env-required:"true"`
}
type Config struct {
	Env        string `yaml:"env" env:"ENV" env-required:"true"`
	HTTPServer `yaml:"http_server" env-required:"true"`
	Database   `yaml:"database" env-required:"true"`
	Key        `yaml:"key" env-required:"true"`
}

func LoadConfig() (*Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "", "path to config file")
		flag.Parse()
		configPath = *flags
		if configPath == "" {
			return nil, fmt.Errorf("config file not specified")
		}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file not found at %v", configPath)
	}
	var config Config
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		return nil, fmt.Errorf("failed to read config file at %v: %w", configPath, err)
	}
	return &config, nil
}