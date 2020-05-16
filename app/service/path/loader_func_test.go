package path_test

import (
	pathService "github.com/n0tm/HsReplayBot/app/service/path"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func TestLoader_Load(t *testing.T) {
	t.Parallel()

	pathLoader := new(pathService.Loader)
	service, err := pathLoader.Load()
	assert.IsType(t, &pathService.Path{}, service)
	assert.Nil(t, err)

	currentDir, _ := os.Getwd()
	assert.Equal(t, service.(*pathService.Path).RootPath, currentDir)
	assert.Equal(t, service.(*pathService.Path).ConfigPath, path.Join(currentDir, "app/service/config/config.json"))
}
