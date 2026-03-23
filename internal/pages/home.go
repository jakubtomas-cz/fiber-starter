package pages

import (
	"fiberstarter/internal/utils"

	"github.com/gofiber/fiber/v3"
)

func (p *Pages) HomePage(c fiber.Ctx) error {
	return c.Render("index", fiber.Map{"Title": utils.SERVER_NAME})
}
