package controller

import (
	"context"
	"net/http"
	
	"github.com/devanfer02/ratemyubprof/internal/app/user/contracts"
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/pkg/http/response"
	"github.com/labstack/echo/v4"
)

func (c *UserController) Register(ectx echo.Context) error {
	ctx, cancel := context.WithTimeout(ectx.Request().Context(), c.timeout)
	defer cancel()

	var (
		req dto.UserRegisterRequest
		resp *response.Response
	)

	if err := ectx.Bind(&req); err != nil {
		return err 
	}

	if err := c.validator.Struct(req); err != nil {
		return err 
	}

	err := c.userSvc.RegisterUser(ctx, &req)
	if err != nil {
		return err  
	}

	resp = response.New(
		"Successfully register user",
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

func (c *UserController) FetchReviews(ectx echo.Context) error {
	ctx, cancel := context.WithTimeout(ectx.Request().Context(), c.timeout)
	defer cancel()

	var (
		pageQuery dto.PaginationQuery
		param dto.FetchReviewParams
	)

	ectx.Bind(&param)
	ectx.Bind(&pageQuery)
	pageQuery.SetDefaultValue()

	if val := ectx.Get("userId").(string); val != "" {
		param.UserId = val
	}

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

func (c *UserController) ResetPassowrd(ectx echo.Context) error {
	ctx, cancel := context.WithTimeout(ectx.Request().Context(), c.timeout)
	defer cancel()

	var (
		req dto.ForgotPasswordRequest
		resp *response.Response
	)

	if err := ectx.Bind(&req); err != nil {
		return err 
	}

	if err := c.validator.Struct(req); err != nil {
		return err 
	}

	err := c.userSvc.ForgotPassword(ctx, &req)
	if err != nil {
		return err 
	}

	resp = response.New(
		"Successfully reset password",
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

func (c *UserController) FetchUserProfile(ectx echo.Context) error {
	ctx, cancel := context.WithTimeout(ectx.Request().Context(), c.timeout)
	defer cancel()

	var (
		req dto.UserProfileRequest
	)

	if val := ectx.Get("userId").(string); val != "" {
		req.UserID = val
	}

	res, err := c.userSvc.FetchUserProfile(ctx, &req)
	if err != nil {
		return err
	}

	resp := response.New(
		"Successfully fetch user profile",
		res,
		nil,
	)

	select {
	case <-ctx.Done():
		return contracts.ErrRequestTimeout
	default:
		return ectx.JSON(http.StatusOK, resp)
	}
}

func (c *UserController) FetchUserProfileByID(ectx echo.Context) error {
	ctx, cancel := context.WithTimeout(ectx.Request().Context(), c.timeout)
	defer cancel()

	var (
		req dto.UserProfileRequest
	)

	req.UserID = ectx.Param("id")

	res, err := c.userSvc.FetchUserProfile(ctx, &req)
	if err != nil {
		return err 
	}

	resp := response.New(
		"Successfully fetch user profile",
		res,
		nil,
	)

	select {
	case <-ctx.Done():
		return contracts.ErrRequestTimeout
	default:
		return ectx.JSON(http.StatusOK, resp)
	}
}