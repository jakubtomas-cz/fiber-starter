package transport

import (
	"fiberstarter/internal/utils"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func (t *Transport) SetStartupMessage() {
	t.App.Hooks().OnPreStartupMessage(func(sm *fiber.PreStartupMessageData) error {
		sm.PreventDefault = true
		return nil
	})

	t.App.Hooks().OnPostStartupMessage(func(sm *fiber.PostStartupMessageData) error {
		if !sm.Disabled && !sm.IsChild {
			port := utils.GetPort()

			fmt.Print(utils.OrderedJSON(
				[]string{"name", "status", "pid", "port", "url"},
				[]any{utils.SERVER_NAME, "running", strconv.Itoa(os.Getpid()), strings.Replace(port, ":", "", 1), "http://localhost" + port},
			))
		}

		return nil
	})
}
