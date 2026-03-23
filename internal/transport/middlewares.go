package transport

import (
	"fiberstarter/internal/utils"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/idempotency"
	"github.com/gofiber/fiber/v3/middleware/logger"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

var recoverConfig = recoverer.Config{
	EnableStackTrace:  true,
	StackTraceHandler: recoverer.ConfigDefault.StackTraceHandler,
}

var loggerConfig = logger.Config{
	DisableColors: true,
	TimeFormat:    time.RFC3339,
	CustomTags: map[string]logger.LogFunc{
		"id": func(output logger.Buffer, c fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
			return output.WriteString(requestid.FromContext(c))
		},
	},
	Format: utils.OrderedJSON(
		[]string{"time", "method", "path", "status", "latency", "ip", "id", "host", "error"},
		[]any{"${time}", "${method}", "${path}", "${status}", "${latency}", "${ip}", "${id}", "${host}", "${error}"},
	),
}

func (t *Transport) SetupMiddlewares() {
	t.App.Use(recoverer.New(recoverConfig))
	t.App.Use(requestid.New())
	t.App.Use(logger.New(loggerConfig))
	t.App.Use(compress.New())
	t.App.Use(idempotency.New())
	t.App.Use(cors.New())
}
