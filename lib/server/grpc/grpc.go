package grpc

import (
	"context"
	"fmt"
	"log"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/hungtg7/backend-app/lib/middleware"
	"google.golang.org/grpc"
)

type grpcConfig struct {
	Addr                   Listen
	UnaryServerChain       []grpc.ServerOption
	UnaryServerInterceptor []grpc.UnaryServerInterceptor
}

// grpcServer wraps grpc.Server setup process.
type grpcServer struct {
	server *grpc.Server
	config *grpcConfig
}

func createDefaultGrpcConfig() *grpcConfig {
	config := &grpcConfig{
		Addr: Listen{
			Host: "0.0.0.0",
			Port: 10443,
		},
		UnaryServerChain: []grpc.ServerOption{
			grpc_middleware.WithUnaryServerChain(
				grpc_auth.UnaryServerInterceptor(middleware.CustomAuthFunc),
				grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			),
		},
	}

	return config
}

func (c *grpcConfig) buildUnaryServerInterceptor() []grpc.UnaryServerInterceptor {
	var placeHolder []grpc.UnaryServerInterceptor

	placeHolder = append(
		placeHolder,
		grpc_auth.UnaryServerInterceptor(middleware.CustomAuthFunc),
	)
	placeHolder = append(
		placeHolder,
		grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
	)
	placeHolder = append(
		placeHolder,
		c.UnaryServerInterceptor...,
	)

	return placeHolder
}

func (c *grpcConfig) serverOptions() []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			c.buildUnaryServerInterceptor()...,
		),
	}
}

func newGrpcServer(cfg *grpcConfig, server ServiceServer) *grpcServer {
	s := grpc.NewServer(
		cfg.serverOptions()...,
	)

	server.RegisterWithServer(s)

	return &grpcServer{
		server: s,
		config: cfg,
	}
}

// Serve implements Server.Server
func (s *grpcServer) Serve() error {
	l, err := s.config.Addr.CreateListener()
	if err != nil {
		return fmt.Errorf("failed to create listener %w", err)
	}

	log.Println("gRPC server is starting ", l.Addr())

	err = s.server.Serve(l)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to serve gRPC server %w", err)
	}
	log.Println("gRPC server ready")

	return nil
}

// Shutdown
func (s *grpcServer) Shutdown(ctx context.Context) {
	s.server.GracefulStop()
}
