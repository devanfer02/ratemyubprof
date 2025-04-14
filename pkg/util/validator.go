package util

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func GetErrorValidationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "alphanum":
		return fmt.Sprintf("%s must be alphanumeric", fe.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s", fe.Field(), fe.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s", fe.Field(), fe.Param())
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	case "reactionType":
		return fmt.Sprintf("%s must be one of the following: like, dislike", fe.Field())
	default:
		return fmt.Sprintf("%s is not valid", fe.Field())
	}
}
