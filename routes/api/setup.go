package routes

import (
	"freight-cote-api/configs"
	"freight-cote-api/routes/v1/endpoints"

	"github.com/gofiber/fiber/v2"
)

var settings = configs.GetSettings()

func RoutesV1(app *fiber.App) {
	v1 := app.Group(settings["API_V1"])

	// middlewares := middlewares.NewMiddlewarePagination()
	// v1.Use(middlewares.ValidatePaginationParameters)

	// hello-world
	hello := endpoints.NewHelloWorld()
	v1.Get("hello-world", hello.GetHelloWorld)
}
