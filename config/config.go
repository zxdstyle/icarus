package config

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
)

type Config struct {
}

func New() (*Config, error) {
	config.AddDriver(yamlv3.Driver)
	if err := config.LoadFiles("config.yaml"); err != nil {
		return nil, err
	}
	return &Config{}, nil
}

func (Config) Scan(key string, pointer any) error {
	return config.BindStruct(key, pointer)
}
