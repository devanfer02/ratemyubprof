package contracts

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/dto"
)

type UserService interface {
	RegisterUser(ctx context.Context, usr *dto.UserRegisterRequest) error 
	LoginUser(ctx context.Context, usr *dto.UserLoginRequest) (dto.UserTokenResponse, error)	
	RefreshAccessToken(ctx context.Context, req dto.RefreshATRequest) (dto.UserTokenResponse, error)
}