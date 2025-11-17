package service

import (
	"github.com/devanfer02/ratemyubprof/internal/app/user/contracts"
	rvwContract "github.com/devanfer02/ratemyubprof/internal/app/review/contracts"

	"github.com/devanfer02/ratemyubprof/pkg/config"
)

type userService struct {
	userRepo contracts.UserRepositoryProvider
	reviewRepo rvwContract.ReviewRepositoryProvider
	jwtHandler *config.JwtHandler
}

func NewUserService(
	userRepo contracts.UserRepositoryProvider, 
	reviewRepo rvwContract.ReviewRepositoryProvider,
	jwtHandler *config.JwtHandler,
) contracts.UserService {
	return &userService{
		jwtHandler: jwtHandler,
		userRepo: userRepo,
		reviewRepo: reviewRepo,
	}
}