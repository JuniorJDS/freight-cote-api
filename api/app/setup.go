package app

import (
	"freight-cote-api/docs"
	routes "freight-cote-api/routes/api"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

func MakeApp() *fiber.App {
	log.SetFlags(log.Ltime | log.Lshortfile)

	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())

	docs.SwaggerInfo.Host = "localhost:5000"
	app.Get("/docs/*", swagger.HandlerDefault) // default

	allowOriginsCors := ""

	app.Use(cors.New(cors.Config{
		AllowOrigins:     allowOriginsCors,
		AllowHeaders:     "*",
		AllowMethods:     "*",
		AllowCredentials: true,
		ExposeHeaders:    "Content-Disposition",
	}))

	routes.RoutesV1(app)

	return app
}
