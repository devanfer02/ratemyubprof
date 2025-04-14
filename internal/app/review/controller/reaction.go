package controller

import (
	"context"
	"net/http"

	"github.com/devanfer02/ratemyubprof/internal/app/auth/contracts"
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/infra/rabbitmq"
	"github.com/devanfer02/ratemyubprof/pkg/http/response"
	"github.com/labstack/echo/v4"
)

func (c *ReviewController) CreateReaction(ectx echo.Context) error {
	ctx, cancel := context.WithTimeout(ectx.Request().Context(), c.timeout)
	defer cancel()

	var (
		req dto.ReviewReactionRequest
	)

	if err := ectx.Bind(&req); err != nil {
		return err 
	}

	if err := c.validator.Struct(req); err != nil {
		return err 
	}

	req.UserID = ectx.Get("userId").(string)

	err := c.reactionSvc.PublishReaction(ctx, rabbitmq.ReactionReviewCreateQueue, &req)

	if err != nil {
		return err 
	}

	res := response.New(
		"successfully published create review reaction to queue",
		nil,
		nil, 
	)

	select {
	case <- ctx.Done():
		return contracts.ErrRequestTimeout
	default:
		return ectx.JSON(http.StatusCreated, res)
	}
}