package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/devanfer02/ratemyubprof/internal/app/professor/contracts"
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/middleware"
	"github.com/devanfer02/ratemyubprof/pkg/http/response"
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

func (c *ProfessorController) FetchAll(ectx echo.Context) error {
	ctx, cancel := context.WithTimeout(ectx.Request().Context(), c.timeout)
	defer cancel()

	var (
		pageQuery dto.PaginationQuery
		queryParam dto.FetchProfessorParam
		resp *response.Response
	)

	ectx.Bind(&pageQuery)
	ectx.Bind(&queryParam)

	if pageQuery.Limit == 0 {
		pageQuery.Limit = 10
	}
	if pageQuery.Page == 0 {
		pageQuery.Page = 1
	}

	professors, meta, err := c.profSvc.FetchAllProfessors(ctx, &queryParam, &pageQuery)
	if err != nil {
		return err 
	}

	resp = response.New(
		"Successfully fetched all professors",
		professors,
		meta,
	)

	select {
	case <-ctx.Done():
		return contracts.ErrRequestTimeout
	default:
		return ectx.JSON(http.StatusOK, resp)
	}
}

func (c *ProfessorController) FetchByID(ectx echo.Context) error {
	ctx, cancel := context.WithTimeout(ectx.Request().Context(), c.timeout)
	defer cancel()

	var (
		resp *response.Response
	)

	profId := ectx.Param("id")

	professor, err := c.profSvc.FetchProfessorByID(ctx, profId)
	if err != nil {
		return err 
	}

	resp = response.New(
		"Successfully fetched professor",
		professor,
		nil,
	)

	select {
	case <-ctx.Done():
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