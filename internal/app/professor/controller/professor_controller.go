package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/devanfer02/ratemyubprof/internal/app/professor/contracts"
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/pkg/http/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ProfessorController struct {
	profSvc contracts.ProfessorService
	validator *validator.Validate
	timeout time.Duration
}

func NewProfessorController(profSvc contracts.ProfessorService, validator *validator.Validate) *ProfessorController {
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
		errChan = make(chan error)
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
			errChan <- err 
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
	case err := <- errChan:
		return err 
	case resp := <- responeChan:
		return ectx.JSON(http.StatusOK, resp)
	}
	
}

func (c *ProfessorController) CreateReview(ectx echo.Context) error {
	ctx, cancel := context.WithTimeout(ectx.Request().Context(), c.timeout)
	defer cancel()

	var (
		responeChan = make(chan response.Response)
		errChan = make(chan error)
	)

	go func () {
		defer close(responeChan)

		var (
			req dto.ProfessorReviewRequest
			profId = ectx.Param("id")
		)

		if err := ectx.Bind(&req); err != nil {
			errChan <- err 
		}

		if err := c.validator.Struct(&req); err != nil {
			errChan <- err 
		}

		req.ProfessorID = profId
		req.UserID = ectx.Get("userId").(string)

		err := c.profSvc.CreateReview(ctx, &req)
		if err != nil {
			errChan <- err 
			return
		}

		responeChan <- *response.New(
			"Successfully create professor review",
			nil,
			nil,
		)
	}()

	select {
	case <-ctx.Done():
		return contracts.ErrRequestTimeout
	case err := <- errChan:
		return err 
	case resp := <- responeChan:
		return ectx.JSON(http.StatusCreated, resp)
	}	
}