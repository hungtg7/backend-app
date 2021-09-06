package main

import (
	"fmt"
	"io/ioutil"
	// "log"
	"os"

	"github.com/hungtran150/api-app/app/app_data_monitoring/config"
	"github.com/hungtran150/api-app/app/app_data_monitoring/service"
	"github.com/hungtran150/api-app/app/gateway"
	"github.com/hungtran150/api-app/lib/server"
	"google.golang.org/grpc/grpclog"
)

var (
	cfg *config.Base
)

func main() {
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)
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
