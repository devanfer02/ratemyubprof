package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	user_ctr "github.com/devanfer02/ratemyubprof/internal/app/user/controller"
	user_repo "github.com/devanfer02/ratemyubprof/internal/app/user/repository"
	user_svc "github.com/devanfer02/ratemyubprof/internal/app/user/service"

	auth_ctr "github.com/devanfer02/ratemyubprof/internal/app/auth/controller"
	auth_svc "github.com/devanfer02/ratemyubprof/internal/app/auth/service"

	prof_ctr "github.com/devanfer02/ratemyubprof/internal/app/professor/controller"
	prof_repo "github.com/devanfer02/ratemyubprof/internal/app/professor/repository"
	prof_svc "github.com/devanfer02/ratemyubprof/internal/app/professor/service"

	review_repo "github.com/devanfer02/ratemyubprof/internal/app/review/repository"
	review_svc "github.com/devanfer02/ratemyubprof/internal/app/review/service"

	"github.com/devanfer02/ratemyubprof/internal/infra/database/postgres"
	"github.com/devanfer02/ratemyubprof/internal/infra/env"
	"github.com/devanfer02/ratemyubprof/internal/middleware"
	"github.com/devanfer02/ratemyubprof/pkg/config"
	logger "github.com/devanfer02/ratemyubprof/pkg/log"
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
	router := config.NewRouter()
	validator := config.NewValidator()

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
	jwtHandler := config.NewJwtHandler(h.Env)
	middleware := middleware.NewMiddleware(jwtHandler)

	userRepo := user_repo.NewUserRepository(h.Database)
	reviewRepo := review_repo.NewReviewRepository(h.Database)
	profRepo := prof_repo.NewProfessorRepository(h.Database)
	
	profSvc := prof_svc.NewProfessorService(profRepo)
	userSvc := user_svc.NewUserService(userRepo, jwtHandler)
	authSvc := auth_svc.NewAuthService(userRepo, jwtHandler)
	reviewSvc := review_svc.NewReviewService(reviewRepo)

	profCtr := prof_ctr.NewProfessorController(profSvc, reviewSvc,h.Validator, middleware)
	userCtr := user_ctr.NewUserController(userSvc, h.Validator, middleware)
	authCtr := auth_ctr.NewAuthController(authSvc, h.Validator, middleware)

	h.Handlers = append(
		h.Handlers, 
		userCtr,
		profCtr,
		authCtr,
	)
}

func (h *httpServer) Start() {
	h.Router.Use(middleware.ErrLogger(h.Logger))
	h.Router.Use(middleware.RequestLogger(h.Logger))
	h.mountHandlers()

	for _, handler := range h.Handlers {
		handler.Mount(h.Router.Group("/api/v1"))
	}
	
	if h.Env.App.Env == "development" {
		h.logRoutes()
	}
	h.Logger.Info("Starting up the application....")
	
	if err := h.Router.Start(":" + h.Env.App.Port); err != nil {
		panic(err)
	}
}

func (h *httpServer) shutdown() {
	h.Logger.Info("Shutting down application...")

	timeoutFunc := time.AfterFunc(5 * time.Second, func() {
		log.Println("Timeout, forcefully shutting down...")
	})

	defer timeoutFunc.Stop()

	operations := map[string]func(ctx context.Context) error {
		"database": func(ctx context.Context) error {
			return h.Database.Close()
		},
		"logger": func(ctx context.Context) error {
			return h.Logger.Sync()
		},
		"router": func(ctx context.Context) error {
			return h.Router.Close()
		},
	}

	var wg sync.WaitGroup
	
	for key, operation := range operations {
		wg.Add(1)
		go func(op func(ctx context.Context) error, name string) {
			defer wg.Done()

			log.Printf("Cleaning up %s...", name)
			if err := op(context.Background()); err != nil {
				log.Printf("Error cleaning up %s: %v", name, err)
			} else {
				log.Printf("%s cleaned up successfully", name)
			}
		}(operation, key)
	}

	wg.Wait()
}

func (h *httpServer) GracefullyShutdown() {
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<- sigChan
		h.shutdown()
	}()
}

func (h *httpServer) logRoutes() {
	h.Logger.Info("------------ Registered Routes ------------")
	for _, route := range h.Router.Routes() {
		h.Logger.Info(route.Method + " " + route.Path)
	}
	h.Logger.Info("-------------------------------------------")
}