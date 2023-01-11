package service

import (
	"context"
	"fmt"

	"github.com/hungtg7/api-app/app/pet/config"
	"github.com/hungtg7/api-app/app/pet/repo"
	petv1 "github.com/hungtg7/api-app/proto/pet"
	"google.golang.org/grpc"
)

type Service struct {
	petv1.UnimplementedPetStoreServer
	repo repo.PetRepo

	config *config.Base
}

func NewService(cfg *config.Base) *Service {
	return &Service{
		config: cfg,
	}
}

// RegisterWithServer implementing service server interface
func (s *Service) RegisterWithServer(server *grpc.Server) {
	petv1.RegisterPetStoreServer(server, s)
}

func (s *Service) GetPet(ctx context.Context, req *petv1.GetPetRequest) (*petv1.GetPetResponse, error) {
	return nil, fmt.Errorf("error")
}
