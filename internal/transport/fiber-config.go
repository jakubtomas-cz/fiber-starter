package transport

import (
	"fiberstarter/internal/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/handlebars/v2"
)

func (t *Transport) FiberConfig() fiber.Config {
	return fiber.Config{
		AppName:                  utils.SERVER_NAME,
		Views:                    handlebars.New("./views", ".hbs"),
		ProxyHeader:              utils.ENV.ProxyHeader,
		StrictRouting:            true,
		CaseSensitive:            true,
		DisableKeepalive:         true,
		EnableSplittingOnParsers: true,
		EnableIPValidation:       true,
		ErrorHandler:             t.ErrorHandler,
		StructValidator:          utils.NewStructValidator(),
	}
}
