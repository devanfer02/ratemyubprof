package contracts

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/entity"
)

type UserRepositoryProvider interface {
	NewClient(tx bool) (UserRepository, error)
}

type UserRepository interface {
	FetchUserByParams(ctx context.Context, params *dto.FetchUserParams) (entity.User, error)
	InsertUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	FetchUserProfile(ctx context.Context, userID string) (dto.UserProfileResponse, error)
}