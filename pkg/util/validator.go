package util

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func GetErrorValidationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	default:
		return fmt.Sprintf("%s is not valid", fe.Field())
	}
}
