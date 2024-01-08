package main

import (
	api "freight-cote-api/api/app"
	"freight-cote-api/configs"
)

var settings = configs.GetSettings()

// @title freight cote - API
// @version 1.0.0
// @description API for freight cote operations
// @BasePath /api/v1
func main() {
	app := api.MakeApp()

	portListen := ":" + settings["PORT"]
	err := app.Listen(portListen)
	if err != nil {
		panic(err)
	}
}
