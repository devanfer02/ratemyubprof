package logger

import (
	"fmt"

	"github.com/devanfer02/ratemyubprof/internal/infra/env"
	"go.uber.org/zap"
)

func NewLogger(env *env.Env) *zap.Logger {
	var (
		cfg zap.Config
		err error 
	)
	
	switch env.Logger.Type {
	case "production":
		cfg = zap.NewProductionConfig()
	case "development":
		cfg = zap.NewDevelopmentConfig()
	default:
		panic("Invalid logger type")
	}

	cfg.OutputPaths = []string{
		"stdout",
		fmt.Sprintf("internal/logs/app-%s.log", env.Logger.Type),
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return logger.WithOptions(zap.AddStacktrace(zap.DPanicLevel))
}