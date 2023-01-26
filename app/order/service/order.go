package service

import (
	"context"

	"github.com/hungtg7/api-app/app/order/config"
	"github.com/hungtg7/api-app/app/order/repo"
)

type Service struct {
	repo *repo.PetRepo

	config *config.Base
}

func NewService(cfg *config.Base, repo *repo.PetRepo) *Service {
	return &Service{
		config: cfg,
		repo:   repo,
	}
}

func (s *Service) Close(ctx context.Context) {
	ctx.Done()
}
