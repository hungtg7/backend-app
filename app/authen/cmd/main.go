package main

import (
	"context"

	"github.com/hungtg7/backend-app/app/authen/config"
	"github.com/hungtg7/backend-app/app/authen/service"
	"github.com/hungtg7/backend-app/lib/logging"
	"github.com/hungtg7/backend-app/lib/server/restapi"
)

var (
	cfg *config.Base
)

func main() {
	ctx := context.Background()
	cfg = config.LoadAppConfig()

	authenService := service.NewService(cfg)
	server := restapi.New(cfg.Server.Addr)

	server.RegisterHandleFunc(
		restapi.HandleFunc{
			Pattern: "/authen",
			Handler: authenService.Authenticate(ctx),
			Method:  []string{"GET"},
		},
	)

	logging.Log.Fatal(server.Serve().Error())
}
