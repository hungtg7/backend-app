package service

import (
	"context"
	"fmt"

	"github.com/hungtg7/api-app/app/pet/config"
	"github.com/hungtg7/api-app/app/pet/entity"
	"github.com/hungtg7/api-app/app/pet/repo"
	petv1 "github.com/hungtg7/api-app/proto/pet"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	petv1.UnimplementedPetStoreServer
	repo *repo.PetRepo

	config *config.Base
}

func NewService(cfg *config.Base, repo *repo.PetRepo) *Service {
	return &Service{
		config: cfg,
		repo:   repo,
	}
}

// RegisterWithServer implementing service server interface
func (s *Service) RegisterWithServer(server *grpc.Server) {
	petv1.RegisterPetStoreServer(server, s)
}

func (s *Service) Close(ctx context.Context) {
	ctx.Done()
}

func (s *Service) GetPet(ctx context.Context, req *petv1.GetPetRequest) (*petv1.GetPetResponse, error) {
	var resp *petv1.GetPetResponse
	pet, err := s.repo.GetPetByID(ctx, req.GetPetId())
	if err != nil {
		return nil, err
	}
	// Serializing the struct and assigning it to body
	p1 := &petv1.Pet{PetType: pet.PetType, Id: pet.ID, CreatedAt: timestamppb.New(pet.CreatedAt)}
	resp.Pet = p1
	return nil, fmt.Errorf("error")
}

func (s *Service) CreatePet(ctx context.Context, req *petv1.CreatePetRequest) (*petv1.CreatePetResponse, error) {
	var resp *petv1.CreatePetResponse
	e := &entity.Pet{PetType: req.PetType.String(), Name: req.Name}
	if err := s.repo.Add(ctx, e); err != nil {
		return nil, err
	}
	resp.Code = 200
	return resp, nil
}