package endpoints

import (
	"freight-cote-api/schemas/input"
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
	quoteInput := new(input.Quote)
	err := c.BodyParser(quoteInput)
	if err != nil {
		return q.errorsHandler.InvalidBody(c, err)
	}

	err = q.validator.Struct(quoteInput)
	if err != nil {
		return q.errorsHandler.InvalidBody(c, err)
	}

	response, err := q.quoteService.Create(*quoteInput)
	if err != nil {
		return q.errorsHandler.InternalServerError(c, err)
	}
	return c.Status(http.StatusOK).JSON(response)
}

func (q *Quote) Metrics(c *fiber.Ctx) error {
	lastQuotes := c.QueryInt("last_quotes", -1)
	result, err := q.quoteService.GetMetrics(int64(lastQuotes))
	if err != nil {
		return q.errorsHandler.InternalServerError(c, err)
	}
	return c.Status(http.StatusOK).JSON(result)
}
