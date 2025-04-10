package controller

import (
	"context"
	"net/http"

	"github.com/devanfer02/ratemyubprof/internal/app/professor/contracts"
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/pkg/http/response"
	"github.com/labstack/echo/v4"
)

func (c *ProfessorController) FetchReviews(ectx echo.Context) error {
	ctx, cancel := context.WithTimeout(ectx.Request().Context(), c.timeout)
	defer cancel()

	var (
		pageQuery dto.PaginationQuery
	)

	idParam := ectx.Param("id")
	ectx.Bind(&pageQuery)
	pageQuery.SetDefaultValue()

	res, meta, err := c.profSvc.FetchProfessorReviews(ctx, idParam, &pageQuery)
	if err != nil {
		return err 
	}

	resp := response.New(
		"Successfully fetch professor reviews",
		res,
		meta,
	)

	select {
	case <- ctx.Done():
		return contracts.ErrRequestTimeout
	default:
		return ectx.JSON(http.StatusOK, resp)
	}
}

func (c *ProfessorController) CreateReview(ectx echo.Context) error {
	ctx, cancel := context.WithTimeout(ectx.Request().Context(), c.timeout)
	defer cancel()

	var (
		req dto.ProfessorReviewRequest
		resp *response.Response
	)

	if err := ectx.Bind(&req); err != nil {
		return err 
	}

	if err := c.validator.Struct(req); err != nil {
		return err 
	}

	req.UserID = ectx.Get("userId").(string)

	err := c.profSvc.CreateReview(ctx, &req)
	if err != nil {
		return err 
	}

	resp = response.New(
		"Successfully create professor review",
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