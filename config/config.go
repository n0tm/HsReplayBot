package config

import (
	"encoding/json"
	"os"
	"path"
	"runtime"
)

const configFileName = "config.json"

var (
	_, baseFilename, _, _ = runtime.Caller(0)
	baseDirPath           = path.Dir(baseFilename)
	pathToConfig          = path.Join(baseDirPath, configFileName)
)

func Get() (*Config, error) {
	config := new(Config)
	err := config.ParseByPath(pathToConfig)
	return config, err
}

func (c *Config) ParseByPath(path string) error {
	file, _ := os.Open(path)
	decoder := json.NewDecoder(file)
	return decoder.Decode(&c)
}

type Config struct {
	Bot struct {
		AccessToken string `json:"access_token"`
		GroupId     int    `json:"group_id"`
	}
}
