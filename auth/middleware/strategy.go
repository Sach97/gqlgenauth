package middleware

import (
	"context"
	"net/http"
)

type Strategy interface {
	GetUserID(ctx context.Context) string
	AuthMiddleware(next http.Handler) http.Handler
}

type RouterStrategy struct {
	Strategy Strategy
}

func (s *RouterStrategy) GetUserID(ctx context.Context) string {
	return s.Strategy.GetUserID(ctx)
}

func (s *RouterStrategy) AuthMiddleware(next http.Handler) http.Handler {
	return s.Strategy.AuthMiddleware(next)
}
