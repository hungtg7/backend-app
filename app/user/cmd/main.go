package main

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/hungtg7/backend-app/app/user/config"
	"github.com/hungtg7/backend-app/app/user/repo"
	"github.com/hungtg7/backend-app/app/user/service"
	"github.com/hungtg7/backend-app/lib/logging"
	"github.com/hungtg7/backend-app/lib/server/restapi"
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

	db, err := gorm.Open(postgres.Open(connectStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		NowFunc:        func() time.Time { return time.Now().Local() },
	})

	if err != nil {
		logging.Log.Fatal(err.Error())
		return
	}

	repo := &repo.UserRepo{Db: db}
	authenService := service.NewService(cfg, repo)
	server := restapi.New(cfg.Server.Addr)

	server.RegisterHandleFunc(
		restapi.HandleFunc{
			Pattern: "/auth/google/login",
			Handler: authenService.GoogleLogin(ctx),
			Method:  []string{"GET"},
		},
		restapi.HandleFunc{
			Pattern: "/auth/google/callback",
			Handler: authenService.GoogleCallback(ctx),
			Method:  []string{"GET"},
		},
	)

	logging.Log.Fatal(server.Serve().Error())
}
