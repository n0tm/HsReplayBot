package main

import (
	"github.com/n0tm/HsReplayBot/app"
	"log"
)

func main() {
	application := app.Init()
	log.Print(application.Path.RootPath)
}
