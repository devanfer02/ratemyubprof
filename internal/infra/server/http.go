package server

import (
	"os"
	"os/signal"
	"syscall"

	handler "github.com/devanfer02/presentia-api/internal/app/user/handler/http"
	"github.com/devanfer02/presentia-api/internal/app/user/repository"
	"github.com/devanfer02/presentia-api/internal/app/user/service"
	"github.com/devanfer02/presentia-api/internal/infra/database"
	"github.com/devanfer02/presentia-api/internal/infra/env"
	"github.com/devanfer02/presentia-api/internal/middleware"
	logger "github.com/devanfer02/presentia-api/pkg/log"
	validate "github.com/devanfer02/presentia-api/pkg/validator"
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

	return &httpServer{
		Env: env,
		Logger: logger,
		Router: router,
		Database: db,
		Validator: validator,
		Handlers: make([]httpHandler, 0),
	}
}

func (h *httpServer) MountHandlers() {
	userRepo := repository.NewUserRepository(h.Database)
	userSvc := service.NewUserService(userRepo)
	userCtr := handler.NewUserHandler(userSvc, h.Validator)

	h.Handlers = append(h.Handlers, userCtr)
}

func (h *httpServer) Start() {
	h.Router.Use(middleware.ErrLogger(h.Logger))
	h.Router.Use(middleware.RequestLogger(h.Logger))
	h.MountHandlers()

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