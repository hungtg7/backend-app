package service

import (
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