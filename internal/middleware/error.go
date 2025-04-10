package middleware

import (
	apperr "github.com/devanfer02/ratemyubprof/pkg/http/errors"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func ErrLogger(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			
			err := next(c)

			if err != nil {

				if appErr, ok := err.(*apperr.AppError); ok {
					logger.Error("Error",
						zap.String("Path", c.Path()),
						zap.String("Method", c.Request().Method),
						zap.Any("Query Params", c.QueryParams()),
						zap.Strings("Path Params", c.ParamValues()),
						zap.String("Error", err.Error()),
						zap.String("File", appErr.File),
						zap.String("Action", appErr.Action),
						zap.Int("Line", appErr.Line),
					)
				} else {
					logger.Error("Error",
						zap.String("Path", c.Path()),
						zap.String("Method", c.Request().Method),
						zap.Any("Query Params", c.QueryParams()),
						zap.Strings("Path Params", c.ParamValues()),
						zap.String("Error", err.Error()),
					)
				}
			}
			return err
		}
	}
}
