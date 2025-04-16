package service

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/entity"
	"github.com/oklog/ulid/v2"

	"github.com/devanfer02/ratemyubprof/pkg/siam"
	"github.com/devanfer02/ratemyubprof/pkg/util"
)

func (s *userService) RegisterUser(ctx context.Context, usr *dto.UserRegisterRequest) error {
	authMgr := siam.NewSiamAuthManager()

	err := authMgr.Authenticate(usr.NIM, usr.Password)

	if err != nil {
		return err 
	}

	repoClient, err := s.userRepo.NewClient(false)
	if err != nil {
		return err 
	}

	hashed, err := util.HashPassword(usr.NewPassword) 
	if err != nil {
		return err 
	}

	err = repoClient.InsertUser(ctx, &entity.User{
		ID: ulid.Make().String(),
		NIM: usr.NIM,
		Username: usr.Username,
		Password: hashed,
	})

	if err != nil {
		return err 
	}

	return nil 
}

