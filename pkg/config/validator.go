package config

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)


func NewValidator() *validator.Validate {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	validations := map[string]func(validator.FieldLevel) bool{
		
	}

	for tag, method := range validations {
		if err := v.RegisterValidation(tag, method); err != nil {
			zap.S().Errorf("failed to register %s validation: %v", tag, err)
		}
	}

	return v
}