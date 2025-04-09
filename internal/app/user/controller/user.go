package controller

import (
	"context"
	"net/http"
	
	"github.com/devanfer02/ratemyubprof/internal/app/user/contracts"
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/pkg/http/response"
	"github.com/labstack/echo/v4"
)

func (c *UserController) Register(ectx echo.Context) error {
	ctx, cancel := context.WithTimeout(ectx.Request().Context(), c.timeout)
	defer cancel()

	var (
		req dto.UserRegisterRequest
		resp *response.Response
	)

	if err := ectx.Bind(&req); err != nil {
		return err 
	}

	if err := c.validator.Struct(req); err != nil {
		return err 
	}

	err := c.userSvc.RegisterUser(ctx, &req)
	if err != nil {
		return err  
	}

	resp = response.New(
		"Successfully register user",
		nil,
		nil,
	)

	select {
	case <-ctx.Done():
		return contracts.ErrRequestTimeout
	default:
		return ectx.JSON(http.StatusCreated, resp)
	}
}