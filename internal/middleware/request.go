package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func RequestLogger(logger *zap.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:      true,
		LogStatus:   true,
		LogMethod:   true,
		LogLatency:  true,
		LogRemoteIP: true,
		LogError:    true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			logger.Info("Request",
				zap.String("remote_ip", values.RemoteIP),
				zap.String("method", values.Method),
				zap.Int("status", values.Status),
				zap.String("uri", values.URI),
				zap.String("latency", values.Latency.String()),
			)
			return nil
		},
	})
}