package controller

import (
	"time"

	"github.com/devanfer02/ratemyubprof/internal/app/auth/contracts"
	"github.com/devanfer02/ratemyubprof/internal/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	userSvc contracts.AuthService
	validator *validator.Validate
	mdlwr *middleware.Middleware
	timeout time.Duration
}

func NewAuthController(userSvc contracts.AuthService, validator *validator.Validate, mdlwr *middleware.Middleware) *AuthController {
	return &AuthController{
		userSvc: userSvc,
		mdlwr: mdlwr,
		timeout: 5 * time.Second,
		validator: validator,
	}
}

func (c *AuthController) Mount(r *echo.Group) {
	authR := r.Group("/auth")

	authR.POST("/login", c.Login)
	authR.POST("/refresh" ,c.RefreshToken)
}