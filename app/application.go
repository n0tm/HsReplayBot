package app

import (
	configService "github.com/n0tm/HsReplayBot/app/service/config"
	pathService "github.com/n0tm/HsReplayBot/app/service/path"
)

// TODO: Написать на это тест. Сделать нормальные моки
func Init() *application {
	application := new(application)

	application.Services.load()

	return application
}

type application struct {
	Services Services
}

type Services struct {
	Config *configService.Config
	Path   *pathService.Path
}

func (s *Services) load() {
	pathLoader := new(pathService.Loader)

	loadedPathService, err := pathLoader.Load()
	if err != nil {
		panic(err.Error())
	}

	s.Path = loadedPathService.(*pathService.Path)

	configLoader := configService.Loader{ConfigPath: s.Path.ConfigPath}
	loadedConfigService, err := configLoader.Load()
	if err != nil {
		panic(err.Error())
	}

	s.Config = loadedConfigService.(*configService.Config)
}
