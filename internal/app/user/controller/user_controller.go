package controller

import (
	"context"
	"time"

	"github.com/devanfer02/ratemyubprof/internal/app/user/contracts"
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/pkg/http/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userSvc contracts.UserService
	validator *validator.Validate
	timeout time.Duration
}

func NewUserController(userSvc contracts.UserService, validator *validator.Validate) *UserController {
	return &UserController{
		userSvc: userSvc,
		timeout: 10 * time.Second,
		validator: validator,
	}
}

func (c *UserController) Mount(r *echo.Group) {
	userR := r.Group("/users")

	userR.POST("/register", c.Register)	
	userR.POST("/login", c.Login)
}

func (c *UserController) Register(ectx echo.Context) error {
	ctx, cancel := context.WithTimeout(ectx.Request().Context(), c.timeout)
	defer cancel()

	var (
		responeChan = make(chan response.Response)
		errChan = make(chan error)
	)

	go func () {
		var (
			req dto.UserRegisterRequest
		)

		if err := ectx.Bind(&req); err != nil {
			errChan <- err
			return
		}

		if err := c.validator.Struct(req); err != nil {
			errChan <- err
			return
		}

		err := c.userSvc.RegisterUser(ctx, &req)
		if err != nil {
			errChan <- err 
			return
		}

		responeChan <- *response.New(
			"Successfully register user",
			nil,
			nil,
		)
	}()

	select {
	case <-ctx.Done():
		return contracts.ErrRequestTimeout
	case err := <- errChan:
		return err 
	case resp := <- responeChan:
		return ectx.JSON(resp.Code, resp)
	}
}

func (c *UserController) Login(ectx echo.Context) error {
	ctx, cancel := context.WithTimeout(ectx.Request().Context(), c.timeout)
	defer cancel()

	var (
		responeChan = make(chan response.Response)
		errChan = make(chan error)
	)

	go func () {
		var (
			req dto.UserLoginRequest
		)

		if err := ectx.Bind(&req); err != nil {
			errChan <- err
			return
		}

		if err := c.validator.Struct(req); err != nil {
			errChan <- err
			return
		}

		token, err := c.userSvc.LoginUser(ctx, &req)
		if err != nil {
			errChan <- err 
			return
		}

		responeChan <- *response.New(
			"Successfully register user",
			echo.Map{
				"token": token,
			},
			nil,
		)
	}()

	select {
	case <-ctx.Done():
		return contracts.ErrRequestTimeout
	case err := <- errChan:
		return err 
	case resp := <- responeChan:
		return ectx.JSON(resp.Code, resp)
	}
}