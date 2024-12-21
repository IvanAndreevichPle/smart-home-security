// internal/config/config.go
package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Camera   CameraConfig   `yaml:"camera"`
	Detector DetectorConfig `yaml:"detector"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cfg := &Config{}
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
