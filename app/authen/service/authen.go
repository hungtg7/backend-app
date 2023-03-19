package service

import (
	"context"
	"net/http"

	"github.com/hungtg7/backend-app/app/authen/config"
)

type Service struct {
	config *config.Base
}

func NewService(cfg *config.Base) *Service {
	return &Service{
		config: cfg,
	}
}

func (s *Service) Authenticate(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}
}
