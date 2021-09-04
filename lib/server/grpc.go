package server

import (
	"context"
	"fmt"
	"log"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"

	"github.com/hungtran150/api-app/ssl"
	"github.com/hungtran150/api-app/lib/middleware"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/credentials"
)

type grpcConfig struct {
	Addr             Listen
	UnaryServerChain []grpc.ServerOption
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
			),
			// grpc.Creds(credentials.NewServerTLSFromCert(&ssl.Cert)),
		},
	}

	return config
}

func (c *grpcConfig) ServerOptions() []grpc.ServerOption {
	return []grpc.ServerOption{
			grpc_middleware.WithUnaryServerChain(
				grpc_auth.UnaryServerInterceptor(middleware.CustomAuthFunc),
			),
			// grpc.Creds(credentials.NewServerTLSFromCert(&ssl.Cert)),
		}
}

func newGrpcServer(cfg *grpcConfig, servers []ServiceServer) *grpcServer {
	s := grpc.NewServer(
		cfg.ServerOptions()...
	)
	for _, svr := range servers {
		svr.RegisterWithServer(s)
	}
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
