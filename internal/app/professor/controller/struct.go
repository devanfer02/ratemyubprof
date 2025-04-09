package controller

import (
	"time"

	"github.com/devanfer02/ratemyubprof/internal/app/professor/contracts"
	"github.com/devanfer02/ratemyubprof/internal/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ProfessorController struct {
	profSvc contracts.ProfessorService
	validator *validator.Validate
	mdlwr *middleware.Middleware
	timeout time.Duration
}

func NewProfessorController(profSvc contracts.ProfessorService, validator *validator.Validate, mdlwr *middleware.Middleware) *ProfessorController {
	return &ProfessorController{
		profSvc: profSvc,
		timeout: 5 * time.Second,
		mdlwr: mdlwr,
		validator: validator,
	}
}

func (c *ProfessorController) Mount(r *echo.Group) {
	profR := r.Group("/professors")

	profR.GET("", c.FetchAll)
	profR.GET("/:id", c.FetchByID)
	profR.POST("/:id/reviews", c.CreateReview, c.mdlwr.Authenticate())
}