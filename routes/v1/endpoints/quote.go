package endpoints

import (
	"freight-cote-api/schemas"
	"freight-cote-api/services"
	"freight-cote-api/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Quote struct {
	errorsHandler utils.ErrorsHandler
	validator     *validator.Validate
	quoteService  services.QuoteService
}

func NewQuote() *Quote {
	return &Quote{
		errorsHandler: *utils.NewErrorsHandler(),
		validator:     validator.New(),
		quoteService:  *services.NewQuoteService(),
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

	response, _ := q.quoteService.QuoteFreightV3(*quoteInputDTO)
	return c.Status(http.StatusOK).JSON(response)
}

func (q *Quote) Metrics(c *fiber.Ctx) error {
	return c.SendString("metricas")
}
