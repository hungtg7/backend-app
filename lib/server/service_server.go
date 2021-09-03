package server

import (
	"google.golang.org/grpc"
)

// ServiceServer
type ServiceServer interface {
	RegisterWithServer(*grpc.Server)
}
