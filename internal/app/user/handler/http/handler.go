package handler

import (
	"github.com/devanfer02/ratemyubprof/internal/app/user/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userSvc service.UserService
	validator *validator.Validate
}

func NewUserHandler(userSvc service.UserService, validator *validator.Validate) *UserHandler {
	return &UserHandler{
		userSvc: userSvc,
		validator: validator,
	}
}

func (h *UserHandler) Mount(r *echo.Group) {
	userR := r.Group("/users")

	userR.GET("", h.Login)
}