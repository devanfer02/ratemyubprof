package service

import (
	"github.com/devanfer02/ratemyubprof/internal/app/user/contracts"

	"github.com/devanfer02/ratemyubprof/pkg/config"
)

type userService struct {
	userRepo contracts.UserRepositoryProvider
	jwtHandler *config.JwtHandler
}

func NewUserService(
	userRepo contracts.UserRepositoryProvider, 
	jwtHandler *config.JwtHandler,
) contracts.UserService {
	return &userService{
		jwtHandler: jwtHandler,
		userRepo: userRepo,
	}
}