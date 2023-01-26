package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"

	"github.com/hungtg7/api-app/app/order/config"
	"github.com/hungtg7/api-app/app/order/repo"
	"github.com/hungtg7/api-app/lib/logging"
	"github.com/hungtg7/api-app/lib/server/restapi"
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

	_ = &repo.PetRepo{Db: db}
	// petServer := service.NewService(cfg, repo)
	server := restapi.New()
	if err != nil {
		logging.Log.Fatal(err.Error())
	}

	server.Serve()
	logging.Log.Fatal(err.Error())
}
