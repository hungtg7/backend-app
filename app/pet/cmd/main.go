package main

import (
	"fmt"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/hungtg7/api-app/app/pet/config"
	"github.com/hungtg7/api-app/app/pet/service"
	"github.com/hungtg7/api-app/lib/logging"
	"github.com/hungtg7/api-app/lib/server"
)

var (
	cfg *config.Base
)

func main() {
	// initialize logger
	if err := logging.Init(0, ""); err != nil {
		fmt.Printf("failed to initialize logger: %v", err)
	}
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	cfg = config.LoadAppConfig()

	alertServer := service.NewService(cfg)
	server, err := server.New(
		server.WithGrpcAddrListen(cfg.Server.GRPC),
		server.WithServiceServer(alertServer),
		server.WithServerInterceptor(
			grpc_zap.UnaryServerInterceptor(logging.Log),
		),
	)
	if err != nil {
		logging.Log.Fatal(err.Error())
	}

	server.Serve()
	logging.Log.Fatal(err.Error())
}
