package middleware

import (
	"net/http"
	"strings"

	"github.com/devanfer02/ratemyubprof/pkg/config"
	"github.com/labstack/echo/v4"
)

func (m *Middleware) Authenticate() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			header := c.Request().Header.Get("Authorization")

			if header == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing authorization header")
			}

			parts := strings.Split(header, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid authorization format")
			}

			userID, err := m.jwtHandler.ValidateToken(parts[1], config.AccessToken)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired token")
			}

			c.Set("userId", userID)

			return next(c)
		}
	}	
}