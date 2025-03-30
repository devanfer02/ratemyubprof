package util

import (
	"time"

	"github.com/devanfer02/ratemyubprof/internal/infra/env"
	"github.com/golang-jwt/jwt/v5"
)

type JwtHandler struct {
	SecretKey   string
	ExpiredTime time.Duration
}

func NewJwtHandler(env *env.Env) *JwtHandler {
	return &JwtHandler{
		SecretKey:   env.Jwt.SecretKey,
		ExpiredTime: time.Duration(env.Jwt.ExpiredTime) * time.Hour,
	}
}

func (j *JwtHandler) GenerateToken(id string) (string, error) {
	claim := &jwt.RegisteredClaims{
		ID: id,
		Issuer: "dvnnfrr/ratemyubprof",
		Subject: "user",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
	
		return "", err
	}

	return tokenString, nil
}

func (j *JwtHandler) ValidateToken(reqToken string) (string, error) {
	var claims jwt.RegisteredClaims

	token, err := jwt.ParseWithClaims(reqToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})

	if err != nil {
		return "", err 
	}

	if !token.Valid {
		return "", jwt.ErrSignatureInvalid
	}

	return claims.ID, nil 
}