package main

import (
	"fmt"
	"log"

	"github.com/hungtran150/api-app/app/app_data_monitoring/config"
	"github.com/hungtran150/api-app/app/app_data_monitoring/service"
	"github.com/hungtran150/api-app/app/gateway"
	"github.com/hungtran150/api-app/lib/server"
)

var (
	cfg *config.Base
)

func main() {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	cfg = config.LoadAppConfig()

	alertServer := service.NewService(cfg)
	server, err := server.New(
		server.WithGrpcAddrListen(cfg.Server.GRPC),
		server.WithServiceServer(alertServer),
	)
	if err != nil {
		log.Fatal(err)
	}

	server.Serve()
	// TODO: split gateway server and grpc server
	err = gateway.Run(
		"dns:///" + fmt.Sprintf("%s:%d", cfg.Server.GRPC.Host, cfg.Server.GRPC.Port),
	)
	log.Fatalln(err)
}
