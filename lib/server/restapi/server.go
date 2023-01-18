package restapi

import "net/http"

type Server struct {
	Add string
	mux *http.ServeMux
}

func New() *Server {
	return &Server{
	mux: http.NewServeMux(),
}}

type HandleFunc struct {
	pattern string
	handler func(http.ResponseWriter, *http.Request)
	method []string
}

func (s *Server) RegisterHandleFunc(function ...HandleFunc) {}

func (s *Server) Serve() error {
	return http.ListenAndServe(s.Add, s.mux)
}
