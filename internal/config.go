package internal

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config is a struct that holds the configuration for the server and REST gateway
type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`

	RestGateway struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"rest_gateway"`

	Database struct {
		Host       string            `yaml:"host"`
		Port       int               `yaml:"port"`
		User       string            `yaml:"user"`
		Password   string            `yaml:"password"`
		Properties map[string]string `yaml:"properties"`
	} `yaml:"database"`
}

// NewConfigFromPath reads a YAML file from the given path and returns a Config struct
func NewConfigFromPath(path string) (*Config, error) {
	config := &Config{}
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
