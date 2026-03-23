package transport

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	fiberutils "github.com/gofiber/utils/v2"
)

func (t *Transport) ErrorHandler(c fiber.Ctx, err error) error {
	id := requestid.FromContext(c)
	status := fiber.StatusInternalServerError
	statusMessage := fiberutils.StatusMessage(status)
	message := fiber.ErrInternalServerError.Message

	var e *fiber.Error
	isFiberErr := errors.As(err, &e)

	if isFiberErr {
		status = e.Code
		statusMessage = fiberutils.StatusMessage(status)
		message = e.Error()
	}

	return c.Status(status).JSON(fiber.Map{
		"status":  status,
		"error":   statusMessage,
		"message": message,
		"id":      id,
	})
}
