package contracts

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/dto"
)

type UserService interface {
	RegisterUser(ctx context.Context, usr *dto.UserRegisterRequest) error 
}