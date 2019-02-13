package middleware

import (
	"context"
	"net/http"
)

type Strategy interface {
	GetUserID(ctx context.Context) (string, error)
	AuthMiddleware(next http.Handler) http.Handler
}

type Strategy struct {
	Strategy Strategy
}

func (s *Strategy) GetUserID(ctx context.Context) (string, error) {
	return s.Strategy.GetUserID(ctx)
}

func (s *Strategy) AuthMiddleware(next http.Handler) http.Handler {
	return s.Strategy.AuthMiddleware(next)
}
