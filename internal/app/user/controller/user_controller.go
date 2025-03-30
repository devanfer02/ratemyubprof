package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/devanfer02/ratemyubprof/internal/app/user/contracts"
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/pkg/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userSvc contracts.UserService
	timeout time.Duration
}

func NewUserController(userSvc contracts.UserService, validator *validator.Validate) *UserController {
	return &UserController{
		userSvc: userSvc,
		timeout: 10 * time.Second,
	}
}

func (c *UserController) Mount(r *echo.Group) {
	userR := r.Group("/users")

	userR.POST("/register", c.Register)	
}

func (c *UserController) Register(ectx echo.Context) error {
	ctx, cancel := context.WithTimeout(ectx.Request().Context(), c.timeout)
	defer cancel()

	var (
		responeChan = make(chan response.Response)
	)

	go func () {
		defer close(responeChan)

		var (
			req dto.UserRegisterRequest
		)

		if err := ectx.Bind(&req); err != nil {
			responeChan <- *response.New(
				"Failed to bind request data",
				nil,
				nil,
			).
			WithErr(err).
			WithCode(http.StatusBadRequest)
			return
		}

		err := c.userSvc.RegisterUser(ctx, &req)
		if err != nil {
			responeChan <- *response.New(
				"Failed to register user",
				nil,
				nil,
			).
			WithErr(err). 
			WithCode(http.StatusInternalServerError)
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
	case resp := <- responeChan:
		return ectx.JSON(resp.Code, resp)
	}
}