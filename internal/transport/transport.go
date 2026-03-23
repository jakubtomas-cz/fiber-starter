package transport

import (
	"fiberstarter/internal/handlers"
	"fiberstarter/internal/pages"
	"fiberstarter/internal/repository"
	"fiberstarter/internal/service"
	"fiberstarter/internal/utils"

	"github.com/gofiber/fiber/v3"
)

type Transport struct {
	App      *fiber.App
	Handlers *handlers.Handlers
	Pages    *pages.Pages
}

func New(service *service.Service, repository *repository.Repository) error {
	if service == nil {
		return utils.ErrNoService
	}

	if repository == nil {
		return utils.ErrNoRepository
	}

	transport := &Transport{}
	transport.App = fiber.New(transport.FiberConfig())
	transport.Handlers = handlers.New(service, repository)
	transport.Pages = pages.New(service, repository)

	transport.SetStartupMessage()
	transport.SetupMiddlewares()
	transport.SetupEndpoints()

	return transport.App.Listen(utils.GetPort(), fiber.ListenConfig{
		ListenerNetwork: "tcp4",
	})
}
