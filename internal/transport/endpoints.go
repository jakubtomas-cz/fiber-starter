package transport

import (
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/limiter"
)

func (t *Transport) SetupEndpoints() {
	t.HealthEndpoints()
	t.PagesEndpoints()
}

func (t *Transport) HealthEndpoints() {
	t.App.Get("/health", limiter.New(), healthcheck.New())
}

func (t *Transport) PagesEndpoints() {
	t.App.Get("/", t.Pages.HomePage)
}

func (t *Transport) APIEndpoints() {
	t.App.Get("/api/hello", t.Handlers.HelloHandler)
}
