package middleware

import "github.com/devanfer02/ratemyubprof/pkg/config"

type Middleware struct {
	jwtHandler *config.JwtHandler
}

func NewMiddleware(jwtHandler *config.JwtHandler) *Middleware {
	return &Middleware{
		jwtHandler: jwtHandler,
	}
}