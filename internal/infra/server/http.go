package server

import (
	"os"
	"os/signal"
	"syscall"

	user_ctr "github.com/devanfer02/ratemyubprof/internal/app/user/controller"
	user_repo "github.com/devanfer02/ratemyubprof/internal/app/user/repository"
	user_svc "github.com/devanfer02/ratemyubprof/internal/app/user/service"

	prof_svc "github.com/devanfer02/ratemyubprof/internal/app/professor/service"
	prof_ctr "github.com/devanfer02/ratemyubprof/internal/app/professor/controller"

	"github.com/devanfer02/ratemyubprof/internal/infra/database"
	"github.com/devanfer02/ratemyubprof/internal/infra/env"
	"github.com/devanfer02/ratemyubprof/internal/middleware"
	"github.com/devanfer02/ratemyubprof/pkg/config"
	logger "github.com/devanfer02/ratemyubprof/pkg/log"
	validate "github.com/devanfer02/ratemyubprof/pkg/validator"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"go.uber.org/zap"
)

type httpHandler interface {
	Mount(e *echo.Group)
}

type httpServer struct {
	Env *env.Env 
	Router *echo.Echo
	Database *sqlx.DB 
	Logger *zap.Logger
	Validator *validator.Validate
	Handlers []httpHandler
}

func NewHttpServer() *httpServer {
	env := env.NewEnv()
	db := database.NewDatabase(env)
	logger := logger.NewLogger(env)
	router := echo.New()
	validator := validate.NewValidator()

	router.JSONSerializer = config.NewSonicJSONSerializer()

	return &httpServer{
		Env: env,
		Logger: logger,
		Router: router,
		Database: db,
		Validator: validator,
		Handlers: make([]httpHandler, 0),
	}
}

func (h *httpServer) mountHandlers() {
	userRepo := user_repo.NewUserRepository(h.Database)
	userSvc := user_svc.NewUserService(userRepo, h.Logger)
	userCtr := user_ctr.NewUserController(userSvc, h.Validator)

	profSvc := prof_svc.NewProfessorService()
	profCtr := prof_ctr.NewProfessorController(profSvc)

	h.Handlers = append(
		h.Handlers, 
		userCtr,
		profCtr,
	)
}

func (h *httpServer) Start() {
	h.Router.Use(middleware.ErrLogger(h.Logger))
	h.Router.Use(middleware.RequestLogger(h.Logger))
	h.mountHandlers()

	for _, handler := range h.Handlers {
		handler.Mount(h.Router.Group("/api/v1"))
	}
	
	h.Logger.Info("Starting up the application....")
	h.Router.Start(":" + h.Env.App.Port)
}

func (h *httpServer) shutdown() {
	h.Logger.Info("Shutting down application...")
	h.Database.Close()
	h.Logger.Sync()
	h.Router.Close()
}

func (h *httpServer) GracefullyShutdown() {
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<- sigChan
		h.shutdown()
	}()
}