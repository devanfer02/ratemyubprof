package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *UserHandler) Login(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message":"Hello world!",
	})
}