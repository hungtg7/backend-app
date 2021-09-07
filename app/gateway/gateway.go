package gateway

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	// "google.golang.org/grpc/credentials"

	"github.com/hungtran150/api-app/lib/logging"
	alert_bp "github.com/hungtran150/api-app/proto/v1/app_data_monitoring_bp"
	"github.com/hungtran150/api-app/ssl"
)

// Run runs the gRPC-Gateway, dialling the provided address.
func Run(dialAddr string) error {
	log := logging.Log
	// Create a client connection to the gRPC Server we just started.
	// This is where the gRPC-Gateway proxies the requests.
	// Make sure gRPC server work proberly and accesible
	conn, err := grpc.DialContext(
		context.Background(),
		dialAddr,
		// grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(ssl.CertPool, "")),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		return fmt.Errorf("failed to dial server: %w", err)
	}

	gwmux := runtime.NewServeMux()
	// Register AlertServiceHandler
	// TODO Refactor RegisterAlertServiceHandler to RegisterAlertServiceHandlerEndPoint
	err = alert_bp.RegisterSlackAlertServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		return fmt.Errorf("failed to register gateway: %w", err)
	}
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "11000"
	}
	gatewayAddr := "0.0.0.0:" + port
	gwServer := &http.Server{
		Addr: gatewayAddr,
		Handler: FormWrapper(gwmux, logging.Log),
	}
	// Empty parameters mean use the TLS Config specified with the server.
	if strings.ToLower(os.Getenv("SERVE_HTTP")) == "true" {
		log.Info(fmt.Sprint("Serving gRPC-Gateway and OpenAPI Documentation on http://", gatewayAddr))
		return fmt.Errorf("serving gRPC-Gateway server: %w", gwServer.ListenAndServe())
	}

	gwServer.TLSConfig = &tls.Config{
		Certificates: []tls.Certificate{ssl.Cert},
	}
	log.Info(fmt.Sprint("Serving gRPC-Gateway and OpenAPI Documentation on https://", gatewayAddr))
	return fmt.Errorf("serving gRPC-Gateway server use TLS Config: %w", gwServer.ListenAndServeTLS("", ""))
}
