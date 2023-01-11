package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc/reflection"
)

// Server is the framework instance.
type Server struct {
	grpcServer *grpcServer
	config     *Config
}

// New creates a server intstance.
func New(opts ...Option) (*Server, error) {
	cfg := createConfig(opts)

	log.Println("Create grpc server")
	grpcServer := newGrpcServer(cfg.Grpc, cfg.ServiceServer)
	reflection.Register(grpcServer.server)

	return &Server{
		grpcServer: grpcServer,
		config:     cfg,
	}, nil
}

// Serve starts gRPC and Gateway servers.
func (s *Server) Serve() error {
	// TODO add signal stop (Grateful shutdown)
	stop := make(chan os.Signal, 1)
	errch := make(chan error)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := s.grpcServer.Serve(); err != nil {
			log.Fatal(err)
			errch <- err
		}
	}()

	for {
		select {
		case <-stop:
			log.Println("Shutting down server")

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			s.grpcServer.Shutdown(context.Background())
			s.config.ServiceServer.Close(ctx)

			s.grpcServer.Shutdown(ctx)
			return nil
		case err := <-errch:
			return err
		}
	}
}
