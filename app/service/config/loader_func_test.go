package config_test

import (
	"github.com/n0tm/HsReplayBot/app/service/config"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func TestLoader_LoadWhenRootPathExist(t *testing.T) {
	t.Parallel()

	currentDirectoryPath, _ := os.Getwd()
	pathToConfig := path.Join(currentDirectoryPath, "config.test.json")

	configLoader := config.Loader{ConfigPath: pathToConfig}
	service, err := configLoader.Load()

	assert.IsType(t, &config.Config{}, service)
	assert.Nil(t, err)
}

func TestLoader_LoadWhenRootPathNotExist(t *testing.T) {
	t.Parallel()

	configLoader := config.Loader{ConfigPath: "::not existing root path"}
	service, err := configLoader.Load()

	assert.IsType(t, &config.Config{}, service)
	assert.Error(t, err, "invalid argument")
}
