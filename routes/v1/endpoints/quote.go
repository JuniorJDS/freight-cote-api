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

// Create godoc
// @Summary 	Create quote
// @Description Route to receive input data and generate a fictional quote using the Fast Freight API.
// @Tags 		Quote
// @Accept      application/json
// @Produce 	application/json
// @Param 		quote body 	   input.Quote  true "Quote Datas"
// @Success 	200   {object} response.QuoteResponse
// @Failure 	400   {object} response.InfoResponse
// @Failure 	500   {object} response.InfoResponse
// @Router 		/indicators [post]
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

// Metrics godoc
// @Summary 	Get Metrics
// @Description Gets all available metrics.
// @Tags 		Quote
// @Produce 	application/json
// @Param 		last_quotes   		   query    int   				       false "quantidade de cotações (ordem decrescente)"
// @Success 	200   {object} response.Metrics
// @Failure 	500   {object} response.InfoResponse
// @Router 		/quote/metrics [get]
func (q *Quote) Metrics(c *fiber.Ctx) error {
	lastQuotes := c.QueryInt("last_quotes", -1)
	result, err := q.quoteService.GetMetrics(int64(lastQuotes))
	if err != nil {
		return q.errorsHandler.InternalServerError(c, err)
	}
	return c.Status(http.StatusOK).JSON(result)
}
