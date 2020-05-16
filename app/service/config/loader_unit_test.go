package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func TestLoader_parseConfigByPathWhenPathExist(t *testing.T) {
	t.Parallel()

	baseDir, _ := os.Getwd()
	pathToConfig := path.Join(baseDir, "config.test.json")

	config := new(Config)
	assert.Empty(t, config.Bot)
	assert.Nil(t, parseConfigByPath(pathToConfig, config))

	assert.Equal(t, config.Bot.AccessToken, "::test token::")
	assert.Equal(t, 12345, config.Bot.GroupId)
}

func TestLoader_parseConfigByPathWhenPathNotExist(t *testing.T) {
	t.Parallel()

	config := new(Config)
	assert.Empty(t, config.Bot)

	nonExistingPath := "::non existing path::"
	assert.EqualError(t, parseConfigByPath(nonExistingPath, config), "invalid argument")
	assert.Empty(t, config.Bot)
}
