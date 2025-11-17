package contracts

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/dto"
)

type UserService interface {
	RegisterUser(ctx context.Context, usr *dto.UserRegisterRequest) error 
	ForgotPassword(ctx context.Context, req *dto.ForgotPasswordRequest) error
	FetchUserProfile(ctx context.Context, usr *dto.FetchUserParams) (dto.UserProfileResponse, error)
}