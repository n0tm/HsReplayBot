package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	t.Parallel()

	path := new(Path)
	assert.Empty(t, path.RootPath)

	path.Load()

	currentDir, _ := os.Getwd()
	assert.Equal(t, path.RootPath, currentDir)
}
