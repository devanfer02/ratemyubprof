package middleware

import (
	"errors"

	"github.com/devanfer02/presentia-api/pkg/response"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func ErrLogger(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)

			if err != nil {
				var custom *response.Response
	
				if errors.As(err, &custom) {
					logger.Warn("ERR",
						zap.String("Error", custom.Err.Error()),
						zap.String("File Location", custom.Location),
					)
					
				} else {
					logger.Info("ERR",
						zap.String("Error", err.Error()),
					)
				}
			}
			return err 
		}
	}
}