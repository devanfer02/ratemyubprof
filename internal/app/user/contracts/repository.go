package contracts

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/entity"
)

type UserRepositoryProvider interface {
	NewClient(tx bool) (UserRepository, error)
}

type UserRepository interface {
	InsertUser(ctx context.Context, user *entity.User) error 	
}