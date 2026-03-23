package middlewares

import (
	"errors"
	"fiberstarter/internal/utils"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
)

func Timeout(handler fiber.Handler, timeout time.Duration) fiber.Handler {
	if timeout <= 0 {
		return handler
	}

	return func(ctx fiber.Ctx) error {
		panicChannel := make(chan error, 1)
		responseChannel := make(chan error, 1)

		go func() {
			defer func() {
				err := recover()
				if err != nil {
					fmt.Print(utils.OrderedJSON(
						[]string{"panic"},
						[]any{err},
					))

					str, strOK := err.(string)
					if strOK {
						err := errors.New(str)
						panicChannel <- err
					}

					if !strOK {
						err, errOK := err.(error)
						if errOK {
							panicChannel <- err
						}
					}
				}
			}()

			responseChannel <- handler(ctx)
		}()

		select {
		case err, _ := <-responseChannel:
			return err
		case err, _ := <-panicChannel:
			return err
		case <-time.After(timeout):
			return fiber.ErrRequestTimeout
		}
	}
}
