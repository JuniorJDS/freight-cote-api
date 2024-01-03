package main

import (
	api "freight-cote-api/api/app"
	"freight-cote-api/configs"
)

var settings = configs.GetSettings()

func main() {
	app := api.MakeApp()

	portListen := ":" + settings["PORT"]
	err := app.Listen(portListen)
	if err != nil {
		panic(err)
	}
}
