package controller

import (
	"time"

	"github.com/devanfer02/ratemyubprof/internal/app/professor/contracts"
	review "github.com/devanfer02/ratemyubprof/internal/app/review/contracts"
	"github.com/devanfer02/ratemyubprof/internal/middleware"
	"github.com/devanfer02/ratemyubprof/pkg/config"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ProfessorController struct {
	profSvc contracts.ProfessorService
	reviewSvc review.ReviewService
	validator *validator.Validate
	mdlwr *middleware.Middleware
	timeout time.Duration
}

func NewProfessorController(profSvc contracts.ProfessorService, reviewSvc review.ReviewService, validator *validator.Validate, mdlwr *middleware.Middleware) *ProfessorController {
	return &ProfessorController{
		profSvc: profSvc,
		reviewSvc: reviewSvc,
		timeout: 5 * time.Second,
		mdlwr: mdlwr,
		validator: validator,
	}
}

func (c *ProfessorController) Mount(r *echo.Group) {
	profR := r.Group("/professors")

	profR.GET("", c.FetchAll, config.FetchLimiter)
	profR.GET("/:id", c.FetchByID, config.FetchLimiter)

	profR.GET("/:profId/reviews", c.FetchReviews )
	profR.POST("/:id/reviews", c.CreateReview, c.mdlwr.Authenticate(), config.PostLimiter)
	profR.PUT("/:id/reviews", c.UpdateReview, c.mdlwr.Authenticate(), config.PostLimiter)
	profR.DELETE("/:profId/reviews", c.DeleteReview, c.mdlwr.Authenticate(), config.PostLimiter)
}