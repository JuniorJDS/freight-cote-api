package endpoints

import (
	"freight-cote-api/schemas"
	"freight-cote-api/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Quote struct {
	errorsHandler utils.ErrorsHandler
	validator     *validator.Validate
}

func NewQuote() *Quote {
	return &Quote{
		errorsHandler: *utils.NewErrorsHandler(),
		validator:     validator.New(),
	}
}

func (q *Quote) Create(c *fiber.Ctx) error {
	quoteInputDTO := new(schemas.QuoteInputDTO)
	err := c.BodyParser(quoteInputDTO)
	if err != nil {
		return q.errorsHandler.InvalidBody(c, err)
	}

	err = q.validator.Struct(quoteInputDTO)
	if err != nil {
		return q.errorsHandler.InvalidBody(c, err)
	}
	return nil
}

func (q *Quote) Metrics(c *fiber.Ctx) error {
	return c.SendString("metricas")
}
