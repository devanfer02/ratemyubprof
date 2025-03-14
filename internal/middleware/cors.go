package middleware

import (
	"github.com/labstack/echo/v4"
	echoMw "github.com/labstack/echo/v4/middleware"
)

func CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.DELETE, echo.POST, echo.PUT, echo.OPTIONS},
	})(next)
}