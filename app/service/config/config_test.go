package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

const pathToTestConfig = "config.test.json"

func TestParseByPathWhenPathExist(t *testing.T) {
	t.Parallel()

	baseDir, _ := os.Getwd()
	pathToConfig := path.Join(baseDir, pathToTestConfig)

	config := new(Config)
	assert.Empty(t, config.Bot)
	assert.Nil(t, config.parseByPath(pathToConfig))

	assert.Equal(t, config.Bot.AccessToken, "::test token::")
	assert.Equal(t, 12345, config.Bot.GroupId)
}

func TestParseByPathWhenPathNotExist(t *testing.T) {
	t.Parallel()

	config := new(Config)
	assert.Empty(t, config.Bot)

	nonExistingPath := "::non existing path::"
	assert.EqualError(t, config.parseByPath(nonExistingPath), "invalid argument")
	assert.Empty(t, config.Bot)
}

func TestLoad(t *testing.T) {
	t.Parallel()

	config := new(Config)
	assert.Empty(t, config.Bot)

	config.Load()

	anotherConfig := new(Config)
	assert.Empty(t, anotherConfig.Bot)
	assert.Nil(t, anotherConfig.parseByPath(pathToConfig))

	assert.Equal(t, anotherConfig, config)
}
