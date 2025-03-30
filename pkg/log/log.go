package logger

import (
	"github.com/devanfer02/ratemyubprof/internal/infra/env"
	"go.uber.org/zap"
)

func NewLogger(env *env.Env) *zap.Logger {
	var (
		logger *zap.Logger
		err error 
	)	

	if env.Logger.Type == "production" {
		logger, err = zap.NewProduction()
	} else if env.Logger.Type == "development" {
		logger, err = zap.NewDevelopment()	
	} else {
		panic("Invalid logger type")
	}

	if err != nil {
		panic(err)
	}

	return logger.WithOptions(zap.AddStacktrace(zap.DPanicLevel))
}