package controller

import (
	"context"
	"time"

	"github.com/devanfer02/presentia-api/internal/app/professor/contracts"
	"github.com/devanfer02/presentia-api/internal/dto"
	"github.com/devanfer02/presentia-api/pkg/response"
	"github.com/labstack/echo/v4"
)

type ProfessorController struct {
	profSvc contracts.ProfessorService
	timeout time.Duration
}

func NewProfessorController(profSvc contracts.ProfessorService) *ProfessorController {
	return &ProfessorController{
		profSvc: profSvc,
		timeout: 5 * time.Second,
	}
}

func (c *ProfessorController) Mount(r *echo.Group) {
	profR := r.Group("/professors")

	profR.GET("/static", c.FetchStaticProfessorData)
}

func (c *ProfessorController) FetchStaticProfessorData(ectx echo.Context) error {
	ctx, cancel := context.WithTimeout(ectx.Request().Context(), c.timeout)
	defer cancel()

	var (
		responeChan = make(chan response.Response)
	)

	go func () {
		defer close(responeChan)

		var (
			nameQuery = ectx.QueryParam("name")
			facultyQuery = ectx.QueryParam("faculty")
			prodiQuery = ectx.QueryParam("prodi")
		)

		fetchQuery := dto.FetchProfessorParam{
			Name: nameQuery,
			Faculty: facultyQuery,
			Prodi: prodiQuery,
		}

		professors, err := c.profSvc.FetchStaticProfessorData(&fetchQuery)
		if err != nil {
			responeChan <- *response.New(
				"Failed to fetch data from static file",
				nil,
				nil,
			).WithErr(err)
			return
		}

		responeChan <- *response.New(
			"Successfully fetch professors data from static file",
			professors,
			nil,
		)
	}()

	select {
	case <-ctx.Done():
		return contracts.ErrRequestTimeout
	case resp := <- responeChan:
		return ectx.JSON(resp.Code, resp)
	}
	
}
