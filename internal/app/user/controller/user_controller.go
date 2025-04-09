package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/devanfer02/ratemyubprof/internal/app/user/contracts"
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/middleware"
	"github.com/devanfer02/ratemyubprof/pkg/http/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userSvc contracts.UserService
	validator *validator.Validate
	mdlwr *middleware.Middleware
	timeout time.Duration
}

func NewUserController(userSvc contracts.UserService, validator *validator.Validate, mdlwr *middleware.Middleware) *UserController {
	return &UserController{
		userSvc: userSvc,
		mdlwr: mdlwr,
		timeout: 10 * time.Second,
		validator: validator,
	}
}

func (c *UserController) Mount(r *echo.Group) {
	userR := r.Group("/users")

	userR.POST("/register", c.Register)	
	userR.POST("/login", c.Login)
	userR.POST("/refresh" ,c.RefreshToken)
}

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

func (c *UserController) Login(ectx echo.Context) error {
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

func (c *UserController) RefreshToken(ectx echo.Context) error {
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