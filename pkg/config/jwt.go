package config

import (
	"time"

	"github.com/devanfer02/ratemyubprof/internal/infra/env"
	"github.com/golang-jwt/jwt/v5"
)

type TokenType = string

var (
	AccessToken  TokenType = "accessToken"
	RefreshToken TokenType = "refreshToken"
)

type JwtHandler struct {
	ATSecretKey   string
	RTSecretKey   string
	ATExpiredTime time.Duration
	RTExpiredTime time.Duration
}

func NewJwtHandler(env *env.Env) *JwtHandler {
	return &JwtHandler{
		ATSecretKey:   env.Jwt.ATSecretKey,
		RTSecretKey:   env.Jwt.RTSecretKey,
		ATExpiredTime: time.Duration(env.Jwt.ATExpiredTime) * time.Hour,
		RTExpiredTime: time.Duration(env.Jwt.ATExpiredTime) * time.Hour,
	}
}

func (j *JwtHandler) GenerateToken(id string, tokenType TokenType) (string, error) {
	expiration := j.ATExpiredTime
	secret := j.ATSecretKey

	if tokenType == RefreshToken {
		expiration = j.RTExpiredTime
		secret = j.RTSecretKey
	}

	claim := jwt.RegisteredClaims{
		ID:        id,
		Issuer:    "dvnnfrr/ratemyubprof",
		Subject:   "user",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(secret))
}


func (j *JwtHandler) ValidateToken(reqToken string, tokenType TokenType) (string, error) {
	var claims jwt.RegisteredClaims
	secret := j.ATSecretKey

	if tokenType == RefreshToken {
		secret = j.RTSecretKey
	}

	token, err := jwt.ParseWithClaims(reqToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", jwt.ErrSignatureInvalid
	}

	return claims.ID, nil
}
