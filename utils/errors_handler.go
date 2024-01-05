package utils

import (
	r "freight-cote-api/schemas/response"
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
		JSON(r.InfoResponse{Message: err.Error()})
}
