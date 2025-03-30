package service

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"

	"github.com/devanfer02/ratemyubprof/internal/app/user/contracts"
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/pkg/response"
	"go.uber.org/zap"
)

type userService struct {
	userRepo contracts.UserRepository
	logger *zap.Logger
}

func NewUserService(userRepo contracts.UserRepository, logger *zap.Logger) contracts.UserService {
	return &userService{
		userRepo: userRepo,
		logger: logger,
	}
}

func (s *userService) RegisterUser(ctx context.Context, usr *dto.UserRegisterRequest) error {
	formData := fmt.Sprintf("username=%s&password=%s=login=Masuk", usr.NIM, usr.Password)
	req, _ := http.NewRequest("POST", "https://siam.ub.ac.id/login", strings.NewReader(formData))
	jar, _ := cookiejar.New(nil)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")

	client := &http.Client{Jar: jar, Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		s.logger.Error("Failed to send request to SIAM", zap.Error(err))
		return err 
	}
	defer resp.Body.Close()

	//  Silahkan memasukkan ulang password Anda! 
	//  Silahkan ulangi LOGIN beberapa saat lagi!
	body, _ := io.ReadAll(resp.Body)
	content := string(body)

	if strings.Contains(content, "Silahkan memasukkan ulang password Anda!") {
		return response.NewErr(400, "Wrong credentials")
	}


	if strings.Contains(content, "Silahkan ulangi LOGIN beberapa saat lagi!") {
		return response.NewErr(400, "Wrong credentials")
	}
	fmt.Println(content)

	return nil 
}