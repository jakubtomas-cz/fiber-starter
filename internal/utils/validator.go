package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type StructValidator struct {
	validate *validator.Validate
}

func (v *StructValidator) Validate(out any) error {
	var result []error

	err := v.validate.Struct(out)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			result = append(
				result,
				fiber.NewError(
					fiber.StatusBadRequest,
					strings.Trim(
						fmt.Sprintf("%s: %v %v %v", VALIDATION_ERR_MSG, strings.ToLower(err.Field()), err.ActualTag(), err.Param()),
						" ",
					),
				),
			)
		}
	}

	if len(result) == 0 {
		return nil
	}

	return result[0]
}

func NewStructValidator() *StructValidator {
	return &StructValidator{
		validate: validator.New(),
	}
}
