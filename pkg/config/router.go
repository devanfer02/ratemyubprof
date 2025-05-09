package config

import (
	"net/http"

	apperr "github.com/devanfer02/ratemyubprof/pkg/http/errors"
	"github.com/devanfer02/ratemyubprof/pkg/util"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

var (
	FetchLimiter = middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(5000)))
	PostLimiter  = middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(500)))
)

func NewRouter() *echo.Echo {
	router := echo.New()

	router.JSONSerializer = newSonicJSONSerializer()
	router.HTTPErrorHandler = errHandler()
	router.Use(middleware.Recover())

	return router
}

func errHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if ce, ok := err.(*apperr.AppError); ok {
			c.JSON(ce.Code, echo.Map{
				"message": ce.Message,
				"error":   ce.Error(),
			})
			return
		}

		if ve, ok := err.(validator.ValidationErrors); ok {
			out := make(map[string]string)
			for _, e := range ve {
				out[e.Field()] = util.GetErrorValidationMessage(e)
			}
			c.JSON(http.StatusBadRequest, echo.Map{
				"message": "validation error",
				"error":   out,
			})
			return
		}

		if ee, ok := err.(*echo.HTTPError); ok {
			c.JSON(ee.Code, echo.Map{
				"message": ee.Message,
				"error":   ee.Message,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
			"error":   err.Error(),
		})
	}
}
