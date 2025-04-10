package service

import (

	"github.com/devanfer02/ratemyubprof/internal/app/auth/contracts"
	userCts "github.com/devanfer02/ratemyubprof/internal/app/user/contracts"

	"github.com/devanfer02/ratemyubprof/pkg/config"
)

type authService struct {
	userRepo userCts.UserRepositoryProvider
	jwtHandler *config.JwtHandler
}

func NewAuthService(
	userRepo userCts.UserRepositoryProvider, 
	jwtHandler *config.JwtHandler,
) contracts.AuthService {
	return &authService{
		jwtHandler: jwtHandler,
		userRepo: userRepo,
	}
}