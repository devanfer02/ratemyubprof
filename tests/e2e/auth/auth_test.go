package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	// "github.com/bytedance/sonic"
	"github.com/devanfer02/ratemyubprof/internal/infra/server"
	"github.com/devanfer02/ratemyubprof/tests/fixtures"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAuthLoginFlow(t *testing.T) {
	var (
		mqUrl, cleanUpMq = fixtures.NewRabbitMq()
		db, cleanUpDb    = fixtures.NewDB()
	)

	defer func() {
		cleanUpDb()
		cleanUpMq()
	}()

	var (
		srv = server.NewHttpServerWithConfig(server.CustomServerConfig{
			MQURL: mqUrl,
			DB: db,
		})
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
	t.Log("Login Response Body:", rec.Body.String())
	assert.Equal(t, http.StatusCreated, rec.Code)

	// Test user login 
	req = httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", strings.NewReader(loginPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Accept", echo.MIMEApplicationJSON)
	req.Header.Set("X-API-KEY", "apikeyvalue")
	rec = httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	// Assert status code
	t.Log("Register Response Body:", rec.Body.String())
	assert.Equal(t, http.StatusOK, rec.Code)
}	

