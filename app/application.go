package app

import (
	configService "github.com/n0tm/HsReplayBot/app/service/config"
	pathService "github.com/n0tm/HsReplayBot/app/service/path"
)

// TODO: Написать на это тест. Сделать нормальные моки
func Init() *application {
	application := new(application)

	application.Path.Load()
	application.Config.Load()

	return application
}

type application struct {
	Config configService.Config
	Path   pathService.Path
}
