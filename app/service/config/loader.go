package config

import (
	"encoding/json"
	"os"
)

type Loader struct {
	ConfigPath string
}

func (l Loader) Load() (interface{}, error) {
	service := new(Config)
	err := parseConfigByPath(l.ConfigPath, service)
	return service, err
}

func parseConfigByPath(path string, config *Config) error {
	file, _ := os.Open(path)
	decoder := json.NewDecoder(file)
	return decoder.Decode(&config)
}
