package utils

import (
	"freight-cote-api/schemas/responses"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ErrorsHandler struct {
	Message string
}

func NewErrorsHandler() *ErrorsHandler {
	return &ErrorsHandler{}
}

func (e ErrorsHandler) InvalidBody(c *fiber.Ctx, err error) error {
	return c.Status(http.StatusBadRequest).
		JSON(responses.InfoResponse{Message: err.Error()})
}
