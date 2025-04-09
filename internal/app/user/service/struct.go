package service

import (
	"github.com/devanfer02/ratemyubprof/internal/app/user/contracts"

	"github.com/devanfer02/ratemyubprof/pkg/config"
	"go.uber.org/zap"
)

type userService struct {
	userRepo contracts.UserRepositoryProvider
	jwtHandler *config.JwtHandler
	logger *zap.Logger
}

func NewUserService(
	userRepo contracts.UserRepositoryProvider, 
	jwtHandler *config.JwtHandler,
	logger *zap.Logger,
) contracts.UserService {
	return &userService{
		jwtHandler: jwtHandler,
		userRepo: userRepo,
		logger: logger,
	}
}