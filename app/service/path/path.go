package config

import (
	"os"
)

type Path struct {
	RootPath string
}

// TODO: Root путь задаётся неверно, нужно либо выпилить сервис path, либо сделать его корректным. (Root задаётся по текущей директории)
func (p *Path) Load() {
	rootPath, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	p.RootPath = rootPath
}
