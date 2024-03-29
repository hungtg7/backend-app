package main

import (
	"fmt"
	"time"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"gorm.io/driver/postgres"

	"github.com/hungtg7/backend-app/app/pet/config"
	"github.com/hungtg7/backend-app/app/pet/repo"
	"github.com/hungtg7/backend-app/app/pet/service"
	"github.com/hungtg7/backend-app/lib/logging"
	"github.com/hungtg7/backend-app/lib/server/grpc"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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

	var connectStr = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		"postgres",
		"postgres",
		"postgres",
		"postgres",
		"5432",
	)

	db, err := gorm.Open(postgres.Open(connectStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		NowFunc:        func() time.Time { return time.Now().Local() },
	})

	repo := &repo.PetRepo{Db: db}
	petServer := service.NewService(cfg, repo)
	server, err := grpc.New(
		grpc.WithGrpcAddrListen(cfg.Server.GRPC),
		grpc.WithServiceServer(petServer),
		grpc.WithServerInterceptor(
			grpc_zap.UnaryServerInterceptor(logging.Log),
		),
	)
	if err != nil {
		logging.Log.Fatal(err.Error())
	}

	server.Serve()
	logging.Log.Fatal(err.Error())
}
