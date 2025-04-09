package service

import (

	"github.com/devanfer02/ratemyubprof/internal/app/auth/contracts"
	userCts "github.com/devanfer02/ratemyubprof/internal/app/user/contracts"

	"github.com/devanfer02/ratemyubprof/pkg/config"
	"go.uber.org/zap"
)

type authService struct {
	userRepo userCts.UserRepositoryProvider
	jwtHandler *config.JwtHandler
	logger *zap.Logger
}

func NewAuthService(
	userRepo userCts.UserRepositoryProvider, 
	jwtHandler *config.JwtHandler,
	logger *zap.Logger,
) contracts.AuthService {
	return &authService{
		jwtHandler: jwtHandler,
		userRepo: userRepo,
		logger: logger,
	}
}