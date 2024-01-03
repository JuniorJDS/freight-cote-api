package endpoints

import "github.com/gofiber/fiber/v2"

type HelloWorld struct {
}

func NewHelloWorld() *HelloWorld {
	return &HelloWorld{}
}

func (h *HelloWorld) GetHelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
