package middleware

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func RequestLogger(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			err := next(c)

			logger.Info("Request",
				zap.String("Remote IP", c.RealIP()),
				zap.String("Host", req.Host),
				zap.String("URI", req.RequestURI),
				zap.String("Method", req.Method),
				zap.Int("Status", res.Status),
				zap.Int64("Size", res.Size),
				zap.String("User Agent", req.UserAgent()),
			)

			return err
		}
	}
}