package app

import (
	routes "freight-cote-api/routes/api"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func MakeApp() *fiber.App {
	log.SetFlags(log.Ltime | log.Lshortfile)

	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())
	// app.Use(middlewares.SecurityNew(middlewares.SecurityConfig{
	//	ContentSecurityPolicy: "default-src 'self'; base-uri 'self'; block-all-mixed-content; style-src 'self' 'unsafe-inline' https://fonts.googleapis.com; img-src 'self' data: validator.swagger.io; script-src 'self' 'unsafe-inline'; form-action 'self'; frame-ancestors 'self'; font-src 'self' https://fonts.gstatic.com",
	//	CacheControlMaxAge:    600,
	//	XFrameOptions:         "DENY",
	// }))

	// docs.SwaggerInfo.Host = configs.GetSettings()["HOST"]
	// app.Get("/docs/*", swagger.HandlerDefault) // default
	// app.Use(middlewares.MiddlewareAuthenticator)

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
