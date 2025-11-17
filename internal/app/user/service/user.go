package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/devanfer02/ratemyubprof/internal/app/user/contracts"
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/entity"
	"github.com/oklog/ulid/v2"

	"github.com/devanfer02/ratemyubprof/pkg/siam"
	"github.com/devanfer02/ratemyubprof/pkg/util"
	"github.com/devanfer02/ratemyubprof/pkg/util/formatter"
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
		CreatedAt: time.Now(),
	})

	if err != nil {
		return err 
	}

	return nil 
}

func (s *userService) ForgotPassword(ctx context.Context, req *dto.ForgotPasswordRequest) error {
	repoClient, err := s.userRepo.NewClient(false)
	if err != nil {
		return err 
	}

	// Check if user already registered or not before using SIAM Auth
	user, err := repoClient.FetchUserByParams(ctx, &dto.FetchUserParams{Username: req.Username, NIM: req.NIM})
	if err != nil {
		return err 
	}

	// Check if recently user has reseted their password. Max Reset: 7Day/One Reset
	if user.ForgotPasswordAt.Valid && user.ForgotPasswordAt.Time.Add(7 * 24 * time.Hour).After(time.Now()) {
		return contracts.ErrResetPasswordLimit
	}

	authMgr := siam.NewSiamAuthManager()

	err = authMgr.Authenticate(req.NIM, req.Password)

	if err != nil {
		return err 
	}

	hashed, err := util.HashPassword(req.NewPassword) 
	if err != nil {
		return err 
	}

	err = repoClient.UpdateUser(ctx, &entity.User{
		NIM: req.NIM,
		Username: req.Username,
		Password: hashed,
		ForgotPasswordAt: sql.NullTime{Time: time.Now(), Valid: true},
	})

	if err != nil {
		return err 
	}

	return nil 
}

func (s *userService) FetchUserProfile(ctx context.Context, usr *dto.FetchUserParams) (dto.UserProfileResponse, error) {
	repoClient, err := s.userRepo.NewClient(false)
	if err != nil {
		return dto.UserProfileResponse{}, err
	}

	reviewRepoClient, err := s.reviewRepo.NewClient(false)

	user, err := repoClient.FetchUserByParams(ctx, usr)

	if err != nil {
		return dto.UserProfileResponse{}, err
	}

	reviews, err := reviewRepoClient.FetchReviewsByParams(context.TODO(), &dto.FetchReviewParams{UserId: user.ID}, &dto.PaginationQuery{})
	if err != nil {
		return dto.UserProfileResponse{}, err 
	}

	userProfile := formatter.FormatToUserProfile(&user, reviews)

	return userProfile, nil
}

