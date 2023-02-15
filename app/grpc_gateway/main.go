package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	petv1 "github.com/hungtg7/backend-app/pkg/proto_file/pet"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint        = flag.String("grpc-server-endpoint", "pet:10550", "gRPC server endpoint")
	options            string = "OPTIONS"
	allow_origin       string = "Access-Control-Allow-Origin"
	allow_methods      string = "Access-Control-Allow-Methods"
	allow_headers      string = "Access-Control-Allow-Headers"
	allow_credentials  string = "Access-Control-Allow-Credentials"
	expose_headers     string = "Access-Control-Expose-Headers"
	credentials        string = "true"
	origin             string = "Origin"
	methods            string = "POST, GET, OPTIONS, PUT, DELETE, HEAD, PATCH"

	// If you want to expose some other headers add it here
	headers string = "Access-Control-Allow-Origin, Accept, Accept-Encoding, Authorization, Content-Length, Content-Type, X-CSRF-Token"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := petv1.RegisterPetStoreHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":9090", CORS(mux))
}

// Handler will allow cross-origin HTTP requests
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set allow origin to match origin of our request or fall back to *
		if o := r.Header.Get(origin); o != "" {
			w.Header().Set(allow_origin, o)
		} else {
			w.Header().Set(allow_origin, "*")
		}

		// Set other headers
		w.Header().Set(allow_headers, headers)
		w.Header().Set(allow_methods, methods)
		w.Header().Set(allow_credentials, credentials)
		w.Header().Set(expose_headers, headers)

		// If this was preflight options request let's write empty ok response and return
		if r.Method == options {
			w.WriteHeader(http.StatusOK)
			w.Write(nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
