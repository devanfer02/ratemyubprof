package middleware

import (
	"github.com/devanfer02/ratemyubprof/internal/infra/env"
	"github.com/labstack/echo/v4"
)

func ApiKey(env *env.Env) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			apiKey := c.Request().Header.Get(env.App.ApiKeyHeader)
			if apiKey != env.App.ApiKey {
				return echo.ErrUnauthorized
			}
			return next(c)
		}
	}
}