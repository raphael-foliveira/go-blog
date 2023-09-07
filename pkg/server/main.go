package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type appServer struct {
	s *http.Server
}

func New(addr string, router *chi.Mux) *appServer {
	return &appServer{&http.Server{Addr: addr}}
}

func (as *appServer) Run() error {
	return as.s.ListenAndServe()
}
