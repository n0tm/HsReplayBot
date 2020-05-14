package config

import (
	"github.com/stretchr/testify/assert"
	"path"
	"testing"
)

var testPathToConfig = path.Join(baseDirPath, "config.test.json")

func TestParseByPathWhenPathExist(t *testing.T) {
	t.Parallel()
	config := new(Config)
	assert.Empty(t, config.Bot)
	assert.Nil(t, config.ParseByPath(testPathToConfig))

	assert.Equal(t, config.Bot.AccessToken, "::test token::")
	assert.Equal(t, 12345, config.Bot.GroupId)
}

func TestParseByPathWhenPathNotExist(t *testing.T) {
	t.Parallel()
	config := new(Config)
	assert.Empty(t, config.Bot)
	nonExistingPath := "::non existing path::"
	assert.EqualError(t, config.ParseByPath(nonExistingPath), "invalid argument")
	assert.Empty(t, config.Bot)
}
