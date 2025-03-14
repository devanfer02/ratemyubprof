package service

import (
	contracts "github.com/devanfer02/presentia-api/internal/contracts/repositories"
)

type UserService interface {

}

type userService struct {
	userRepo contracts.UserRepository
}

func NewUserService(userRepo contracts.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}