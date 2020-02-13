package main

import (
	"github.com/AlexFrazer/films-server/app"
	"github.com/AlexFrazer/films-server/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
