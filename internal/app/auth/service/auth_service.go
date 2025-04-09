package service

import (
	"context"

	"github.com/devanfer02/ratemyubprof/internal/app/auth/contracts"
	userCts "github.com/devanfer02/ratemyubprof/internal/app/user/contracts"
	"github.com/devanfer02/ratemyubprof/internal/dto"

	"github.com/devanfer02/ratemyubprof/pkg/config"
	"github.com/devanfer02/ratemyubprof/pkg/util"
	"go.uber.org/zap"
)

type authService struct {
	userRepo userCts.UserRepositoryProvider
	jwtHandler *config.JwtHandler
	logger *zap.Logger
}

func NewAuthService(
	userRepo userCts.UserRepositoryProvider, 
	jwtHandler *config.JwtHandler,
	logger *zap.Logger,
) contracts.AuthService {
	return &authService{
		jwtHandler: jwtHandler,
		userRepo: userRepo,
		logger: logger,
	}
}

func (s *authService) LoginUser(ctx context.Context, usr *dto.UserLoginRequest) (dto.UserTokenResponse, error) {
	repoClient, err := s.userRepo.NewClient(false)
	if err != nil {
		return dto.UserTokenResponse{}, err 
	}

	user, err := repoClient.FetchUserByUsername(ctx, usr.Username)
	if err != nil {
		return dto.UserTokenResponse{}, err 
	}

	if !util.CheckPasswordHash(usr.Password, user.Password) {
		return dto.UserTokenResponse{}, contracts.ErrInvalidCredential
	}

	atToken, err := s.jwtHandler.GenerateToken(user.ID, config.AccessToken)

	if err != nil {
		return dto.UserTokenResponse{}, err 
	}

	rtToken, err := s.jwtHandler.GenerateToken(user.ID, config.RefreshToken)

	if err != nil {
		return dto.UserTokenResponse{}, err 
	}

	return dto.UserTokenResponse{
		AccessToken: atToken,
		RefreshToken: rtToken,
	}, nil 
}

func (s *authService) RefreshAccessToken(ctx context.Context, req dto.RefreshATRequest) (dto.UserTokenResponse, error) {
	userId, err := s.jwtHandler.ValidateToken(req.RefreshToken, config.RefreshToken)

	if err != nil {
		return dto.UserTokenResponse{},contracts.ErrInvalidToken
	}

	atToken, err := s.jwtHandler.GenerateToken(userId, config.AccessToken)

	if err != nil {
		return dto.UserTokenResponse{}, err 
	}

	return dto.UserTokenResponse{
		AccessToken: atToken,
		RefreshToken: req.RefreshToken,
	}, nil 
}