package restapi

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	Add        string
	httpServer *http.Server
	mux        *mux.Router
}

func New() *Server {
	return &Server{
		mux: mux.NewRouter(),
	}
}

type HandleFunc struct {
	pattern string
	handler func(http.ResponseWriter, *http.Request)
	method  []string
}

func (s *Server) RegisterHandleFunc(function ...HandleFunc) {
	for _, f := range function {
		s.mux.HandleFunc(f.pattern, f.handler).Methods(f.method...)
	}
}

func (s *Server) Serve() error {
	// TODO add signal stop (Grateful shutdown)
	stop := make(chan os.Signal, 1)
	errch := make(chan error)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := http.ListenAndServe(s.Add, s.mux); err != nil {
			log.Fatal(err)
			errch <- err
		}
	}()

	for {
		select {
		case <-stop:
			log.Println("Shutting down server")

			_, cancel := context.WithTimeout(context.Background(), 30*time.Second)

			defer cancel()
			return nil
		case err := <-errch:
			return err
		}
	}
}
