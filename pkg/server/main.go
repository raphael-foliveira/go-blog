package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type appServer struct {
	s *http.Server
}

func New(addr string, router *chi.Mux) *appServer {
	return &appServer{&http.Server{Addr: addr, Handler: router}}
}

func (as *appServer) Run() error {
	log.Printf("about to listen on port %v\n", as.s.Addr)
	return as.s.ListenAndServe()
}
