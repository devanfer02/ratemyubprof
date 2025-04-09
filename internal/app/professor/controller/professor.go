package controller

import (
	"context"
	"net/http"

	"github.com/devanfer02/ratemyubprof/internal/app/professor/contracts"
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/pkg/http/response"
	"github.com/labstack/echo/v4"
)

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

