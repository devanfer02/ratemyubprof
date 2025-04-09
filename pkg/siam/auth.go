package siam

import (
	"bytes"
	"mime/multipart"
	"strings"
	"time"

	"github.com/Noooste/azuretls-client"
	"github.com/devanfer02/ratemyubprof/internal/app/auth/contracts"
	"go.uber.org/zap"
)

type SiamAuthManager struct {
	session *azuretls.Session
	logger *zap.Logger
}

func NewSiamAuthManager() *SiamAuthManager {
	return &SiamAuthManager{}
}

func (s *SiamAuthManager) getTokenAndHeaders() (string, map[string]string, error) {
	s.session = azuretls.NewSession()
	defer s.session.Close()
	
	s.session.OrderedHeaders = azuretls.OrderedHeaders{
		{"User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36"},
		{"Accept", "*/*"},
		{"Connection", "keep-alive"},
	}

    response, err := s.session.Head("https://siam.ub.ac.id")
    if err != nil {
        s.logger.Error("Failed to get siam.ub.ac.id headers", zap.Error(err))
		return "", nil, err 
    }


	return response.Cookies["PHPSESSID"], map[string]string{
		"Cf-Ray": strings.Join(response.Header["Cf-Ray"], ","),
		"Cf-Cache-Status": strings.Join(response.Header["Cf-Cache-Status"], ","),
		"Cf-Apo-Via": strings.Join(response.Header["Cf-Apo-Via"], ","),
	}, nil 
}

func (s *SiamAuthManager) Authenticate(username, password string) error {
	token, headers, err := s.getTokenAndHeaders()
	if err != nil {
		return err 
	}

	s.session = azuretls.NewSession()
	
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
    
    _ = writer.WriteField("username", username)
    _ = writer.WriteField("password", password)
    _ = writer.WriteField("login", "Masuk")

	reqHeaders := azuretls.OrderedHeaders{
		{"User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36"},
		{"Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"},
		{"Cf-Ray", headers["Cf-Ray"]},
		{"Cf-Cache-Status", headers["Cf-Cache-Status"]},
		{"Cf-Apo-Via", headers["Cf-Apo-Via"]},
		{"Server", "Cloudflare"},
		{"Content-Type", writer.FormDataContentType()},
		{"Cookie", token},
	}

	s.session.OrderedHeaders = reqHeaders
	s.session.SetTimeout(10 * time.Second)

	response, err := s.session.Post("https://siam.ub.ac.id/index.php", &requestBody)
    if err != nil {
        s.logger.Error("Failed to send login request to siam.ub.ac.id", zap.Error(err))
		return err 
    }
	
	if strings.Contains(string(response.Body), "Silahkan memasukkan ulang password Anda!") {
		return contracts.ErrInvalidCredential
	} else if strings.Contains(string(response.Body), "User belum terdaftar di database") {
		return contracts.ErrInvalidCredential
	} 

	s.session.Close()

	s.logout(reqHeaders)
	

	return nil 
}

func (s *SiamAuthManager) logout(headers azuretls.OrderedHeaders) {
	s.session = azuretls.NewSession()
	defer s.session.Close()
	
	s.session.OrderedHeaders = headers

	s.session.Get("https://siam.ub.ac.id/logout.php", nil)	
}