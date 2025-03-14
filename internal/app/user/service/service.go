package service

import "github.com/devanfer02/ratemyubprof/internal/app/user/contracts"

type UserService interface {

}

type userService struct {
	userRepo contracts.UserRepository
}

func NewUserService(userRepo contracts.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}