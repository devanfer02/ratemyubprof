package service

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/app/user/contracts"
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/entity"
	"github.com/oklog/ulid/v2"

	"github.com/devanfer02/ratemyubprof/pkg/siam"
	"github.com/devanfer02/ratemyubprof/pkg/util"
	"go.uber.org/zap"
)

type userService struct {
	userRepo contracts.UserRepositoryProvider
	jwtHandler *util.JwtHandler
	logger *zap.Logger
}

func NewUserService(
	userRepo contracts.UserRepositoryProvider, 
	jwtHandler *util.JwtHandler,
	logger *zap.Logger,
) contracts.UserService {
	return &userService{
		jwtHandler: jwtHandler,
		userRepo: userRepo,
		logger: logger,
	}
}

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

	hashed, err := util.HashPassword(usr.Password)
	if err != nil {
		return err 
	}

	err = repoClient.InsertUser(ctx, &entity.User{
		ID: ulid.Make().String(),
		Username: usr.Username,
		Password: hashed,
	})

	if err != nil {
		return err 
	}

	return nil 
}

func (s *userService) LoginUser(ctx context.Context, usr *dto.UserLoginRequest) (string, error) {
	repoClient, err := s.userRepo.NewClient(false)
	if err != nil {
		return "", err 
	}

	user, err := repoClient.FetchUserByUsername(ctx, usr.Username)
	if err != nil {
		return "", err 
	}

	if !util.CheckPasswordHash(usr.Password, user.Password) {
		return "", contracts.ErrInvalidCredential
	}

	token, err := s.jwtHandler.GenerateToken(user.ID)

	if err != nil {
		return "", err 
	}

	return token, nil 
}