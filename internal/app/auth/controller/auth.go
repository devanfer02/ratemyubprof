package controller

import (
	"context"
	"net/http"

	"github.com/devanfer02/ratemyubprof/internal/app/auth/contracts"
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/pkg/http/response"
	"github.com/labstack/echo/v4"
)


func (c *AuthController) Login(ectx echo.Context) error {
	ctx, cancel := context.WithTimeout(ectx.Request().Context(), c.timeout)
	defer cancel()

	var (
		req dto.UserLoginRequest
		resp *response.Response
	)

	if err := ectx.Bind(&req); err != nil {
		return err 
	}

	if err := c.validator.Struct(req); err != nil {
		return err 
	}

	token, err := c.userSvc.LoginUser(ctx, &req)
	if err != nil {
		return err  
	}

	resp = response.New(
		"Successfully register user",
		token,
		nil,
	)

	select {
	case <-ctx.Done():
		return contracts.ErrRequestTimeout
	default:
		return ectx.JSON(http.StatusOK, resp)
	}
}

func (c *AuthController) RefreshToken(ectx echo.Context) error {
	ctx, cancel := context.WithTimeout(ectx.Request().Context(), c.timeout)
	defer cancel()

	var (
		req dto.RefreshATRequest
		resp *response.Response
	)

	if err := ectx.Bind(&req); err != nil {
		return err  			
	}

	if err := c.validator.Struct(&req); err != nil {
		return err  
	}

	token, err := c.userSvc.RefreshAccessToken(ctx, req)

	if err != nil {
		return err  
	}

	resp = response.New(
		"Successfully refresh access token",
		token,
		nil,
	)

	select {
	case <- ctx.Done():
		return contracts.ErrRequestTimeout
	default:
		return ectx.JSON(http.StatusCreated, resp) 
	}

}