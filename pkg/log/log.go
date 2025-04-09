package logger

import (
	"time"

	"github.com/devanfer02/ratemyubprof/internal/infra/env"
	"go.uber.org/zap"
)

func NewLogger(env *env.Env) *zap.Logger {
	var (
		cfg zap.Config
		err error 
	)	

	if env.Logger.Type == "production" {
		cfg = zap.NewProductionConfig()
	} else if env.Logger.Type == "development" {
		cfg = zap.NewDevelopmentConfig()
	} else {
		panic("Invalid logger type")
	}

	cfg.OutputPaths = []string{
		"stdout",
		"stderr",
		"internal/logs/" + time.Now().Format("2006-01-02") + ".log",
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return logger.WithOptions(zap.AddStacktrace(zap.DPanicLevel))
}