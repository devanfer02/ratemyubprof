package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	// "github.com/bytedance/sonic"
	"github.com/devanfer02/ratemyubprof/internal/infra/server"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAuthLoginFlow(t *testing.T) {
	var (
		srv = server.NewHttpServer()
		loginPayload = `{"username": "testuser", "password": "testpass"}`
		registerPayload = `{"nim": "225150200111", "username": "testuser", "password": "testpass", "newPassword": "testpass"}`
	)

	srv.MountHandlers()
	srv.Bootstrap()

	e := srv.GetRouter().(*echo.Echo)

	// Test user registration 
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/register", strings.NewReader(registerPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Accept", echo.MIMEApplicationJSON)
	req.Header.Set("X-API-KEY", "apikeyvalue")
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	// Assert status code
	t.Log("response:", rec.Body.String())
	assert.Equal(t, http.StatusCreated, rec.Code)

	// Test user login 
	req = httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", strings.NewReader(loginPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Accept", echo.MIMEApplicationJSON)
	req.Header.Set("X-API-KEY", "apikeyvalue")
	rec = httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	// Assert status code
	assert.Equal(t, http.StatusOK, rec.Code)

	t.Log("response:", rec.Body.String())
}	

