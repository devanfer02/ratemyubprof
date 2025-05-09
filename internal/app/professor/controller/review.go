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
		param dto.FetchReviewParams
	)

	ectx.Bind(&param)
	ectx.Bind(&pageQuery)
	pageQuery.SetDefaultValue()

	userId, _ := ectx.Get("userId").(string)
	param.SignedUser = userId
	param.ID = ectx.QueryParam("reviewId")

	res, meta, err := c.reviewSvc.FetchReviewsByParams(ctx, &param, &pageQuery)
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

func (c *ProfessorController) UpdateReview(ectx echo.Context) error {
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

	err := c.profSvc.UpdateProfessorReview(ctx, &req)
	if err != nil {
		return err 
	}

	resp = response.New(
		"Successfully update professor review",
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

func (c *ProfessorController) DeleteReview(ectx echo.Context) error {
	ctx, cancel := context.WithTimeout(ectx.Request().Context(), c.timeout)
	defer cancel()

	var (
		params dto.FetchReviewParams
		resp *response.Response
	)

	ectx.Bind(&params)

	params.UserId = ectx.Get("userId").(string)

	err := c.profSvc.DeleteProfessorReview(ctx, &params)
	if err != nil {
		return err 
	}

	resp = response.New(
		"Successfully delete professor review",
		nil,
		nil,
	)
	
	select {
	case <-ctx.Done():
		return contracts.ErrRequestTimeout
	default:
		return ectx.JSON(http.StatusOK, resp)
	}		
}