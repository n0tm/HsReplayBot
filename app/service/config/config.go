package config

import (
	"encoding/json"
	"os"
	"path"
)

const pathToConfig = "config.json"

func (c *Config) parseByPath(path string) error {
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

func (c *Config) Load() {
	baseDir, _ := os.Getwd()
	pathToConfig := path.Join(baseDir, pathToConfig)

	err := c.parseByPath(pathToConfig)
	if err != nil {
		panic(err.Error())
	}
}
