package config

import (
	"github.com/gokafka/service"
	"gopkg.in/yaml.v3"

	"os"
)

var (
	Conf *Config
)

type SectionHTTP struct {
	Port int `yaml:"port"`
}

// Config is config structure.
type Config struct {
	Service service.Config `yaml:"service"`
	HTTP    SectionHTTP    `yaml:"http"`
}

// LoadConfig load config from file
func LoadConfig(confPath string) (*Config, error) {
	configFile, err := os.ReadFile(confPath)
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
