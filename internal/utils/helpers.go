package utils

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

type env struct {
	DatabaseURL string `validate:"required"`
	Port        string `validate:"omitempty,number"`
	ProxyHeader string `validate:"omitempty"`
}

var ENV *env

func Init() error {
	godotenv.Load()

	env := &env{
		DatabaseURL: os.Getenv(DATABASE_URL_KEY),
		Port:        os.Getenv(PORT_KEY),
		ProxyHeader: os.Getenv(PROXY_HEADER_KEY),
	}

	if env.Port == "" {
		env.Port = strconv.Itoa(DEFAULT_PORT)
	}

	validator := NewStructValidator()
	if err := validator.Validate(env); err != nil {
		return err
	}

	ENV = env

	return nil
}

func GetPort() string {
	return fmt.Sprintf(":%s", ENV.Port)
}

func Trim(src string) string {
	return strings.Trim(src, " ")
}

func IsTrue(src string) bool {
	return src == "true" || src == "1"
}

// Useful for validating query params and such, where user input is always gathered
// as string such as quety params and .env properties
func ConvertStringToInt(key string, src string, min, max int) (int, error) {
	num, err := strconv.Atoi(src)
	if err != nil {
		return 0, NewValidationError(fmt.Sprintf("%s: %v %v", VALIDATION_ERR_MSG, key, "number"))
	}

	if min != -1 && num < min {
		return 0, NewValidationError(fmt.Sprintf("%s: %v %v %v", VALIDATION_ERR_MSG, key, "min", min))
	}

	if max != -1 && num > max {
		return 0, NewValidationError(fmt.Sprintf("%s: %v %v %v", VALIDATION_ERR_MSG, key, "max", max))
	}

	return num, nil
}

func NewValidationError(message string) *fiber.Error {
	return fiber.NewError(http.StatusBadRequest, message)
}

func OrderedJSON(keys []string, values []any) string {
	var start strings.Builder
	start.WriteString("{")
	end := "}\n"
	length := len(keys)

	for index, key := range keys {
		fmt.Fprintf(&start, "\"%s\": \"%s\"", key, values[index])
		if index+1 < length {
			fmt.Fprint(&start, ", ")
		}
	}

	return start.String() + end
}
