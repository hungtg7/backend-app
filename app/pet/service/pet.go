package service

import (
	"context"
	"strconv"

	"github.com/hungtg7/backend-app/app/pet/config"
	"github.com/hungtg7/backend-app/app/pet/entity"
	"github.com/hungtg7/backend-app/app/pet/repo"
	"github.com/hungtg7/backend-app/lib/logging"
	petv1 "github.com/hungtg7/backend-app/pkg/proto_file/pet"
	"google.golang.org/grpc"
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
	resp := &petv1.GetPetResponse{}
	id, err := strconv.Atoi(req.GetPetId())
	if err != nil {
		return nil, err
	}
	p1 := &petv1.Pet{Id: int32(id)}
	resp.Pet = p1
	return resp, nil
}

func (s *Service) GetAllPet(ctx context.Context, req *petv1.GetAllPetRequest) (*petv1.GetAllPetResponse, error) {
	logging.Log.Info("Getting All pet...")
	resp := &petv1.GetAllPetResponse{}

	total := s.repo.CountAllPet(ctx)
	pet, err := s.repo.GetPets(ctx, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}
	// Serializing the struct and assigning it to body
	// p1 := &petv1.Pet{
	// 	Id:      1234,
	// 	Name:    "Heo",
	// 	PetType: "cat",
	// }
	// p2 := &petv1.Pet{
	// 	Id:      12,
	// 	Name:    "Bo",
	// 	PetType: "cat",
	// }
	respPet := []*petv1.Pet{}

	for _, p := range pet {
		petPb := &petv1.Pet{}
		petPb.Id = p.ID
		petPb.Name = p.Name
		petPb.PetType = p.PetType
		respPet = append(respPet, petPb)
	}
	// resp.Pet = []*petv1.Pet{p1, p2}
	resp.Pet = respPet
	resp.Total = total
	return resp, nil
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
