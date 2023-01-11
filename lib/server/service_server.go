package server

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// ServiceServer
type ServiceServer interface {
	RegisterWithServer(*grpc.Server)
	Close(context.Context)
}
