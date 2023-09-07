package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/raphael-foliveira/blog-backend/pkg/server"
)

func main() {
	router := startRoutes()
	appServer := server.New(":3000", router)
	appServer.Run()
}

func startRoutes() *chi.Mux {
	router := chi.NewRouter()
	attachMiddleware(router)

	return router
}

func attachMiddleware(r *chi.Mux) {
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"*"},
		AllowedHeaders: []string{"*"},
	}))
	r.Use(middleware.Logger)
}
