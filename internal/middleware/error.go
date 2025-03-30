package middleware

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func ErrLogger(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			
			err := next(c)

			if err != nil {

				logger.
				Error("Error",
					zap.String("Path", c.Path()),
					zap.String("Method", c.Request().Method),
					zap.Any("Query Params", c.QueryParams()),
					zap.Strings("Path Params", c.ParamValues()),
					zap.String("Error", err.Error()),
				)
			}
			return err
		}
	}
}
