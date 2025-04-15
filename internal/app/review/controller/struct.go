package controller

import (
	"time"

	reaction "github.com/devanfer02/ratemyubprof/internal/app/reaction/contracts"
	"github.com/devanfer02/ratemyubprof/internal/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ReviewController struct {
	reactionSvc reaction.ReviewReactionService
	validator *validator.Validate
	mdlwr *middleware.Middleware
	timeout time.Duration
}

func NewReviewController(reactionSvc reaction.ReviewReactionService, validator *validator.Validate, mdlwr *middleware.Middleware) *ReviewController {
	return &ReviewController{
		reactionSvc: reactionSvc,
		timeout: 5 * time.Second,
		mdlwr: mdlwr,
		validator: validator,
	}
}

func (c *ReviewController) Mount(r *echo.Group) {
	profR := r.Group("/reviews")

	profR.POST("/:id/reactions", c.CreateReaction, c.mdlwr.Authenticate())
	profR.DELETE("/:id/reactions", c.DeleteReaction, c.mdlwr.Authenticate())
}