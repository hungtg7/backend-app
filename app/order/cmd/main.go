package main

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"

	"github.com/hungtg7/api-app/app/order/config"
	"github.com/hungtg7/api-app/app/order/repo"
	"github.com/hungtg7/api-app/app/order/service"
	"github.com/hungtg7/api-app/lib/logging"
	"github.com/hungtg7/api-app/lib/server/restapi"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	cfg *config.Base
)

func main() {
	ctx := context.Background()
	cfg = config.LoadAppConfig()

	var connectStr = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		"postgres",
		"postgres",
		"postgres",
		"postgres",
		"5432",
	)

	fmt.Println("Connecting DB...")
	db, err := gorm.Open(postgres.Open(connectStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		NowFunc:        func() time.Time { return time.Now().Local() },
	})
	if err != nil {
		fmt.Printf("%v", err)
	}

	repo := &repo.OrderRepo{Db: db}
	orderService := service.NewService(cfg, repo)
	server := restapi.New(cfg.Server.Add)
	if err != nil {
		logging.Log.Fatal(err.Error())
	}
	server.RegisterHandleFunc(
		restapi.HandleFunc{
			Pattern: "/order",
			Handler: orderService.GetOrder(ctx),
			Method:  []string{"GET"},
		},
	)

	logging.Log.Fatal(server.Serve().Error())
}
