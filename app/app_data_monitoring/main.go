package main

import (
	"fmt"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/hungtran150/api-app/app/app_data_monitoring/config"
	"github.com/hungtran150/api-app/app/app_data_monitoring/service"
	"github.com/hungtran150/api-app/app/gateway"
	"github.com/hungtran150/api-app/lib/logging"
	"github.com/hungtran150/api-app/lib/server"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/codes"
)

var (
	cfg *config.Base
)

func main() {
	// initialize logger
	if err := logging.Init(-1, ""); err != nil {
		fmt.Errorf("failed to initialize logger: %v", err)
	}
	zapOption := []grpc_zap.Option{
		grpc_zap.WithLevels(codeToLevel),
	}
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	cfg = config.LoadAppConfig()

	alertServer := service.NewService(cfg)
	server, err := server.New(
		server.WithGrpcAddrListen(cfg.Server.GRPC),
		server.WithServiceServer(alertServer),
		server.WithServerInterceptor(
			grpc_zap.UnaryServerInterceptor(logging.Log, zapOption...),
		),
	)
	if err != nil {
		logging.Log.Fatal(err.Error())
	}

	server.Serve()
	// TODO: split gateway server and grpc server
	err = gateway.Run(
		"dns:///" + fmt.Sprintf("%s:%d", cfg.Server.GRPC.Host, cfg.Server.GRPC.Port),
	)
	logging.Log.Fatal(err.Error())
}

// codeToLevel redirects OK to DEBUG level logging instead of INFO
// This is example how you can log several gRPC code results
func codeToLevel(code codes.Code) zapcore.Level {
	if code == codes.OK {
		// It is DEBUG
		return zap.DebugLevel
	}
	return grpc_zap.DefaultCodeToLevel(code)
}