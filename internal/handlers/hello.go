package handlers

import "github.com/gofiber/fiber/v3"

func (h *Handlers) HelloHandler(c fiber.Ctx) error {
	return c.SendString("Hello World")
}
