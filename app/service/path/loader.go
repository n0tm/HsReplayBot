package path

import (
	"os"
	"path"
)

type Loader struct{}

func (l Loader) Load() (interface{}, error) {
	service := new(Path)

	rootPath, err := os.Getwd()

	service.RootPath = rootPath
	service.ConfigPath = path.Join(rootPath, "app/service/config/config.json")

	return service, err
}
