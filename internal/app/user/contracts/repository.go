package contracts

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/entity"
)

type UserRepositoryProvider interface {
	NewClient(tx bool) (UserRepository, error)
}

type UserRepository interface {
	FetchUserByUsername(ctx context.Context, username string) (*entity.User, error)
	InsertUser(ctx context.Context, user *entity.User) error 	
}