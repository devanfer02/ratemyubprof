package config

import (
	"reflect"
	"strings"

	"github.com/devanfer02/ratemyubprof/internal/entity"
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
		"reactionType": ValidateReactionType,	
	}

	for tag, method := range validations {
		if err := v.RegisterValidation(tag, method); err != nil {
			zap.S().Errorf("failed to register %s validation: %v", tag, err)
		}
	}

	return v
}

func ValidateReactionType(fe validator.FieldLevel) bool {
	reactionType := fe.Field().String()
	reaction := entity.ToReactionType(reactionType)

	if reaction == 0 {
		return false
	}

	return true
}