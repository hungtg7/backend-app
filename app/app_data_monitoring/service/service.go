package service

import (
	"github.com/hungtran150/api-app/app/app_data_monitoring/config"
	"github.com/hungtran150/api-app/proto/v1/app_data_monitoring_bp"
	"google.golang.org/grpc"
)

type Service struct {
	app_data_monitoring_bp.UnimplementedAlertServiceServer

	config     *config.Base
}

func NewService(cfg *config.Base,) *Service {
	return &Service{
		config:     cfg,
	}
}

// RegisterWithServer implementing service server interface
func (s *Service) RegisterWithServer(server *grpc.Server) {
	app_data_monitoring_bp.RegisterAlertServiceServer(server, s)
}