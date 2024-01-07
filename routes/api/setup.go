package routes

import (
	"freight-cote-api/configs"
	"freight-cote-api/routes/v1/endpoints"

	"github.com/gofiber/fiber/v2"
)

var settings = configs.GetSettings()

func RoutesV1(app *fiber.App) {
	v1 := app.Group(settings["API_V1"])

	// Quote
	quote := endpoints.NewQuote()
	v1.Post("/quote", quote.Create)
	v1.Get("/quote/metrics", quote.Metrics)
}
