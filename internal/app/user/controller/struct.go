package controller

import (
	"time"

	"github.com/devanfer02/ratemyubprof/internal/app/user/contracts"
	review "github.com/devanfer02/ratemyubprof/internal/app/review/contracts"
	"github.com/devanfer02/ratemyubprof/internal/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userSvc contracts.UserService
	reviewSvc review.ReviewService
	validator *validator.Validate
	mdlwr *middleware.Middleware
	timeout time.Duration
}

func NewUserController(userSvc contracts.UserService, reviewSvc review.ReviewService, validator *validator.Validate, mdlwr *middleware.Middleware) *UserController {
	return &UserController{
		userSvc: userSvc,
		reviewSvc: reviewSvc,
		mdlwr: mdlwr,
		timeout: 10 * time.Second,
		validator: validator,
	}
}

func (c *UserController) Mount(r *echo.Group) {
	userR := r.Group("/users")

	userR.GET("/:userId/reviews", c.FetchReviews)
	userR.GET("/reviews", c.FetchReviews, c.mdlwr.Authenticate())
	userR.POST("/register", c.Register)	
}